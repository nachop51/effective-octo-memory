package main

import (
	"server/config"
	"server/db"
	"server/routes"
	"server/users"
)

func main() {
	config.Init()

	db.AutoMigrate(config.Conn, &users.User{})

	userStore := users.NewUserStore(config.Conn)
	userService := users.NewUserService(userStore, config.JwtKey)
	userHandler := users.NewUserHandler(userService)

	app := routes.Setup(&routes.AppHandlers{
		UserHandler: userHandler,
	})

	app.Listen("0.0.0.0:1234")
}
