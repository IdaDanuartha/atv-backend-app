package services

import (
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

// EntertainmentPackageService EntertainmentPackageService struct
type EntertainmentPackageService struct {
	repository repositories.EntertainmentPackageRepository
}

// NewEntertainmentPackageService : returns the EntertainmentPackageService struct instance
func NewEntertainmentPackageService(r repositories.EntertainmentPackageRepository) EntertainmentPackageService {
	return EntertainmentPackageService{
		repository: r,
	}
}

// Save -> calls Entertainment Package repo save method
func (p EntertainmentPackageService) Save(entertainmentPackage models.EntertainmentPackage) error {
	return p.repository.Save(entertainmentPackage)
}

// FindAll -> calls Entertainment Package repo find all method
func (p EntertainmentPackageService) FindAll(entertainmentPackage models.EntertainmentPackage, search string) (*[]models.EntertainmentPackage, int64, error) {
	return p.repository.FindAll(entertainmentPackage, search)
}

// Update -> calls Entertainment Package repo update method
func (p EntertainmentPackageService) Update(entertainmentPackage models.EntertainmentPackage) error {
	return p.repository.Update(entertainmentPackage)
}

// Delete -> calls Entertainment Package repo delete method
func (p EntertainmentPackageService) Delete(id string) error {
	var entertainmentPackage models.EntertainmentPackage
	entertainmentPackage.ID = id
	return p.repository.Delete(entertainmentPackage)
}

// Find -> calls Entertainment Package repo find method
func (p EntertainmentPackageService) Find(entertainmentPackage models.EntertainmentPackage) (models.EntertainmentPackage, error) {
	return p.repository.Find(entertainmentPackage)
}
