package config

import (
	"os"

	"server/config"

	"github.com/joho/godotenv"
)

// GetTestConfig returns a configuration suitable for testing
func GetTestConfig() *config.Config {
	godotenv.Load(".env.test")

	return &config.Config{
		Server: config.ServerConfig{
			Host: "localhost",
			Port: 8080,
			Env:  "test",
		},
		Database: config.DatabaseConfig{
			Host:     getEnvOrDefault("TEST_DB_HOST", "localhost"),
			Port:     5432,
			Username: getEnvOrDefault("TEST_DB_USERNAME", "user"),
			Password: getEnvOrDefault("TEST_DB_PASSWORD", "password"),
			Name:     getEnvOrDefault("TEST_DB_NAME", "test_db"),
			SSLMode:  "disable",
		},
		JWTConfig: config.JWTConfig{
			SecretKey: []byte("test-jwt-secret-key-for-testing"),
		},
	}
}

// GetTestDatabaseURL returns a test database URL
func GetTestDatabaseURL() string {
	testDBURL := os.Getenv("TEST_DATABASE_URL")
	if testDBURL == "" {
		testDBURL = "postgres://user:password@localhost:5432/test_db?sslmode=disable"
	}
	return testDBURL
}

// getEnvOrDefault gets an environment variable or returns a default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// IsTestEnvironment checks if we're running in test environment
func IsTestEnvironment() bool {
	return os.Getenv("GO_ENV") == "test" || os.Getenv("SERVER_ENV") == "test"
}
