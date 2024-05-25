package models

import (
	"github.com/IdaDanuartha/atv-backend-app/app/enums"
)

// User Model
type User struct {
	Base
	Username    string     `gorm:"size:100" json:"username"`
	Email       string     `gorm:"size:100" json:"email"`
	Password    string     `gorm:"size:100" json:"password"`
	Role        enums.Role `gorm:"type:enum('admin', 'staff', 'instructor', 'customer')" json:"role"`
	ProfilePath *string    `gorm:"size:100;" json:"profile_path"`
}

// TableName method sets table name for Bus model
func (User *User) TableName() string {
	return "users"
}

// ResponseMap -> response map method of Entertainment Category
func (User *User) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = User.ID
	resp["username"] = User.Username
	resp["email"] = User.Email
	resp["password"] = User.Password
	resp["role"] = User.Role
	resp["profile_path"] = User.ProfilePath
	resp["created_at"] = User.CreatedAt
	resp["updated_at"] = User.UpdatedAt
	resp["deleted_at"] = User.DeletedAt

	return resp
}
