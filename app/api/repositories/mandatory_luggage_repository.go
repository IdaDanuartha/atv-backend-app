package repositories

import (
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type MandatoryLuggageRepository interface {
	FindAll(mandatoryLuggage models.MandatoryLuggage, search string, currentPage int, pageSize int) ([]models.MandatoryLuggage, int64, int, error)
	Find(ID string) (models.MandatoryLuggage, error)
	Save(mandatoryLuggage models.MandatoryLuggage) (models.MandatoryLuggage, error)
	Update(mandatoryLuggage models.MandatoryLuggage) (models.MandatoryLuggage, error)
	Delete(mandatoryLuggage models.MandatoryLuggage) (models.MandatoryLuggage, error)
}

type mandatoryLuggageRepository struct {
	db config.Database
}

// NewMandatoryLuggageRepository : fetching database
func NewMandatoryLuggageRepository(db config.Database) mandatoryLuggageRepository {
	return mandatoryLuggageRepository{db}
}

// FindAll -> Method for fetching all Mandatory Luggage from database
func (r mandatoryLuggageRepository) FindAll(mandatoryLuggage models.MandatoryLuggage, search string, currentPage int, pageSize int) ([]models.MandatoryLuggage, int64, int, error) {
	var mandatory_luggages []models.MandatoryLuggage
	var totalRows int64 = 0

	queryBuilder := r.db.DB.Order("created_at desc").Model(&models.MandatoryLuggage{})

	// Search parameter
	if search != "" {
		querySearch := "%" + search + "%"
		queryBuilder = queryBuilder.Where(
			r.db.DB.Where("mandatory_luggages.name LIKE ? ", querySearch))
	}

	if pageSize > 0 {
		// count the total number of rows
		err := queryBuilder.
			Where(mandatoryLuggage).
			Count(&totalRows).Error

		// Apply offset and limit to fetch paginated results
		err = queryBuilder.
			Where(mandatoryLuggage).
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).
			Find(&mandatory_luggages).Error
		return mandatory_luggages, totalRows, currentPage, err
	} else {
		err := queryBuilder.
			Where(mandatoryLuggage).
			Find(&mandatory_luggages).
			Count(&totalRows).Error
		return mandatory_luggages, 0, 0, err
	}
}

// Find -> Method for fetching Mandatory Luggage by id
func (r mandatoryLuggageRepository) Find(ID string) (models.MandatoryLuggage, error) {
	var mandatory_luggages models.MandatoryLuggage
	err := r.db.DB.
		Debug().
		Model(&models.MandatoryLuggage{}).
		Where("id = ?", ID).
		Find(&mandatory_luggages).Error
	return mandatory_luggages, err
}

// Save -> Method for saving Mandatory Luggage to database
func (r mandatoryLuggageRepository) Save(mandatoryLuggage models.MandatoryLuggage) (models.MandatoryLuggage, error) {
	err := r.db.DB.Create(&mandatoryLuggage).Error
	if err != nil {
		return mandatoryLuggage, err
	}

	return mandatoryLuggage, nil
}

// Update -> Method for updating Mandatory Luggage
func (r *mandatoryLuggageRepository) Update(mandatoryLuggage models.MandatoryLuggage) (models.MandatoryLuggage, error) {
	err := r.db.DB.Save(&mandatoryLuggage).Error

	if err != nil {
		return mandatoryLuggage, err
	}

	return mandatoryLuggage, nil
}

// Delete -> Method for deleting Mandatory Luggage
func (r mandatoryLuggageRepository) Delete(mandatoryLuggage models.MandatoryLuggage) (models.MandatoryLuggage, error) {
	err := r.db.DB.Delete(&mandatoryLuggage).Error

	if err != nil {
		return mandatoryLuggage, err
	}

	return mandatoryLuggage, nil
}
