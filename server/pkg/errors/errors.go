package errors

import (
	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Code:    fiber.StatusBadRequest,
		Message: message,
	}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Code:    fiber.StatusUnauthorized,
		Message: message,
	}
}

func NewForbiddenError(message string) *AppError {
	return &AppError{
		Code:    fiber.StatusForbidden,
		Message: message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    fiber.StatusNotFound,
		Message: message,
	}
}

func NewInternalServerError(message string) *AppError {
	return &AppError{
		Code:    fiber.StatusInternalServerError,
		Message: message,
	}
}

func NewConflictError(message string) *AppError {
	return &AppError{
		Code:    fiber.StatusConflict,
		Message: message,
	}
}

func NewValidationError(message string, details any) *AppError {
	return &AppError{
		Code:    fiber.StatusUnprocessableEntity,
		Message: message,
		Details: details,
	}
}

func NewUnprocessableEntityError(message string, details any) *AppError {
	return &AppError{
		Code:    fiber.StatusUnprocessableEntity,
		Message: message,
		Details: details,
	}
}

func IsAppError(err error) (*AppError, bool) {
	appErr, ok := err.(*AppError)
	return appErr, ok
}
