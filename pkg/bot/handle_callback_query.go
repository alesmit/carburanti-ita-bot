package bot

import (
	"errors"

	"github.com/alesmit/fuel-master/pkg/dataset"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func handleCallbackQuery(update *tgbotapi.Update, api *tgbotapi.BotAPI) {
	chatId := update.CallbackQuery.Message.Chat.ID

	if err := dataset.SyncDatasets(); err != nil {
		sendError(errors.New("unable to sync datasets"), chatId, api)
		return
	}
	if err := sendStationInfo(update.CallbackQuery.Data, chatId, api); err != nil {
		sendError(err, chatId, api)
		return
	}
}

func sendStationInfo(stationId string, chatId int64, api *tgbotapi.BotAPI) error {
	station, err := dataset.GetStationById(stationId)
	if err != nil {
		return err
	}

	// send station details
	msg := tgbotapi.MessageConfig{
		Text:      formatStation(station),
		ParseMode: tgbotapi.ModeMarkdown,
		BaseChat: tgbotapi.BaseChat{
			ChatID: chatId,
		},
	}

	api.Send(msg)

	// send station position
	location := tgbotapi.NewLocation(chatId, station.Lat, station.Lon)
	api.Send(location)
	return nil
}
