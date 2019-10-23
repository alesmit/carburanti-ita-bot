package bot

import (
	"github.com/alesmit/fuel-master/pkg/i18n"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func handleDefaultMessage(update *tgbotapi.Update, api *tgbotapi.BotAPI) {
	lang := "en"
	if update.Message != nil {
		lang = update.Message.From.LanguageCode
	}

	text := i18n.T(lang, i18n.TextPleaseSendLocation)
	api.Send(tgbotapi.NewMessage(update.Message.Chat.ID, text))
}
