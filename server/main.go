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

	app := routes.Setup()

	app.Listen("0.0.0.0:1234")
}
