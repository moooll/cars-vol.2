package server

import (
	"bytes"
	"cars/dto"
	"log"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/pquerna/ffjson/ffjson"
	"go.uber.org/zap"
)

type carDateSync struct {
	date time.Time
	mu sync.Mutex
}

func sortSlice(date time.Time) (todaysCars []dto.CarInformation, err error) {
	var carDate carDateSync
	carDate.date = date
	todaysCars, err = carDate.readFromFileVol2()
	if err != nil {
		return []dto.CarInformation{}, err
	}

	sort.Slice(todaysCars, func(i, j int) bool {
		return todaysCars[i].Speed < todaysCars[j].Speed
	})
	return todaysCars, nil
}

func (carDate *carDateSync) readFromFileVol2() (todaysCars []dto.CarInformation, err error) {
	var filename = "storage/" + carDate.date.Format("2006-01-02") + ".txt"
//	var fileMutex *sync.Mutex

	carDate.mu.Lock()
	dataFromFile, err := os.ReadFile(filename)
	if err != nil {
		return []dto.CarInformation{}, err
	}
	carDate.mu.Unlock()
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
