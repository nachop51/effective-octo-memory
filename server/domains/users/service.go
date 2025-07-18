package users

import (
	"time"

	"server/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nrednav/cuid2"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	store  UserRepository
	config *config.Config
}

func NewUserService(store UserRepository, config *config.Config) *UserService {
	return &UserService{
		store:  store,
		config: config,
	}
}

type CustomClaims struct {
	jwt.RegisteredClaims
	User map[string]string `json:"user"`
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
		ID:        cuid2.Generate(),
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

func (s *UserService) GetUserByID(id string) (*User, error) {
	return s.store.GetUserByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*User, error) {
	return s.store.GetUserByEmail(email)
}

func (s *UserService) GenerateJWT(user *User) (string, error) {
	now := time.Now()

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = &CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID,
			Issuer:    "effective-octo-memory",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(30 * 24 * time.Hour)),
		},
		User: map[string]string{
			"id":        user.ID,
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"email":     user.Email,
		},
	}

	return token.SignedString(s.config.Auth.SecretKey)
}
