package main

import (
	"radar/api"
)

func main() {
	server := api.NewRadarServer()
	defer server.Close()

	server.Setup()
	server.Run()
}
