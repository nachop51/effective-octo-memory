package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Server    ServerConfig   `envPrefix:"SERVER_"`
	Database  DatabaseConfig `envPrefix:"DB_"`
	JWTConfig JWTConfig
	Security  SecurityConfig
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

type JWTConfig struct {
	SecretKey []byte
}

type SecurityConfig struct {
	UnprotectedRoutes map[string]bool
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

	if err := cfg.parseJWTConfig(); err != nil {
		return nil, err
	}

	cfg.parseUnprotectedRoutes()

	return cfg, nil
}

func (c *Config) parseJWTConfig() error {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return fmt.Errorf("JWT_SECRET_KEY is not set")
	}
	c.JWTConfig.SecretKey = []byte(secretKey)
	return nil
}

func (c *Config) parseUnprotectedRoutes() {
	unprotectedRoutes := os.Getenv("UNPROTECTED_ROUTES")
	if unprotectedRoutes == "" {
		c.Security.UnprotectedRoutes = make(map[string]bool)
		return
	}

	routes := strings.Split(unprotectedRoutes, ",")
	c.Security.UnprotectedRoutes = make(map[string]bool, len(routes))
	for _, route := range routes {
		route = strings.TrimSpace(route)
		if route != "" {
			c.Security.UnprotectedRoutes[route] = true
		}
	}
	if len(c.Security.UnprotectedRoutes) == 0 {
		c.Security.UnprotectedRoutes = make(map[string]bool)
	}
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
