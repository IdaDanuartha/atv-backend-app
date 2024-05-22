package repositories

import (
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

// MandatoryLuggageRepository -> MandatoryLuggageRepository
type MandatoryLuggageRepository struct {
	db config.Database
}

// NewMandatoryLuggageRepository : fetching database
func NewMandatoryLuggageRepository(db config.Database) MandatoryLuggageRepository {
	return MandatoryLuggageRepository{
		db: db,
	}
}

// Save -> Method for saving Mandatory Luggage to database
func (p MandatoryLuggageRepository) Save(mandatoryLuggage models.MandatoryLuggage) error {
	return p.db.DB.Create(&mandatoryLuggage).Error
}

// FindAll -> Method for fetching all mandatory luggage from database
func (p MandatoryLuggageRepository) FindAll(mandatoryLuggage models.MandatoryLuggage, search string) (*[]models.MandatoryLuggage, int64, error) {
	var mandatory_luggages []models.MandatoryLuggage
	var totalRows int64 = 0

	queryBuider := p.db.DB.Order("created_at desc").Model(&models.MandatoryLuggage{})

	// Search parameter
	if search != "" {
		querySearch := "%" + search + "%"
		queryBuider = queryBuider.Where(
			p.db.DB.Where("mandatory_luggages.name LIKE ? ", querySearch))
	}

	err := queryBuider.
		Where(mandatoryLuggage).
		Find(&mandatory_luggages).
		Count(&totalRows).Error
	return &mandatory_luggages, totalRows, err
}

// Update -> Method for updating Mandatory Luggage
func (p MandatoryLuggageRepository) Update(mandatoryLuggage models.MandatoryLuggage) error {
	return p.db.DB.Save(&mandatoryLuggage).Error
}

// Find -> Method for fetching Mandatory Luggage by id
func (p MandatoryLuggageRepository) Find(mandatoryLuggage models.MandatoryLuggage) (models.MandatoryLuggage, error) {
	var mandatory_luggages models.MandatoryLuggage
	err := p.db.DB.
		Debug().
		Model(&models.MandatoryLuggage{}).
		Where(&mandatoryLuggage).
		Take(&mandatory_luggages).Error
	return mandatory_luggages, err
}

// Delete -> Method for deleting Mandatory Luggage
func (p MandatoryLuggageRepository) Delete(bus models.MandatoryLuggage) error {
	return p.db.DB.Delete(&bus).Error
}
