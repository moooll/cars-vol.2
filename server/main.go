package server

import (
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	err := fasthttp.ListenAndServe(":8080", func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/exceeded-cars":
			exceededCarsHandler(ctx)
		case "/createCars":
			createCarsHandler(ctx)
		case "min-and-max-speed":
			minAndMaxSpeedHandler(ctx)
		}
	})
	if err != nil {
		log.Fatal("server is not up")
	}
}
