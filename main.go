package main

import (
	"log"
	"net/http"
	"os"

	"github.com/alesmit/fuel-master/pkg/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// get env variables
	port := os.Getenv("PORT")
	token := os.Getenv("TG_BOT_TOKEN")
	url := os.Getenv("HEROKU_URL")
	debug := os.Getenv("DEBUG")

	// init the tgbotapi lib
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// set debug mode
	api.Debug = debug == "1"

	// use webhook
	_, err = api.SetWebhook(tgbotapi.NewWebhook(url + api.Token))
	if err != nil {
		log.Fatal(err)
	}

	// start server
	go http.ListenAndServe(":"+port, nil)

	// handle updates
	updates := api.ListenForWebhook("/" + api.Token)
	for update := range updates {
		bot.HandleUpdate(&update, api)
	}
}
