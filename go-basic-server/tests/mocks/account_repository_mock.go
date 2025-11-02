package mocks

import (
	"server/domains/accounts"

	"github.com/stretchr/testify/mock"
)

// MockAccountRepository is a mock implementation of the AccountInterface
type MockAccountRepository struct {
	mock.Mock
}

// GetAccounts mocks the GetAccounts method
func (m *MockAccountRepository) GetAccounts() ([]*accounts.Account, error) {
	args := m.Called()
	return args.Get(0).([]*accounts.Account), args.Error(1)
}

// CreateAccount mocks the CreateAccount method
func (m *MockAccountRepository) CreateAccount(account *accounts.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

// GetAccountsByUserID mocks the GetAccountsByUserID method if it exists
func (m *MockAccountRepository) GetAccountsByUserID(userID uint) ([]*accounts.Account, error) {
	args := m.Called(userID)
	return args.Get(0).([]*accounts.Account), args.Error(1)
}
