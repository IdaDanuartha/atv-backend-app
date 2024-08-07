package seeders

import (
	"log"
	"time"

	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"github.com/IdaDanuartha/atv-backend-app/app/utils"
	"gorm.io/gorm"
)

func SeedBooking(db *gorm.DB) {
	// AutoMigrate the Booking model and the relations
	db.AutoMigrate(&models.Booking{})
	db.AutoMigrate(&models.BookingDetail{})

	var services []models.EntertainmentService
	db.Find(&services)

	var customer models.Customer
	db.First(&customer)

	// Seed data into the Booking table
	bookings := []models.Booking{
		{
			CustomerID:    customer.ID,
			Code:          utils.GenerateFormattedString(),
			Name:          customer.Name,
			PhoneNumber:   customer.PhoneNumber,
			PaymentMethod: "BCA VA",
			Date:          time.Now(),
			Details: []models.BookingDetail{
				{
					EntertainmentServiceID: services[0].ID,
					Price:                  services[0].Price,
					Qty:                    1,
				},
				{
					EntertainmentServiceID: services[1].ID,
					Price:                  services[0].Price,
					Qty:                    1,
				},
			},
		},
	}

	for _, booking := range bookings {
		var existingBooking models.Booking
		// Check if the package already exists
		db.Where("name = ?", booking.Name).Find(&existingBooking)
		if existingBooking.ID == "" {
			// If the package does not exist, create it
			if err := db.Create(&booking).Error; err != nil {
				log.Fatalf("Error seeding booking: %v", err)
			}
		}
	}

	log.Println("Booking seeding completed successfully.")
}
