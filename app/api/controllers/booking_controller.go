package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/IdaDanuartha/atv-backend-app/app/api/formatters"
	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/services"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"github.com/IdaDanuartha/atv-backend-app/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
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
	currentPage, err := strconv.Atoi(ctx.Query("current_page"))
	if err != nil {
		currentPage = 1
	}

	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil {
		pageSize = 0
	}

	instructorID := ctx.Query("instructor_id")
	customerID := ctx.Query("customer_id")

	booking, total, _, err := h.service.FindAll(bookings, search, currentPage, pageSize, instructorID, customerID)

	if err != nil {
		response := utils.APIResponse("Failed to find booking", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if pageSize > 0 {
		response := utils.APIResponseWithPagination("Booking result set", http.StatusOK, "success", total, currentPage, pageSize, formatters.FormatBookings(booking))
		ctx.JSON(http.StatusOK, response)
	} else {
		response := utils.APIResponse("Booking result set", http.StatusOK, "success", formatters.FormatBookings(booking))
		ctx.JSON(http.StatusOK, response)
	}
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

func (h *BookingController) ExportToExcel(ctx *gin.Context) {
	var booking models.Booking

	bookings, _, _, err := h.service.FindAll(booking, "", 1, 0, "", "")
	if err != nil {
		response := utils.APIResponse("Failed to find booking", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Create a new Excel file
	file := excelize.NewFile()
	sheetName := "Sheet1"

	// Set the header row
	headers := []string{"ID", "Code", "Name", "Phone Number", "Total Price", "Total Pay", "Total Change", "Payment Method"}
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		file.SetCellValue(sheetName, cell, header)
	}

	// Fill the Excel file with data
	for i, booking := range bookings {
		file.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), i+1)
		file.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), booking.Code)
		file.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), booking.Name)
		file.SetCellValue(sheetName, fmt.Sprintf("D%d", i+2), booking.PhoneNumber)
		file.SetCellValue(sheetName, fmt.Sprintf("E%d", i+2), utils.FormatRupiah(int64(booking.TotalPrice)))
		file.SetCellValue(sheetName, fmt.Sprintf("F%d", i+2), utils.FormatRupiah(int64(booking.TotalPay)))
		file.SetCellValue(sheetName, fmt.Sprintf("G%d", i+2), utils.FormatRupiah(int64(booking.TotalChange)))
		file.SetCellValue(sheetName, fmt.Sprintf("H%d", i+2), booking.PaymentMethod)
	}

	// Write the file to a temporary buffer
	buffer, err := file.WriteToBuffer()
	if err != nil {
		response := utils.APIResponse("Could not create the Excel file", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Set the appropriate headers and return the file
	ctx.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Header("Content-Disposition", "attachment; filename=bookings.xlsx")

	ctx.DataFromReader(http.StatusOK, int64(buffer.Len()), "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buffer, nil)
}