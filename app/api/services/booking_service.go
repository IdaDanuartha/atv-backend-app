package services

import (
	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"github.com/IdaDanuartha/atv-backend-app/app/utils"
)

type BookingService interface {
	FindAll(model models.Booking, search string, currentPage int, pageSize int) ([]models.Booking, int64, int, error)
	Find(input inputs.GetBookingDetailInput) (models.Booking, error)
	Save(input inputs.BookingInput) (models.Booking, error)
	Update(inputID inputs.GetBookingDetailInput, input inputs.BookingInput) (models.Booking, error)
	Delete(inputID inputs.GetBookingDetailInput) (models.Booking, error)
}

// BookingService BookingService struct
type bookingService struct {
	repository repositories.BookingRepository
}

// NewBookingService : returns the BookingService struct instance
func NewBookingService(repository repositories.BookingRepository) bookingService {
	return bookingService{repository}
}

// FindAll -> calls Booking repo find all method
func (s bookingService) FindAll(model models.Booking, search string, currentPage int, pageSize int) ([]models.Booking, int64, int, error) {
	bookings, total, currentPage, err := s.repository.FindAll(model, search, currentPage, pageSize)
	if err != nil {
		return bookings, total, currentPage, err
	}

	return bookings, total, currentPage, nil
}

// Find -> calls Booking repo find method
func (s bookingService) Find(input inputs.GetBookingDetailInput) (models.Booking, error) {
	booking, err := s.repository.Find(input.ID, true)

	if err != nil {
		return booking, err
	}

	return booking, nil
}

// Save -> calls Booking repository save method
func (s bookingService) Save(input inputs.BookingInput) (models.Booking, error) {
	booking := models.Booking{}

	booking.CustomerID = input.CustomerID
	booking.Name = input.Name
	booking.Code = utils.GenerateFormattedString()
	booking.PhoneNumber = input.PhoneNumber
	booking.PaymentMethod = input.PaymentMethod
	booking.Date = input.Date
	booking.TotalPrice = input.TotalPrice
	booking.TotalPay = input.TotalPay
	booking.TotalChange = input.TotalChange
	booking.Details = input.Details

	newBooking, err := s.repository.Save(booking)
	if err != nil {
		return newBooking, err
	}

	return newBooking, nil
}

// Update -> calls Booking repo update method
func (s bookingService) Update(inputID inputs.GetBookingDetailInput, input inputs.BookingInput) (models.Booking, error) {
	booking, err := s.repository.Find(inputID.ID, false)
	if err != nil {
		return booking, err
	}

	booking.Name = input.Name
	booking.PhoneNumber = input.PhoneNumber
	booking.PaymentMethod = input.PaymentMethod
	booking.Date = input.Date
	booking.TotalPrice = input.TotalPrice
	booking.TotalPay = input.TotalPay
	booking.TotalChange = input.TotalChange
	booking.Details = input.Details

	updatedBooking, err := s.repository.Update(booking)
	if err != nil {
		return updatedBooking, err
	}

	return updatedBooking, nil
}

// Delete -> calls Booking repo delete method
func (s bookingService) Delete(inputID inputs.GetBookingDetailInput) (models.Booking, error) {
	booking, err := s.repository.Find(inputID.ID, true)
	if err != nil {
		return booking, err
	}

	deletedBooking, err := s.repository.Delete(booking)
	if err != nil {
		return deletedBooking, err
	}

	return deletedBooking, nil
}
