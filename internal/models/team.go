package models

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	ID uint

	Title       string
	Description string

	Users []User `gorm:"many2many:user_teams"`

	Tasks []Task `gorm:"constraint:OnDelete:CASCADE"`

	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
