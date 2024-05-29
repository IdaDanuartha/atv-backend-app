package formatters

import "github.com/IdaDanuartha/atv-backend-app/app/models"

func FormatCustomer(customer models.Customer) models.Customer {
	customerFormatter := models.Customer{}
	customerFormatter.ID = customer.ID
	customerFormatter.Name = customer.Name
	customerFormatter.CreatedAt = customer.CreatedAt
	customerFormatter.UpdatedAt = customer.UpdatedAt
	customerFormatter.DeletedAt = customer.DeletedAt

	customerFormatter.User.ID = customer.User.ID
	customerFormatter.User.Username = customer.User.Username
	customerFormatter.User.Email = customer.User.Email
	customerFormatter.User.Password = customer.User.Password
	customerFormatter.User.Role = customer.User.Role
	customerFormatter.User.CreatedAt = customer.User.CreatedAt
	customerFormatter.User.UpdatedAt = customer.User.UpdatedAt
	customerFormatter.User.DeletedAt = customer.User.DeletedAt

	return customerFormatter
}

func FormatCustomers(customers []models.Customer) []models.Customer {
	customersFormatter := []models.Customer{}

	for _, customer := range customers {
		customer := FormatCustomer(customer)
		customersFormatter = append(customersFormatter, customer)
	}

	return customersFormatter
}
