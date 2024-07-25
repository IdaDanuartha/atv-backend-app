package seeders

import (
	"log"
	"time"

	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"gorm.io/gorm"
)

func SeedEntertainmentPackage(db *gorm.DB) {
	// AutoMigrate the Entertainment Package model and the relations
	db.AutoMigrate(&models.EntertainmentPackage{})
	db.AutoMigrate(&models.EntertainmentPackageDetail{})

	var services []models.EntertainmentService
	db.Find(&services)

	// Seed data into the Entertainment Package table
	entertainmentPackages := []models.EntertainmentPackage{
		{
			Name:                    "ATV + Cycling Package",
			Description:             "Description of the ATV + Cycling Package",
			Price:                   650000,
			Duration:                240,
			ExpiredAt: time.Now().AddDate(0, 1, 0), // to get time next month (year, month, date)
			Services: []models.EntertainmentPackageDetail{
				{
					EntertainmentServiceID: services[0].ID,
				},
				{
					EntertainmentServiceID: services[1].ID,
				},
			},
		},
	}

	for _, entertainmentPackage := range entertainmentPackages {
		var existingPackage models.EntertainmentPackage
		// Check if the package already exists
		db.Where("name = ?", entertainmentPackage.Name).Find(&existingPackage)
		if existingPackage.ID == "" {
			// If the package does not exist, create it
			if err := db.Create(&entertainmentPackage).Error; err != nil {
				log.Fatalf("Error seeding entertainment package: %v", err)
			}
		}
	}

	log.Println("Entertainment package seeding completed successfully.")
}
