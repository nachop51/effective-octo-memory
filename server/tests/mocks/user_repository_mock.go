package mocks

import (
	"server/domains/users"

	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of the UserRepository interface
type MockUserRepository struct {
	mock.Mock
}

// GetUsers mocks the GetUsers method
func (m *MockUserRepository) GetUsers() ([]*users.User, error) {
	args := m.Called()
	return args.Get(0).([]*users.User), args.Error(1)
}

// CreateUser mocks the CreateUser method
func (m *MockUserRepository) CreateUser(user *users.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// GetUserByEmail mocks the GetUserByEmail method
func (m *MockUserRepository) GetUserByEmail(email string) (*users.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*users.User), args.Error(1)
}

// GetUserByID mocks the GetUserByID method
func (m *MockUserRepository) GetUserByID(id string) (*users.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*users.User), args.Error(1)
}
