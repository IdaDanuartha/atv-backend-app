package services

import (
	"strconv"
	"time"

	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type EntertainmentPackageService interface {
	FindAll(model models.EntertainmentPackage, search string) ([]models.EntertainmentPackage, int64, error)
	Find(input inputs.GetEntertainmentPackageDetailInput) (models.EntertainmentPackage, error)
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
func (s entertainmentPackageService) FindAll(model models.EntertainmentPackage, search string) ([]models.EntertainmentPackage, int64, error) {
	entertainmentPackages, total, err := s.repository.FindAll(model, search)
	if err != nil {
		return entertainmentPackages, total, err
	}

	return entertainmentPackages, total, nil
}

// Find -> calls Entertainment Package repo find method
func (s entertainmentPackageService) Find(input inputs.GetEntertainmentPackageDetailInput) (models.EntertainmentPackage, error) {
	entertainmentPackage, err := s.repository.Find(input.ID)

	if err != nil {
		return entertainmentPackage, err
	}

	return entertainmentPackage, nil
}

// Save -> calls Entertainment Package repository save method
func (s entertainmentPackageService) Save(input inputs.EntertainmentPackageInput) (models.EntertainmentPackage, error) {
	entertainmentPackage := models.EntertainmentPackage{}

	price, err := strconv.ParseInt(input.Price, 10, 32)
	if err!= nil {
        return entertainmentPackage, err
    }

	priceInt32 := int32(price)
	parseTime, err := time.Parse(time.RFC1123, input.ExpiredAt)
    if err != nil {
      panic(err)
    }

	entertainmentPackage.Name = input.Name
	entertainmentPackage.Description = input.Description
	entertainmentPackage.Price = priceInt32
	entertainmentPackage.ExpiredAt = parseTime

	newEntertainmentPackage, err := s.repository.Save(entertainmentPackage)
	if err != nil {
		return newEntertainmentPackage, err
	}

	return newEntertainmentPackage, nil
}

// Update -> calls Entertainment Package repo update method
func (s entertainmentPackageService) Update(inputID inputs.GetEntertainmentPackageDetailInput, input inputs.EntertainmentPackageInput) (models.EntertainmentPackage, error) {
	entertainmentPackage, err := s.repository.Find(inputID.ID)
	if err != nil {
		return entertainmentPackage, err
	}

	price, err := strconv.ParseInt(input.Price, 10, 32)
	if err!= nil {
        return entertainmentPackage, err
    }

	priceInt32 := int32(price)
	parseTime, err := time.Parse(time.RFC1123, input.ExpiredAt)
    if err != nil {
      panic(err)
    }

	entertainmentPackage.Name = input.Name
	entertainmentPackage.Description = input.Description
	entertainmentPackage.Price = priceInt32
	entertainmentPackage.ExpiredAt = parseTime

	updatedEntertainmentPackage, err := s.repository.Update(entertainmentPackage)
	if err != nil {
		return updatedEntertainmentPackage, err
	}

	return updatedEntertainmentPackage, nil
}

// Delete -> calls Entertainment Package repo delete method
func (s entertainmentPackageService) Delete(inputID inputs.GetEntertainmentPackageDetailInput) (models.EntertainmentPackage, error) {
	entertainmentPackage, err := s.repository.Find(inputID.ID)
	if err != nil {
		return entertainmentPackage, err
	}

	deletedEntertainmentPackage, err := s.repository.Delete(entertainmentPackage)
	if err != nil {
		return deletedEntertainmentPackage, err
	}

	return deletedEntertainmentPackage, nil
}
