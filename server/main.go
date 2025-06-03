package main

import (
	"log"
	"server/config"
	"server/routes"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file")
	}

	db := config.InitDB()

	app := routes.Setup(db)

	app.Listen("0.0.0.0:1234")
}
