package users

type UserRepository interface {
	GetUsers() ([]*User, error)
	GetUserByID(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) error
}
