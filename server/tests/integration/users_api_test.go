package integration_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"server/app"
	"server/config"
	"server/domains/users"
	test_config "server/tests/config"
	"server/tests/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserAPITestSuite struct {
	suite.Suite
	app *app.App
	db  *gorm.DB
}

func (suite *UserAPITestSuite) SetupSuite() {
	helpers.LoadEnv()

	// Setup test database
	testDBURL := os.Getenv("TEST_DATABASE_URL")

	database, err := gorm.Open(postgres.Open(testDBURL), &gorm.Config{})
	suite.Require().NoError(err)

	suite.db = database

	// Run migrations
	err = database.AutoMigrate(&users.User{})
	suite.Require().NoError(err)

	// Setup test config
	appConfig := test_config.GetTestConfig()

	// Create dependencies
	deps := &config.Dependencies{
		DB:       database,
		Validate: validator.New(),
		Config:   appConfig,
	}

	// Create and setup app
	suite.app = app.New(deps)
	suite.app.Setup()
}

func (suite *UserAPITestSuite) TearDownSuite() {
	if suite.db != nil {
		sqlDB, _ := suite.db.DB()
		sqlDB.Close()
	}
}

func (suite *UserAPITestSuite) SetupTest() {
	// Clean up database before each test
	suite.db.Exec("DELETE FROM users")
}

func (suite *UserAPITestSuite) TestCreateUser_Success() {
	// Arrange
	userBody := helpers.CreateUserBody()
	jsonBody, _ := json.Marshal(userBody)

	// Act
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.app.GetServer().Test(req)

	// Assert
	suite.Require().NoError(err)
	suite.Equal(http.StatusCreated, resp.StatusCode)

	bodyBytes, _ := io.ReadAll(resp.Body)
	var createdUser users.User
	err = json.Unmarshal(bodyBytes, &createdUser)
	suite.Require().NoError(err)

	suite.Equal(userBody.FirstName, createdUser.FirstName)
	suite.Equal(userBody.LastName, createdUser.LastName)
	suite.Equal(userBody.Email, createdUser.Email)
	suite.NotEmpty(createdUser.ID)
	suite.NotEmpty(createdUser.CreatedAt)
	suite.NotEmpty(createdUser.UpdatedAt)

	// Verify user was saved in database
	var dbUser users.User
	err = suite.db.Where("email = ?", userBody.Email).First(&dbUser).Error
	suite.Require().NoError(err)
	suite.Equal(createdUser.ID, dbUser.ID)
}

func (suite *UserAPITestSuite) TestCreateUser_ValidationError() {
	// Arrange - Invalid user body (missing required fields)
	userBody := users.UserBody{
		FirstName:            "", // Missing required field
		LastName:             "Doe",
		Email:                "invalid-email", // Invalid email format
		Password:             "123",           // Too short
		PasswordConfirmation: "123",
	}
	jsonBody, _ := json.Marshal(userBody)

	// Act
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.app.GetServer().Test(req)

	// Assert
	suite.Require().NoError(err)
	suite.Equal(http.StatusBadRequest, resp.StatusCode)

	// Verify user was not saved in database
	var count int64
	suite.db.Model(&users.User{}).Count(&count)
	suite.Equal(int64(0), count)
}

func (suite *UserAPITestSuite) TestCreateUser_DuplicateEmail() {
	// Arrange - Create first user
	userBody := helpers.CreateUserBody()
	jsonBody, _ := json.Marshal(userBody)

	// Create first user
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.app.GetServer().Test(req)
	suite.Require().NoError(err)
	suite.Equal(http.StatusCreated, resp.StatusCode)

	// Act - Try to create second user with same email
	req, _ = http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err = suite.app.GetServer().Test(req)

	// Assert
	suite.Require().NoError(err)
	suite.Equal(http.StatusConflict, resp.StatusCode)

	// Verify only one user exists in database
	var count int64
	suite.db.Model(&users.User{}).Count(&count)
	suite.Equal(int64(1), count)
}

func (suite *UserAPITestSuite) TestGetUsers_Success() {
	// Arrange - Create test users
	testUsers := []*users.User{
		helpers.CreateTestUserWithEmail("user1@example.com"),
		helpers.CreateTestUserWithEmail("user2@example.com"),
	}

	for _, user := range testUsers {
		suite.db.Create(user)
	}

	// Act
	req, _ := http.NewRequest("GET", "/users", nil)
	resp, err := suite.app.GetServer().Test(req)

	// Assert
	suite.Require().NoError(err)
	suite.Equal(http.StatusOK, resp.StatusCode)

	bodyBytes, _ := io.ReadAll(resp.Body)
	var returnedUsers []*users.User
	err = json.Unmarshal(bodyBytes, &returnedUsers)
	suite.Require().NoError(err)

	suite.Len(returnedUsers, 2)
	suite.Equal("user1@example.com", returnedUsers[0].Email)
	suite.Equal("user2@example.com", returnedUsers[1].Email)
}

func (suite *UserAPITestSuite) TestGetUsers_Empty() {
	// Act
	req, _ := http.NewRequest("GET", "/users", nil)
	resp, err := suite.app.GetServer().Test(req)

	// Assert
	suite.Require().NoError(err)
	suite.Equal(http.StatusOK, resp.StatusCode)

	bodyBytes, _ := io.ReadAll(resp.Body)
	var returnedUsers []*users.User
	err = json.Unmarshal(bodyBytes, &returnedUsers)
	suite.Require().NoError(err)

	suite.Empty(returnedUsers)
}

func (suite *UserAPITestSuite) TestLogin_Success() {
	// Arrange - Create a user first
	userBody := helpers.CreateUserBody()
	jsonBody, _ := json.Marshal(userBody)

	// Create user
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.app.GetServer().Test(req)
	suite.Require().NoError(err)
	suite.Equal(http.StatusCreated, resp.StatusCode)

	// Prepare login request
	loginBody := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    userBody.Email,
		Password: userBody.Password,
	}
	loginJSON, _ := json.Marshal(loginBody)

	// Act
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(loginJSON))
	req.Header.Set("Content-Type", "application/json")

	resp, err = suite.app.GetServer().Test(req)

	// Assert
	suite.Require().NoError(err)
	suite.Equal(http.StatusOK, resp.StatusCode)

	bodyBytes, _ := io.ReadAll(resp.Body)
	var loginResponse struct {
		Token string `json:"token"`
	}
	err = json.Unmarshal(bodyBytes, &loginResponse)
	suite.Require().NoError(err)

	suite.NotEmpty(loginResponse.Token)
}

func (suite *UserAPITestSuite) TestLogin_InvalidCredentials() {
	// Arrange
	loginBody := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    "nonexistent@example.com",
		Password: "wrongpassword",
	}
	loginJSON, _ := json.Marshal(loginBody)

	// Act
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(loginJSON))
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.app.GetServer().Test(req)

	// Assert
	suite.Require().NoError(err)
	suite.Equal(http.StatusUnauthorized, resp.StatusCode)
}

func TestUserAPITestSuite(t *testing.T) {
	suite.Run(t, new(UserAPITestSuite))
}
