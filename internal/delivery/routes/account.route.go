package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/account-service-api/internal/config/database"
	"github.com/savanyv/account-service-api/internal/delivery/handlers"
	"github.com/savanyv/account-service-api/internal/repository"
	"github.com/savanyv/account-service-api/internal/usecase"
)

func accountRoutes(app fiber.Router) {
	accountRepo := repository.NewAccountRepository(database.DB)
	transactionRepo := repository.NewTransactionRepository(database.DB)
	usecase := usecase.NewAccountUsecase(accountRepo, transactionRepo)
	handler := handlers.NewAccountHandler(usecase)

	app.Post("/register", handler.Register)
	app.Post("/deposit", handler.Deposit)
	app.Post("/withdraw", handler.Withdraw)
	app.Get("/balance/:account_no", handler.GetBalance)
}
