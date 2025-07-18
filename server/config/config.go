package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig   `envPrefix:"SERVER_"`
	Database DatabaseConfig `envPrefix:"DB_"`
	Auth     AuthConfig
}

type ServerConfig struct {
	Host string `env:"HOST" envDefault:"0.0.0.0"`
	Port int    `env:"PORT" envDefault:"1234"`
	Env  string `env:"ENV" envDefault:"development"`
}

type DatabaseConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     int    `env:"PORT" envDefault:"5432"`
	Username string `env:"USERNAME" envDefault:"user"`
	Password string `env:"PASSWORD" envDefault:"password"`
	Name     string `env:"NAME" envDefault:"mydb"`
	SSLMode  string `env:"SSL_MODE" envDefault:"disable"`
}
type AuthConfig struct {
	SecretKey  []byte
	CookieName string `env:"COOKIE_NAME" envDefault:"effective_octo_auth_token"`
	// CookieExpires time.Duration
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			return nil, err // Only fail for real errors
		}
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	if err := cfg.ParseJWTConfig(); err != nil {
		return nil, err
	}

	// if err := cfg.parseCookieExpires(); err != nil {
	// 	return nil, err
	// }

	return cfg, nil
}

// func (c *Config) parseCookieExpires() error {
// 	timeInSeconds := os.Getenv("COOKIE_EXPIRES")
// 	if timeInSeconds == "" {
// 		return fmt.Errorf("COOKIE_EXPIRES is not set")
// 	}

// 	val, err := strconv.ParseInt(timeInSeconds, 10, 64)
// 	if err != nil {
// 		return fmt.Errorf("invalid COOKIE_EXPIRES value: %v", err)
// 	}

// 	c.Auth.CookieExpires = time.Duration(val) * time.Second

// 	return nil
// }

func (c *Config) ParseJWTConfig() error {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return fmt.Errorf("JWT_SECRET_KEY is not set")
	}
	c.Auth.SecretKey = []byte(secretKey)
	return nil
}

func (c *Config) IsDevelopment() bool {
	return c.Server.Env == "development"
}

func (c *Config) GetDBDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.Username,
		c.Database.Password,
		c.Database.Name,
		c.Database.SSLMode,
	)
}
