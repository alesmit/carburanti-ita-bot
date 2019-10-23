package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"errors"
	"strconv"
	"strings"
)

func handleCallbackQuery(update *tgbotapi.Update, api *tgbotapi.BotAPI) error {
	data := strings.Split(update.CallbackQuery.Data, ";")

	// lat
	lat, err := strconv.ParseFloat(data[0], 64)
	if err != nil {
		return errors.New("unable to parse latitude")
	}

	lon, err := strconv.ParseFloat(data[1], 64)
	if err != nil {
		return errors.New("unable to parse longitude")
	}

	location := tgbotapi.NewLocation(update.CallbackQuery.Message.Chat.ID, lat, lon)
	api.Send(location)
	return nil
}
