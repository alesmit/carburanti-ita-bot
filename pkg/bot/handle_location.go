package bot

import (
	"errors"
	"github.com/alesmit/fuel-master/pkg/i18n"

	"github.com/alesmit/fuel-master/pkg/dataset"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func handleLocation(update *tgbotapi.Update, api *tgbotapi.BotAPI) {
	chatId := update.Message.Chat.ID

	if err := dataset.SyncDatasets(); err != nil {
		sendError(errors.New("unable to sync datasets"), chatId, api)
		return
	}
	if err := sendClosestStationsWithPrices(update, api); err != nil {
		sendError(err, chatId, api)
		return
	}
}

func sendClosestStationsWithPrices(update *tgbotapi.Update, api *tgbotapi.BotAPI) error {
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

		btn := tgbotapi.NewInlineKeyboardButtonData(i18n.T(update.Message.From.LanguageCode, i18n.BtnGetMap), swp.Station.Id)
		row := tgbotapi.NewInlineKeyboardRow(btn)
		markup := tgbotapi.NewInlineKeyboardMarkup(row)

		msg := tgbotapi.MessageConfig{
			Text:      formatStationWithPrices(&swp),
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
