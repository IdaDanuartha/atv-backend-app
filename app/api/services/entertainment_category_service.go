package services

import (
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

//EntertainmentCategoryService EntertainmentCategoryService struct
type EntertainmentCategoryService struct {
    repository repositories.EntertainmentCategoryRepository
}

//NewEntertainmentCategoryService : returns the EntertainmentCategoryService struct instance
func NewEntertainmentCategoryService(r repositories.EntertainmentCategoryRepository) EntertainmentCategoryService {
    return EntertainmentCategoryService{
        repository: r,
    }
}

//Save -> calls Entertainment Category repository save method
func (p EntertainmentCategoryService) Save(entertainmentCategory models.EntertainmentCategory) error {
    return p.repository.Save(entertainmentCategory)
}

//FindAll -> calls Entertainment Category repo find all method
func (p EntertainmentCategoryService) FindAll(entertainmentCategory models.EntertainmentCategory, keyword string) (*[]models.EntertainmentCategory, int64, error) {
    return p.repository.FindAll(entertainmentCategory, keyword)
}

// Update -> calls Entertainment Category repo update method
func (p EntertainmentCategoryService) Update(entertainmentCategory models.EntertainmentCategory) error {
    return p.repository.Update(entertainmentCategory)
}

// Delete -> calls bus repo delete method
func (p EntertainmentCategoryService) Delete(id string) error {
    var entertainmentCategory models.EntertainmentCategory
    entertainmentCategory.ID = id
    return p.repository.Delete(entertainmentCategory)
}

// Find -> calls bus repo find method
func (p EntertainmentCategoryService) Find(entertainmentCategory models.EntertainmentCategory) (models.EntertainmentCategory, error) {
    return p.repository.Find(entertainmentCategory)
}