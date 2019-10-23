package bot

import (
	"encoding/json"
	"github.com/alesmit/fuel-master/pkg/dataset"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"

	"errors"
)

func HandleUpdate(update *tgbotapi.Update, api *tgbotapi.BotAPI) {

	// debug
	if api.Debug && update.Message != nil {
		updateJson, _ := json.Marshal(update)
		api.Send(tgbotapi.NewMessage(update.Message.Chat.ID, string(updateJson)))
	}

	// handle location
	if update.Message != nil && update.Message.Location != nil {
		if err := dataset.SyncDatasets(); err != nil {
			handleError(errors.New("unable to sync datasets"), update, api)
			return
		}
		if err := handleLocation(update, api); err != nil {
			handleError(err, update, api)
		}

		return
	}

	// handle query
	if update.CallbackQuery != nil {
		updateJson, _ := json.Marshal(update)
		log.Println("RECEIVED JSON:", string(updateJson))
		/*
			if err := dataset.SyncDatasets(); err != nil {
				handleError(errors.New("unable to sync datasets"), update, api)
				return
			}
			if err := handleCallbackQuery(update, api); err != nil {
				handleError(err, update, api)
			}
		*/

		return
	}

	handleDefault(update, api)
}

func handleError(err error, update *tgbotapi.Update, api *tgbotapi.BotAPI) {
	text := capitalize(err.Error()) + ". Please try again later."
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	api.Send(msg)
}
