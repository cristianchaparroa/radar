package main

import (
	"github.com/joho/godotenv"

	"log"
	"radar/api"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := api.NewRadarServer()
	defer server.Close()

	server.Setup()
	server.Run()
}
