package db

import (
	"log"

	"server/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

// Connect connects to the database.
func Connect() {
	dsn := config.GetDBDSN()
	var err error
	Conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Database connected")
}

// AutoMigrate migrates the database with the given models.
func AutoMigrate(models ...any) {
	if err := Conn.AutoMigrate(models...); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrated")
}
