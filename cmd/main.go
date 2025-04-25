package main

import (
	"log"

	"github.com/savanyv/account-service-api/internal/app"
	"github.com/savanyv/account-service-api/internal/config"
)

func main() {
	// Load Config
	cfg := config.Load()

	// Start Server
	server := app.NewServer(cfg)
	if err := server.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
