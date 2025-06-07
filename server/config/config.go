package config

import (
	"log"
	"os"

	"server/db"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	Validate = validator.New()
	JwtKey   []byte
	Conn     *gorm.DB
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file")
	}

	JwtKey = []byte(os.Getenv("JWT_KEY"))

	if len(JwtKey) == 0 {
		log.Fatalln("JWT_KEY is not set")
	}

	Conn = db.Connect()
}
