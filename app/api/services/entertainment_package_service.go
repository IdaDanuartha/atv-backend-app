package services

import (
	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type EntertainmentPackageService interface {
	FindAll(model models.EntertainmentPackage, search string, currentPage int, pageSize int) ([]models.EntertainmentPackage, int64, int, error)
	Find(input inputs.GetEntertainmentPackageDetailInput) (models.EntertainmentPackage, error)
	SaveImage(ID string, fileLocation string) (models.EntertainmentPackage, error)
	Save(input inputs.EntertainmentPackageInput) (models.EntertainmentPackage, error)
	Update(inputID inputs.GetEntertainmentPackageDetailInput, input inputs.EntertainmentPackageInput) (models.EntertainmentPackage, error)
	Delete(inputID inputs.GetEntertainmentPackageDetailInput) (models.EntertainmentPackage, error)
}

// EntertainmentPackageService EntertainmentPackageService struct
type entertainmentPackageService struct {
	repository repositories.EntertainmentPackageRepository
}

// NewEntertainmentPackageService : returns the EntertainmentPackageService struct instance
func NewEntertainmentPackageService(repository repositories.EntertainmentPackageRepository) entertainmentPackageService {
	return entertainmentPackageService{repository}
}

// FindAll -> calls Entertainment Package repo find all method
func (s entertainmentPackageService) FindAll(model models.EntertainmentPackage, search string, currentPage int, pageSize int) ([]models.EntertainmentPackage, int64, int, error) {
	entertainmentPackages, total, currentPage, err := s.repository.FindAll(model, search, currentPage, pageSize)
	if err != nil {
		return entertainmentPackages, total, currentPage, err
	}

	return entertainmentPackages, total, currentPage, nil
}

// Find -> calls Entertainment Package repo find method
func (s entertainmentPackageService) Find(input inputs.GetEntertainmentPackageDetailInput) (models.EntertainmentPackage, error) {
	entertainmentPackage, err := s.repository.Find(input.ID, true)

	if err != nil {
		return entertainmentPackage, err
	}

	return entertainmentPackage, nil
}

func (s entertainmentPackageService) SaveImage(ID string, fileLocation string) (models.EntertainmentPackage, error) {
	entertainment_package, err := s.repository.Find(ID, false)
	if err != nil {
		return entertainment_package, err
	}

	entertainment_package.ImagePath = &fileLocation

	updatedEntertainmentPackage, err := s.repository.Update(entertainment_package)
	if err != nil {
		return updatedEntertainmentPackage, err
	}

	return updatedEntertainmentPackage, nil
}

// Save -> calls Entertainment Package repository save method
func (s entertainmentPackageService) Save(input inputs.EntertainmentPackageInput) (models.EntertainmentPackage, error) {
	entertainmentPackage := models.EntertainmentPackage{}

	entertainmentPackage.Name = input.Name
	entertainmentPackage.Description = input.Description
	entertainmentPackage.Duration = input.Duration
	entertainmentPackage.Price = input.Price
	entertainmentPackage.ExpiredAt = input.ExpiredAt
	entertainmentPackage.Services = input.Services

	newEntertainmentPackage, err := s.repository.Save(entertainmentPackage)
	if err != nil {
		return newEntertainmentPackage, err
	}

	return newEntertainmentPackage, nil
}

// Update -> calls Entertainment Package repo update method
func (s entertainmentPackageService) Update(inputID inputs.GetEntertainmentPackageDetailInput, input inputs.EntertainmentPackageInput) (models.EntertainmentPackage, error) {
	entertainmentPackage, err := s.repository.Find(inputID.ID, false)
	if err != nil {
		return entertainmentPackage, err
	}

	entertainmentPackage.Name = input.Name
	entertainmentPackage.Description = input.Description
	entertainmentPackage.Duration = input.Duration
	entertainmentPackage.Price = input.Price
	entertainmentPackage.ExpiredAt = input.ExpiredAt
	entertainmentPackage.Services = input.Services

	updatedEntertainmentPackage, err := s.repository.Update(entertainmentPackage)
	if err != nil {
		return updatedEntertainmentPackage, err
	}

	return updatedEntertainmentPackage, nil
}

// Delete -> calls Entertainment Package repo delete method
func (s entertainmentPackageService) Delete(inputID inputs.GetEntertainmentPackageDetailInput) (models.EntertainmentPackage, error) {
	entertainmentPackage, err := s.repository.Find(inputID.ID, true)
	if err != nil {
		return entertainmentPackage, err
	}

	deletedEntertainmentPackage, err := s.repository.Delete(entertainmentPackage)
	if err != nil {
		return deletedEntertainmentPackage, err
	}

	return deletedEntertainmentPackage, nil
}
