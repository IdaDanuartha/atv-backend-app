package repositories

import (
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type EntertainmentCategoryRepository interface {
	FindAll(entertainmentCategory models.EntertainmentCategory, search string, currentPage int, pageSize int) ([]models.EntertainmentCategory, int64, int, error)
	Find(ID string) (models.EntertainmentCategory, error)
	Save(entertainmentCategory models.EntertainmentCategory) (models.EntertainmentCategory, error)
	Update(entertainmentCategory models.EntertainmentCategory) (models.EntertainmentCategory, error)
	Delete(entertainmentCategory models.EntertainmentCategory) (models.EntertainmentCategory, error)
}

type entertainmentCategoryRepository struct {
	db config.Database
}

// NewEntertainmentCategoryRepository : fetching database
func NewEntertainmentCategoryRepository(db config.Database) entertainmentCategoryRepository {
	return entertainmentCategoryRepository{db}
}

// FindAll -> Method for fetching all Entertainment Category from database
func (r entertainmentCategoryRepository) FindAll(entertainmentCategory models.EntertainmentCategory, search string, currentPage int, pageSize int) ([]models.EntertainmentCategory, int64, int, error) {
	var entertainment_categories []models.EntertainmentCategory
	var totalRows int64 = 0

	queryBuilder := r.db.DB.Order("created_at desc").Model(&models.EntertainmentCategory{})

	// Search parameter
	if search != "" {
		querySearch := "%" + search + "%"
		queryBuilder = queryBuilder.Where(
			r.db.DB.Where("entertainment_categories.name LIKE ? ", querySearch))
	}

	if pageSize > 0 {
		// count the total number of rows
		err := queryBuilder.
			Where(entertainmentCategory).
			Count(&totalRows).Error

		// Apply offset and limit to fetch paginated results
		err = queryBuilder.
			Where(entertainmentCategory).
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).
			Find(&entertainment_categories).Error
		return entertainment_categories, totalRows, currentPage, err
	} else {
		err := queryBuilder.
			Where(entertainmentCategory).
			Find(&entertainment_categories).
			Count(&totalRows).Error
		return entertainment_categories, 0, 0, err
	}
}

// Find -> Method for fetching Entertainment Category by id
func (r entertainmentCategoryRepository) Find(ID string) (models.EntertainmentCategory, error) {
	var entertainment_categories models.EntertainmentCategory
	err := r.db.DB.
		Debug().
		Model(&models.EntertainmentCategory{}).
		Where("id = ?", ID).
		Find(&entertainment_categories).Error
	return entertainment_categories, err
}

// Save -> Method for saving Entertainment Category to database
func (r entertainmentCategoryRepository) Save(entertainmentCategory models.EntertainmentCategory) (models.EntertainmentCategory, error) {
	err := r.db.DB.Create(&entertainmentCategory).Error
	if err != nil {
		return entertainmentCategory, err
	}

	return entertainmentCategory, nil
}

// Update -> Method for updating Entertainment Category
func (r *entertainmentCategoryRepository) Update(entertainmentCategory models.EntertainmentCategory) (models.EntertainmentCategory, error) {
	err := r.db.DB.Save(&entertainmentCategory).Error

	if err != nil {
		return entertainmentCategory, err
	}

	return entertainmentCategory, nil
}

// Delete -> Method for deleting Entertainment Category
func (r entertainmentCategoryRepository) Delete(entertainmentCategory models.EntertainmentCategory) (models.EntertainmentCategory, error) {
	err := r.db.DB.Delete(&entertainmentCategory).Error

	if err != nil {
		return entertainmentCategory, err
	}

	return entertainmentCategory, nil
}
