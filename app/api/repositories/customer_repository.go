package repositories

import (
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type CustomerRepository interface {
	FindAll(customer models.Customer, search string) ([]models.Customer, int64, error)
	Find(ID string) (models.Customer, error)
}

type customerRepository struct {
	db config.Database
}

// NewCustomerRepository : fetching database
func NewCustomerRepository(db config.Database) customerRepository {
	return customerRepository{db}
}

// FindAll -> Method for fetching all Customer from database
func (r customerRepository) FindAll(customer models.Customer, search string) ([]models.Customer, int64, error) {
	var customers []models.Customer
	var totalRows int64 = 0

	queryBuider := r.db.DB.Order("created_at desc").Model(&models.Customer{})

	// Search parameter
	if search != "" {
		querySearch := "%" + search + "%"
		queryBuider = queryBuider.Where(
			r.db.DB.Where("customers.name LIKE ? ", querySearch))
	}

	err := queryBuider.
		Preload("User").
		Where(customer).
		Find(&customers).
		Count(&totalRows).Error
	return customers, totalRows, err
}

// Find -> Method for fetching Customer by id
func (r customerRepository) Find(ID string) (models.Customer, error) {
	var customer models.Customer
	err := r.db.DB.
		Preload("User").
		Debug().
		Model(&models.Customer{}).
		Where("id = ?", ID).
		Find(&customer).Error
	return customer, err
}
