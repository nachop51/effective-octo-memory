package server

import (
	"server/config"
	"server/routes"
	"server/users"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	app *fiber.App
	db  *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	app := fiber.New(fiber.Config{})
	return &Server{
		app: app,
		db:  db,
	}
}

func (s *Server) registerUserRoutes() {
	userStore := users.NewUserStore(config.Conn)
	userService := users.NewUserService(userStore, config.JwtKey)
	userHandler := users.NewUserHandler(userService)

	userHandler.RegisterRoutes(s.app)
}

func (s *Server) SetupRoutes() {
	routes.SetupMiddlewares(s.app)

	s.registerUserRoutes()
}

func (s *Server) Start(addr string) error {
	return s.app.Listen(addr)
}
