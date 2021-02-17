package main

import (
	"log"

	"cars/server"
)

func main() {
	err := server.Server()
	if err != nil {
		log.Println("error launching server: ", err)
	}
}
