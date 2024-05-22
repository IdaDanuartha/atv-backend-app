package repositories

import (
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

//EntertainmentCategoryRepository -> EntertainmentCategoryRepository
type EntertainmentCategoryRepository struct {
    db config.Database
}

// NewEntertainmentCategoryRepository : fetching database
func NewEntertainmentCategoryRepository(db config.Database) EntertainmentCategoryRepository {
    return EntertainmentCategoryRepository{
        db: db,
    }
}

//Save -> Method for saving Entertainment Category to database
func (p EntertainmentCategoryRepository) Save(entertainmentCategory models.EntertainmentCategory) error {
    return p.db.DB.Create(&entertainmentCategory).Error
}

//FindAll -> Method for fetching all Entertainment Category from database
func (p EntertainmentCategoryRepository) FindAll(entertainmentCategory models.EntertainmentCategory, search string) (*[]models.EntertainmentCategory, int64, error) {
    var entertainment_categories []models.EntertainmentCategory
    var totalRows int64 = 0

    queryBuider := p.db.DB.Order("created_at desc").Model(&models.EntertainmentCategory{})

    // Search parameter
    if search != "" {
        querySearch := "%" + search + "%"
        queryBuider = queryBuider.Where(
            p.db.DB.Where("entertainment_categories.name LIKE ? ", querySearch))
    }

    err := queryBuider.
        Where(entertainmentCategory).
        Find(&entertainment_categories).
        Count(&totalRows).Error
    return &entertainment_categories, totalRows, err
}

//Update -> Method for updating Entertainment Category
func (p EntertainmentCategoryRepository) Update(entertainmentCategory models.EntertainmentCategory) error {
    return p.db.DB.Save(&entertainmentCategory).Error
}

//Find -> Method for fetching Entertainment Category by id
func (p EntertainmentCategoryRepository) Find(entertainmentCategory models.EntertainmentCategory) (models.EntertainmentCategory, error) {
    var entertainment_categories models.EntertainmentCategory
    err := p.db.DB.
        Debug().
        Model(&models.EntertainmentCategory{}).
        Where(&entertainmentCategory).
        Take(&entertainment_categories).Error
    return entertainment_categories, err
}

//Delete -> Method for deleting Entertainment Category
func (p EntertainmentCategoryRepository) Delete(bus models.EntertainmentCategory) error {
    return p.db.DB.Delete(&bus).Error
}