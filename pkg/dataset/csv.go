package dataset

import (
	"bufio"
	"errors"
	"os"
	"sort"
	"strconv"
	"strings"
)

type GetClosestStationRequest struct {
	Lat float64
	Lon float64
	Qty int
}

type Station struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

type Price struct {
	StationId string  `json:"stationId"`
	FuelType  string  `json:"fuelType"`
	Price     float64 `json:"price"`
}

type StationWithPrices struct {
	Station Station `json:"station"`
	Prices  []Price `json:"prices"`
}

func GetClosestStationsWithPrices(req *GetClosestStationRequest) ([]StationWithPrices, error) {
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
	var results []StationWithPrices

	for _, s := range stations {
		var stationPrices []Price

		for _, p := range prices {
			if s.Id == p.StationId {
				stationPrices = append(stationPrices, p)
			}
		}

		results = append(results, StationWithPrices{
			Station: s,
			Prices:  stationPrices,
		})
	}

	return results, nil
}

func parseCsvPricesForStations(stations []Station) ([]Price, error) {
	var ds dataset = datasetPrices
	var prices []Price

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

		// get the price
		p, err := strconv.ParseFloat(cells[2], 64)
		if err != nil {
			p = 0
		}

		if p > 0 {
			prices = append(prices, Price{
				StationId: stationId,
				FuelType:  strings.TrimSpace(cells[1]),
				Price:     p,
			})
		}
	}

	return prices, nil
}

func parseCsvStations() ([]Station, error) {
	var ds dataset = datasetStations
	var stations []Station

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
		name := strings.TrimSpace(cells[2]) + " " + strings.TrimSpace(cells[1]) + strings.TrimSpace(cells[4])
		addr := strings.TrimSpace(cells[5]) + " " + strings.TrimSpace(cells[6]) + strings.TrimSpace(cells[7])

		// push to the stations slice whether lat and lon are valid
		if lat > 0 && lon > 0 {
			stations = append(stations, Station{
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
