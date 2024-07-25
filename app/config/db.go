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

	// users
	db.AutoMigrate(&models.User{})
	seeders.SeedAdmin(db)
	seeders.SeedStaff(db)
	seeders.SeedInstructor(db)
	seeders.SeedCustomer(db)
	// master data
	seeders.SeedFacility(db)
	seeders.SeedMandatoryLuggage(db)
	seeders.SeedRoute(db)
	seeders.SeedEntertainmentCategory(db)
	seeders.SeedEntertainmentService(db)
	// package
	seeders.SeedEntertainmentPackage(db)
	// booking
	seeders.SeedBooking(db)
	// blog
	seeders.SeedBlog(db)

	fmt.Println("Database connection established")
	return Database{
		DB: db,
	}

}
