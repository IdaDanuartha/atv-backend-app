package repositories

import (
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type EntertainmentServiceRepository interface {
	FindAll(entertainmentService models.EntertainmentService, search string) ([]models.EntertainmentService, int64, error)
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
func (r entertainmentServiceRepository) FindAll(entertainmentService models.EntertainmentService, search string) ([]models.EntertainmentService, int64, error) {
	var entertainment_services []models.EntertainmentService
	var totalRows int64 = 0

	queryBuider := r.db.DB.Order("created_at desc").Model(&models.EntertainmentService{})

	// Search parameter
	if search != "" {
		querySearch := "%" + search + "%"
		queryBuider = queryBuider.Where(
			r.db.DB.Where("entertainment_services.name LIKE ? ", querySearch))
	}

	err := queryBuider.
		Preload("EntertainmentCategory").
		Preload("Route").
		Where(entertainmentService).
		Find(&entertainment_services).
		Count(&totalRows).Error
	return entertainment_services, totalRows, err
}

// Find -> Method for fetching Entertainment Service by id
func (r entertainmentServiceRepository) Find(ID string) (models.EntertainmentService, error) {
	var entertainment_services models.EntertainmentService
	err := r.db.DB.
		Debug().
		Model(&models.EntertainmentService{}).
		Where("id = ?", ID).
		Find(&entertainment_services).Error
	return entertainment_services, err
}

// Save -> Method for saving Entertainment Service to database
func (r entertainmentServiceRepository) Save(entertainmentService models.EntertainmentService) (models.EntertainmentService, error) {
	err := r.db.DB.Create(&entertainmentService).Error
	if err != nil {
		return entertainmentService, err
	}

	return entertainmentService, nil
}

// Update -> Method for updating Entertainment Service
func (r *entertainmentServiceRepository) Update(entertainmentService models.EntertainmentService) (models.EntertainmentService, error) {
	err := r.db.DB.Save(&entertainmentService).Error

	if err != nil {
		return entertainmentService, err
	}

	return entertainmentService, nil
}

// Delete -> Method for deleting Entertainment Service
func (r entertainmentServiceRepository) Delete(entertainmentService models.EntertainmentService) (models.EntertainmentService, error) {
	err := r.db.DB.Delete(&entertainmentService).Error

	if err != nil {
		return entertainmentService, err
	}

	return entertainmentService, nil
}
