package bot

import (
	"fmt"
	"github.com/alesmit/fuel-master/pkg/dataset"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"errors"
)

func handleLocation(update *tgbotapi.Update, api *tgbotapi.BotAPI) error {
	req := &dataset.GetClosestStationRequest{
		Lat: update.Message.Location.Latitude,
		Lon: update.Message.Location.Longitude,
		Qty: 3,
	}

	stationsWithPrices, err := dataset.GetClosestStationsWithPrices(req)
	if err != nil {
		return errors.New("unable to get closest stations")
	}

	for _, swp := range stationsWithPrices {

		btn := tgbotapi.NewInlineKeyboardButtonData("Map", fmt.Sprint(swp.Station.Lat, ";", swp.Station.Lon))
		row := tgbotapi.NewInlineKeyboardRow(btn)
		markup := tgbotapi.NewInlineKeyboardMarkup(row)

		msg := tgbotapi.MessageConfig{
			Text:      format(&swp),
			ParseMode: tgbotapi.ModeMarkdown,
			BaseChat: tgbotapi.BaseChat{
				ChatID:      update.Message.Chat.ID,
				ReplyMarkup: markup,
			},
		}

		api.Send(msg)
	}

	return nil
}
