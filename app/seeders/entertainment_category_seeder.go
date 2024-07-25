package seeders

import (
	"log"

	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"gorm.io/gorm"
)

func SeedEntertainmentCategory(db *gorm.DB) {
	// AutoMigrate the Entertainment Category model
	db.AutoMigrate(&models.EntertainmentCategory{})

	// Seed data into the Facility table
	entertainmentCategories := []models.EntertainmentCategory{
		{Name: "ATV"},
		{Name: "Water Cycling"},
		{Name: "Hiking"},
		{Name: "Flying Fox"},
		{Name: "Rafting"},
	}

	for _, category := range entertainmentCategories {
		var existingCategory models.EntertainmentCategory
		// Check if the category already exists
		db.Where("name = ?", category.Name).Find(&existingCategory)
		if existingCategory.ID == "" {
			// If the category does not exist, create it
			if err := db.Create(&category).Error; err != nil {
				log.Fatalf("Error seeding entertainment category: %v", err)
			}
		}
	}

	log.Println("Entertainment category seeding completed successfully.")
}
