package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BindAndValidate[T any](c *fiber.Ctx, validate *validator.Validate) (*T, bool) {
	var payload T

	if err := c.BodyParser(&payload); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err.Error())
		return nil, false
	}

	if err := validate.Struct(payload); err != nil {
		errors := make(fiber.Map)
		for _, e := range err.(validator.ValidationErrors) {
			errors[e.Field()] = e.Tag()
		}
		c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
		return nil, false
	}
	return &payload, true
}
