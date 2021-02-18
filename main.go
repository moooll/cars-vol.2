package main

import (
	"log"

	"go.uber.org/zap"

	"cars/server"
)

func main() {
	err := server.Server()
	if err != nil {
		log.Println("error launching server: ", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

}
