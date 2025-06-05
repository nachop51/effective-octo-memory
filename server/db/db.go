package db

import (
	"log"

	"server/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Connect() {
	dsn := config.GetDBDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	Conn = db

	log.Println("Database connected")
}
