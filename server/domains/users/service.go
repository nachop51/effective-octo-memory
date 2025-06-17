package users

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	store  UserRepository
	jwtKey []byte
}

func NewUserService(store UserRepository, jwtKey []byte) *UserService {
	return &UserService{
		store:  store,
		jwtKey: jwtKey,
	}
}

func (s *UserService) GetUsers() ([]*User, error) {
	return s.store.GetUsers()
}

func (s *UserService) CreateUser(body UserBody) (*User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  password,
	}

	res := s.store.CreateUser(user)

	if res != nil {
		return nil, res
	}

	return user, nil
}

func (s *UserService) GetUser(email string) (*User, error) {
	return s.store.GetUserByEmail(email)
}

func (s *UserService) GenerateJWT(user *User) (string, error) {
	now := time.Now()

	claims := jwt.RegisteredClaims{
		Subject:   strconv.Itoa(int(user.ID)),
		Issuer:    "effective-octo-memory",
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24 * 30)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(s.jwtKey)
}
