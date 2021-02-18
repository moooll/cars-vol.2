package server

import (
	"bytes"
	"cars/dto"
	"io/ioutil"
	"log"
	"sort"
	"strings"
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

func readFromFile(date time.Time) (todaysCars []dto.CarInformation, err error) {
	var filename = "storage/" + date.Format("2006-01-02") + ".json"
	dataFromFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return []dto.CarInformation{}, err
	}
	dataFromFileString := string(dataFromFile)
	var newFileString string
	var counter = 0
	for i := range dataFromFileString {
		if counter != 0 && i != len(dataFromFileString)-1 {
			newFileString = strings.ReplaceAll(dataFromFileString, "]", "")
			newFileString = strings.ReplaceAll(dataFromFileString, "[", ",")
			counter++
		}
	}
	log.Println(newFileString)
	newFileBytes := []byte(newFileString)
	err = ffjson.Unmarshal(newFileBytes, &todaysCars)
	if err != nil {
		return []dto.CarInformation{}, err
	}
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
