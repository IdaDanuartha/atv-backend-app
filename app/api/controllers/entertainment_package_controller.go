package controllers

import (
	"net/http"

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

	entertainmentPackages, _, err := h.service.FindAll(entertainment_packages, search)

	if err != nil {
		response := utils.APIResponse("Failed to find entertainment package", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Entertainment package result set", http.StatusOK, "success", formatters.FormatEntertainmentPackages(entertainmentPackages))
	ctx.JSON(http.StatusOK, response)
}

// GetEntertainmentPackage : get entertainment package by id
func (h *EntertainmentPackageController) GetEntertainmentPackage(c *gin.Context) {
	var input inputs.GetEntertainmentPackageDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of entertainment package", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	entertainmentPackage, err := h.service.Find(input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of entertainment package", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Entertainment package detail", http.StatusOK, "success", formatters.FormatEntertainmentPackage(entertainmentPackage))
	c.JSON(http.StatusOK, response)

}

// AddEntertainmentPackage : AddEntertainmentPackage controller
func (h *EntertainmentPackageController) AddEntertainmentPackage(ctx *gin.Context) {
	var input inputs.EntertainmentPackageInput

	// Check if request body is empty or has no content type
	if ctx.Request.Body == nil || ctx.Request.ContentLength == 0 || ctx.GetHeader("Content-Type") == "" {
		errorMessage := gin.H{"errors": "No fields sent"}
		response := utils.APIResponse("No fields sent", http.StatusBadRequest, "error", errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := utils.APIResponse("Failed to store entertainment package", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newEntertainmentPackage, err := h.service.Save(input)
	if err != nil {
		response := utils.APIResponse("Failed to store entertainment package", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to store entertainment package", http.StatusOK, "success", formatters.FormatEntertainmentPackage(newEntertainmentPackage))
	ctx.JSON(http.StatusOK, response)
}

// UpdateEntertainmentPackage : get update by id
func (h *EntertainmentPackageController) UpdateEntertainmentPackage(ctx *gin.Context) {
	var inputID inputs.GetEntertainmentPackageDetailInput

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to update entertainment package", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData inputs.EntertainmentPackageInput

	err = ctx.ShouldBindJSON(&inputData)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := utils.APIResponse("Failed to update entertainment package", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedEntertainmentPackage, err := h.service.Update(inputID, inputData)
	if err != nil {
		response := utils.APIResponse("Failed to update entertainment package", http.StatusBadRequest, "error", nil)
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
		response := utils.APIResponse("Failed to delete entertainment package", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	deletedEntertainmentPackage, err := h.service.Delete(inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete entertainment package", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to delete entertainment package", http.StatusOK, "success", formatters.FormatEntertainmentPackage(deletedEntertainmentPackage))
	ctx.JSON(http.StatusOK, response)
}
