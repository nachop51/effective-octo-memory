package users

import (
	"server/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) getUsers(c *fiber.Ctx) error {
	users, err := h.service.GetUsers()

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

type UserBody struct {
	FirstName            string `json:"firstName" validate:"required"`
	LastName             string `json:"lastName" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required,eqfield=Password"`
}

func (h *UserHandler) createUserHandler(c *fiber.Ctx) error {
	body, ok := utils.BindAndValidate[UserBody](c)
	if !ok {
		return nil
	}

	user, err := h.service.CreateUser(*body)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

type LoginBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (h *UserHandler) loginHandler(c *fiber.Ctx) error {
	creds, ok := utils.BindAndValidate[LoginBody](c)
	if !ok {
		return nil
	}

	user, err := h.service.GetUser(creds.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	tokenString, err := h.service.GenerateJWT(user)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": tokenString,
	})
}

func (h *UserHandler) RegisterRoutes(app *fiber.App) {
	app.Get("/users", h.getUsers)
	app.Post("/users", h.createUserHandler)
	app.Post("/login", h.loginHandler)
}
