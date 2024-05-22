package repositories

import (
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

// FacilityRepository -> FacilityRepository
type FacilityRepository struct {
	db config.Database
}

// NewFacilityRepository : fetching database
func NewFacilityRepository(db config.Database) FacilityRepository {
	return FacilityRepository{
		db: db,
	}
}

// Save -> Method for saving Facility to database
func (p FacilityRepository) Save(facility models.Facility) error {
	return p.db.DB.Create(&facility).Error
}

// FindAll -> Method for fetching all facility from database
func (p FacilityRepository) FindAll(facility models.Facility, search string) (*[]models.Facility, int64, error) {
	var facilities []models.Facility
	var totalRows int64 = 0

	queryBuider := p.db.DB.Order("created_at desc").Model(&models.Facility{})

	// Search parameter
	if search != "" {
		querySearch := "%" + search + "%"
		queryBuider = queryBuider.Where(
			p.db.DB.Where("facilities.name LIKE ? ", querySearch))
	}

	err := queryBuider.
		Where(facility).
		Find(&facilities).
		Count(&totalRows).Error
	return &facilities, totalRows, err
}

// Update -> Method for updating Facility
func (p FacilityRepository) Update(facility models.Facility) error {
	return p.db.DB.Save(&facility).Error
}

// Find -> Method for fetching Facility by id
func (p FacilityRepository) Find(facility models.Facility) (models.Facility, error) {
	var facilities models.Facility
	err := p.db.DB.
		Debug().
		Model(&models.Facility{}).
		Where(&facility).
		Take(&facilities).Error
	return facilities, err
}

// Delete -> Method for deleting Facility
func (p FacilityRepository) Delete(bus models.Facility) error {
	return p.db.DB.Delete(&bus).Error
}
