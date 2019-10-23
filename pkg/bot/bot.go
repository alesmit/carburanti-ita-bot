package bot

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleUpdate(update *tgbotapi.Update, api *tgbotapi.BotAPI) {

	// debug
	if api.Debug {
		updateJson, _ := json.Marshal(update)
		log.Println("RECEIVED JSON:", string(updateJson))
	}

	// handle location
	if update.Message != nil && update.Message.Location != nil {
		handleLocation(update, api)
		return
	}

	// handle query
	if update.CallbackQuery != nil {
		handleCallbackQuery(update, api)
		return
	}

	if update.Message != nil {
		handleDefaultMessage(update, api)
		return
	}
}
