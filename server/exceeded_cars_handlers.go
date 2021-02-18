package server

import (
	"sort"
	"time"

	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"cars/dto"
)

func exceededCarsHandler(ctx *fasthttp.RequestCtx) {
	req := ctx.PostBody()
	var speedDateReq dto.SpeedDateReq
	err := ffjson.Unmarshal(req, &speedDateReq)
	if err != nil {
		zap.L().Error("could not unmarshal request: " +  err.Error())
		ctx.WriteString("server error, try later")
	}
	
	for i, v := range speedDateReq.Date {
		if v == 'T' {
			speedDateReq.Date = speedDateReq.Date[:i]
		}
	}
	date, err := time.Parse("2006-01-02", speedDateReq.Date)
	if err != nil {
		zap.L().Error("could not parse time: " + err.Error())
		ctx.WriteString("check input and try later")
	}

	var speedDate = dto.SpeedDate{
		Speed: speedDateReq.Speed,
		Date:  date,
	}
	exceededCars, err := findCarsWhichExceeded(speedDate)
	if err != nil {
		zap.L().Error("could not find cars: " + err.Error())
		ctx.WriteString("server error, try later")
	}
	resp, err := ffjson.Marshal(exceededCars)
	if err != nil {
		zap.L().Error("could not marshal response: " + err.Error())
		ctx.WriteString("server error, try later")
	}
	ctx.SetBody(resp)
}

func findCarsWhichExceeded(speedDate dto.SpeedDate) (exceededCars []dto.CarInformation, err error) {
	todaysCars, err := sortSlice(speedDate.Date)
	if err != nil {
		return []dto.CarInformation{}, err
	}

	index := sort.Search(len(todaysCars), func(i int) bool {
		return speedDate.Speed >= todaysCars[i].Speed
	})
	exceededCars = todaysCars[index:]
	return exceededCars, err
}
