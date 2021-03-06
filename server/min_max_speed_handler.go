package server

import (
	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"cars/dto"
)

func minAndMaxSpeedHandler(ctx *fasthttp.RequestCtx) {
	reqBody := ctx.PostBody()
	var minMax dto.MinMaxSpeedReq
	err := ffjson.Unmarshal(reqBody, &minMax)
	if err != nil {
		zap.L().Error("could not unmarshal body")
		ctx.WriteString("check the input and try later")
	}
	minMaxSpeed, err := findMinMaxSpeed(minMax)
	if err != nil {
		zap.L().Error("could not find min-max speed")
		ctx.WriteString("could not find relevant results")
	}
	minMaxBytes, err := ffjson.Marshal(minMaxSpeed)
	if err != nil {
		zap.L().Error("could not marshal response: " + err.Error())
		ctx.WriteString("check the input and try later")
	}
	if string(minMaxBytes) == `{"min_speed":{"date_time":"0001-01-01T00:00:00Z","car_name":"","speed":0},"max_speed":{"date_time":"0001-01-01T00:00:00Z","car_name":"","speed":0}}` {
		ctx.SetBody([]byte("check the input and try later"))
	} else {
		ctx.SetBody(minMaxBytes)
	}
}

func findMinMaxSpeed(minMax dto.MinMaxSpeedReq) (minMaxSpeedResult dto.MinMaxSpeedResp, err error) {
	sortedCars, err := sortSlice(minMax.Date)
	if err != nil {
		return dto.MinMaxSpeedResp{}, err
	}
	minMaxSpeedResult = dto.MinMaxSpeedResp{
		MinSpeed: sortedCars[0],
		MaxSpeed: sortedCars[len(sortedCars)-1],
	}
	return minMaxSpeedResult, nil
}
