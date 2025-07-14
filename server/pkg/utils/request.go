package utils

import (
	"strings"

	"server/pkg/errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BindAndValidate[T any](c *fiber.Ctx, validate *validator.Validate) (*T, error) {
	var payload T

	if err := c.BodyParser(&payload); err != nil {
		return nil, errors.NewBadRequestError("Validation failed")
	}

	if err := validate.Struct(payload); err != nil {
		return nil, formatValidationError(err)
	}

	return &payload, nil
}

func formatValidationError(err error) *errors.AppError {
	var errorMessages []string

	for _, err := range err.(validator.ValidationErrors) {
		errorMessages = append(errorMessages, strings.ToLower(err.Field()))
	}

	return &errors.AppError{
		Code:    422,
		Message: strings.Join(errorMessages, ", "),
		Details: errorMessages,
	}
}
