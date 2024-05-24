package services

import (
	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type MandatoryLuggageService interface {
	FindAll(model models.MandatoryLuggage, search string) ([]models.MandatoryLuggage, int64, error)
	Find(input inputs.GetMandatoryLuggageDetailInput) (models.MandatoryLuggage, error)
	Save(input inputs.MandatoryLuggageInput) (models.MandatoryLuggage, error)
	Update(inputID inputs.GetMandatoryLuggageDetailInput, input inputs.MandatoryLuggageInput) (models.MandatoryLuggage, error)
	Delete(inputID inputs.GetMandatoryLuggageDetailInput) (models.MandatoryLuggage, error)
}

// MandatoryLuggageService MandatoryLuggageService struct
type mandatoryLuggageService struct {
	repository repositories.MandatoryLuggageRepository
}

// NewMandatoryLuggageService : returns the MandatoryLuggageService struct instance
func NewMandatoryLuggageService(repository repositories.MandatoryLuggageRepository) mandatoryLuggageService {
	return mandatoryLuggageService{repository}
}

// FindAll -> calls Entertainment Package repo find all method
func (s mandatoryLuggageService) FindAll(model models.MandatoryLuggage, search string) ([]models.MandatoryLuggage, int64, error) {
	mandatoryLuggages, total, err := s.repository.FindAll(model, search)
	if err != nil {
		return mandatoryLuggages, total, err
	}

	return mandatoryLuggages, total, nil
}

// Find -> calls Entertainment Package repo find method
func (s mandatoryLuggageService) Find(input inputs.GetMandatoryLuggageDetailInput) (models.MandatoryLuggage, error) {
	mandatoryLuggage, err := s.repository.Find(input.ID)

	if err != nil {
		return mandatoryLuggage, err
	}

	return mandatoryLuggage, nil
}

// Save -> calls Entertainment Package repository save method
func (s mandatoryLuggageService) Save(input inputs.MandatoryLuggageInput) (models.MandatoryLuggage, error) {
	mandatoryLuggage := models.MandatoryLuggage{}
	mandatoryLuggage.Name = input.Name

	newMandatoryLuggage, err := s.repository.Save(mandatoryLuggage)
	if err != nil {
		return newMandatoryLuggage, err
	}

	return newMandatoryLuggage, nil
}

// Update -> calls Entertainment Package repo update method
func (s mandatoryLuggageService) Update(inputID inputs.GetMandatoryLuggageDetailInput, input inputs.MandatoryLuggageInput) (models.MandatoryLuggage, error) {
	mandatoryLuggage, err := s.repository.Find(inputID.ID)
	if err != nil {
		return mandatoryLuggage, err
	}

	mandatoryLuggage.Name = input.Name

	updatedMandatoryLuggage, err := s.repository.Update(mandatoryLuggage)
	if err != nil {
		return updatedMandatoryLuggage, err
	}

	return updatedMandatoryLuggage, nil
}

// Delete -> calls Entertainment Package repo delete method
func (s mandatoryLuggageService) Delete(inputID inputs.GetMandatoryLuggageDetailInput) (models.MandatoryLuggage, error) {
	mandatoryLuggage, err := s.repository.Find(inputID.ID)
	if err != nil {
		return mandatoryLuggage, err
	}

	deletedMandatoryLuggage, err := s.repository.Delete(mandatoryLuggage)
	if err != nil {
		return deletedMandatoryLuggage, err
	}

	return deletedMandatoryLuggage, nil
}
