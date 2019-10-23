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
	if api.Debug {
		updateJson, _ := json.Marshal(update)
		log.Println("RECEIVED JSON:", string(updateJson))
	}

	// handle location
	if update.Message != nil && update.Message.Location != nil {
		if err := dataset.SyncDatasets(); err != nil {
			sendError(errors.New("unable to sync datasets"), update.Message.Chat.ID, api)
			return
		}
		if err := handleLocation(update, api); err != nil {
			sendError(err, update.Message.Chat.ID, api)
		}

		return
	}

	// handle query
	if update.CallbackQuery != nil {
		if err := dataset.SyncDatasets(); err != nil {
			sendError(errors.New("unable to sync datasets"), update.CallbackQuery.Message.Chat.ID, api)
			return
		}
		if err := handleCallbackQuery(update, api); err != nil {
			sendError(err, update.CallbackQuery.Message.Chat.ID, api)
		}

		return
	}

	if update.Message != nil {
		handleDefaultMessage(update, api)
	}
}

func sendError(err error, chatId int64, api *tgbotapi.BotAPI) {
	text := capitalize(err.Error()) + ". Please try again later."
	msg := tgbotapi.NewMessage(chatId, text)
	api.Send(msg)
}
