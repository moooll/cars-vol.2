package server

import (
	"github.com/valyala/fasthttp"
)

func Server() error {
	err := fasthttp.ListenAndServe(":8080", func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/exceeded-cars":
			exceededCarsHandler(ctx)
		case "/create-cars":
			createCarsHandler(ctx)
		case "/min-and-max-speed":
			minAndMaxSpeedHandler(ctx)
		}
	})
	if err != nil {
		return err
	}
	return nil
}
