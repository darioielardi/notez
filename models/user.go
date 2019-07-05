package models

import (
	"notez/utils/enums"
)

// User is the user model
type User struct {
	Model
	Email  string     `gorm:"unique;not null" json:"email"`
	Name   string     `gorm:"unique;not null" json:"name"`
	AuthID string     `gorm:"column:auth_id;unique;not null" json:"auth_id"`
	Role   enums.Role `gorm:"not null" sql:"type:user_role" json:"role"`
	
	// Relations
	Notes []Note `json:"notes"`
}
