package test_config

import (
	"os"

	"server/config"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

// GetTestConfig returns a configuration suitable for testing
func GetTestConfig() *config.Config {
	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			return nil
		}
	}

	cfg := &config.Config{}
	if err := env.Parse(cfg); err != nil {
		return nil
	}

	if err := cfg.ParseJWTConfig(); err != nil {
		return nil
	}

	return cfg
}

// GetTestDatabaseURL returns a test database URL
func GetTestDatabaseURL() string {
	testDBURL := os.Getenv("TEST_DATABASE_URL")
	if testDBURL == "" {
		testDBURL = "postgres://user:password@localhost:5432/test_db?sslmode=disable"
	}
	return testDBURL
}

// IsTestEnvironment checks if we're running in test environment
func IsTestEnvironment() bool {
	return os.Getenv("GO_ENV") == "test" || os.Getenv("SERVER_ENV") == "test"
}
