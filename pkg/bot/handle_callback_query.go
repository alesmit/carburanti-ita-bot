package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func handleCallbackQuery(update *tgbotapi.Update, api *tgbotapi.BotAPI) error {
	api.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Soon available"))
	return nil
}
