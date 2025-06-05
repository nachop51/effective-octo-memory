package users

import (
	"server/utils"

	"github.com/gofiber/fiber/v2"
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

	if err := CreateUser(*body); err != nil {
		return err
	}

	return c.JSON(User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
	})
}

func RegisterRoutes(app *fiber.App) {
	app.Post("/users", createUserHandler)
}
