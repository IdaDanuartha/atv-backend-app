package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/api/formatters"
	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/services"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"github.com/IdaDanuartha/atv-backend-app/app/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
	authService services.AuthService
}

func NewUserController(userService services.UserService, authService services.AuthService) *UserController {
	return &UserController{userService, authService}
}

func (h *UserController) RegisterUser(c *gin.Context) {
	var input inputs.RegisterInput

	// Check if request body is empty or has no content type
	if c.Request.Body == nil || c.Request.ContentLength == 0 || c.GetHeader("Content-Type") == "" {
		errorMessage := gin.H{"errors": "No fields sent"}
		response := utils.APIResponse("No fields sent", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		// errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": utils.Customizer.DecryptErrors(err)}

		response := utils.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := utils.APIResponse("Register account failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := utils.APIResponse("Register account failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := formatters.FormatAuthCustomer(newUser, token)

	response := utils.APIResponse("Account registered successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *UserController) Login(c *gin.Context) {
	var input inputs.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		// errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": utils.Customizer.DecryptErrors(err)}

		response := utils.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := utils.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := utils.APIResponse("Login failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := formatters.FormatAuth(loggedinUser, token)

	response := utils.APIResponse("Successfuly loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *UserController) FetchUser(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(models.User)

	formatter := formatters.FormatAuth(currentUser, "")

	response := utils.APIResponse("Successfuly fetch user data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *UserController) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := utils.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(models.User)
	userID := currentUser.ID

	// Check if the old avatar image exists for the user
	_, err = os.Stat(*currentUser.ProfilePath)
	if err == nil {
		// If the old avatar image exists, delete it
		err := os.Remove(*currentUser.ProfilePath)
		if err != nil {
			data := gin.H{"is_uploaded": false}
			response := utils.APIResponse("Failed to delete old avatar image", http.StatusBadRequest, "error", data)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	path := fmt.Sprintf("uploads/users/%s-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := utils.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := utils.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := utils.APIResponse("Avatar successfuly uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

func (h *UserController) UpdateProfile(ctx *gin.Context) {
	var inputData inputs.UpdateProfileInput

	err := ctx.ShouldBindJSON(&inputData)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := utils.APIResponse("Failed to update your profile", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Check if avatar file is uploaded
	// file, _ := ctx.FormFile("avatar")
	// if file != nil {
	// 	// If avatar file is uploaded, call UploadAvatar function
	// 	h.UploadAvatar(ctx)
	// 	return
	// }

	updatedUser, err := h.userService.UpdateUser(inputData, ctx)
	if err != nil {
		response := utils.APIResponse("Failed to update your profile", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to update your profile", http.StatusOK, "success", formatters.FormatUser(updatedUser))
	ctx.JSON(http.StatusOK, response)
}
