package server

import (
	"log"

	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"

	"cars/dto"
)

func minAndMaxSpeedHandler(ctx *fasthttp.RequestCtx) {
	reqBody := ctx.PostBody()
	var minMax dto.MinMaxSpeedReq
	err := ffjson.Unmarshal(reqBody, &minMax)
	minMaxSpeed, err := findMinMaxSpeed(minMax)
	if err != nil {
		log.Println("could not find min-max speed")
		ctx.WriteString("could not find relevant results")
	}
	minMaxBytes, err := ffjson.Marshal(minMaxSpeed)
	if err != nil {
		log.Println("could not marshal response: ", err)
		ctx.WriteString("server error, try later")
	}
	ctx.SetBody(minMaxBytes)
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
