package accounts

import (
	"server/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	service   *AccountService
	validator *validator.Validate
}

func NewAccountHandler(service *AccountService, validator *validator.Validate) *AccountHandler {
	return &AccountHandler{
		service:   service,
		validator: validator,
	}
}

func (h *AccountHandler) getAccountsHandler(c *fiber.Ctx) error {
	accounts, err := h.service.GetAccounts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve accounts",
		})
	}

	return c.Status(fiber.StatusOK).JSON(accounts)
}

type CreateAccountBody struct {
	Name   string `json:"name" validate:"required"`
	UserID uint   `json:"userId" validate:"required"`
}

func (h *AccountHandler) createAccountHandler(c *fiber.Ctx) error {
	body, ok := utils.BindAndValidate[CreateAccountBody](c, h.validator)
	if !ok {
		return nil
	}

	account, err := h.service.CreateAccount(body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create account",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(account)
}

func (h *AccountHandler) RegisterRoutes(app *fiber.App) {
	app.Get("/accounts", h.getAccountsHandler)
	app.Post("/accounts", h.createAccountHandler)
}
