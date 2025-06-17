package accounts

import (
	"time"

	"server/domains/users"
)

type Account struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name" gorm:"not null"`
	UserID    uint       `json:"userId" gorm:"not null"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	User      users.User `json:"-"`
}
