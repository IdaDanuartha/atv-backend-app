package seeders

import (
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"gorm.io/gorm"
	"log"
)

func SeedFacility(db *gorm.DB) {
	// AutoMigrate the Facility model
	db.AutoMigrate(&models.Facility{})

	// Seed data into the Facility table
	facilities := []models.Facility{
		{Name: "Latest ATVs with Well-Maintained Physical Conditions"},
		{Name: "High-quality mountain bikes"},
		{Name: "Restrooms, rest areas, and picnic spots"},
		{Name: "Snacks and Drinks: Refreshments to keep you energized"},
		{Name: "Comprehensive and Reliable Safety Gear"},
	}

	for _, facility := range facilities {
		var existingFacility models.Facility
		// Check if the facility already exists
		db.Where("name = ?", facility.Name).Find(&existingFacility)
		if existingFacility.ID == "" {
			// If the facility does not exist, create it
			if err := db.Create(&facility).Error; err != nil {
				log.Fatalf("Error seeding Facility: %v", err)
			}
		}
	}

	log.Println("Facility seeding completed successfully.")
}
