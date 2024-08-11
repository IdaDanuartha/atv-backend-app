package repositories

import (
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type BookingRepository interface {
	FindAll(booking models.Booking, search string, currentPage int, pageSize int, instructorID string, customerID string) ([]models.Booking, int64, int, error)
	Find(ID string, showRelations bool) (models.Booking, error)
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
func (r bookingRepository) FindAll(booking models.Booking, search string, currentPage int, pageSize int, instructorID string, customerID string) ([]models.Booking, int64, int, error) {
	var bookings []models.Booking
	var totalRows int64 = 0

	queryBuilder := r.db.DB.Order("created_at desc").Model(&models.Booking{})

	// Search parameter
	if search != "" {
		querySearch := "%" + search + "%"
		queryBuilder = queryBuilder.Where(
			r.db.DB.Where("bookings.name LIKE ? ", querySearch).
				Or("bookings.code LIKE ? ", querySearch).
				Or("bookings.phone_number LIKE ? ", querySearch).
				Or("bookings.payment_method LIKE ? ", querySearch).
				Or("bookings.date LIKE ? ", querySearch).
				Or("bookings.total_price LIKE ? ", querySearch).
				Or("bookings.total_pay LIKE ? ", querySearch).
				Or("bookings.total_change LIKE ? ", querySearch))
	}

	if instructorID != "" {
		queryBuilder = queryBuilder.Joins("JOIN booking_details ON booking_details.booking_id = bookings.id").
			Joins("JOIN entertainment_services ON entertainment_services.id = booking_details.entertainment_service_id").
			Joins("JOIN entertainment_service_instructors ON entertainment_service_instructors.entertainment_service_id = entertainment_services.id").
			Where("entertainment_service_instructors.instructor_id = ?", instructorID)
	}

	if customerID != "" {
		queryBuilder = queryBuilder.Where("customer_id = ?", customerID)
	}

	if pageSize > 0 {
		// count the total number of rows
		err := queryBuilder.
			Where(booking).
			Count(&totalRows).Error
		if err != nil {
			return bookings, totalRows, currentPage, err
		}

		// Apply offset and limit to fetch paginated results
		err = queryBuilder.
			Preload("Details.EntertainmentService.EntertainmentCategory").
			Preload("Customer.User").
			Where(booking).
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).
			Find(&bookings).Error
		return bookings, totalRows, currentPage, err
	} else {
		err := queryBuilder.
			Preload("Details.EntertainmentService.EntertainmentCategory").
			Preload("Customer.User").
			Where(booking).
			Find(&bookings).
			Count(&totalRows).Error

		return bookings, 0, 0, err
	}
}

// Find -> Method for fetching Entertainment Package by id
func (r bookingRepository) Find(ID string, showRelations bool) (models.Booking, error) {
	var bookings models.Booking
	
	if(showRelations) {
		err := r.db.DB.
			Preload("Details.EntertainmentService.EntertainmentCategory").
			Preload("Customer.User").
			Debug().
			Model(&models.Booking{}).
			Where("id = ?", ID).
			Find(&bookings).Error
		return bookings, err
	} else {
		err := r.db.DB.
			Debug().
			Model(&models.Booking{}).
			Where("id = ?", ID).
			Find(&bookings).Error
		return bookings, err
	}
}

// Save -> Method for saving Entertainment Package to database
func (r bookingRepository) Save(booking models.Booking) (models.Booking, error) {
	err := r.db.DB.
		Preload("Details.EntertainmentService.EntertainmentCategory").
		Preload("Customer.User").
		Create(&booking).Error
	if err != nil {
		return booking, err
	}

	return booking, nil
}

// Update -> Method for updating Entertainment Package
func (r *bookingRepository) Update(booking models.Booking) (models.Booking, error) {
	err := r.db.DB.Save(&booking).Error

	if err != nil {
		return booking, err
	}

	return booking, nil
}

// Delete -> Method for deleting Entertainment Package
func (r bookingRepository) Delete(booking models.Booking) (models.Booking, error) {
	err := r.db.DB.
		Preload("Details.EntertainmentService.EntertainmentCategory").
		Preload("Customer.User").
		Delete(&booking).Error

	if err != nil {
		return booking, err
	}

	return booking, nil
}
