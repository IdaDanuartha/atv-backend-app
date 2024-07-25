package seeders

import (
	"log"

	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"gorm.io/gorm"
)

func SeedMandatoryLuggage(db *gorm.DB) {
	// AutoMigrate the Mandatory luggage model
	db.AutoMigrate(&models.MandatoryLuggage{})

	// Seed data into the Mandatory luggage table
	mandatoryLuggages := []models.MandatoryLuggage{
		{Name: "Safe & Closed-Toe Shoes"},
		{Name: "Gloves"},
		{Name: "Basic Medications"},
		{Name: "Glasses"},
		{Name: "Sunscreen"},
	}

	for _, mandatoryLuggage := range mandatoryLuggages {
		var existingMandatoryLuggage models.MandatoryLuggage
		// Check if the mandatory luggage already exists
		db.Where("name = ?", mandatoryLuggage.Name).Find(&existingMandatoryLuggage)
		if existingMandatoryLuggage.ID == "" {
			// If the mandatory luggage does not exist, create it
			if err := db.Create(&mandatoryLuggage).Error; err != nil {
				log.Fatalf("Error seeding mandatory luggage: %v", err)
			}
		}
	}

	log.Println("Mandatory luggage seeding completed successfully.")
}
