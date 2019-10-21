package dataset

import (
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strings"
	"time"
)

// check whether the datasets folder exists or not
func folderExist() (bool, error) {
	_, err := os.Stat(datasetFolder)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err
}

// creates the datasets folder if it doesn't exist yet
func createFolderIfNotExist() error {
	exist, err := folderExist()
	if err != nil {
		return err
	}

	if !exist {
		err := os.MkdirAll(datasetFolder, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

// returns a filename that includes the current date. example: prices_2019-10-15.csv
func (ds *dataset) newFilename() string {
	now := time.Now().UTC().String()[:10]
	return datasetFolder + "/" + string(*ds) + "_" + now + ".csv"
}

// extract timestamp from the file name
func timeFromFilename(filename string) time.Time {
	s := filename
	s = strings.Replace(s, datasetPrices+"_", "", -1)
	s = strings.Replace(s, datasetStations+"_", "", -1)
	s = strings.Replace(s, ".csv", "", -1)

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Now()
	}

	t = t.Add(time.Hour * 8)
	return t
}

// get the filename of the most recent dataset
func (ds *dataset) mostRecentFilename() (string, error) {
	files, err := ioutil.ReadDir(datasetFolder)
	if err != nil {
		return "", nil
	}

	var filenames []string

	for _, f := range files {
		if strings.Contains(f.Name(), string(*ds)) {
			filenames = append(filenames, f.Name())
		}
	}

	if len(filenames) == 0 {
		return "", nil
	}

	sort.Slice(filenames, func(i, j int) bool {
		t1 := timeFromFilename(filenames[i])
		t2 := timeFromFilename(filenames[j])
		return t1.Before(t2)
	})

	return filenames[0], nil
}

// check whether the dataset is expired or not
func isExpired(filename string) bool {
	var t time.Time
	last := timeFromFilename(filename)

	// exp time for prices dataset is 1 day
	if strings.Contains(filename, datasetPrices) {
		t = last.AddDate(0, 0, 1)
	}

	// exp time for stations dataset is 15 days
	if strings.Contains(filename, datasetStations) {
		t = last.AddDate(0, 0, 15)
	}

	return t.Before(time.Now())
}

// get distance between two points: (x1, y1) and (x2, y2)
func getDistance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}
