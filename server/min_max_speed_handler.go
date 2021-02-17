package server

import (
	"log"
	"time"

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
	// todo: actually find this in storage
	datetime, err := time.Parse("2006-01-02T15:04:00", "2021-02-14T15:40:00")
	if err != nil {
		return dto.MinMaxSpeedResp{}, err
	}

	minMaxSpeedResult = dto.MinMaxSpeedResp{
		MinSpeed: dto.CarInformation{
			DateTime:  datetime,
			CarNumber: "1234 PM-1",
			Speed:     40.5,
		},
		MaxSpeed: dto.CarInformation{
			DateTime:  datetime,
			CarNumber: "1234 PM-1",
			Speed:     60.5,
		},
	}
	return minMaxSpeedResult, nil
}
