package repositories

import (
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

// EntertainmentPackageRepository -> EntertainmentPackageRepository
type EntertainmentPackageRepository struct {
	db config.Database
}

// NewEntertainmentPackageRepository : fetching database
func NewEntertainmentPackageRepository(db config.Database) EntertainmentPackageRepository {
	return EntertainmentPackageRepository{
		db: db,
	}
}

// Save -> Method for saving Entertainment Package to database
func (p EntertainmentPackageRepository) Save(entertainmentPackage models.EntertainmentPackage) error {
	return p.db.DB.Create(&entertainmentPackage).Error
}

// FindAll -> Method for fetching all Entertainment Package from database
func (p EntertainmentPackageRepository) FindAll(entertainmentPackage models.EntertainmentPackage, search string) (*[]models.EntertainmentPackage, int64, error) {
	var entertainment_packages []models.EntertainmentPackage
	var totalRows int64 = 0

	queryBuider := p.db.DB.Order("created_at desc").Model(&models.EntertainmentPackage{})

	// Search parameter
	if search != "" {
		querySearch := "%" + search + "%"
		queryBuider = queryBuider.Where(
			p.db.DB.Where("entertainment_packages.name LIKE ? ", querySearch))
	}

	err := queryBuider.
		Where(entertainmentPackage).
		Find(&entertainment_packages).
		Count(&totalRows).Error
	return &entertainment_packages, totalRows, err
}

// Update -> Method for updating Entertainment Package
func (p EntertainmentPackageRepository) Update(entertainmentPackage models.EntertainmentPackage) error {
	return p.db.DB.Save(&entertainmentPackage).Error
}

// Find -> Method for fetching Entertainment Package by id
func (p EntertainmentPackageRepository) Find(entertainmentPackage models.EntertainmentPackage) (models.EntertainmentPackage, error) {
	var entertainment_packages models.EntertainmentPackage
	err := p.db.DB.
		Debug().
		Model(&models.EntertainmentPackage{}).
		Where(&entertainmentPackage).
		Take(&entertainment_packages).Error
	return entertainment_packages, err
}

// Delete -> Method for deleting Entertainment Package
func (p EntertainmentPackageRepository) Delete(bus models.EntertainmentPackage) error {
	return p.db.DB.Delete(&bus).Error
}
