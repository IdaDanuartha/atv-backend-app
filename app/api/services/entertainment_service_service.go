package services

import (
	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type EntertainmentServiceService interface {
	FindAll(model models.EntertainmentService, search string, currentPage int, pageSize int) ([]models.EntertainmentService, int64, int, error)
	Find(input inputs.GetEntertainmentServiceDetailInput) (models.EntertainmentService, error)
	SaveImage(ID string, fileLocation string) (models.EntertainmentService, error)
	Save(input inputs.EntertainmentServiceInput) (models.EntertainmentService, error)
	Update(inputID inputs.GetEntertainmentServiceDetailInput, input inputs.EntertainmentServiceInput) (models.EntertainmentService, error)
	Delete(inputID inputs.GetEntertainmentServiceDetailInput) (models.EntertainmentService, error)
}

// EntertainmentServiceService EntertainmentServiceService struct
type entertainmentServiceService struct {
	repository repositories.EntertainmentServiceRepository
}

// NewEntertainmentServiceService : returns the EntertainmentServiceService struct instance
func NewEntertainmentServiceService(repository repositories.EntertainmentServiceRepository) entertainmentServiceService {
	return entertainmentServiceService{repository}
}

// FindAll -> calls Entertainment Service repo find all method
func (s entertainmentServiceService) FindAll(model models.EntertainmentService, search string, currentPage int, pageSize int) ([]models.EntertainmentService, int64, int, error) {
	entertainmentServices, total, currentPage, err := s.repository.FindAll(model, search, currentPage, pageSize)
	if err != nil {
		return entertainmentServices, total, currentPage, err
	}

	return entertainmentServices, total, currentPage, nil
}

// Find -> calls Entertainment Service repo find method
func (s entertainmentServiceService) Find(input inputs.GetEntertainmentServiceDetailInput) (models.EntertainmentService, error) {
	entertainmentService, err := s.repository.Find(input.ID, true)

	if err != nil {
		return entertainmentService, err
	}

	return entertainmentService, nil
}

func (s entertainmentServiceService) SaveImage(ID string, fileLocation string) (models.EntertainmentService, error) {
	// Find the existing entertainment service by ID
	entertainment_service, err := s.repository.Find(ID, false)
	if err != nil {
		return entertainment_service, err
	}

	// Update only the ImagePath field
	err = s.repository.UpdateImagePath(ID, fileLocation)
	if err != nil {
		return entertainment_service, err
	}

	// Fetch the updated entertainment service to return it
	updatedEntertainmentService, err := s.repository.Find(ID, false)
	if err != nil {
		return updatedEntertainmentService, err
	}

	return updatedEntertainmentService, nil
}

// Save -> calls Entertainment Service repository save method
func (s entertainmentServiceService) Save(input inputs.EntertainmentServiceInput) (models.EntertainmentService, error) {
	entertainmentService := models.EntertainmentService{}

	entertainmentService.Name = input.Name
	entertainmentService.Price = input.Price
	entertainmentService.EntertainmentCategoryID = input.EntertainmentCategoryID
	entertainmentService.Routes = input.Routes
	entertainmentService.Facilities = input.Facilities
	entertainmentService.Instructors = input.Instructors
	entertainmentService.MandatoryLuggages = input.MandatoryLuggages

	newEntertainmentService, err := s.repository.Save(entertainmentService)
	if err != nil {
		return newEntertainmentService, err
	}

	return newEntertainmentService, nil
}

// Update -> calls Entertainment Service repo update method
func (s entertainmentServiceService) Update(inputID inputs.GetEntertainmentServiceDetailInput, input inputs.EntertainmentServiceInput) (models.EntertainmentService, error) {
	entertainmentService, err := s.repository.Find(inputID.ID, false)
	if err != nil {
		return entertainmentService, err
	}

	entertainmentService.Name = input.Name
	entertainmentService.Price = input.Price
	entertainmentService.EntertainmentCategoryID = input.EntertainmentCategoryID
	entertainmentService.Routes = input.Routes
	entertainmentService.Facilities = input.Facilities
	entertainmentService.Instructors = input.Instructors
	entertainmentService.MandatoryLuggages = input.MandatoryLuggages

	updatedEntertainmentService, err := s.repository.Update(entertainmentService)
	if err != nil {
		return updatedEntertainmentService, err
	}

	return updatedEntertainmentService, nil
}

// Delete -> calls Entertainment Service repo delete method
func (s entertainmentServiceService) Delete(inputID inputs.GetEntertainmentServiceDetailInput) (models.EntertainmentService, error) {
	entertainmentService, err := s.repository.Find(inputID.ID, true)
	if err != nil {
		return entertainmentService, err
	}

	deletedEntertainmentService, err := s.repository.Delete(entertainmentService)
	if err != nil {
		return deletedEntertainmentService, err
	}

	return deletedEntertainmentService, nil
}