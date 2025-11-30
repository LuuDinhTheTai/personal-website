package main

import (
	"log"
	"personal-website/internal/config"
	server2 "personal-website/internal/server"
)

func main() {
	server := server2.NewServer(config.SERVER_HOST, config.SERVER_PORT)
	err := server.InitializeRoutes()
	if err != nil {
		log.Fatal(err)
	}

	server.Run()
}
