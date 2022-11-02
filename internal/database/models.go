package database

import (
	"time"
)

type User struct {
	ID          uint
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Password    string
	AccessLevel int `gorm:"default:1"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
