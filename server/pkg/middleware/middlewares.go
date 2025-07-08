package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v5"
)

func NewAuthMiddleware(jwtKey []byte, unprotectedRoutes map[string]bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		path := c.Path()

		path = strings.Trim(path, "/")

		if _, ok := unprotectedRoutes[path]; ok {
			return c.Next()
		}

		token := c.Get("Authorization")

		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing or invalid token",
			})
		}

		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}

		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
			// Ensure the token's signing method is valid
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			// Replace with your secret key
			return jwtKey, nil
		})
		if err != nil || !parsedToken.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token",
			})
		}

		return c.Next()
	}
}

func SetupMiddlewares(app *fiber.App, jwtKey []byte, unprotectedRoutes map[string]bool) {
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(NewAuthMiddleware(jwtKey, unprotectedRoutes))
}
