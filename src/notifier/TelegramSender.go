package notifier

import (
	"KworkTasksNotifier/src/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

// SendMessage отправляет информацию в телеграм канал
func SendMessage(data models.KworkResponseModel) (err error) {
	bot, err := tgbotapi.NewBotAPI(GetEnvVariable("BOT_TOKEN"))
	bot.Debug = false

	if err != nil {
		return err
	}

	channelId, _ := strconv.ParseInt(GetEnvVariable("CHANNEL_ID"), 10, 64)
	msg := tgbotapi.NewMessage(channelId, BuildMessageBody(data))
	msg.ParseMode = "HTML"

	var markup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Перейти", "https://kwork.ru"+data.URL),
		),
	)
	msg.ReplyMarkup = markup

	_, err = bot.Send(msg)

	if err != nil {
		return err
	}

	return nil
}

// BuildMessageBody создает строку для отправки уведомления в
// телеграм с информацией о заказе
func BuildMessageBody(model models.KworkResponseModel) (msg string) {
	msg = "<b>🔔 Новый заказ!</b>\n" +
		"📋 <b>Название: " + model.Name + "\n</b>" +
		"✍️ <b>Описание: </b><i>" + model.Description + "</i>\n" +
		"💵 <b>Бюджет: " + model.PriceLimit + " ₽</b>\n" +
		"💰 <b>Допустимый бюджет: " + strconv.Itoa(model.PossiblePriceLimit) + " ₽</b>"

	if len(model.Files) != 0 {
		msg += "\n📂 Прикрепленные файлы:"
		for _, payload := range model.Files {
			msg += "\n<a href='" + payload.Url + "'>" + payload.Name + "</a>"
		}
	}
	return msg
}

// GetEnvVariable используя пакет godot получает значение по ключу
// из .env файла (переменные среды)
func GetEnvVariable(key string) string {
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
