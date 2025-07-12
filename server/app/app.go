package app

import (
	"fmt"

	"server/config"
	"server/domains/accounts"
	"server/domains/users"
	"server/pkg/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type App struct {
	server    *fiber.App
	config    *config.Config
	db        *gorm.DB
	validator *validator.Validate
}

func New(dps *config.Dependencies) *App {
	return &App{
		server:    fiber.New(),
		config:    dps.Config,
		db:        dps.DB,
		validator: dps.Validate,
	}
}

func (a *App) setupMiddleware() {
	middleware.SetupMiddlewares(a.server, a.config.JWTConfig.SecretKey)
}

func (a *App) setupRoutes() {
	a.server.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Users
	userStore := users.NewUserStore(a.db)
	userService := users.NewUserService(userStore, a.config.JWTConfig.SecretKey)
	userHandler := users.NewUserHandler(userService, a.validator)
	userHandler.RegisterRoutes(a.server)

	// Accounts
	accountStore := accounts.NewAccountStore(a.db)
	accountService := accounts.NewAccountService(accountStore)
	accountHandler := accounts.NewAccountHandler(accountService, a.validator)
	accountHandler.RegisterRoutes(a.server)
}

func (a *App) Setup() {
	a.setupMiddleware()
	a.setupRoutes()
}

func (a *App) Start() error {
	addr := fmt.Sprintf("%s:%d", a.config.Server.Host, a.config.Server.Port)
	return a.server.Listen(addr)
}

func (a *App) GetServer() *fiber.App {
	return a.server
}
