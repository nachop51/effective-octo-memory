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

func (s *UserStore) GetUsers() ([]*User, error) {
	var users []*User

	res := s.db.Find(&users)

	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

func (s *UserStore) GetUserByID(id string) (*User, error) {
	var user User

	err := s.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserStore) CreateUser(user *User) error {
	return s.db.Create(user).Error
}

func (s *UserStore) GetUserByEmail(email string) (*User, error) {
	var user User

	err := s.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
