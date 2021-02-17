package main

import (
	"log"

	"cars/server"
)

func main() {
	err := server.Server()
	if err != nil {
		log.Fatal("error launching server: ", err)
	}
}
