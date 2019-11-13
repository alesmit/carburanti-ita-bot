package bot

import (
	"fmt"
	"strings"

	"github.com/alesmit/carburanti-ita-bot/pkg/i18n"
	"github.com/alesmit/carburanti-ita-bot/pkg/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func formatStationWithPrices(s *model.StationWithPrices, lang string) string {
	out := formatStation(&s.Station)

	if len(s.Prices) == 0 {
		out += "\n" + fmt.Sprintln(i18n.T(lang, i18n.UnableToGetPricesInfo))
	}

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
