package repositories

import (
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type EntertainmentServiceRepository interface {
	FindAll(entertainmentService models.EntertainmentService, search string, currentPage int, pageSize int) ([]models.EntertainmentService, int64, int, error)
	Find(ID string) (models.EntertainmentService, error)
	Save(entertainmentService models.EntertainmentService) (models.EntertainmentService, error)
	Update(entertainmentService models.EntertainmentService) (models.EntertainmentService, error)
	Delete(entertainmentService models.EntertainmentService) (models.EntertainmentService, error)
}

type entertainmentServiceRepository struct {
	db config.Database
}

// NewEntertainmentServiceRepository : fetching database
func NewEntertainmentServiceRepository(db config.Database) entertainmentServiceRepository {
	return entertainmentServiceRepository{db}
}

// FindAll -> Method for fetching all Entertainment Service from database
func (r entertainmentServiceRepository) FindAll(entertainmentService models.EntertainmentService, search string, currentPage int, pageSize int) ([]models.EntertainmentService, int64, int, error) {
	var entertainment_services []models.EntertainmentService
	var totalRows int64 = 0

	queryBuilder := r.db.DB.Order("created_at desc").Model(&models.EntertainmentService{})

	// Search parameter
	if search != "" {
		querySearch := "%" + search + "%"
		queryBuilder = queryBuilder.Where(
			r.db.DB.Where("entertainment_services.name LIKE ? ", querySearch))
	}

	if pageSize > 0 {
		// count the total number of rows
		err := queryBuilder.
			Where(entertainmentService).
			Count(&totalRows).Error

		// Apply offset and limit to fetch paginated results
		err = queryBuilder.
			Preload("EntertainmentCategory").
			Preload("Routes.Route").
			Preload("Facilities.Facility").
			Preload("Instructors.Instructor.User").
			Preload("MandatoryLuggages.MandatoryLuggage").
			Where(entertainmentService).
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).
			Find(&entertainment_services).Error
		return entertainment_services, totalRows, currentPage, err
	} else {
		err := queryBuilder.
			Preload("EntertainmentCategory").
			Preload("Routes.Route").
			Preload("Facilities.Facility").
			Preload("Instructors.Instructor.User").
			Preload("MandatoryLuggages.MandatoryLuggage").
			Where(entertainmentService).
			Find(&entertainment_services).
			Count(&totalRows).Error
		return entertainment_services, 0, 0, err
	}
}

// Find -> Method for fetching Entertainment Service by id
func (r entertainmentServiceRepository) Find(ID string) (models.EntertainmentService, error) {
	var entertainment_services models.EntertainmentService
	err := r.db.DB.
		Preload("EntertainmentCategory").
		Preload("Routes.Route").
		Preload("Facilities.Facility").
		Preload("Instructors.Instructor.User").
		Preload("MandatoryLuggages.MandatoryLuggage").
		Debug().
		Model(&models.EntertainmentService{}).
		Where("id = ?", ID).
		Find(&entertainment_services).Error
	return entertainment_services, err
}

// Save -> Method for saving Entertainment Service to database
func (r entertainmentServiceRepository) Save(entertainmentService models.EntertainmentService) (models.EntertainmentService, error) {
	err := r.db.DB.
		Preload("EntertainmentCategory").
		Preload("Routes.Route").
		Preload("Facilities.Facility").
		Preload("Instructors.Instructor.User").
		Preload("MandatoryLuggages.MandatoryLuggage").
		Create(&entertainmentService).Error

	if err != nil {
		return entertainmentService, err
	}

	return entertainmentService, nil
}

// Delete -> Method for deleting Entertainment Service Route
//func (r entertainmentServiceRepository) DeleteRoute(entertainmentServiceRoute models.EntertainmentServiceRoute) (error) {
//	err := r.db.DB.Delete(&entertainmentServiceRoute).Error
//
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

// Update -> Method for updating Entertainment Service
func (r *entertainmentServiceRepository) Update(entertainmentService models.EntertainmentService) (models.EntertainmentService, error) {
	err := r.db.DB.
		Where("entertainment_service_id = ?", entertainmentService.ID).
		Delete(&models.EntertainmentServiceRoute{}).Error

	if err != nil {
		return entertainmentService, err
	}

	err = r.db.DB.
		Where("entertainment_service_id = ?", entertainmentService.ID).
		Delete(&models.EntertainmentServiceFacility{}).Error

	if err != nil {
		return entertainmentService, err
	}

	err = r.db.DB.
		Where("entertainment_service_id = ?", entertainmentService.ID).
		Delete(&models.EntertainmentServiceInstructor{}).Error

	if err != nil {
		return entertainmentService, err
	}

	err = r.db.DB.
		Where("entertainment_service_id = ?", entertainmentService.ID).
		Delete(&models.MandatoryLuggageEntertainmentService{}).Error

	if err != nil {
		return entertainmentService, err
	}

	err = r.db.DB.
		Preload("EntertainmentCategory").
		Preload("Routes.Route").
		Preload("Facilities.Facility").
		Preload("Instructors.Instructor.User").
		Preload("MandatoryLuggages.MandatoryLuggage").
		Save(&entertainmentService).Error

	if err != nil {
		return entertainmentService, err
	}

	return entertainmentService, nil
}

// Delete -> Method for deleting Entertainment Service
func (r entertainmentServiceRepository) Delete(entertainmentService models.EntertainmentService) (models.EntertainmentService, error) {
	err := r.db.DB.
		Preload("EntertainmentCategory").
		Preload("Routes.Route").
		Preload("Facilities.Facility").
		Preload("Instructors.Instructor.User").
		Preload("MandatoryLuggages.MandatoryLuggage").
		Delete(&entertainmentService).Error

	if err != nil {
		return entertainmentService, err
	}

	return entertainmentService, nil
}
