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

// InstructorController -> InstructorController
type InstructorController struct {
	service services.InstructorService
}

// NewInstructorController : NewInstructorController
func NewInstructorController(service services.InstructorService) *InstructorController {
	return &InstructorController{service}
}

// GetInstructors : GetInstructors controller
func (h *InstructorController) GetInstructors(ctx *gin.Context) {
	var instructors models.Instructor

	search := ctx.Query("search")

	getInstructors, _, err := h.service.FindAll(instructors, search)

	if err != nil {
		response := utils.APIResponse("Failed to find instructor", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("instructor result set", http.StatusOK, "success", formatters.FormatInstructors(getInstructors))
	ctx.JSON(http.StatusOK, response)
}

// GetInstructor : get instructor by id
func (h *InstructorController) GetInstructor(c *gin.Context) {
	var input inputs.GetInstructorDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of instructor", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	instructor, err := h.service.Find(input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of instructor", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("instructor detail", http.StatusOK, "success", formatters.FormatInstructor(instructor))
	c.JSON(http.StatusOK, response)

}

// AddInstructor : AddInstructor controller
func (h *InstructorController) AddInstructor(ctx *gin.Context) {
	var input inputs.InstructorInput

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

		response := utils.APIResponse("Failed to store instructor", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newInstructor, err := h.service.Save(input)
	if err != nil {
		response := utils.APIResponse("Failed to store instructor", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to store instructor", http.StatusOK, "success", formatters.FormatInstructor(newInstructor))
	ctx.JSON(http.StatusOK, response)
}

// UpdateInstructor : get update by id
func (h *InstructorController) UpdateInstructor(ctx *gin.Context) {
	var inputID inputs.GetInstructorDetailInput

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to update instructor", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData inputs.InstructorInput

	err = ctx.ShouldBindJSON(&inputData)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := utils.APIResponse("Failed to update instructor", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedInstructor, err := h.service.Update(inputID, inputData)
	if err != nil {
		response := utils.APIResponse("Failed to update instructor", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to update instructor", http.StatusOK, "success", formatters.FormatInstructor(updatedInstructor))
	ctx.JSON(http.StatusOK, response)
}

// DeleteInstructor : Deletes instructor
func (h *InstructorController) DeleteInstructor(ctx *gin.Context) {
	var inputID inputs.GetInstructorDetailInput

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete instructor", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	deletedInstructor, err := h.service.Delete(inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete instructor", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to delete instructor", http.StatusOK, "success", formatters.FormatInstructor(deletedInstructor))
	ctx.JSON(http.StatusOK, response)
}
