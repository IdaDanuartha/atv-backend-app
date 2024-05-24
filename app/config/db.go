package config

import (
	"fmt"
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Database struct
type Database struct {
    DB *gorm.DB
}

//NewDatabase : intializes and returns mysql db
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
	db.AutoMigrate(&models.Staff{})
	db.AutoMigrate(&models.Instructor{})
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.EntertainmentCategory{})
    db.AutoMigrate(&models.EntertainmentPackage{})
    db.AutoMigrate(&models.Facility{})
    db.AutoMigrate(&models.MandatoryLuggage{})

    fmt.Println("Database connection established")
    return Database{
        DB: db,
    }

}