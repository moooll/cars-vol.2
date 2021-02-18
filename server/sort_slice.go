package server

import (
	"bytes"
	"cars/dto"
	"io/ioutil"
	"log"
	"sort"
	"time"

	"github.com/pquerna/ffjson/ffjson"
	"go.uber.org/zap"
)

func sortSlice(date time.Time) (todaysCars []dto.CarInformation, err error) {
	todaysCars, err = readFromFileVol2((date))
	if err != nil {
		return []dto.CarInformation{}, err
	}

	sort.Slice(todaysCars, func(i, j int) bool {
		return todaysCars[i].Speed < todaysCars[j].Speed
	})
	return todaysCars, nil
}

func readFromFileVol2(date time.Time) (todaysCars []dto.CarInformation, err error) {
	var filename = "storage/" + date.Format("2006-01-02") + ".txt"
	dataFromFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return []dto.CarInformation{}, err
	}

	zap.L().Info(string(string(dataFromFile)))
	fileDataWithoutSuffix := bytes.TrimSuffix(dataFromFile, []byte("\n"))
	sliceFromFile := bytes.Split(fileDataWithoutSuffix, []byte("\n"))
	var todaysCar dto.CarInformation
	for _, v := range sliceFromFile {
		zap.L().Info(string(v))
		err = ffjson.Unmarshal(v, &todaysCar)
		if err != nil {
			return []dto.CarInformation{}, err
		}
		todaysCars = append(todaysCars, todaysCar)

	}
	log.Println(todaysCars)
	return todaysCars, nil
}
