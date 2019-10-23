package i18n

import "strings"

const (
	BtnGetMap = iota + 1
	TextPleaseSendLocation
)

var translations = map[string]map[int]string{
	"en": {
		BtnGetMap:              "Map",
		TextPleaseSendLocation: "Please send your location",
	},
	"it": {
		BtnGetMap:              "Mappa",
		TextPleaseSendLocation: "Invia la tua posizione",
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
