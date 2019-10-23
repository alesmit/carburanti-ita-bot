package model

import "time"

type Price struct {
	StationId string    `json:"stationId"`
	FuelType  string    `json:"fuelType"`
	Price     float64   `json:"price"`
	UpdatedAt time.Time `json:"updatedAt"`
}
