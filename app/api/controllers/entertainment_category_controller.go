package controllers

import (
	"net/http"
	"strconv"

	"github.com/IdaDanuartha/atv-backend-app/app/api/formatters"
	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/services"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"github.com/IdaDanuartha/atv-backend-app/app/utils"
	"github.com/gin-gonic/gin"
)

// EntertainmentCategoryController -> EntertainmentCategoryController
type EntertainmentCategoryController struct {
	service services.EntertainmentCategoryService
}

// NewEntertainmentCategoryController : NewEntertainmentCategoryController
func NewEntertainmentCategoryController(service services.EntertainmentCategoryService) *EntertainmentCategoryController {
	return &EntertainmentCategoryController{service}
}

// GetEntertainmentCategories : GetEntertainmentCategories controller
func (h *EntertainmentCategoryController) GetEntertainmentCategories(ctx *gin.Context) {
	var entertainment_categories models.EntertainmentCategory

	search := ctx.Query("search")
	currentPage, err := strconv.Atoi(ctx.Query("current_page"))
	if err != nil {
		currentPage = 1
	}

	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil {
		pageSize = 0
	}

	entertainmentCategories, total, _, err := h.service.FindAll(entertainment_categories, search, currentPage, pageSize)

	if err != nil {
		response := utils.APIResponse("Failed to find entertainment category", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	if pageSize > 0 {
		response := utils.APIResponseWithPagination("Entertainment categories result set", http.StatusOK, "success", total, currentPage, pageSize, formatters.FormatEntertainmentCategories(entertainmentCategories))
		ctx.JSON(http.StatusOK, response)
	} else {
		response := utils.APIResponse("Entertainment categories result set", http.StatusOK, "success", formatters.FormatEntertainmentCategories(entertainmentCategories))
		ctx.JSON(http.StatusOK, response)
	}
}

// GetEntertainmentCategory : get entertainment category by id
func (h *EntertainmentCategoryController) GetEntertainmentCategory(c *gin.Context) {
	var input inputs.GetEntertainmentCategoryDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of entertainment category", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	entertainmentCategory, err := h.service.Find(input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of entertainment category", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Entertainment category detail", http.StatusOK, "success", formatters.FormatEntertainmentCategory(entertainmentCategory))
	c.JSON(http.StatusOK, response)

}

// AddEntertainmentCategory : AddEntertainmentCategory controller
func (h *EntertainmentCategoryController) AddEntertainmentCategory(ctx *gin.Context) {
	var input inputs.EntertainmentCategoryInput
	customizer := g.Validator(inputs.EntertainmentCategoryInput{})

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

		response := utils.APIResponse("Failed to store entertainment category", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newEntertainmentCategory, err := h.service.Save(input)
	if err != nil {
		response := utils.APIResponse("Failed to store entertainment category", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to store entertainment category", http.StatusCreated, "success", formatters.FormatEntertainmentCategory(newEntertainmentCategory))
	ctx.JSON(http.StatusCreated, response)
}

// UpdateEntertainmentCategory : get update by id
func (h *EntertainmentCategoryController) UpdateEntertainmentCategory(ctx *gin.Context) {
	var inputID inputs.GetEntertainmentCategoryDetailInput
	customizer := g.Validator(inputs.EntertainmentCategoryInput{})

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to update entertainment category", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData inputs.EntertainmentCategoryInput

	err = ctx.ShouldBindJSON(&inputData)
	if err != nil {
		// errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": customizer.DecryptErrors(err)}

		response := utils.APIResponse("Failed to update entertainment category", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedEntertainmentCategory, err := h.service.Update(inputID, inputData)
	if err != nil {
		response := utils.APIResponse("Failed to update entertainment category", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to update entertainment category", http.StatusOK, "success", formatters.FormatEntertainmentCategory(updatedEntertainmentCategory))
	ctx.JSON(http.StatusOK, response)
}

// DeleteEntertainmentCategory : Deletes Entertainment Category
func (h *EntertainmentCategoryController) DeleteEntertainmentCategory(ctx *gin.Context) {
	var inputID inputs.GetEntertainmentCategoryDetailInput

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete entertainment category", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	deletedEntertainmentCategory, err := h.service.Delete(inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete entertainment category", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to delete entertainment category", http.StatusOK, "success", formatters.FormatEntertainmentCategory(deletedEntertainmentCategory))
	ctx.JSON(http.StatusOK, response)
}
