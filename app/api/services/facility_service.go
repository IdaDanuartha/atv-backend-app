package services

import (
	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"github.com/gin-gonic/gin"
)

type FacilityService interface {
	FindAll(model models.Facility, search string, currentPage int, pageSize int) ([]models.Facility, int64, int, error)
	ExportToExcel(model models.Facility, ctx *gin.Context) error
	Find(input inputs.GetFacilityDetailInput) (models.Facility, error)
	Save(input inputs.FacilityInput) (models.Facility, error)
	Update(inputID inputs.GetFacilityDetailInput, input inputs.FacilityInput) (models.Facility, error)
	Delete(inputID inputs.GetFacilityDetailInput) (models.Facility, error)
}

// FacilityService FacilityService struct
type facilityService struct {
	repository repositories.FacilityRepository
}

// NewFacilityService : returns the FacilityService struct instance
func NewFacilityService(repository repositories.FacilityRepository) facilityService {
	return facilityService{repository}
}

// FindAll -> calls Facility repo find all method
func (s facilityService) FindAll(model models.Facility, search string, currentPage int, pageSize int) ([]models.Facility, int64, int, error) {
	facilities, total, currentPage, err := s.repository.FindAll(model, search, currentPage, pageSize)
	if err != nil {
		return facilities, total, currentPage, err
	}

	return facilities, total, currentPage, nil
}

// ExportToExcel -> calls Facility repo ExportToExcel method
func (s facilityService) ExportToExcel(model models.Facility, ctx *gin.Context) error {
	err := s.repository.ExportToExcel(model, ctx)
	if err != nil {
		return err
	}

	return nil
}

// Find -> calls Facility repo find method
func (s facilityService) Find(input inputs.GetFacilityDetailInput) (models.Facility, error) {
	facility, err := s.repository.Find(input.ID)

	if err != nil {
		return facility, err
	}

	return facility, nil
}

// Save -> calls Facility repository save method
func (s facilityService) Save(input inputs.FacilityInput) (models.Facility, error) {
	facility := models.Facility{}
	facility.Name = input.Name

	newFacility, err := s.repository.Save(facility)
	if err != nil {
		return newFacility, err
	}

	return newFacility, nil
}

// Update -> calls Facility repo update method
func (s facilityService) Update(inputID inputs.GetFacilityDetailInput, input inputs.FacilityInput) (models.Facility, error) {
	facility, err := s.repository.Find(inputID.ID)
	if err != nil {
		return facility, err
	}

	facility.Name = input.Name

	updatedFacility, err := s.repository.Update(facility)
	if err != nil {
		return updatedFacility, err
	}

	return updatedFacility, nil
}

// Delete -> calls Facility repo delete method
func (s facilityService) Delete(inputID inputs.GetFacilityDetailInput) (models.Facility, error) {
	facility, err := s.repository.Find(inputID.ID)
	if err != nil {
		return facility, err
	}

	deletedFacility, err := s.repository.Delete(facility)
	if err != nil {
		return deletedFacility, err
	}

	return deletedFacility, nil
}
