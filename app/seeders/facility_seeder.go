package seeders

//import (
//	"github.com/IdaDanuartha/atv-backend-app/app/models"
//	"gorm.io/gorm"
//	"log"
//)
//
//func FacilitySeeder(db *gorm.DB) {
//	// AutoMigrate the Facility model
//	if err := db.AutoMigrate(&models.Facility{}).Error; err != nil {
//		log.Fatalf("Error migrating Facility model: %v", err)
//	}
//
//	// Seed data into the Facility table
//	facilities := []models.Facility{
//		{Name: "Fasilitas 1"},
//		{Name: "Fasilitas 2"},
//		{Name: "Fasilitas 3"},
//	}
//
//	for _, f := range facilities {
//		if err := db.Create(&f).Error; err != nil {
//			log.Fatalf("Error seeding Facility: %v", err)
//		}
//	}
//
//	log.Println("Facility seeding completed successfully.")
//}
