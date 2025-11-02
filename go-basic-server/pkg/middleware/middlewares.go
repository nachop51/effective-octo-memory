package middleware

import (
	"strings"

	"server/pkg/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v5"
)

var unprotectedRoutes = map[string]bool{
	"/login":  true,
	"/signup": true,
	"/health": true,
}

var ErrorUnauthorized = errors.NewUnauthorizedError("Missing or invalid token")

func NewAuthMiddleware(jwtKey []byte, cookieName string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		path := c.Path()

		path = strings.TrimRight(path, "/")

		if _, ok := unprotectedRoutes[path]; ok {
			return c.Next()
		}

		token := c.Cookies(cookieName)
		if token == "" {
			token = c.Get("Authorization")
		}

		if token == "" {
			return ErrorUnauthorized
		}

		token = strings.TrimPrefix(token, "Bearer ")

		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
			// Ensure the token's signing method is valid
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrorUnauthorized
			}
			// Replace with your secret key
			return jwtKey, nil
		})
		if err != nil || !parsedToken.Valid {
			return ErrorUnauthorized
		}

		claims := parsedToken.Claims.(jwt.MapClaims)

		userId, err := claims.GetSubject()
		if err != nil {
			return ErrorUnauthorized
		}

		c.Locals("UserId", userId)

		return c.Next()
	}
}

func SetupMiddlewares(app *fiber.App, jwtKey []byte, cookieName string) {
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
	}))
	app.Use(NewAuthMiddleware(jwtKey, cookieName))
}
