package utils

import (
	"server/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

func GetUserID(c *fiber.Ctx) (string, error) {
	UserId, ok := c.Locals("UserId").(string)

	if !ok {
		return "", errors.NewUnauthorizedError("User ID not found in context")
	}

	return UserId, nil
}
