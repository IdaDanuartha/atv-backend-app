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

// FacilityController -> FacilityController
type FacilityController struct {
	service services.FacilityService
}

// NewFacilityController : NewFacilityController
func NewFacilityController(service services.FacilityService) *FacilityController {
	return &FacilityController{service}
}

// GetFacilities : GetFacilities controller
func (h *FacilityController) GetFacilities(ctx *gin.Context) {
	var facilities models.Facility

	search := ctx.Query("search")

	getFacilities, _, err := h.service.FindAll(facilities, search)

	if err != nil {
		response := utils.APIResponse("Failed to find facility", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("facilities result set", http.StatusOK, "success", formatters.FormatFacilities(getFacilities))
	ctx.JSON(http.StatusOK, response)
}

// GetFacility : get facility by id
func (h *FacilityController) GetFacility(c *gin.Context) {
	var input inputs.GetFacilityDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of facility", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	Facility, err := h.service.Find(input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of facility", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("facility detail", http.StatusOK, "success", formatters.FormatFacility(Facility))
	c.JSON(http.StatusOK, response)

}

// AddFacility : AddFacility controller
func (h *FacilityController) AddFacility(ctx *gin.Context) {
	var input inputs.FacilityInput

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
		errorMessage := gin.H{"errors": utils.Customizer.DecryptErrors(err)}

		response := utils.APIResponse("Failed to store facility", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newFacility, err := h.service.Save(input)
	if err != nil {
		response := utils.APIResponse("Failed to store facility", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to store facility", http.StatusOK, "success", formatters.FormatFacility(newFacility))
	ctx.JSON(http.StatusOK, response)
}

// UpdateFacility : get update by id
func (h *FacilityController) UpdateFacility(ctx *gin.Context) {
	var inputID inputs.GetFacilityDetailInput

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to update facility", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData inputs.FacilityInput

	err = ctx.ShouldBindJSON(&inputData)
	if err != nil {
		// errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": utils.Customizer.DecryptErrors(err)}

		response := utils.APIResponse("Failed to update facility", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedFacility, err := h.service.Update(inputID, inputData)
	if err != nil {
		response := utils.APIResponse("Failed to update facility", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to update facility", http.StatusOK, "success", formatters.FormatFacility(updatedFacility))
	ctx.JSON(http.StatusOK, response)
}

// DeleteFacility : Deletes facility
func (h *FacilityController) DeleteFacility(ctx *gin.Context) {
	var inputID inputs.GetFacilityDetailInput

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete facility", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	deletedFacility, err := h.service.Delete(inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete facility", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to delete facility", http.StatusOK, "success", formatters.FormatFacility(deletedFacility))
	ctx.JSON(http.StatusOK, response)
}
