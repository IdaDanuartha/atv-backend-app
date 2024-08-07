package seeders

import (
	"log"

	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"gorm.io/gorm"
)

func SeedBlog(db *gorm.DB) {
	// AutoMigrate the Blog model and the relations
	db.AutoMigrate(&models.Blog{})

	// Seed data into the Booking table
	blogs := []models.Blog{
		{
			Title: "Beginner's Guide: Starting Water Sports Safely and Excitingly",
			Slug: "beginners-guide-starting-water-sports-safely-and-excitingly",
			Description: "A brief description of the blog",
			Content: "Content of the blog",
		},
		{
			Title: "Water Sports Safety: Essential Tips and Tricks for Safe Adventures",
			Slug: "water-sports-safety-essential-tips-and-tricks-for-safe-adventures",
			Description: "A brief description of the blog",
			Content: "Content of the blog",
		},
	}

	for _, blog := range blogs {
		var existingBlog models.Blog
		// Check if the package already exists
		db.Where("title = ?", blog.Title).Find(&existingBlog)
		if existingBlog.ID == "" {
			// If the package does not exist, create it
			if err := db.Create(&blog).Error; err != nil {
				log.Fatalf("Error seeding blog: %v", err)
			}
		}
	}

	log.Println("Blog seeding completed successfully.")
}
