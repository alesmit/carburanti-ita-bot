package model

type Station struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

type StationWithPrices struct {
	Station Station `json:"station"`
	Prices  []Price `json:"prices"`
}
