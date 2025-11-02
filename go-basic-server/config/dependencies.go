package config

import (
	"server/pkg/db"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Dependencies struct {
	DB       *gorm.DB
	Validate *validator.Validate
	Config   *Config
}

func NewDependencies(cfg *Config) (*Dependencies, error) {
	dbConn, err := db.Connect(cfg.GetDBDSN())
	if err != nil {
		return nil, err
	}

	return &Dependencies{
		DB:       dbConn,
		Validate: validator.New(),
		Config:   cfg,
	}, nil
}
