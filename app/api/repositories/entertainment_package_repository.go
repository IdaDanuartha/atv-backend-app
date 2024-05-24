package repositories

import (
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type EntertainmentPackageRepository interface {
	FindAll(entertainmentPackage models.EntertainmentPackage, search string) ([]models.EntertainmentPackage, int64, error)
	Find(ID string) (models.EntertainmentPackage, error)
	Save(entertainmentPackage models.EntertainmentPackage) (models.EntertainmentPackage, error)
	Update(entertainmentPackage models.EntertainmentPackage) (models.EntertainmentPackage, error)
	Delete(entertainmentPackage models.EntertainmentPackage) (models.EntertainmentPackage, error)
}

type entertainmentPackageRepository struct {
	db config.Database
}

// NewEntertainmentPackageRepository : fetching database
func NewEntertainmentPackageRepository(db config.Database) entertainmentPackageRepository {
	return entertainmentPackageRepository{db}
}

// FindAll -> Method for fetching all Entertainment Package from database
func (r entertainmentPackageRepository) FindAll(entertainmentPackage models.EntertainmentPackage, search string) ([]models.EntertainmentPackage, int64, error) {
	var entertainment_packages []models.EntertainmentPackage
	var totalRows int64 = 0

	queryBuider := r.db.DB.Order("created_at desc").Model(&models.EntertainmentPackage{})

	// Search parameter
	if search != "" {
		querySearch := "%" + search + "%"
		queryBuider = queryBuider.Where(
			r.db.DB.Where("entertainment_packages.name LIKE ? ", querySearch))
	}

	err := queryBuider.
		Where(entertainmentPackage).
		Find(&entertainment_packages).
		Count(&totalRows).Error
	return entertainment_packages, totalRows, err
}

// Find -> Method for fetching Entertainment Package by id
func (r entertainmentPackageRepository) Find(ID string) (models.EntertainmentPackage, error) {
	var entertainment_packages models.EntertainmentPackage
	err := r.db.DB.
		Debug().
		Model(&models.EntertainmentPackage{}).
		Where("id = ?", ID).
		Find(&entertainment_packages).Error
	return entertainment_packages, err
}

// Save -> Method for saving Entertainment Package to database
func (r entertainmentPackageRepository) Save(entertainmentPackage models.EntertainmentPackage) (models.EntertainmentPackage, error) {
	err := r.db.DB.Create(&entertainmentPackage).Error
	if err != nil {
		return entertainmentPackage, err
	}

	return entertainmentPackage, nil
}

// Update -> Method for updating Entertainment Package
func (r *entertainmentPackageRepository) Update(entertainmentPackage models.EntertainmentPackage) (models.EntertainmentPackage, error) {
	err := r.db.DB.Save(&entertainmentPackage).Error

	if err != nil {
		return entertainmentPackage, err
	}

	return entertainmentPackage, nil
}

// Delete -> Method for deleting Entertainment Package
func (r entertainmentPackageRepository) Delete(entertainmentPackage models.EntertainmentPackage) (models.EntertainmentPackage, error) {
	err := r.db.DB.Delete(&entertainmentPackage).Error

	if err != nil {
		return entertainmentPackage, err
	}

	return entertainmentPackage, nil
}
