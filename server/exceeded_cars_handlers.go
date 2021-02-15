package server

import (
	"cars/dto"
	"log"

	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"
)


func exceededCarsHandler(ctx *fasthttp.RequestCtx) {
	req := ctx.PostBody()
	var speedDate dto.SpeedDateReq
	err := ffjson.Unmarshal(req, &speedDate)
	if err != nil {
		log.Fatal("could not unmarshal request\n")
	}
	exceededCars := findCarsWhichExceeded(speedDate)
	resp, err := ffjson.Marshal(exceededCars)
	if err != nil {
		log.Fatal("could not marshal response\n")
	}
	ctx.SetBody(resp)
}

func findCarsWhichExceeded(speedDate dto.SpeedDateReq) (exceededCars dto.SpeedDateResp) {
	//todo: find Cars in storage
	return exceededCars
}