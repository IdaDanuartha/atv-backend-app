package seeders

import (
	"log"

	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedStaff(db *gorm.DB) {
	// AutoMigrate the Staff model
	db.AutoMigrate(&models.Staff{})
	passwordHash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	// Seed data into the Staff table
	staffs := []models.Staff{
		{
			Name: "Staff 1",
			EmployeeCode: "ST123",
			User: models.User{
				Username: "staff1",
				Email:    "staff1@gmail.com",
				Password: string(passwordHash),
				Role:     "staff",
			},
		},
		{
			Name: "Staff 2",
			EmployeeCode: "ST234",
			User: models.User{
				Username: "staff2",
				Email:    "staff2@gmail.com",
				Password: string(passwordHash),
				Role:     "staff",
			},
		},
	}

	for _, staff := range staffs {
		var existingStaff models.Staff
		// Check if the staff already exists
		db.Where("name = ?", staff.Name).Find(&existingStaff)
		if existingStaff.ID == "" {
			// If the staff does not exist, create it
			if err := db.Create(&staff).Error; err != nil {
				log.Fatalf("Error seeding staff: %v", err)
			}
		}
	}

	log.Println("Staff seeding completed successfully.")
}
