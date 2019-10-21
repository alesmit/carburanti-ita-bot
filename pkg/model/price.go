package model

type Price struct {
	StationId string  `json:"stationId"`
	FuelType  string  `json:"fuelType"`
	Price     float64 `json:"price"`
}
