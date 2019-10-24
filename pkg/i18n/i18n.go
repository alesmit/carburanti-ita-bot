package i18n

import "strings"

const (
	BtnGetMap = iota + 1
	TextPleaseSendLocation
	UnableToGetPricesInfo
)

var translations = map[string]map[int]string{
	"en": {
		BtnGetMap:              "Map",
		TextPleaseSendLocation: "Please send your location",
		UnableToGetPricesInfo:  "Unable to get prices information",
	},
	"it": {
		BtnGetMap:              "Mappa",
		TextPleaseSendLocation: "Invia la tua posizione",
		UnableToGetPricesInfo:  "Informazioni sul prezzo non disponibili",
	},
}

func T(lang string, key int) string {
	lang = strings.ToLower(lang)

	// default to english if unsupported lang
	if translations[lang] == nil {
		lang = "en"
	}

	return translations[lang][key]
}
