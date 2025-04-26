package main

import (
	"log"

	"github.com/savanyv/account-service-api/internal/app"
	"github.com/savanyv/account-service-api/internal/config"
)

func main() {
	// Load Config
	config := config.LoadConfig()

	// Start Server
	server := app.NewServer(config)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
