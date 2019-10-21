package main

import (
	"github.com/alesmit/fuel-master/pkg/dataset"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

func main() {
	// get env variables
	port := os.Getenv("PORT")
	token := os.Getenv("TG_BOT_TOKEN")
	url := os.Getenv("HEROKU_URL")

	// init the tgbotapi lib
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	// using webhook
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(url + bot.Token))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe(":"+port, nil)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msgText := "please send your location"

		if update.Message.Location != nil {
			stations, err := processLocation(update.Message.Location)
			if err != nil {
				msgText = wrapError(err)
			}

			j, err := json.Marshal(stations)
			if err != nil {
				msgText = wrapError(err)
			}

			msgText = string(j)

		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
		bot.Send(msg)
	}
}

func processLocation(location *tgbotapi.Location) ([]dataset.StationWithPrices, error) {
	if err := dataset.SyncDatasets(); err != nil {
		return nil, errors.New("unable to sync datasets")
	}

	req := &dataset.GetClosestStationRequest{
		Lat: location.Latitude,
		Lon: location.Longitude,
		Qty: 3,
	}

	stationsWithPrices, err := dataset.GetClosestStationsWithPrices(req)
	if err != nil {
		return nil, errors.New("unable to get closest stations")
	}

	return stationsWithPrices, nil
}

func wrapError(e error) string {
	return e.Error() + ". please try again later"
}
