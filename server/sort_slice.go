package server

import (
	"cars/dto"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/pquerna/ffjson/ffjson"
)

func sortSlice(date time.Time) (todaysCars []dto.CarInformation, err error) {
	var filename = "storage/" + date.Format("2006-01-02") + ".json"
	dataFromFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return []dto.CarInformation{}, err
	}
	dataFromFileString := string(dataFromFile)
	var newFileString string
	for i := range dataFromFileString {
		if i != 0 && i != len(dataFromFileString)-1 {
			newFileString = strings.ReplaceAll(dataFromFileString, "]", "")
			newFileString = strings.ReplaceAll(dataFromFileString, "[", ",")
		}
	}
	newFileBytes := []byte(newFileString)
	err = ffjson.Unmarshal(newFileBytes, &todaysCars)
	if err != nil {
		return []dto.CarInformation{}, err
	}
	log.Println(todaysCars)
	sort.Slice(todaysCars, func(i, j int) bool {
		return todaysCars[i].Speed < todaysCars[j].Speed
	})
	return todaysCars, nil
}
