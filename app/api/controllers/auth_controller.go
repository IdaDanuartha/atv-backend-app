package controllers

import (
	"net/http"

	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/services"
	"github.com/IdaDanuartha/atv-backend-app/app/enums"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"github.com/IdaDanuartha/atv-backend-app/utils"
	"github.com/gin-gonic/gin"
)

// AuthController -> AuthController
type AuthController struct {
	service services.AuthService
}

type RegisterFormatter struct {
	ID    string      `json:"id"`
	Name  string      `json:"name"`
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

type LoginFormatter struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     enums.Role `json:"role"`
	Token    string `json:"token"`
}

func FormatRegister(customer models.Customer, token string) RegisterFormatter {
	formatter := RegisterFormatter{
		ID:    customer.ID,
		Name:  customer.Name,
		Token: token,
		User:  customer.User,
	}

	return formatter
}

func FormatLogin(user models.User, token string) LoginFormatter {
	formatter := LoginFormatter{
		ID:    user.ID,
		Token: token,
		Username: user.Username,
		Email: user.Email,
		Role: user.Role,
	}

	return formatter
}

// NewAuthController : NewAuthController
func NewAuthController(s services.AuthService) AuthController {
	return AuthController{
		service: s,
	}
}

func (h *AuthController) RegisterUser(ctx *gin.Context) {
	var input inputs.RegisterInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create user account")
		return
	}

	newUser, err := h.service.RegisterUser(input)

	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Register account failed")
		return
	}

	token, err := h.service.GenerateToken(newUser.ID)
	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Token invalid")
		return
	}

	formatter := FormatRegister(newUser, token)

	ctx.JSON(http.StatusOK, &utils.Response{
		Success: true,
		Message: "User registered successfully",
		Data:    formatter,
	})
}

func (h *AuthController) Login(ctx *gin.Context) {
	var input inputs.LoginInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Login Failed")
		return
	}

	loggedInUser, err := h.service.Login(input)

	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Login Failed! Check your credentials")
		return
	}

	token, err := h.service.GenerateToken(loggedInUser.ID)
	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Token invalid")
		return
	}

	formatter := FormatLogin(loggedInUser, token)

	ctx.JSON(http.StatusOK, &utils.Response{
		Success: true,
		Message: "User logged in successfully",
		Data:    formatter,
	})

}
