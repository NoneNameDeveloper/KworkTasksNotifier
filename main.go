package main

import (
	"KworkTasksNotifier/src/notifier"
)

var exit = make(chan bool)

func main() {
	categoryIdScripts := 41 // боты и скрипты
	// categoryIdDessktop := 80 // десктоп программированеи
	go notifier.SheduleTask(categoryIdScripts)
	// go notifier.SheduleTask(categoryIdDessktop)

	<-exit

}
