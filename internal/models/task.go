package models

type Task struct {
	ID uint

	Title       string
	Description string

	TeamID uint
	Team   Team

	AuthorID uint
	Author   User `gorm:"foreignKey:AuthorID"`

	AssigneeID uint
	Assignee   User `gorm:"foreignKey:AssigneeID"`

	Changelogs []TaskChangelog `gorm:"constraint:OnDelete:CASCADE"`

	Comments []Comment `gorm:"constraint:OnDelete:CASCADE"`
}
