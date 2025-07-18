package helpers

import (
	"log"
	"os"
	"regexp"
	"time"

	"server/domains/users"

	"github.com/joho/godotenv"
)

// CreateTestUser creates a test user with default values
func CreateTestUser() *users.User {
	return &users.User{
		ID:        "random cuid2",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  []byte("$2a$10$test.hash.password"),
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// CreateTestUserWithEmail creates a test user with a specific email
func CreateTestUserWithEmail(email string) *users.User {
	user := CreateTestUser()
	user.Email = email
	return user
}

// CreateTestUserWithID creates a test user with a specific ID
func CreateTestUserWithID(id string) *users.User {
	user := CreateTestUser()
	user.ID = id
	return user
}

// CreateTestUsers creates multiple test users
func CreateTestUsers(count int) []*users.User {
	users := make([]*users.User, count)
	for i := range count {
		users[i] = CreateTestUserWithID("random cuid2" + string(i))
		users[i].Email = "user" + string(rune(i+'0')) + "@example.com"
	}
	return users
}

// CreateUserBody creates a test UserBody for requests
func CreateUserBody() users.UserBody {
	return users.UserBody{
		FirstName:            "John",
		LastName:             "Doe",
		Email:                "john@example.com",
		Password:             "password123",
		PasswordConfirmation: "password123",
	}
}

// CreateUserBodyWithEmail creates a test UserBody with specific email
func CreateUserBodyWithEmail(email string) users.UserBody {
	body := CreateUserBody()
	body.Email = email
	return body
}

// TestJWTSecret returns a consistent JWT secret for testing
func TestJWTSecret() []byte {
	return []byte("test-jwt-secret-key-for-testing")
}

const projectDirName = "server"

// LoadEnv loads env vars from .env
func LoadEnv() {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env.test`)
	if err != nil {
		log.Print(map[string]any{
			"cause": err,
			"cwd":   cwd,
		})

		os.Exit(-1)
	}
}
