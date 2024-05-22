package services

import (
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

// MandatoryLuggageService MandatoryLuggageService struct
type MandatoryLuggageService struct {
	repository repositories.MandatoryLuggageRepository
}

// NewMandatoryLuggageService : returns the MandatoryLuggageService struct instance
func NewMandatoryLuggageService(r repositories.MandatoryLuggageRepository) MandatoryLuggageService {
	return MandatoryLuggageService{
		repository: r,
	}
}

// Save -> calls Mandatory Luggage repository save method
func (p MandatoryLuggageService) Save(mandatoryLuggage models.MandatoryLuggage) error {
	return p.repository.Save(mandatoryLuggage)
}

// FindAll -> calls Mandatory Luggage repo find all method
func (p MandatoryLuggageService) FindAll(mandatoryLuggage models.MandatoryLuggage, search string) (*[]models.MandatoryLuggage, int64, error) {
	return p.repository.FindAll(mandatoryLuggage, search)
}

// Update -> calls Mandatory Luggage repo update method
func (p MandatoryLuggageService) Update(mandatoryLuggage models.MandatoryLuggage) error {
	return p.repository.Update(mandatoryLuggage)
}

// Delete -> calls Mandatory Luggage repo delete method
func (p MandatoryLuggageService) Delete(id string) error {
	var mandatoryLuggage models.MandatoryLuggage
	mandatoryLuggage.ID = id
	return p.repository.Delete(mandatoryLuggage)
}

// Find -> calls Mandatory Luggage repo find method
func (p MandatoryLuggageService) Find(mandatoryLuggage models.MandatoryLuggage) (models.MandatoryLuggage, error) {
	return p.repository.Find(mandatoryLuggage)
}
