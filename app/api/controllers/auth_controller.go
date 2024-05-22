package controllers

import (
	"net/http"

	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/services"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"github.com/IdaDanuartha/atv-backend-app/utils"
	"github.com/gin-gonic/gin"
)

// AuthController -> AuthController
type AuthController struct {
	service services.AuthService
}

type UserFormatter struct {
	ID         	string    `json:"id"`
	Name       	string `json:"name"`
	Token      	string `json:"token"`
	User       	models.User `json:"user"`
}

func FormatUser(customer models.Customer, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         customer.ID,
		Name: 		customer.Name,
		Token:      token,
		User: 		customer.User,
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

	formatter := FormatUser(newUser, token)

	ctx.JSON(http.StatusOK, &utils.Response{
		Success: true,
		Message: "User registered successfully",
		Data:    formatter,
	})
}

// func (h *userHandler) Login(c *gin.Context) {
// 	var input user.LoginInput

// 	err := c.ShouldBindJSON(&input)
// 	if err != nil {
// 		errors := helper.FormatValidationError(err)
// 		errorMessage := gin.H{"errors": errors}

// 		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	loggedinUser, err := h.userService.Login(input)

// 	if err != nil {
// 		errorMessage := gin.H{"errors": err.Error()}

// 		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	token, err := h.authService.GenerateToken(loggedinUser.ID)
// 	if err != nil {
// 		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	formatter := user.FormatUser(loggedinUser, token)

// 	response := helper.APIResponse("Successfuly loggedin", http.StatusOK, "success", formatter)

// 	c.JSON(http.StatusOK, response)

// }
