package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model

	TaskID uint
	Task   Task

	UserID uint
	User   User

	Content string
}
