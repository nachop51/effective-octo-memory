package accounts_test

import (
	"testing"
	"time"

	"server/domains/accounts"
	"server/tests/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAccountService_GetAccounts(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockAccountRepository)
	service := accounts.NewAccountService(mockRepo)

	expectedAccounts := []*accounts.Account{
		{
			ID:        1,
			Name:      "Test Account 1",
			UserID:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "Test Account 2",
			UserID:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	mockRepo.On("GetAccounts").Return(expectedAccounts, nil)

	// Act
	result, err := service.GetAccounts()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedAccounts, result)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

func TestAccountService_GetAccounts_Error(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockAccountRepository)
	service := accounts.NewAccountService(mockRepo)

	mockRepo.On("GetAccounts").Return(([]*accounts.Account)(nil), assert.AnError)

	// Act
	result, err := service.GetAccounts()

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestAccountService_CreateAccount_Success(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockAccountRepository)
	service := accounts.NewAccountService(mockRepo)

	accountBody := &accounts.CreateAccountBody{
		Name:   "New Test Account",
		UserID: 1,
	}

	mockRepo.On("CreateAccount", mock.AnythingOfType("*accounts.Account")).Return(nil)

	// Act
	result, err := service.CreateAccount(accountBody)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, accountBody.Name, result.Name)
	assert.Equal(t, accountBody.UserID, result.UserID)
	mockRepo.AssertExpectations(t)
}

func TestAccountService_CreateAccount_RepositoryError(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockAccountRepository)
	service := accounts.NewAccountService(mockRepo)

	accountBody := &accounts.CreateAccountBody{
		Name:   "New Test Account",
		UserID: 1,
	}

	mockRepo.On("CreateAccount", mock.AnythingOfType("*accounts.Account")).Return(assert.AnError)

	// Act
	result, err := service.CreateAccount(accountBody)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestAccountService_CreateAccount_ValidatesInput(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockAccountRepository)
	service := accounts.NewAccountService(mockRepo)

	testCases := []struct {
		name        string
		accountBody *accounts.CreateAccountBody
		expectError bool
	}{
		{
			name: "Valid input",
			accountBody: &accounts.CreateAccountBody{
				Name:   "Valid Account",
				UserID: 1,
			},
			expectError: false,
		},
		{
			name: "Empty name",
			accountBody: &accounts.CreateAccountBody{
				Name:   "",
				UserID: 1,
			},
			expectError: false, // Service doesn't validate, handler does
		},
		{
			name: "Zero UserID",
			accountBody: &accounts.CreateAccountBody{
				Name:   "Test Account",
				UserID: 0,
			},
			expectError: false, // Service doesn't validate, handler does
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if !tc.expectError {
				mockRepo.On("CreateAccount", mock.AnythingOfType("*accounts.Account")).Return(nil).Once()
			}

			// Act
			result, err := service.CreateAccount(tc.accountBody)

			// Assert
			if tc.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tc.accountBody.Name, result.Name)
				assert.Equal(t, tc.accountBody.UserID, result.UserID)
			}
		})
	}

	mockRepo.AssertExpectations(t)
}
