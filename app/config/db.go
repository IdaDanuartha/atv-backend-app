package config

import (
	"fmt"
	// "log"
	"os"

	// "golang.org/x/crypto/bcrypt"

	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"github.com/IdaDanuartha/atv-backend-app/app/seeders"
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

	db, err := gorm.Open(mysql.Open(URL), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")

	}

	db.AutoMigrate(&models.User{})
	seeders.SeedAdmin(db)

	db.AutoMigrate(&models.Staff{})
	db.AutoMigrate(&models.Instructor{})
	db.AutoMigrate(&models.Customer{})
	seeders.SeedFacility(db)
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
	db.AutoMigrate(&models.Blog{})

	fmt.Println("Database connection established")
	return Database{
		DB: db,
	}

}
