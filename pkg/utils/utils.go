package utils

import (
	"github.com/alesmit/fuel-master/pkg/model"

	"fmt"
)

func Format(s *model.StationWithPrices) string {
	out := s.Station.Name
	out += "\n_" + s.Station.Address + "_\n"

	for _, p := range s.Prices {
		out += "\n" + p.FuelType + ": € " + fmt.Sprintf("%f", p.Price) + "/lt."
	}

	return out
}
