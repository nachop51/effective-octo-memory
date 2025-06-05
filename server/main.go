package main

import (
	"log"

	"server/db"
	"server/routes"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file")
	}

	db.Connect()

	app := routes.Setup()

	app.Listen("0.0.0.0:1234")
}
