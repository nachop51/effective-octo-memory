package main

import (
	"server/config"
	"server/db"
	"server/server"
	"server/users"
)

func main() {
	config.Init()

	db.AutoMigrate(config.Conn, &users.User{})

	app := server.NewServer(config.Conn)

	app.SetupRoutes()

	app.Start("0.0.0.0:1234")
}
