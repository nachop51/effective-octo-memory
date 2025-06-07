package users

import (
	"server/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserBody struct {
	FirstName            string `json:"firstName" validate:"required"`
	LastName             string `json:"lastName" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required,eqfield=Password"`
}

func createUserHandler(c *fiber.Ctx) error {
	body, ok := utils.BindAndValidate[UserBody](c)
	if !ok {
		return nil
	}

	user, err := CreateUser(*body)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

type LoginBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func loginHandler(c *fiber.Ctx) error {
	creds, ok := utils.BindAndValidate[LoginBody](c)
	if !ok {
		return nil
	}

	user, err := GetUser(creds.Email)
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

	tokenString, err := GenerateJWT(user)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": tokenString,
	})
}

func RegisterRoutes(app *fiber.App) {
	app.Post("/users", createUserHandler)
	app.Post("/login", loginHandler)
}
