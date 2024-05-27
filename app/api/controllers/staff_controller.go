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

// StaffController -> StaffController
type StaffController struct {
	service services.StaffService
}

// NewStaffController : NewStaffController
func NewStaffController(service services.StaffService) *StaffController {
	return &StaffController{service}
}

// GetStaffs : GetStaffs controller
func (h *StaffController) GetStaffs(ctx *gin.Context) {
	var staffs models.Staff

	search := ctx.Query("search")

	getStaffs, _, err := h.service.FindAll(staffs, search)

	if err != nil {
		response := utils.APIResponse("Failed to find staff", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("staff result set", http.StatusOK, "success", formatters.FormatStaffs(getStaffs))
	ctx.JSON(http.StatusOK, response)
}

// GetStaff : get staff by id
func (h *StaffController) GetStaff(c *gin.Context) {
	var input inputs.GetStaffDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of staff", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	staff, err := h.service.Find(input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of staff", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("staff detail", http.StatusOK, "success", formatters.FormatStaff(staff))
	c.JSON(http.StatusOK, response)

}

// AddStaff : AddStaff controller
func (h *StaffController) AddStaff(ctx *gin.Context) {
	var input inputs.StaffInput

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

		response := utils.APIResponse("Failed to store staff", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newStaff, err := h.service.Save(input)
	if err != nil {
		response := utils.APIResponse("Failed to store staff", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to store staff", http.StatusOK, "success", formatters.FormatStaff(newStaff))
	ctx.JSON(http.StatusOK, response)
}

// UpdateStaff : get update by id
func (h *StaffController) UpdateStaff(ctx *gin.Context) {
	var inputID inputs.GetStaffDetailInput

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to update staff", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData inputs.StaffInput

	err = ctx.ShouldBindJSON(&inputData)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := utils.APIResponse("Failed to update staff", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedStaff, err := h.service.Update(inputID, inputData)
	if err != nil {
		response := utils.APIResponse("Failed to update staff", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to update staff", http.StatusOK, "success", formatters.FormatStaff(updatedStaff))
	ctx.JSON(http.StatusOK, response)
}

// DeleteStaff : Deletes staff
func (h *StaffController) DeleteStaff(ctx *gin.Context) {
	var inputID inputs.GetStaffDetailInput

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete staff", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	deletedStaff, err := h.service.Delete(inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete staff", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to delete staff", http.StatusOK, "success", formatters.FormatStaff(deletedStaff))
	ctx.JSON(http.StatusOK, response)
}
