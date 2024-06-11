package repositories

import (
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type EntertainmentPackageRepository interface {
	FindAll(entertainmentPackage models.EntertainmentPackage, search string, currentPage int, pageSize int) ([]models.EntertainmentPackage, int64, int, error)
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
func (r entertainmentPackageRepository) FindAll(entertainmentPackage models.EntertainmentPackage, search string, currentPage int, pageSize int) ([]models.EntertainmentPackage, int64, int, error) {
	var entertainment_packages []models.EntertainmentPackage
	var totalRows int64 = 0

	queryBuilder := r.db.DB.Order("created_at desc").Model(&models.EntertainmentPackage{})

	// Search parameter
	if search != "" {
		querySearch := "%" + search + "%"
		queryBuilder = queryBuilder.Where(
			r.db.DB.Where("entertainment_packages.name LIKE ? ", querySearch))
	}

	if pageSize > 0 {
		// count the total number of rows
		err := queryBuilder.
			Where(entertainmentPackage).
			Count(&totalRows).Error

		// Apply offset and limit to fetch paginated results
		err = queryBuilder.
			Preload("EntertainmentPackageDetails.EntertainmentService").
			Where(entertainmentPackage).
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).
			Find(&entertainment_packages).Error
		return entertainment_packages, totalRows, currentPage, err
	} else {
		err := queryBuilder.
			Preload("EntertainmentPackageDetails.EntertainmentService").
			Where(entertainmentPackage).
			Find(&entertainment_packages).
			Count(&totalRows).Error
		return entertainment_packages, 0, 0, err
	}
}

// Find -> Method for fetching Entertainment Package by id
func (r entertainmentPackageRepository) Find(ID string) (models.EntertainmentPackage, error) {
	var entertainment_packages models.EntertainmentPackage
	err := r.db.DB.
		Preload("EntertainmentPackageDetails.EntertainmentService").
		Debug().
		Model(&models.EntertainmentPackage{}).
		Where("id = ?", ID).
		Find(&entertainment_packages).Error
	return entertainment_packages, err
}

// Save -> Method for saving Entertainment Package to database
func (r entertainmentPackageRepository) Save(entertainmentPackage models.EntertainmentPackage) (models.EntertainmentPackage, error) {
	err := r.db.DB.
		Preload("EntertainmentPackageDetails.EntertainmentService").
		Create(&entertainmentPackage).Error
	if err != nil {
		return entertainmentPackage, err
	}

	return entertainmentPackage, nil
}

// Update -> Method for updating Entertainment Package
func (r *entertainmentPackageRepository) Update(entertainmentPackage models.EntertainmentPackage) (models.EntertainmentPackage, error) {
	err := r.db.DB.
		Preload("EntertainmentPackageDetails.EntertainmentService").
		Save(&entertainmentPackage).Error

	if err != nil {
		return entertainmentPackage, err
	}

	return entertainmentPackage, nil
}

// Delete -> Method for deleting Entertainment Package
func (r entertainmentPackageRepository) Delete(entertainmentPackage models.EntertainmentPackage) (models.EntertainmentPackage, error) {
	err := r.db.DB.
		Preload("EntertainmentPackageDetails.EntertainmentService").
		Delete(&entertainmentPackage).Error

	if err != nil {
		return entertainmentPackage, err
	}

	return entertainmentPackage, nil
}
