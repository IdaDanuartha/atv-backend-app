package services

import (
	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type EntertainmentServiceService interface {
	FindAll(model models.EntertainmentService, search string) ([]models.EntertainmentService, int64, error)
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
func (s entertainmentServiceService) FindAll(model models.EntertainmentService, search string) ([]models.EntertainmentService, int64, error) {
	entertainmentServices, total, err := s.repository.FindAll(model, search)
	if err != nil {
		return entertainmentServices, total, err
	}

	return entertainmentServices, total, nil
}

// Find -> calls Entertainment Service repo find method
func (s entertainmentServiceService) Find(input inputs.GetEntertainmentServiceDetailInput) (models.EntertainmentService, error) {
	entertainmentService, err := s.repository.Find(input.ID)

	if err != nil {
		return entertainmentService, err
	}

	return entertainmentService, nil
}

func (s entertainmentServiceService) SaveImage(ID string, fileLocation string) (models.EntertainmentService, error) {
	entertainment_service, err := s.repository.Find(ID)
	if err != nil {
		return entertainment_service, err
	}

	entertainment_service.ImagePath = &fileLocation

	updatedEntertainmentService, err := s.repository.Update(entertainment_service)
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
	entertainmentService.EntertainmentServiceFacilities = input.Facilities
	entertainmentService.EntertainmentServiceInstructors = input.Instructors
	entertainmentService.MandatoryLuggageEntertainmentServices = input.MandatoryLuggages

	newEntertainmentService, err := s.repository.Save(entertainmentService)
	if err != nil {
		return newEntertainmentService, err
	}

	return newEntertainmentService, nil
}

// Update -> calls Entertainment Service repo update method
func (s entertainmentServiceService) Update(inputID inputs.GetEntertainmentServiceDetailInput, input inputs.EntertainmentServiceInput) (models.EntertainmentService, error) {
	entertainmentService, err := s.repository.Find(inputID.ID)
	if err != nil {
		return entertainmentService, err
	}

	entertainmentService.Name = input.Name
	entertainmentService.Price = input.Price
	entertainmentService.EntertainmentCategoryID = input.EntertainmentCategoryID
	entertainmentService.Routes = input.Routes
	entertainmentService.EntertainmentServiceFacilities = input.Facilities
	entertainmentService.EntertainmentServiceInstructors = input.Instructors
	entertainmentService.MandatoryLuggageEntertainmentServices = input.MandatoryLuggages

	updatedEntertainmentService, err := s.repository.Update(entertainmentService)
	if err != nil {
		return updatedEntertainmentService, err
	}

	return updatedEntertainmentService, nil
}

// Delete -> calls Entertainment Service repo delete method
func (s entertainmentServiceService) Delete(inputID inputs.GetEntertainmentServiceDetailInput) (models.EntertainmentService, error) {
	entertainmentService, err := s.repository.Find(inputID.ID)
	if err != nil {
		return entertainmentService, err
	}

	deletedEntertainmentService, err := s.repository.Delete(entertainmentService)
	if err != nil {
		return deletedEntertainmentService, err
	}

	return deletedEntertainmentService, nil
}
