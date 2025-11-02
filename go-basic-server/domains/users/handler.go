package users

import (
	"log"

	"server/config"
	"server/pkg/errors"
	"server/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	service   *UserService
	validator *validator.Validate
	config    *config.Config
}

func NewUserHandler(service *UserService, validator *validator.Validate, config *config.Config) *UserHandler {
	return &UserHandler{
		service:   service,
		validator: validator,
		config:    config,
	}
}

func (h *UserHandler) getUsers(c *fiber.Ctx) error {
	users, err := h.service.GetUsers()
	if err != nil {
		log.Println("Error retrieving users:", err)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve users",
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

type UserBody struct {
	FirstName            string `json:"firstName" validate:"required,min=2,max=30"`
	LastName             string `json:"lastName" validate:"required,min=2,max=30"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required,min=6,max=40"`
	PasswordConfirmation string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

func (h *UserHandler) handleSignUp(c *fiber.Ctx) error {
	body, err := utils.BindAndValidate[UserBody](c, h.validator)
	if err != nil {
		return err
	}

	existingUser, err := h.service.GetUserByEmail(body.Email)
	if err != nil {
		log.Println("Error checking existing user:", err)
		return errors.NewInternalServerError("Failed to check existing user")
	}

	if existingUser != nil {
		return errors.NewConflictError("User already exists")
	}

	user, err := h.service.CreateUser(*body)
	if err != nil {
		log.Println("Error creating user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	token, err := h.service.GenerateJWT(user)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": token,
		"user":         user,
	})
}

type LoginBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (h *UserHandler) loginHandler(c *fiber.Ctx) error {
	creds, err := utils.BindAndValidate[LoginBody](c, h.validator)
	if err != nil {
		return err
	}

	user, err := h.service.GetUserByEmail(creds.Email)
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

	token, err := h.service.GenerateJWT(user)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": token,
		"user":         user,
	})
}

func (h *UserHandler) checkAuth(c *fiber.Ctx) error {
	userId, err := utils.GetUserID(c)
	if err != nil {
		return errors.NewUnauthorizedError("Unauthorized")
	}

	user, err := h.service.store.GetUserByID(userId)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.NewUnauthorizedError("User not found")
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *UserHandler) RegisterRoutes(app *fiber.App) {
	app.Get("/users", h.getUsers)
	app.Post("/login", h.loginHandler)
	app.Post("/signup", h.handleSignUp)
	app.Get("/check", h.checkAuth)
}
