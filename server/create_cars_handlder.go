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
		ctx.WriteString("server error, try later")
	}
	const layout = "2006-01-02T15:04:00Z07:00"
	var carData = make([]dto.CarInformation, len(carReq))
	for i, v := range carReq {
		carData[i].DateTime, err = time.Parse(layout, v.DateTime)
		if err != nil {
			zap.L().Error("could not parse time" + err.Error())
			ctx.WriteString("check input: please input date in format yyyy-mm-ddThh:mm:ssZhh:mm (Zhh:mm for time zone)")
		}
		carData[i].CarNumber = v.CarNumber
		carData[i].Speed = v.Speed
	}

	err = writeToFileVol2(carData)
	if err != nil {
		zap.L().Error("could not write to file: " + err.Error())
		ctx.WriteString("error saving data, try later")
	}
}

//todo: goroutines here??

// func writeToFile(carData []dto.CarInformation) (err error) {
// 	var counter int
// 	for i, v := range carData {
// 		var filename = v.DateTime.Format("2006-01-02") + ".txt"
// 		var fileNotExist bool
// 		if _, er := os.Stat(filename); os.IsNotExist(er) {
// 			fileNotExist = true
// 		}
// 		file, err := os.OpenFile("storage/"+filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
// 		if err != nil {
// 			return err
// 		}
// 		defer file.Close()
// 		dataMrshld, err := ffjson.Marshal(v)
// 		if err != nil {
// 			return err
// 		}
// 		if fileNotExist && counter == 0 {
// 			_, err = file.WriteString("[")
// 			counter++
// 		}
// 		_, err = file.Write(dataMrshld)
// 		_, err = file.WriteString(",")
// 		if i == len(carData)-1 {
// 			_, err = file.WriteString("]")
// 		}
// 	}
// 	return nil
// }

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
