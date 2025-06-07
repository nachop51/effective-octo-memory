package users

import (
	"time"

	"server/config"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(body UserBody) (*User, error) {
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

	res := config.Conn.Create(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func GetUser(email string) (*User, error) {
	var user User

	err := config.Conn.Where("email = ?", email).First(&user)

	if err.Error != nil {
		return nil, err.Error
	}

	return &user, nil
}

func GenerateJWT(user *User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	return token.SignedString(config.JwtKey)
}
