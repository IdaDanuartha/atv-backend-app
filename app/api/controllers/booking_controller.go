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

// BookingController -> BookingController
type BookingController struct {
	service services.BookingService
}

// NewBookingController : NewBookingController
func NewBookingController(service services.BookingService) *BookingController {
	return &BookingController{service}
}

// GetBookings : GetBookings controller
func (h *BookingController) GetBookings(ctx *gin.Context) {
	var bookings models.Booking

	search := ctx.Query("search")

	booking, _, err := h.service.FindAll(bookings, search)

	if err != nil {
		response := utils.APIResponse("Failed to find booking", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Booking result set", http.StatusOK, "success", formatters.FormatBookings(booking))
	ctx.JSON(http.StatusOK, response)
}

// GetBooking : get booking by id
func (h *BookingController) GetBooking(c *gin.Context) {
	var input inputs.GetBookingDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of booking", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	booking, err := h.service.Find(input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of booking", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("booking detail", http.StatusOK, "success", formatters.FormatBooking(booking))
	c.JSON(http.StatusOK, response)

}

// AddBooking : AddBooking controller
func (h *BookingController) AddBooking(ctx *gin.Context) {
	var input inputs.BookingInput
	customizer := g.Validator(inputs.BookingInput{})

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

		response := utils.APIResponse("Failed to store booking", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newBooking, err := h.service.Save(input)
	if err != nil {
		response := utils.APIResponse("Failed to store booking", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to store booking", http.StatusCreated, "success", formatters.FormatBooking(newBooking))
	ctx.JSON(http.StatusCreated, response)
}

// UpdateBooking : get update by id
func (h *BookingController) UpdateBooking(ctx *gin.Context) {
	var inputID inputs.GetBookingDetailInput
	customizer := g.Validator(inputs.BookingInput{})

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to update booking", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData inputs.BookingInput

	err = ctx.ShouldBindJSON(&inputData)
	if err != nil {
		// errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": customizer.DecryptErrors(err)}

		response := utils.APIResponse("Failed to update booking", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedBooking, err := h.service.Update(inputID, inputData)
	if err != nil {
		response := utils.APIResponse("Failed to update booking", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to update booking", http.StatusOK, "success", formatters.FormatBooking(updatedBooking))
	ctx.JSON(http.StatusOK, response)
}

// DeleteBooking : Deletes booking
func (h *BookingController) DeleteBooking(ctx *gin.Context) {
	var inputID inputs.GetBookingDetailInput

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete booking", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	deletedBooking, err := h.service.Delete(inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete booking", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to delete booking", http.StatusOK, "success", formatters.FormatBooking(deletedBooking))
	ctx.JSON(http.StatusOK, response)
}
