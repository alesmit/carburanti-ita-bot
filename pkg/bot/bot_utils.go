package bot

import (
	"fmt"
	"strings"

	"github.com/alesmit/fuel-master/pkg/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func formatStationWithPrices(s *model.StationWithPrices) string {
	out := formatStation(&s.Station)

	for _, p := range s.Prices {
		out += "\n" + p.FuelType + ": € " + fmt.Sprintf("%.3f", p.Price) + "/lt."
	}

	return out
}

func formatStation(s *model.Station) string {
	out := fmt.Sprintln("*", s.Name, "*")
	out += fmt.Sprintln("_", s.Address, "_")
	return out
}

func capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func sendError(err error, chatId int64, api *tgbotapi.BotAPI) {
	text := capitalize(err.Error()) + ". Please try again later."
	msg := tgbotapi.NewMessage(chatId, text)
	api.Send(msg)
}
