package routes

import (
	"server/users"

	"github.com/gofiber/fiber/v2"
)

type AppHandlers struct {
	UserHandler *users.UserHandler
}

func Setup(handlers *AppHandlers) *fiber.App {
	app := fiber.New(fiber.Config{})

	setupMiddlewares(app)

	handlers.UserHandler.RegisterRoutes(app)

	return app
}
