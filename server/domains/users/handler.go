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
	FirstName            string `json:"firstName" validate:"required"`
	LastName             string `json:"lastName" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required,eqfield=Password"`
}

func (h *UserHandler) createUserHandler(c *fiber.Ctx) error {
	body, err := utils.BindAndValidate[UserBody](c, h.validator)
	if err != nil {
		return err
	}

	user, err := h.service.CreateUser(*body)
	if err != nil {
		log.Println("Error creating user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
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

	// Set auth cookie with configuration-based settings
	c.Cookie(&fiber.Cookie{
		Name:     h.config.Auth.CookieName,
		Value:    tokenString,
		HTTPOnly: true,
		Secure:   !h.config.IsDevelopment(),
		SameSite: fiber.CookieSameSiteStrictMode,
		Path:     "/",
		MaxAge:   h.config.Auth.CookieMaxAge,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": tokenString,
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
	app.Post("/signup", h.createUserHandler)
	app.Get("/check", h.checkAuth)
}
