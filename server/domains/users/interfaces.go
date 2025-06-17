package users

type UserRepository interface {
	GetUsers() ([]*User, error)
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
}
