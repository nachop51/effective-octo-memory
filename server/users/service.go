package users

import "server/db"

func CreateUser(body UserBody) error {
	user := User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
	}

	res := db.Conn.Create(&user)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
