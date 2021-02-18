package server

import (
	"os"
	"sync"
	"time"

	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"cars/dto"
)

type carInfoSync struct {
	carInfo []dto.CarInformation
	mu      sync.Mutex
}

func createCarsHandler(ctx *fasthttp.RequestCtx) {
	reqBody := ctx.PostBody()
	var carReq []dto.CreateCars
	err := ffjson.Unmarshal(reqBody, &carReq)
	if err != nil {
		zap.L().Error("could not unmarshal response: " + err.Error())
		ctx.WriteString("server error, try later\n")
	}
	const layout = "2006-01-02T15:04:00Z07:00"
	var carData = make([]dto.CarInformation, len(carReq))
	var timeErr error
	for i, v := range carReq {
		carData[i].DateTime, timeErr = time.Parse(layout, v.DateTime)
		if timeErr != nil {
			zap.L().Error("could not parse time" + timeErr.Error())
		}
		carData[i].CarNumber = v.CarNumber
		carData[i].Speed = v.Speed
	}
	if timeErr != nil {
		ctx.WriteString("check input: please input date in format yyyy-mm-ddThh:mm:ssZhh:mm (Zhh:mm for time zone)\n")
	}
	var carInfo carInfoSync
	carInfo.carInfo = carData
	err = carInfo.writeToFileVol2()
	if err != nil {
		zap.L().Error("could not write to file: " + err.Error())
		ctx.WriteString("error saving data, try later\n")
	}
}

func (carInfo *carInfoSync) writeToFileVol2() (err error) {
	for _, v := range carInfo.carInfo {
		var filename = v.DateTime.Format("2006-01-02") + ".txt"
		file, err := os.OpenFile("storage/"+filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
		if err != nil {
			return err
		}

		defer file.Close()
		carDataByte, err := ffjson.Marshal(v)
		if err != nil {
			return err
		}
		carInfo.mu.Lock()
		file.Write(carDataByte)
		file.WriteString("\n")
		carInfo.mu.Unlock()
	}
	return nil
}
