package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/IdaDanuartha/atv-backend-app/app/api/formatters"
	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/services"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"github.com/IdaDanuartha/atv-backend-app/app/utils"
	"github.com/gin-gonic/gin"
)

// EntertainmentPackageController -> EntertainmentPackageController
type EntertainmentPackageController struct {
	service services.EntertainmentPackageService
}

// NewEntertainmentPackageController : NewEntertainmentPackageController
func NewEntertainmentPackageController(service services.EntertainmentPackageService) *EntertainmentPackageController {
	return &EntertainmentPackageController{service}
}

// GetEntertainmentPackages : GetEntertainmentPackages controller
func (h *EntertainmentPackageController) GetEntertainmentPackages(ctx *gin.Context) {
	var entertainment_packages models.EntertainmentPackage

	search := ctx.Query("search")
	currentPage, err := strconv.Atoi(ctx.Query("current_page"))
	if err != nil {
		currentPage = 1
	}

	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil {
		pageSize = 0
	}

	entertainmentPackages, total, _, err := h.service.FindAll(entertainment_packages, search, currentPage, pageSize)

	if err != nil {
		response := utils.APIResponse("Failed to find entertainment package", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if pageSize > 0 {
		response := utils.APIResponseWithPagination("Entertainment packages result set", http.StatusOK, "success", total, currentPage, pageSize, formatters.FormatEntertainmentPackages(entertainmentPackages))
		ctx.JSON(http.StatusOK, response)
	} else {
		response := utils.APIResponse("Entertainment packages result set", http.StatusOK, "success", formatters.FormatEntertainmentPackages(entertainmentPackages))
		ctx.JSON(http.StatusOK, response)
	}

}

// GetEntertainmentPackage : get entertainment package by id
func (h *EntertainmentPackageController) GetEntertainmentPackage(c *gin.Context) {
	var input inputs.GetEntertainmentPackageDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of entertainment package", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	entertainmentPackage, err := h.service.Find(input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of entertainment package", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Entertainment package detail", http.StatusOK, "success", formatters.FormatEntertainmentPackage(entertainmentPackage))
	c.JSON(http.StatusOK, response)

}

func (h *EntertainmentPackageController) UploadImage(ctx *gin.Context) {
	var inputID inputs.GetEntertainmentPackageDetailInput

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to get entertainment package id", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		data := gin.H{"message": err.Error()}
		response := utils.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	entertainmentPackage, err := h.service.Find(inputID)
	if err != nil {
		data := gin.H{"message": err.Error()}
		response := utils.APIResponse("Failed to get entertainment package id", http.StatusBadRequest, "error", data)

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	ID := entertainmentPackage.ID

	// Check if ImagePath is not nil before proceeding
	if entertainmentPackage.ImagePath != nil {
		// Check if the old avatar image exists for the user
		_, err := os.Stat(*entertainmentPackage.ImagePath)
		if err == nil {
			// If the old avatar image exists, delete it
			err := os.Remove(*entertainmentPackage.ImagePath)
			if err != nil {
				data := gin.H{
					"is_uploaded": false,
					"message":     err.Error(),
				}
				response := utils.APIResponse("Failed to delete old avatar image", http.StatusBadRequest, "error", data)
				ctx.JSON(http.StatusBadRequest, response)
				return
			}
		} else if !os.IsNotExist(err) {
			// Handle other possible errors from os.Stat
			data := gin.H{
				"is_uploaded": false,
				"message":     err.Error(),
			}
			response := utils.APIResponse("Error checking old avatar image", http.StatusInternalServerError, "error", data)
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}
	}

	path := fmt.Sprintf("uploads/entertainment_packages/%s-%s", ID, file.Filename)

	err = ctx.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := utils.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.SaveImage(ID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := utils.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := utils.APIResponse("Image successfuly uploaded", http.StatusOK, "success", data)

	ctx.JSON(http.StatusOK, response)
}

// AddEntertainmentPackage : AddEntertainmentPackage controller
func (h *EntertainmentPackageController) AddEntertainmentPackage(ctx *gin.Context) {
	var input inputs.EntertainmentPackageInput
	customizer := g.Validator(inputs.EntertainmentPackageInput{})

	// Check if request body is empty or has no content type
	if ctx.Request.Body == nil || ctx.Request.ContentLength == 0 || ctx.GetHeader("Content-Type") == "" {
		errorMessage := gin.H{"errors": "No fields sent"}
		response := utils.APIResponse("No fields sent", http.StatusBadRequest, "error", errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		// errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": customizer.DecryptErrors(err)}

		response := utils.APIResponse("Failed to store entertainment package", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newEntertainmentPackage, err := h.service.Save(input)
	if err != nil {
		response := utils.APIResponse("Failed to store entertainment package", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to store entertainment package", http.StatusCreated, "success", formatters.FormatEntertainmentPackage(newEntertainmentPackage))
	ctx.JSON(http.StatusCreated, response)
}

// UpdateEntertainmentPackage : get update by id
func (h *EntertainmentPackageController) UpdateEntertainmentPackage(ctx *gin.Context) {
	var inputID inputs.GetEntertainmentPackageDetailInput
	customizer := g.Validator(inputs.EntertainmentPackageInput{})

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to update entertainment package", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData inputs.EntertainmentPackageInput

	err = ctx.ShouldBindJSON(&inputData)
	if err != nil {
		// errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": customizer.DecryptErrors(err)}

		response := utils.APIResponse("Failed to update entertainment package", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedEntertainmentPackage, err := h.service.Update(inputID, inputData)
	if err != nil {
		response := utils.APIResponse("Failed to update entertainment package", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to update entertainment package", http.StatusOK, "success", formatters.FormatEntertainmentPackage(updatedEntertainmentPackage))
	ctx.JSON(http.StatusOK, response)
}

// DeleteEntertainmentPackage : Deletes Entertainment package
func (h *EntertainmentPackageController) DeleteEntertainmentPackage(ctx *gin.Context) {
	var inputID inputs.GetEntertainmentPackageDetailInput

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete entertainment package", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	deletedEntertainmentPackage, err := h.service.Delete(inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete entertainment package", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to delete entertainment package", http.StatusOK, "success", formatters.FormatEntertainmentPackage(deletedEntertainmentPackage))
	ctx.JSON(http.StatusOK, response)
}
