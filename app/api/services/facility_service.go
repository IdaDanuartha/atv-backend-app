package services

import (
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

// FacilityService FacilityService struct
type FacilityService struct {
	repository repositories.FacilityRepository
}

// NewFacilityService : returns the FacilityService struct instance
func NewFacilityService(r repositories.FacilityRepository) FacilityService {
	return FacilityService{
		repository: r,
	}
}

// Save -> calls Facility repository save method
func (p FacilityService) Save(facility models.Facility) error {
	return p.repository.Save(facility)
}

// FindAll -> calls Facility repo find all method
func (p FacilityService) FindAll(facility models.Facility, search string) (*[]models.Facility, int64, error) {
	return p.repository.FindAll(facility, search)
}

// Update -> calls Facility repo update method
func (p FacilityService) Update(facility models.Facility) error {
	return p.repository.Update(facility)
}

// Delete -> calls Facility repo delete method
func (p FacilityService) Delete(id string) error {
	var facility models.Facility
	facility.ID = id
	return p.repository.Delete(facility)
}

// Find -> calls Facility repo find method
func (p FacilityService) Find(facility models.Facility) (models.Facility, error) {
	return p.repository.Find(facility)
}
