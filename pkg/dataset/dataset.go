package dataset

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	datasetFolder   = "./tmp"
	datasetPrices   = "prices"
	datasetStations = "stations"
)

type dataset string

func init() {
	if err := createFolderIfNotExist(); err != nil {
		fmt.Println("unable to create datasets folder")
	}
}

// download the latest datasets of prices and stations if they are missing or out-to-date
func SyncDatasets() error {
	datasets := []dataset{datasetPrices, datasetStations}

	for _, ds := range datasets {
		filename, err := ds.mostRecentFilename()
		if err != nil {
			return err
		}

		// if no dataset is present: download it
		if filename == "" {
			if err := ds.download(); err != nil {
				return err
			}
			continue
		}

		// if dataset is expired: remove it and download a new one
		if isExpired(filename) {
			if err := os.Remove(datasetFolder + "/" + filename); err != nil {
				return err
			}

			if err = ds.download(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ds *dataset) download() error {
	var url string

	switch *ds {
	case datasetPrices:
		url = "https://www.mise.gov.it/images/exportCSV/prezzo_alle_8.csv"
		break

	case datasetStations:
		url = "https://www.mise.gov.it/images/exportCSV/anagrafica_impianti_attivi.csv"
		break

	default:
		return errors.New("invalid dataset type")
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create(ds.newFilename())
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err

}
