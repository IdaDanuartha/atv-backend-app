package seeders

import (
	"log"

	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"gorm.io/gorm"
)

func SeedRoute(db *gorm.DB) {
	// AutoMigrate the Route model
	db.AutoMigrate(&models.Route{})

	// Seed data into the Route table
	routes := []models.Route{
		{
			Name: "Basecamp",
			Address: "Address of the Basecamp",
		},
		{
			Name: "Joyful Forest",
			Address: "Address of the Joyful Forest",
		},
		{
			Name: "Seri River",
			Address: "Address of the Seri River",
		},
		{
			Name: "Root Bridge",
			Address: "Address of the Root Bridge",
		},
		{
			Name: "Green Peak",
			Address: "Address of the Green Peak",
		},
	}

	for _, route := range routes {
		var existingRoute models.Route
		// Check if the route already exists
		db.Where("name = ?", route.Name).Find(&existingRoute)
		if existingRoute.ID == "" {
			// If the route does not exist, create it
			if err := db.Create(&route).Error; err != nil {
				log.Fatalf("Error seeding route: %v", err)
			}
		}
	}

	log.Println("Route seeding completed successfully.")
}
