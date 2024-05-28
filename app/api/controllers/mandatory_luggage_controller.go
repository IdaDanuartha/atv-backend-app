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

// MandatoryLuggageController -> MandatoryLuggageController
type MandatoryLuggageController struct {
	service services.MandatoryLuggageService
}

// NewMandatoryLuggageController : NewMandatoryLuggageController
func NewMandatoryLuggageController(service services.MandatoryLuggageService) *MandatoryLuggageController {
	return &MandatoryLuggageController{service}
}

// GetMandatoryLuggages : GetMandatoryLuggages controller
func (h *MandatoryLuggageController) GetMandatoryLuggages(ctx *gin.Context) {
	var mandatory_luggages models.MandatoryLuggage

	search := ctx.Query("search")

	mandatoryLuggages, _, err := h.service.FindAll(mandatory_luggages, search)

	if err != nil {
		response := utils.APIResponse("Failed to find mandatory luggage", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Mandatory luggage result set", http.StatusOK, "success", formatters.FormatMandatoryLuggages(mandatoryLuggages))
	ctx.JSON(http.StatusOK, response)
}

// GetMandatoryLuggage : get mandatory luggage by id
func (h *MandatoryLuggageController) GetMandatoryLuggage(c *gin.Context) {
	var input inputs.GetMandatoryLuggageDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of mandatory luggage", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	mandatoryLuggage, err := h.service.Find(input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of mandatory luggage", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("mandatory luggage detail", http.StatusOK, "success", formatters.FormatMandatoryLuggage(mandatoryLuggage))
	c.JSON(http.StatusOK, response)

}

// AddMandatoryLuggage : AddMandatoryLuggage controller
func (h *MandatoryLuggageController) AddMandatoryLuggage(ctx *gin.Context) {
	var input inputs.MandatoryLuggageInput

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

		response := utils.APIResponse("Failed to store mandatory luggage", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newMandatoryLuggage, err := h.service.Save(input)
	if err != nil {
		response := utils.APIResponse("Failed to store mandatory luggage", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to store mandatory luggage", http.StatusOK, "success", formatters.FormatMandatoryLuggage(newMandatoryLuggage))
	ctx.JSON(http.StatusOK, response)
}

// UpdateMandatoryLuggage : get update by id
func (h *MandatoryLuggageController) UpdateMandatoryLuggage(ctx *gin.Context) {
	var inputID inputs.GetMandatoryLuggageDetailInput

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to update mandatory luggage", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData inputs.MandatoryLuggageInput

	err = ctx.ShouldBindJSON(&inputData)
	if err != nil {
		// errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": utils.Customizer.DecryptErrors(err)}

		response := utils.APIResponse("Failed to update mandatory luggage", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedMandatoryLuggage, err := h.service.Update(inputID, inputData)
	if err != nil {
		response := utils.APIResponse("Failed to update mandatory luggage", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to update mandatory luggage", http.StatusOK, "success", formatters.FormatMandatoryLuggage(updatedMandatoryLuggage))
	ctx.JSON(http.StatusOK, response)
}

// DeleteMandatoryLuggage : Deletes mandatory luggage
func (h *MandatoryLuggageController) DeleteMandatoryLuggage(ctx *gin.Context) {
	var inputID inputs.GetMandatoryLuggageDetailInput

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete mandatory luggage", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	deletedMandatoryLuggage, err := h.service.Delete(inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete mandatory luggage", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to delete mandatory luggage", http.StatusOK, "success", formatters.FormatMandatoryLuggage(deletedMandatoryLuggage))
	ctx.JSON(http.StatusOK, response)
}
