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
	var carReq []dto.CreateCars
	err := ffjson.Unmarshal(reqBody, &carReq)
	log.Println(carReq)
	if err != nil {
		log.Println("could not unmarshal response: ", err)
		ctx.WriteString("server error, try later")
	}
	const layout = "2006-01-02T15:04:00Z03:00"
	var carData = make([]dto.CarInformation, len(carReq))
	for i, v := range carReq {
		carData[i].DateTime, err = time.Parse(layout, v.DateTime)
		if err != nil {
			log.Println("could not parse time", err)
			ctx.WriteString("check input and try later")
		}
		carData[i].CarNumber = v.CarNumber
		carData[i].Speed = v.Speed
	}
	
	err = writeToFile(carData)
		if err != nil {
			log.Println("could not write to file: ", err)
			ctx.WriteString("error saving data, try later")
		}
}
//todo: goroutines here??

func writeToFile(carData []dto.CarInformation) (err error) {
	var counter int
	for i, v := range carData {
		var filename = v.DateTime.Format("2006-01-02") + ".json"
		var fileNotExist bool
		if _, er := os.Stat(filename); os.IsNotExist(er) {
			fileNotExist = true
		}
		file, err := os.OpenFile("storage/" + filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
		if err != nil {
			return err
		}
		defer file.Close()
		dataMrshld, err := ffjson.Marshal(v)
		if err != nil {
			return err
		}
		if fileNotExist && counter == 0 {
			_, err = file.WriteString("[")
			counter++
		}
		_, err = file.Write(dataMrshld)
		_, err = file.WriteString(",")
		if i ==  len(carData) - 1 {
			_, err = file.WriteString("]")
		}
	}
	return nil
}