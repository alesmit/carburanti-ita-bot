package utils

import (
	"github.com/alesmit/fuel-master/pkg/model"

	"fmt"
)

func Format(s *model.StationWithPrices) string {
	out := fmt.Sprintln("*", s.Station.Name, "*")
	out += fmt.Sprintln("_", s.Station.Address, "_")

	for _, p := range s.Prices {
		out += "\n" + p.FuelType + ": € " + fmt.Sprintf("%.3f", p.Price) + "/lt."
	}

	return out
}
