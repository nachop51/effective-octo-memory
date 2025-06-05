package routes

import (
	"server/users"

	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App {
	app := fiber.New(fiber.Config{})

	setupMiddlewares(app)

	users.RegisterRoutes(app)

	return app
}
