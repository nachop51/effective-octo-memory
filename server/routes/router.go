package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *fiber.App {
	app := fiber.New(fiber.Config{})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return app
}
