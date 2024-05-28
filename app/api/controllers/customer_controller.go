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

// CustomerController -> CustomerController
type CustomerController struct {
	service services.CustomerService
}

// NewCustomerController : NewCustomerController
func NewCustomerController(service services.CustomerService) *CustomerController {
	return &CustomerController{service}
}

// GetCustomers : GetCustomers controller
func (h *CustomerController) GetCustomers(ctx *gin.Context) {
	var customers models.Customer

	search := ctx.Query("search")

	getCustomer, _, err := h.service.FindAll(customers, search)

	if err != nil {
		response := utils.APIResponse("Failed to find customer", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("customer result set", http.StatusOK, "success", formatters.FormatCustomers(getCustomer))
	ctx.JSON(http.StatusOK, response)
}

// GetCustomer : get customer by id
func (h *CustomerController) GetCustomer(c *gin.Context) {
	var input inputs.GetCustomerDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of customer", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	customer, err := h.service.Find(input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of customer", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("customer detail", http.StatusOK, "success", formatters.FormatCustomer(customer))
	c.JSON(http.StatusOK, response)

}