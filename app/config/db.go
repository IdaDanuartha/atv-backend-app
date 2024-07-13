package config

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database struct
type Database struct {
	DB *gorm.DB
}

// NewDatabase : intializes and returns mysql db
func NewDatabase() Database {
	USER := os.Getenv("DB_USERNAME")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_DATABASE")

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS,
		HOST, PORT, DBNAME)
	fmt.Println(URL)

	db, err := gorm.Open(mysql.Open(URL))

	if err != nil {
		panic("Failed to connect to database!")

	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Admin{})

	var adminCount int64
	var admin models.Admin
	err = db.Model(&models.Admin{}).Where(admin).Count(&adminCount).Error

	if err != nil {
		log.Fatalf("Failed to count admins: %v", err)
	}

	if adminCount == 0 {
		// Admin count is 0, create a new user and admin
		passwordHash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}

		user := models.User{
			Username: "admin1",
			Email:    "admin1@gmail.com",
			Password: string(passwordHash),
			Role:     "admin",
		}

		admin := models.Admin{
			Name: "Admin",
			User: user,
		}

		// Save the user and admin in the database
		result := db.Create(&admin)
		if result.Error != nil {
			log.Fatalf("Failed to create admin: %v", result.Error)
		}
	}

	db.AutoMigrate(&models.Staff{})
	db.AutoMigrate(&models.Instructor{})
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Facility{})
	db.AutoMigrate(&models.MandatoryLuggage{})
	db.AutoMigrate(&models.Route{})
	db.AutoMigrate(&models.EntertainmentCategory{})
	db.AutoMigrate(&models.EntertainmentService{})
	db.AutoMigrate(&models.EntertainmentServiceRoute{})
	db.AutoMigrate(&models.EntertainmentServiceFacility{})
	db.AutoMigrate(&models.EntertainmentServiceInstructor{})
	db.AutoMigrate(&models.MandatoryLuggageEntertainmentService{})
	db.AutoMigrate(&models.EntertainmentPackage{})
	db.AutoMigrate(&models.EntertainmentPackageDetail{})
	db.AutoMigrate(&models.Booking{})
	db.AutoMigrate(&models.BookingDetail{})

	fmt.Println("Database connection established")
	return Database{
		DB: db,
	}

}
