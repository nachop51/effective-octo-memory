package users_test

import (
	"strings"
	"testing"
	"time"

	"server/domains/users"
	test_config "server/tests/config"
	"server/tests/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestUserService_GetUsers(t *testing.T) {
	// Arrange
	cfg := test_config.GetTestConfig()
	mockRepo := new(mocks.MockUserRepository)
	service := users.NewUserService(mockRepo, cfg)

	expectedUsers := []*users.User{
		{
			ID:        "random cuid2",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@example.com",
			Role:      "user",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "random cuid2 - 1",
			FirstName: "Jane",
			LastName:  "Smith",
			Email:     "jane@example.com",
			Role:      "user",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	mockRepo.On("GetUsers").Return(expectedUsers, nil)

	// Act
	result, err := service.GetUsers()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, result)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUsers_Error(t *testing.T) {
	// Arrange
	cfg := test_config.GetTestConfig()
	mockRepo := new(mocks.MockUserRepository)
	service := users.NewUserService(mockRepo, cfg)

	mockRepo.On("GetUsers").Return(([]*users.User)(nil), assert.AnError)

	// Act
	result, err := service.GetUsers()

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUserService_CreateUser_Success(t *testing.T) {
	// Arrange
	cfg := test_config.GetTestConfig()
	mockRepo := new(mocks.MockUserRepository)
	service := users.NewUserService(mockRepo, cfg)

	userBody := users.UserBody{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  "password123",
	}

	mockRepo.On("CreateUser", mock.AnythingOfType("*users.User")).Return(nil)

	// Act
	result, err := service.CreateUser(userBody)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, userBody.FirstName, result.FirstName)
	assert.Equal(t, userBody.LastName, result.LastName)
	assert.Equal(t, userBody.Email, result.Email)
	assert.NotEmpty(t, result.Password)

	// Verify password is hashed
	err = bcrypt.CompareHashAndPassword(result.Password, []byte(userBody.Password))
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUserService_CreateUser_RepositoryError(t *testing.T) {
	// Arrange
	cfg := test_config.GetTestConfig()
	mockRepo := new(mocks.MockUserRepository)
	service := users.NewUserService(mockRepo, cfg)

	userBody := users.UserBody{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  "password123",
	}

	mockRepo.On("CreateUser", mock.AnythingOfType("*users.User")).Return(assert.AnError)

	// Act
	result, err := service.CreateUser(userBody)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUser_Success(t *testing.T) {
	// Arrange
	cfg := test_config.GetTestConfig()
	mockRepo := new(mocks.MockUserRepository)
	service := users.NewUserService(mockRepo, cfg)

	expectedUser := &users.User{
		ID:        "random cuid2 - 2",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	email := "john@example.com"
	mockRepo.On("GetUserByEmail", email).Return(expectedUser, nil)

	// Act
	result, err := service.GetUserByEmail(email)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUser_NotFound(t *testing.T) {
	// Arrange
	cfg := test_config.GetTestConfig()
	mockRepo := new(mocks.MockUserRepository)
	service := users.NewUserService(mockRepo, cfg)

	email := "nonexistent@example.com"
	mockRepo.On("GetUserByEmail", email).Return((*users.User)(nil), assert.AnError)

	// Act
	result, err := service.GetUserByEmail(email)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GenerateJWT_Success(t *testing.T) {
	// Arrange
	cfg := test_config.GetTestConfig()
	mockRepo := new(mocks.MockUserRepository)
	service := users.NewUserService(mockRepo, cfg)

	user := &users.User{
		ID:        "random cuid2 - 3",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Role:      "user",
	}

	// Act
	token, err := service.GenerateJWT(user)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.IsType(t, "", token)

	// Basic JWT structure validation (header.payload.signature)
	tokenParts := len(strings.Split(token, "."))
	assert.Equal(t, 3, tokenParts)
}

func TestUserService_GenerateJWT_EmptySecret(t *testing.T) {
	// Arrange
	cfg := test_config.GetTestConfig()
	mockRepo := new(mocks.MockUserRepository)
	service := users.NewUserService(mockRepo, cfg)

	user := &users.User{
		ID:        "random cuid2 - 4",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Role:      "user",
	}

	// Act
	token, err := service.GenerateJWT(user)

	// Assert
	assert.NoError(t, err) // JWT will still be generated with empty secret
	assert.NotEmpty(t, token)
}
