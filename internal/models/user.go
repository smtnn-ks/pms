package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint

	Email        string
	PasswordHash []byte

	Teams []Team `gorm:"many2many:user_teams"`

	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
