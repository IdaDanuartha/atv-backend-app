package services

import (
	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type CustomerService interface {
	FindAll(model models.Customer, search string) ([]models.Customer, int64, error)
	Find(input inputs.GetCustomerDetailInput) (models.Customer, error)
}

// CustomerService CustomerService struct
type customerService struct {
	repository repositories.CustomerRepository
}

// NewCustomerService : returns the CustomerService struct instance
func NewCustomerService(repository repositories.CustomerRepository) customerService {
	return customerService{repository}
}

// FindAll -> calls Customer repo find all method
func (s customerService) FindAll(model models.Customer, search string) ([]models.Customer, int64, error) {
	customers, total, err := s.repository.FindAll(model, search)
	if err != nil {
		return customers, total, err
	}

	return customers, total, nil
}

// Find -> calls Customer repo find method
func (s customerService) Find(input inputs.GetCustomerDetailInput) (models.Customer, error) {
	customer, err := s.repository.Find(input.ID)

	if err != nil {
		return customer, err
	}

	return customer, nil
}