package main

import (
	"log"

	"server/db"
	"server/routes"
	"server/users"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file")
	}

	db.Connect()
	db.AutoMigrate(&users.User{})

	app := routes.Setup()

	app.Listen("0.0.0.0:1234")
}
