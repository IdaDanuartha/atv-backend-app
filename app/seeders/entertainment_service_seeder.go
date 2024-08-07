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

	var category models.EntertainmentCategory
	db.First(&category)

	var route models.Route
	db.First(&route)

	var facility models.Facility
	db.First(&facility)

	var instructor models.Instructor
	db.First(&instructor)

	var mandatoryLuggage models.MandatoryLuggage
	db.First(&mandatoryLuggage)

	// Seed data into the Entertainment Service table
	entertainmentServices := []models.EntertainmentService{
		{
			Name:                    "ATV Adventure",
			Price:                   400000,
			Duration:                120,
			Description:             "Description of the ATV Adventure",
			EntertainmentCategoryID: category.ID,
			Routes: []models.EntertainmentServiceRoute{
				{
					RouteID: route.ID,
				},
			},
			Facilities: []models.EntertainmentServiceFacility{
				{
					FacilityID: facility.ID,
				},
			},
			Instructors: []models.EntertainmentServiceInstructor{
				{
					InstructorID: instructor.ID,
				},
			},
			MandatoryLuggages: []models.MandatoryLuggageEntertainmentService{
				{
					MandatoryLuggageID: mandatoryLuggage.ID,
				},
			},
		},
		{
			Name:                    "ATV Extreme",
			Price:                   500000,
			Duration:                120,
			Description:             "Description of the ATV Extreme",
			EntertainmentCategoryID: category.ID,
			Routes: []models.EntertainmentServiceRoute{
				{
					RouteID: route.ID,
				},
			},
			Facilities: []models.EntertainmentServiceFacility{
				{
					FacilityID: facility.ID,
				},
			},
			Instructors: []models.EntertainmentServiceInstructor{
				{
					InstructorID: instructor.ID,
				},
			},
			MandatoryLuggages: []models.MandatoryLuggageEntertainmentService{
				{
					MandatoryLuggageID: mandatoryLuggage.ID,
				},
			},
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
