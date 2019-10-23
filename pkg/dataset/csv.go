package dataset

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/alesmit/fuel-master/pkg/model"
)

type GetClosestStationRequest struct {
	Lat float64
	Lon float64
	Qty int
}

func GetClosestStationsWithPrices(req *GetClosestStationRequest) ([]model.StationWithPrices, error) {
	stations, err := parseCsvStations()
	if err != nil {
		return nil, errors.New("unable to parse stations csv")
	}

	sort.Slice(stations, func(i, j int) bool {
		d1 := getDistance(stations[i].Lat, stations[i].Lon, req.Lat, req.Lon)
		d2 := getDistance(stations[j].Lat, stations[j].Lon, req.Lat, req.Lon)
		return d1 < d2
	})

	// closest stations
	if len(stations) >= req.Qty {
		stations = stations[:req.Qty]
	}

	// get prices for those stations
	prices, err := parseCsvPricesForStations(stations)
	if err != nil {
		return nil, errors.New("unable to parse prices csv")
	}

	// build the resulting slice
	var results []model.StationWithPrices

	for _, s := range stations {
		var stationPrices []model.Price

		for _, p := range prices {
			if s.Id == p.StationId {
				stationPrices = append(stationPrices, p)
			}
		}

		results = append(results, model.StationWithPrices{
			Station: s,
			Prices:  stationPrices,
		})
	}

	return results, nil
}

func GetStationById(stationId string) (*model.Station, error) {
	stations, err := parseCsvStations()
	if err != nil {
		return nil, errors.New("unable to parse stations csv")
	}

	for _, station := range stations {
		if station.Id == stationId {
			return &station, nil
		}
	}

	return nil, errors.New("unable to find station")
}

func parseCsvPricesForStations(stations []model.Station) ([]model.Price, error) {
	var ds dataset = datasetPrices
	var prices []model.Price

	filename, err := ds.mostRecentFilename()
	if err != nil || filename == "" {
		return nil, errors.New("unable to find prices dataset")
	}

	csvfile, err := os.Open(datasetFolder + "/" + filename)
	if err != nil {
		return nil, err
	}

	// build a dot-separated string of ids
	stationIds := ""
	for _, s := range stations {
		stationIds += "." + s.Id
	}

	scanner := bufio.NewScanner(csvfile)
	i := 1

	for scanner.Scan() {
		i++

		// skip the first 2 lines
		if i < 2 {
			continue
		}

		// get text of a line and split by ';'
		row := scanner.Text()
		cells := strings.Split(row, ";")

		// only process prices of the interested stations
		stationId := strings.TrimSpace(cells[0])
		if !strings.Contains(stationIds, stationId) {
			continue
		}

		// get the price, skip invalid prices
		p, err := strconv.ParseFloat(cells[2], 64)
		if err != nil {
			continue
		}

		// parse the last updated date, skip invalid dates
		t, err := time.Parse("02/01/2006 15:04:05", cells[4])
		if err != nil {
			continue
		}

		prices = append(prices, model.Price{
			StationId: stationId,
			FuelType:  strings.TrimSpace(cells[1]),
			Price:     p,
			UpdatedAt: t,
		})
	}

	return prices, nil
}

func parseCsvStations() ([]model.Station, error) {
	var ds dataset = datasetStations
	var stations []model.Station

	filename, err := ds.mostRecentFilename()
	if err != nil || filename == "" {
		return nil, errors.New("unable to find stations dataset")
	}

	csvfile, err := os.Open(datasetFolder + "/" + filename)
	if err != nil {
		return nil, err
	}

	defer csvfile.Close()

	scanner := bufio.NewScanner(csvfile)
	i := -1

	for scanner.Scan() {
		i++

		// skip the first 2 lines
		if i < 2 {
			continue
		}

		// get text of a line and split by ';'
		row := scanner.Text()
		cells := strings.Split(row, ";")

		// parse fields holding latitude and longitude as float64
		lat, err := strconv.ParseFloat(cells[8], 64)
		if err != nil {
			lat = 0
		}

		lon, err := strconv.ParseFloat(cells[9], 64)
		if err != nil {
			lon = 0
		}

		// build the station's name and address
		name := fmt.Sprint(strings.TrimSpace(cells[2]), " ", strings.TrimSpace(cells[4]))
		addr := fmt.Sprint(strings.TrimSpace(cells[5]), " ", strings.TrimSpace(cells[6]))

		// push to the stations slice whether lat and lon are valid
		if lat > 0 && lon > 0 {
			stations = append(stations, model.Station{
				Id:      strings.TrimSpace(cells[0]),
				Name:    name,
				Address: addr,
				Lat:     lat,
				Lon:     lon,
			})
		}

	}

	return stations, nil
}
