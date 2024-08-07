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

// RouteController -> RouteController
type RouteController struct {
	service services.RouteService
}

// NewRouteController : NewRouteController
func NewRouteController(service services.RouteService) *RouteController {
	return &RouteController{service}
}

// GetRoutes : GetRoutes controller
func (h *RouteController) GetRoutes(ctx *gin.Context) {
	var routes models.Route

	search := ctx.Query("search")
	currentPage, err := strconv.Atoi(ctx.Query("current_page"))
	if err != nil {
		currentPage = 1
	}

	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil {
		pageSize = 0
	}

	getRoutes, total, _, err := h.service.FindAll(routes, search, currentPage, pageSize)

	if err != nil {
		response := utils.APIResponse("Failed to find route", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if pageSize > 0 {
		response := utils.APIResponseWithPagination("Routes result set", http.StatusOK, "success", total, currentPage, pageSize, formatters.FormatRoutes(getRoutes))
		ctx.JSON(http.StatusOK, response)
	} else {
		response := utils.APIResponse("Routes result set", http.StatusOK, "success", formatters.FormatRoutes(getRoutes))
		ctx.JSON(http.StatusOK, response)
	}
}

// GetRoute : get facility by id
func (h *RouteController) GetRoute(c *gin.Context) {
	var input inputs.GetRouteDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of route", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	route, err := h.service.Find(input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of route", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("route detail", http.StatusOK, "success", formatters.FormatRoute(route))
	c.JSON(http.StatusOK, response)

}

// AddRoute : AddRoute controller
func (h *RouteController) AddRoute(ctx *gin.Context) {
	var input inputs.RouteInput
	customizer := g.Validator(inputs.RouteInput{})

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

		response := utils.APIResponse("Failed to store route", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newRoute, err := h.service.Save(input)
	if err != nil {
		response := utils.APIResponse("Failed to store route", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to store route", http.StatusCreated, "success", formatters.FormatRoute(newRoute))
	ctx.JSON(http.StatusCreated, response)
}

// UpdateRoute : get update by id
func (h *RouteController) UpdateRoute(ctx *gin.Context) {
	var inputID inputs.GetRouteDetailInput
	customizer := g.Validator(inputs.RouteInput{})

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to update route", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData inputs.RouteInput

	err = ctx.ShouldBindJSON(&inputData)
	if err != nil {
		// errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": customizer.DecryptErrors(err)}

		response := utils.APIResponse("Failed to update route", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedRoute, err := h.service.Update(inputID, inputData)
	if err != nil {
		response := utils.APIResponse("Failed to update route", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to update route", http.StatusOK, "success", formatters.FormatRoute(updatedRoute))
	ctx.JSON(http.StatusOK, response)
}

// DeleteRoute : Deletes route
func (h *RouteController) DeleteRoute(ctx *gin.Context) {
	var inputID inputs.GetRouteDetailInput

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete route", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	deletedRoute, err := h.service.Delete(inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete route", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to delete route", http.StatusOK, "success", formatters.FormatRoute(deletedRoute))
	ctx.JSON(http.StatusOK, response)
}
