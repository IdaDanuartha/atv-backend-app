package services

import (
	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type EntertainmentCategoryService interface {
	FindAll(model models.EntertainmentCategory, search string, currentPage int, pageSize int) ([]models.EntertainmentCategory, int64, int, error)
	Find(input inputs.GetEntertainmentCategoryDetailInput) (models.EntertainmentCategory, error)
	Save(input inputs.EntertainmentCategoryInput) (models.EntertainmentCategory, error)
	Update(inputID inputs.GetEntertainmentCategoryDetailInput, input inputs.EntertainmentCategoryInput) (models.EntertainmentCategory, error)
	Delete(inputID inputs.GetEntertainmentCategoryDetailInput) (models.EntertainmentCategory, error)
}

// EntertainmentCategoryService EntertainmentCategoryService struct
type entertainmentCategoryService struct {
	repository repositories.EntertainmentCategoryRepository
}

// NewEntertainmentCategoryService : returns the EntertainmentCategoryService struct instance
func NewEntertainmentCategoryService(repository repositories.EntertainmentCategoryRepository) entertainmentCategoryService {
	return entertainmentCategoryService{repository}
}

// FindAll -> calls Entertainment Category repo find all method
func (s entertainmentCategoryService) FindAll(model models.EntertainmentCategory, search string, currentPage int, pageSize int) ([]models.EntertainmentCategory, int64, int, error) {
	entertainmentCategories, total, currentPage, err := s.repository.FindAll(model, search, currentPage, pageSize)
	if err != nil {
		return entertainmentCategories, total, currentPage, err
	}

	return entertainmentCategories, total, currentPage, nil
}

// Find -> calls Entertainment Category repo find method
func (s entertainmentCategoryService) Find(input inputs.GetEntertainmentCategoryDetailInput) (models.EntertainmentCategory, error) {
	entertainmentCategory, err := s.repository.Find(input.ID)

	if err != nil {
		return entertainmentCategory, err
	}

	return entertainmentCategory, nil
}

// Save -> calls Entertainment Category repository save method
func (s entertainmentCategoryService) Save(input inputs.EntertainmentCategoryInput) (models.EntertainmentCategory, error) {
	entertainmentCategory := models.EntertainmentCategory{}
	entertainmentCategory.Name = input.Name

	newEntertainmentCategory, err := s.repository.Save(entertainmentCategory)
	if err != nil {
		return newEntertainmentCategory, err
	}

	return newEntertainmentCategory, nil
}

// Update -> calls Entertainment Category repo update method
func (s entertainmentCategoryService) Update(inputID inputs.GetEntertainmentCategoryDetailInput, input inputs.EntertainmentCategoryInput) (models.EntertainmentCategory, error) {
	entertainmentCategory, err := s.repository.Find(inputID.ID)
	if err != nil {
		return entertainmentCategory, err
	}

	entertainmentCategory.Name = input.Name

	updatedEntertainmentCategory, err := s.repository.Update(entertainmentCategory)
	if err != nil {
		return updatedEntertainmentCategory, err
	}

	return updatedEntertainmentCategory, nil
}

// Delete -> calls Entertainment Category repo delete method
func (s entertainmentCategoryService) Delete(inputID inputs.GetEntertainmentCategoryDetailInput) (models.EntertainmentCategory, error) {
	entertainmentCategory, err := s.repository.Find(inputID.ID)
	if err != nil {
		return entertainmentCategory, err
	}

	deletedEntertainmentCategory, err := s.repository.Delete(entertainmentCategory)
	if err != nil {
		return deletedEntertainmentCategory, err
	}

	return deletedEntertainmentCategory, nil
}
