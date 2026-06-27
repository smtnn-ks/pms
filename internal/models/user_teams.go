package models

type UserRole int

const (
	UserRoleInvalid UserRole = iota
	UserRoleOwner
	UserRoleAdmin
	UserRoleParticipant
)

type UserTeam struct {
	UserID uint `gorm:"primaryKey"`
	User   User

	TeamID uint `gorm:"primaryKey"`
	Team   Team

	UserRole UserRole
}
