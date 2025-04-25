package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/account-service-api/internal/config"
	"github.com/savanyv/account-service-api/internal/config/database"
	"github.com/savanyv/account-service-api/internal/utils"
)

type Server struct {
	App *fiber.App
	Config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		App: fiber.New(),
		Config: config,
	}
}

func (s *Server) Run(cfg *config.Config) error {
	// initialize logger
	utils.InitLogger()

	// initialize database
	if err := database.ConnectDB(s.Config); err != nil {
		log.Println(err)
	}

	// Start Server
	if err := s.App.Listen(":8080"); err != nil {
		utils.LogCritical("SERVER", "Failed to start server: %v", err)
		log.Fatal(err)
		return err
	}

	utils.LogInfo("SERVER", "Server started successfully")
	return nil
}
