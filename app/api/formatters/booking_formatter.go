package formatters

import (
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

func FormatBooking(booking models.Booking) models.Booking {
	bookingFormatter := models.Booking{}
	bookingFormatter.ID = booking.ID
	bookingFormatter.CustomerID = booking.CustomerID
	bookingFormatter.Name = booking.Name
	bookingFormatter.PhoneNumber = booking.PhoneNumber
	bookingFormatter.PaymentMethod = booking.PaymentMethod
	bookingFormatter.Date = booking.Date
	bookingFormatter.TotalPrice = booking.TotalPrice
	bookingFormatter.TotalPay = booking.TotalPay
	bookingFormatter.TotalChange = booking.TotalChange
	bookingFormatter.CreatedAt = booking.CreatedAt
	bookingFormatter.UpdatedAt = booking.UpdatedAt
	bookingFormatter.DeletedAt = booking.DeletedAt

	bookingDetails := make([]models.BookingDetail, 0)

	for _, packageDetail := range booking.BookingDetails {
		newBookingDetail := models.BookingDetail{}

		newBookingDetail.ID = packageDetail.ID
		newBookingDetail.Price = packageDetail.Price
		newBookingDetail.Qty = packageDetail.Qty
		newBookingDetail.CreatedAt = packageDetail.CreatedAt
		newBookingDetail.UpdatedAt = packageDetail.UpdatedAt
		newBookingDetail.DeletedAt = packageDetail.DeletedAt

		newBookingDetail.EntertainmentService.ID = packageDetail.EntertainmentService.ID
		newBookingDetail.EntertainmentService.Name = packageDetail.EntertainmentService.Name
		newBookingDetail.EntertainmentService.Price = packageDetail.EntertainmentService.Price
		newBookingDetail.EntertainmentService.CreatedAt = packageDetail.EntertainmentService.CreatedAt
		newBookingDetail.EntertainmentService.UpdatedAt = packageDetail.EntertainmentService.UpdatedAt
		newBookingDetail.EntertainmentService.DeletedAt = packageDetail.EntertainmentService.DeletedAt

		bookingDetails = append(bookingDetails, newBookingDetail)
	}
	bookingFormatter.BookingDetails = bookingDetails

	return bookingFormatter
}

func FormatBookings(bookings []models.Booking) []models.Booking {
	bookingsFormatter := []models.Booking{}

	for _, booking := range bookings {
		booking := FormatBooking(booking)
		bookingsFormatter = append(bookingsFormatter, booking)
	}

	return bookingsFormatter
}
