package server

import (
	"cars/dto"
	"log"

	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"
)

func minAndMaxSpeedHandler(ctx *fasthttp.RequestCtx) {
	//todo: write this handler
	minMaxSpeed := findMinMaxSpeed(minMax)
	minMaxBytes, err := ffjson.Marshal(minMaxSpeed)
	if err != nil {
		log.Fatal("could not marshal response\n")
	}
	ctx.SetBody(minMaxBytes)
}

func findMinMaxSpeed(minMax dto.MinMaxSpeedReq) (minMaxSpeedResult dto.MinMaxSpeedResp) {
	// todo: actually find this in storage
	return minMaxSpeedResult
}
