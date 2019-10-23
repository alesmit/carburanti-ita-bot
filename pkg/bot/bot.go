package bot

import (
	"encoding/json"
	"github.com/alesmit/fuel-master/pkg/dataset"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"errors"
)

func HandleUpdate(update *tgbotapi.Update, api *tgbotapi.BotAPI) {
	if update.Message == nil {
		return
	}

	if api.Debug {
		updateJson, _ := json.Marshal(update)
		api.Send(tgbotapi.NewMessage(update.Message.Chat.ID, string(updateJson)))
	}

	if update.Message.Location != nil {
		if err := dataset.SyncDatasets(); err != nil {
			handleError(errors.New("unable to sync datasets"), update, api)
			return
		}
		if err := handleLocation(update, api); err != nil {
			handleError(err, update, api)
		}

		return
	}

	/*
		if update.CallbackQuery != nil {
			if err := dataset.SyncDatasets(); err != nil {
				handleError(errors.New("unable to sync datasets"), update, api)
				return
			}
			if err := handleCallbackQuery(update, api); err != nil {
				handleError(err, update, api)
			}

			return
		}
	*/

	handleDefault(update, api)
}

func handleError(err error, update *tgbotapi.Update, api *tgbotapi.BotAPI) {
	text := capitalize(err.Error()) + ". Please try again later."
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	api.Send(msg)
}
