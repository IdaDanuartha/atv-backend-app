package models

import (
	"github.com/IdaDanuartha/atv-backend-app/app/enums"
)

// User Model
type User struct {
	Base
	Username    string     `gorm:"size:100" json:"username"`
	Email       string     `gorm:"size:100" json:"email"`
	Password    string     `gorm:"size:100" json:"password,omitempty"`
	Role        enums.Role `gorm:"type:enum('admin', 'staff', 'instructor', 'customer')" json:"role"`
	ProfilePath *string    `gorm:"size:100;" json:"profile_path"`
}

// TableName method sets table name for Bus model
func (User *User) TableName() string {
	return "users"
}