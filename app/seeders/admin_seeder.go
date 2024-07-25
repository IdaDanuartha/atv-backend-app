package seeders

import (
	"log"

	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	// db.AutoMigrate(&models.Admin{})

	// var adminCount int64
	// var admin models.Admin
	// err := db.Model(&models.Admin{}).Where(admin).Count(&adminCount).Error

	// if err != nil {
	// 	log.Fatalf("Failed to count admins: %v", err)
	// }

	// if adminCount == 0 {
	// 	// Admin count is 0, create a new user and admin
	// 	passwordHash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	// 	if err != nil {
	// 		log.Fatalf("Failed to hash password: %v", err)
	// 	}

	// 	user := models.User{
	// 		Username: "admin1",
	// 		Email:    "admin1@gmail.com",
	// 		Password: string(passwordHash),
	// 		Role:     "admin",
	// 	}

	// 	admin := models.Admin{
	// 		Name: "Admin",
	// 		User: user,
	// 	}

	// 	// Save the user and admin in the database
	// 	result := db.Create(&admin)
	// 	if result.Error != nil {
	// 		log.Fatalf("Failed to create admin: %v", result.Error)
	// 	}
	// }

	// AutoMigrate the Admin model
	db.AutoMigrate(&models.Admin{})
	passwordHash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	// Seed data into the Admin table
	admins := []models.Admin{
		{
			Name: "Admin", 
			User: models.User{
				Username: "admin1",
				Email:    "admin1@gmail.com",
				Password: string(passwordHash),
				Role:     "admin",
			},
		},
	}

	for _, admin := range admins {
		var existingAdmin models.Admin
		// Check if the admin already exists
		db.Where("name = ?", admin.Name).Find(&existingAdmin)
		if existingAdmin.ID == "" {
			// If the admin does not exist, create it
			if err := db.Create(&admin).Error; err != nil {
				log.Fatalf("Error seeding admin: %v", err)
			}
		}
	}

	log.Println("Admin seeding completed successfully.")
}