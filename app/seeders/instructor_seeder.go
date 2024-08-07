package seeders

import (
	"log"

	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedInstructor(db *gorm.DB) {
	// AutoMigrate the Instructor model
	db.AutoMigrate(&models.Instructor{})
	passwordHash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	// Seed data into the Staff table
	instructors := []models.Instructor{
		{
			Name: "Instructor 1",
			EmployeeCode: "INS123",
			User: models.User{
				Username: "instructor1",
				Email:    "instructor1@gmail.com",
				Password: string(passwordHash),
				Role:     "instructor",
			},
		},
		{
			Name: "Instructor 2",
			EmployeeCode: "INS234",
			User: models.User{
				Username: "instructor2",
				Email:    "instructor2@gmail.com",
				Password: string(passwordHash),
				Role:     "instructor",
			},
		},
	}

	for _, instructor := range instructors {
		var existingInstructor models.Instructor
		// Check if the instructor already exists
		db.Where("name = ?", instructor.Name).Find(&existingInstructor)
		if existingInstructor.ID == "" {
			// If the instructor does not exist, create it
			if err := db.Create(&instructor).Error; err != nil {
				log.Fatalf("Error seeding instructor: %v", err)
			}
		}
	}

	log.Println("Instructor seeding completed successfully.")
}
