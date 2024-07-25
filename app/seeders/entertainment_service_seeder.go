package seeders

import (
	"log"

	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"gorm.io/gorm"
)

func SeedEntertainmentService(db *gorm.DB) {
	// AutoMigrate the Entertainment Service model and the relations
	db.AutoMigrate(&models.EntertainmentService{})
	db.AutoMigrate(&models.EntertainmentServiceRoute{})
	db.AutoMigrate(&models.EntertainmentServiceFacility{})
	db.AutoMigrate(&models.EntertainmentServiceInstructor{})
	db.AutoMigrate(&models.MandatoryLuggageEntertainmentService{})

	// Seed data into the Entertainment Service table
	entertainmentServices := []models.EntertainmentService{
		{
			Name:                    "",
			Price:                   0,
			Duration:                0,
			Description:             "",
			EntertainmentCategoryID: "",
		},
	}

	for _, service := range entertainmentServices {
		var existingService models.EntertainmentService
		// Check if the service already exists
		db.Where("name = ?", service.Name).Find(&existingService)
		if existingService.ID == "" {
			// If the service does not exist, create it
			if err := db.Create(&service).Error; err != nil {
				log.Fatalf("Error seeding entertainment service: %v", err)
			}
		}
	}

	log.Println("Entertainment service seeding completed successfully.")
}
