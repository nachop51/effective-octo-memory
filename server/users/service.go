package users

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	store  *UserStore
	jwtKey []byte
}

func NewUserService(store *UserStore, jwtKey []byte) *UserService {
	return &UserService{
		store:  store,
		jwtKey: jwtKey,
	}
}
func (s *UserService) CreateUser(body UserBody) (*User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  string(password),
	}

	res := s.store.CreateUser(&user)

	if res != nil {
		return nil, res
	}

	return &user, nil
}

func (s *UserService) GetUser(email string) (*User, error) {

	user, err := s.store.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GenerateJWT(user *User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	return token.SignedString(s.jwtKey)
}
