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

// EntertainmentServiceController -> EntertainmentServiceController
type EntertainmentServiceController struct {
	service services.EntertainmentServiceService
}

// NewEntertainmentServiceController : NewEntertainmentServiceController
func NewEntertainmentServiceController(service services.EntertainmentServiceService) *EntertainmentServiceController {
	return &EntertainmentServiceController{service}
}

// GetEntertainmentServices : GetEntertainmentServices controller
func (h *EntertainmentServiceController) GetEntertainmentServices(ctx *gin.Context) {
	var entertainment_services models.EntertainmentService

	search := ctx.Query("search")

	entertainmentServices, _, err := h.service.FindAll(entertainment_services, search)

	if err != nil {
		response := utils.APIResponse("Failed to find entertainment service", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Entertainment service result set", http.StatusOK, "success", formatters.FormatEntertainmentServices(entertainmentServices))
	ctx.JSON(http.StatusOK, response)
}

// GetEntertainmentService : get entertainment service by id
func (h *EntertainmentServiceController) GetEntertainmentService(c *gin.Context) {
	var input inputs.GetEntertainmentServiceDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of entertainment service", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	entertainmentService, err := h.service.Find(input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of entertainment service", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Entertainment service detail", http.StatusOK, "success", formatters.FormatEntertainmentService(entertainmentService))
	c.JSON(http.StatusOK, response)

}

// AddEntertainmentService : AddEntertainmentService controller
func (h *EntertainmentServiceController) AddEntertainmentService(ctx *gin.Context) {
	var input inputs.EntertainmentServiceInput
	customizer := g.Validator(inputs.EntertainmentServiceInput{})

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

		response := utils.APIResponse("Failed to store entertainment service", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newEntertainmentService, err := h.service.Save(input)
	if err != nil {
		response := utils.APIResponse("Failed to store entertainment service", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to store entertainment service", http.StatusCreated, "success", formatters.FormatEntertainmentService(newEntertainmentService))
	ctx.JSON(http.StatusCreated, response)
}

// UpdateEntertainmentService : get update by id
func (h *EntertainmentServiceController) UpdateEntertainmentService(ctx *gin.Context) {
	var inputID inputs.GetEntertainmentServiceDetailInput
	customizer := g.Validator(inputs.EntertainmentServiceInput{})

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to update entertainment service", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData inputs.EntertainmentServiceInput

	err = ctx.ShouldBindJSON(&inputData)
	if err != nil {
		// errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": customizer.DecryptErrors(err)}

		response := utils.APIResponse("Failed to update entertainment service", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedEntertainmentService, err := h.service.Update(inputID, inputData)
	if err != nil {
		response := utils.APIResponse("Failed to update entertainment service", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to update entertainment service", http.StatusOK, "success", formatters.FormatEntertainmentService(updatedEntertainmentService))
	ctx.JSON(http.StatusOK, response)
}

// DeleteEntertainmentService : Deletes Entertainment service
func (h *EntertainmentServiceController) DeleteEntertainmentService(ctx *gin.Context) {
	var inputID inputs.GetEntertainmentServiceDetailInput

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete entertainment service", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	deletedEntertainmentService, err := h.service.Delete(inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete entertainment service", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to delete entertainment service", http.StatusOK, "success", formatters.FormatEntertainmentService(deletedEntertainmentService))
	ctx.JSON(http.StatusOK, response)
}
