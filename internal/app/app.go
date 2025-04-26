package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/account-service-api/internal/config"
	"github.com/savanyv/account-service-api/internal/config/database"
	"github.com/savanyv/account-service-api/internal/delivery/routes"
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

func (s *Server) Run() error {
	// Setup Logger
	utils.InitLogger()

	// Setup Database Connection
	_, err := database.ConnectDB(s.Config)
	if err != nil {
		utils.LogCritical("SERVER", "Failed to connect to database: %v", err)
		return err
	}

	// Setup Routes
	routes.RegisterRoutes(s.App)

	// Start Server
	if err := s.App.Listen(":8080"); err != nil {
		utils.LogError("SERVER", "Failed to start server: %v", err)
		return err
	}
	utils.LogInfo("SERVER", "Server started on port 8080")
	return nil
}
