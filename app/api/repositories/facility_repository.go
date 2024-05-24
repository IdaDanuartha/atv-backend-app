package repositories

import (
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type FacilityRepository interface {
	FindAll(facility models.Facility, search string) ([]models.Facility, int64, error)
	Find(ID string) (models.Facility, error)
	Save(facility models.Facility) (models.Facility, error)
	Update(facility models.Facility) (models.Facility, error)
	Delete(facility models.Facility) (models.Facility, error)
}

type facilityRepository struct {
	db config.Database
}

// NewFacilityRepository : fetching database
func NewFacilityRepository(db config.Database) facilityRepository {
	return facilityRepository{db}
}

// FindAll -> Method for fetching all Facility from database
func (r facilityRepository) FindAll(facility models.Facility, search string) ([]models.Facility, int64, error) {
	var facilities []models.Facility
	var totalRows int64 = 0

	queryBuider := r.db.DB.Order("created_at desc").Model(&models.Facility{})

	// Search parameter
	if search != "" {
		querySearch := "%" + search + "%"
		queryBuider = queryBuider.Where(
			r.db.DB.Where("facilities.name LIKE ? ", querySearch))
	}

	err := queryBuider.
		Where(facility).
		Find(&facilities).
		Count(&totalRows).Error
	return facilities, totalRows, err
}

// Find -> Method for fetching Facility by id
func (r facilityRepository) Find(ID string) (models.Facility, error) {
	var facilities models.Facility
	err := r.db.DB.
		Debug().
		Model(&models.Facility{}).
		Where("id = ?", ID).
		Find(&facilities).Error
	return facilities, err
}

// Save -> Method for saving Facility to database
func (r facilityRepository) Save(facility models.Facility) (models.Facility, error) {
	err := r.db.DB.Create(&facility).Error
	if err != nil {
		return facility, err
	}

	return facility, nil
}

// Update -> Method for updating Facility
func (r *facilityRepository) Update(facility models.Facility) (models.Facility, error) {
	err := r.db.DB.Save(&facility).Error

	if err != nil {
		return facility, err
	}

	return facility, nil
}

// Delete -> Method for deleting Facility
func (r facilityRepository) Delete(facility models.Facility) (models.Facility, error) {
	err := r.db.DB.Delete(&facility).Error

	if err != nil {
		return facility, err
	}

	return facility, nil
}
