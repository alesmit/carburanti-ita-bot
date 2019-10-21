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
	Id  string  `json:"id"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func GetClosestStations(req *GetClosestStationRequest) ([]Station, error) {
	stations, err := parseCsvStations()
	if err != nil || stations == nil {
		return nil, errors.New("unable to parse stations csv")
	}

	sort.Slice(stations, func(i, j int) bool {
		d1 := getDistance(stations[i].Lat, stations[i].Lon, req.Lat, req.Lon)
		d2 := getDistance(stations[j].Lat, stations[j].Lon, req.Lat, req.Lon)
		return d1 < d2
	})

	if len(stations) >= req.Qty {
		return stations[:req.Qty], nil
	}

	return stations, nil
}

func parseCsvStations() ([]Station, error) {
	var ds dataset = datasetStations
	var stations []Station

	filename, err := ds.mostRecentFilename()
	if err != nil || filename == "" {
		return nil, errors.New("unable to find dataset")
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

		// push to the stations slice whether lat and lon are valid
		if lat != 0 && lon != 0 {
			stations = append(stations, Station{
				Id:  strings.TrimSpace(cells[0]),
				Lat: lat,
				Lon: lon,
			})
		}

	}

	return stations, nil
}
