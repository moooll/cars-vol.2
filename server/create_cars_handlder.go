package server

import (
	"os"
	"time"

	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"cars/dto"
)

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
	err = writeToFileVol2(carData)
	if err != nil {
		zap.L().Error("could not write to file: " + err.Error())
		ctx.WriteString("error saving data, try later\n")
	}
}

//todo: goroutines here??

func writeToFileVol2(carData []dto.CarInformation) (err error) {
	for _, v := range carData {
		var filename = v.DateTime.Format("2006-01-02") + ".txt"
		file, err := os.OpenFile("storage/"+filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
		if err != nil {
			return err
		}
		carDataByte, err := ffjson.Marshal(v)
		file.Write(carDataByte)
		file.WriteString("\n")
	}
	return nil
}
