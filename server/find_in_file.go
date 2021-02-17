package server

import (
	"cars/dto"
	"io/ioutil"
	"sort"

	"github.com/pquerna/ffjson/ffjson"
)

func sortSlice(speedDate dto.SpeedDateReq) (todaysCars []dto.CarInformation, err error) {
	var filename = speedDate.Date.Format("2006-01-02")
	dataFromFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return []dto.CarInformation{}, err
	}

	err = ffjson.Unmarshal(dataFromFile, &todaysCars)
	if err != nil {
		return []dto.CarInformation{}, err
	}

	sort.Slice(todaysCars, func(i, j int) bool {
		return todaysCars[i].Speed < todaysCars[j].Speed
	})
	return todaysCars, nil
}
