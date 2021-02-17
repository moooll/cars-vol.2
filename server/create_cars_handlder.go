package server

import (
	"log"
	"os"
	"time"

	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"

	"cars/dto"
)

func createCarsHandler(ctx *fasthttp.RequestCtx) {
	reqBody := ctx.PostBody()
	var carReq dto.CreateCars
	var carData dto.CarInformation
	err := ffjson.Unmarshal(reqBody, &carReq)
	if err != nil {
		log.Fatal("could not unmarshal response: ", err)
	}
	const layout = "2006-01-02T15:04:00"
	carData.DateTime, err = time.Parse(layout, carReq.DateTime)
	if err != nil {
		log.Fatal("could not parse time", err)
	}
	carData.CarNumber = carReq.CarNumber
	carData.Speed = carReq.Speed
	err = writeToFile(carData)
	if err != nil {
		log.Fatal("could not write to file response: ", err)
	}
}

func writeToFile(carData dto.CarInformation) (err error) {
	var filename = carData.DateTime.Format("2006-01-02") + ".json"
	file, err := os.OpenFile("storage/" + filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	dataMrshld, err := ffjson.Marshal(carData)
	if err != nil {
		return err
	}
	_, err = file.Write(dataMrshld)
	return nil
}
