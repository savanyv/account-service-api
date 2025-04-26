package handlers

import (
	"github.com/gofiber/fiber/v2"
	dtos "github.com/savanyv/account-service-api/internal/dto"
	"github.com/savanyv/account-service-api/internal/usecase"
	"github.com/savanyv/account-service-api/internal/utils"
)

type AccountHandler struct {
	usecase usecase.AccountUsecase
	validator *utils.CustomerValidator
}

func NewAccountHandler(usecase usecase.AccountUsecase) *AccountHandler {
	return &AccountHandler{
		usecase: usecase,
		validator: utils.NewValidator(),
	}
}

func (h *AccountHandler) Register(c *fiber.Ctx) error {
	var req dtos.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		utils.LogError("HANDLER", "Failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"remark": "Failed to parse request body",
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		utils.LogError("HANDLER", "Failed to validate request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"remark": "Failed to validate request",
		})
	}

	resp, err := h.usecase.Register(&req)
	if err != nil {
		utils.LogError("HANDLER", "Failed to register customer: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"remark": "Failed to register customer",
		})
	}

	utils.LogInfo("HANDLER", "Customer registered successfully: %s", resp.AccountNo)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": resp,
		"remark": "Customer registered successfully",
	})
}

func (h *AccountHandler) Deposit(c *fiber.Ctx) error {
	var req dtos.DepositRequest
	if err := c.BodyParser(&req); err != nil {
		utils.LogError("HANDLER", "Failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"remark": "Failed to parse request body",
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		utils.LogError("HANDLER", "Failed to validate request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"remark": "Failed to validate request",
		})
	}

	resp, err := h.usecase.Deposit(&req)
	if err != nil {
		utils.LogError("HANDLER", "Failed to deposit: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"remark": "Failed to deposit",
		})
	}

	utils.LogInfo("HANDLER", "Deposit successful: %s", resp.AccountNo)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": resp,
		"remark": "Deposit successful",
	})
}
