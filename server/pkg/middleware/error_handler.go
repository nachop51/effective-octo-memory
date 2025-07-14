package middleware

import (
	"log"

	"server/pkg/errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if appErr, ok := errors.IsAppError(err); ok {
		return c.Status(appErr.Code).JSON(fiber.Map{
			"code":    appErr.Code,
			"error":   appErr.Message,
			"details": appErr.Details,
		})
	}

	if err == gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Resource not found",
			"code":  fiber.StatusNotFound,
		})
	}

	// Handle Fiber errors
	if fiberErr, ok := err.(*fiber.Error); ok {
		return c.Status(fiberErr.Code).JSON(fiber.Map{
			"error": fiberErr.Message,
			"code":  fiberErr.Code,
		})
	}

	// Log unexpected errors
	log.Printf("Unexpected error: %v", err)

	// Default to 500 for unknown errors
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Internal server error",
		"code":  fiber.StatusInternalServerError,
	})
}
