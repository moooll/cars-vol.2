package server

import (
	"log"
	"sort"

	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"

	"cars/dto"
)

func exceededCarsHandler(ctx *fasthttp.RequestCtx) {
	req := ctx.PostBody()
	var speedDate dto.SpeedDateReq
	err := ffjson.Unmarshal(req, &speedDate)
	if err != nil {
		log.Fatal("could not unmarshal request: ", err)
	}
	exceededCars, err := findCarsWhichExceeded(speedDate)
	if err != nil {
		log.Fatal("could not find cars: ", err)
	}
	resp, err := ffjson.Marshal(exceededCars)
	if err != nil {
		log.Fatal("could not marshal response: ", err)
	}
	ctx.SetBody(resp)
}

func findCarsWhichExceeded(speedDate dto.SpeedDateReq) (exceededCars []dto.CarInformation, err error) {

	todaysCars, err := sortSlice(speedDate)
	if err != nil {
		return []dto.CarInformation{}, err
	}

	index := sort.Search(len(todaysCars), func(i int) bool {
		return speedDate.Speed >= todaysCars[i].Speed
	})
	exceededCars = todaysCars[index:]
	return exceededCars, err
}
