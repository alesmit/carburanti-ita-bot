package i18n

var translations = map[string]map[string]string{
	"en": {
		"asdasd": "sasd",
	},
}

func T(lang string, key string) string {
	return translations[lang][key]
}
