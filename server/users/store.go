package users

import "gorm.io/gorm"

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) CreateUser(user *User) error {
	return s.db.Create(user).Error
}

func (s *UserStore) GetUserByEmail(email string) (*User, error) {
	var user User

	err := s.db.Where("email = ?", email).First(&user)

	if err.Error != nil {
		return nil, err.Error
	}

	return &user, nil
}
