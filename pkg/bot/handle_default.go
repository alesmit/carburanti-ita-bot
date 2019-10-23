package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func handleDefaultMessage(update *tgbotapi.Update, api *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please send your location")
	api.Send(msg)
}
