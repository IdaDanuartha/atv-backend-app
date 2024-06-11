package repositories

import (
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type CustomerRepository interface {
	FindAll(customer models.Customer, search string, currentPage int, pageSize int) ([]models.Customer, int64, int, error)
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
func (r customerRepository) FindAll(customer models.Customer, search string, currentPage int, pageSize int) ([]models.Customer, int64, int, error) {
	var customers []models.Customer
	var totalRows int64 = 0

	queryBuilder := r.db.DB.Order("created_at desc").Model(&models.Customer{})

	// Search parameter
	if search != "" {
		querySearch := "%" + search + "%"
		queryBuilder = queryBuilder.Where(
			r.db.DB.Where("customers.name LIKE ? ", querySearch))
	}

	if pageSize > 0 {
		// count the total number of rows
		err := queryBuilder.
			Where(customer).
			Count(&totalRows).Error

		// Apply offset and limit to fetch paginated results
		err = queryBuilder.
			Preload("User").
			Where(customer).
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).
			Find(&customers).Error
		return customers, totalRows, currentPage, err
	} else {
		err := queryBuilder.
			Preload("User").
			Where(customer).
			Find(&customers).
			Count(&totalRows).Error

		return customers, 0, 0, err
	}
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
