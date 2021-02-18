package server

import (
	"errors"
	"fmt"
	"time"

	"github.com/pquerna/ffjson/ffjson"
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

	config, err := os.ReadFile("config.json")
	if err != nil {
		return err
	}

	type config struct {
		Start int
		End int
	}
	var conf config
	err = ffjson.Unmarshal(config, conf) 
	if err != nil {
		return err
	}

	if time.Now().Hour() < conf.Start || time.Now().Hour() > conf.End {
		return errors.New(fmt.Sprintf("requests are processed only from %v to %v, pls try later", conf.Start, conf.End))
	}
	return nil
}
