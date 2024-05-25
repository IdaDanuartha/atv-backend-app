package inputs

import "github.com/IdaDanuartha/atv-backend-app/app/enums"

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UpdateProfileInput struct {
	Name        string     `json:"name" binding:"required"`
	Username    string     `json:"username" binding:"required"`
	Email       string     `json:"email" binding:"required,email"`
	Role        enums.Role `json:"role" binding:"required"`
	ProfilePath string     `json:"profile_path"`
}
