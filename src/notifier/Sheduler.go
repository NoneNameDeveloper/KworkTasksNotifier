package notifier

import (
	"KworkTasksNotifier/src/engine"
	"KworkTasksNotifier/src/models"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func SheduleTask(categoryId int) {
	for {

		objects, err := engine.GetData(categoryId)

		if err != nil {
			fmt.Println(err)
		} else {
			// fmt.Println(objects[0].Name + " " + objects[0].DateCreate)
			if IsNewOrder(objects[1]) == true {
				err := SendMessage(objects[1])
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		time.Sleep(time.Second * 10) // спим 10 секунд между проверками
	}
}

func IsNewOrder(model models.KworkResponseModel) bool {
	var empty bool

	check, err := os.Stat("src/data/cache.txt")

	if check.Size() == 0 {
		empty = true
	}
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("file does not exist, creating")
			_, _ = os.Create("data/cache.txt")
		} else {
			fmt.Println("Unhandled error")
		}
	}

	file, _ := os.OpenFile("data/cache.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	defer file.Close()

	if empty == true {
		_, err := file.WriteString(strconv.Itoa(model.ID) + "\n")
		if err != nil {
			fmt.Println(err)
		}

		return true
	}

	fileScanner := bufio.NewScanner(file)
	var flag bool
	for fileScanner.Scan() {
		if strconv.Itoa(model.ID) != strings.TrimSpace(fileScanner.Text()) {
			flag = true
		} else {
			flag = false
		}
		if err := fileScanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	if flag == true {
		file.WriteString(strconv.Itoa(model.ID) + "\n")
		return true
	} else {
		return false
	}

}
