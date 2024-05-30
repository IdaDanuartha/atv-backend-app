package repositories

import (
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type BookingRepository interface {
	FindAll(booking models.Booking, search string) ([]models.Booking, int64, error)
	Find(ID string) (models.Booking, error)
	Save(booking models.Booking) (models.Booking, error)
	Update(booking models.Booking) (models.Booking, error)
	Delete(booking models.Booking) (models.Booking, error)
}

type bookingRepository struct {
	db config.Database
}

// NewBookingRepository : fetching database
func NewBookingRepository(db config.Database) bookingRepository {
	return bookingRepository{db}
}

// FindAll -> Method for fetching all Entertainment Package from database
func (r bookingRepository) FindAll(booking models.Booking, search string) ([]models.Booking, int64, error) {
	var bookings []models.Booking
	var totalRows int64 = 0

	queryBuider := r.db.DB.Order("created_at desc").Model(&models.Booking{})

	// Search parameter
	if search != "" {
		querySearch := "%" + search + "%"
		queryBuider = queryBuider.Where(
			r.db.DB.Where("bookings.name LIKE ? ", querySearch))
	}

	err := queryBuider.
		Preload("BookingDetails.EntertainmentService").
		Where(booking).
		Find(&bookings).
		Count(&totalRows).Error
	return bookings, totalRows, err
}

// Find -> Method for fetching Entertainment Package by id
func (r bookingRepository) Find(ID string) (models.Booking, error) {
	var bookings models.Booking
	err := r.db.DB.
		Preload("BookingDetails.EntertainmentService").
		Debug().
		Model(&models.Booking{}).
		Where("id = ?", ID).
		Find(&bookings).Error
	return bookings, err
}

// Save -> Method for saving Entertainment Package to database
func (r bookingRepository) Save(booking models.Booking) (models.Booking, error) {
	err := r.db.DB.
		Preload("BookingDetails.EntertainmentService").
		Create(&booking).Error
	if err != nil {
		return booking, err
	}

	return booking, nil
}

// Update -> Method for updating Entertainment Package
func (r *bookingRepository) Update(booking models.Booking) (models.Booking, error) {
	err := r.db.DB.
		Preload("BookingDetails.EntertainmentService").
		Save(&booking).Error

	if err != nil {
		return booking, err
	}

	return booking, nil
}

// Delete -> Method for deleting Entertainment Package
func (r bookingRepository) Delete(booking models.Booking) (models.Booking, error) {
	err := r.db.DB.
		Preload("BookingDetails.EntertainmentService").
		Delete(&booking).Error

	if err != nil {
		return booking, err
	}

	return booking, nil
}
