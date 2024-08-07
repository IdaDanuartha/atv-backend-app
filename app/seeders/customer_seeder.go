package seeders

import (
	"log"

	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedCustomer(db *gorm.DB) {
	// AutoMigrate the Customer model
	db.AutoMigrate(&models.Customer{})
	passwordHash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	// Seed data into the Staff table
	customers := []models.Customer{
		{
			Name:         "Customer 1",
			PhoneNumber: "08123456789",
			User: models.User{
				Username: "customer1",
				Email:    "customer1@gmail.com",
				Password: string(passwordHash),
				Role:     "customer",
			},
		},
		{
			Name:         "Customer 2",
			PhoneNumber: "08987654321",
			User: models.User{
				Username: "customer2",
				Email:    "customer2@gmail.com",
				Password: string(passwordHash),
				Role:     "customer",
			},
		},
	}

	for _, customer := range customers {
		var existingCustomer models.Customer
		// Check if the customer already exists
		db.Where("name = ?", customer.Name).Find(&existingCustomer)
		if existingCustomer.ID == "" {
			// If the customer does not exist, create it
			if err := db.Create(&customer).Error; err != nil {
				log.Fatalf("Error seeding customer: %v", err)
			}
		}
	}

	log.Println("Customer seeding completed successfully.")
}
