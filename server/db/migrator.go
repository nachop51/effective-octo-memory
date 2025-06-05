package db

import "server/users"

func Migrate() {
	Conn.AutoMigrate(&users.User{})
}
