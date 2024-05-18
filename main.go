package main

import (
	"fmt"
	"os"

	"github.com/IdaDanuartha/teman-bus-backend-app/app/api/controllers"
	"github.com/IdaDanuartha/teman-bus-backend-app/app/api/repositories"
	"github.com/IdaDanuartha/teman-bus-backend-app/app/api/routes"
	"github.com/IdaDanuartha/teman-bus-backend-app/app/api/services"
	"github.com/IdaDanuartha/teman-bus-backend-app/app/config"
	"github.com/IdaDanuartha/teman-bus-backend-app/app/models"
)

func init() {
	config.LoadEnv()
}

func main() {
	router := config.NewGinRouter()                           //router has been initialized and configured
	db := config.NewDatabase()                                // databse has been initialized and configured
	busRepository := repositories.NewBusRepository(db)        // repository are being setup
	busService := services.NewBusService(busRepository)       // service are being setup
	busController := controllers.NewBusController(busService) // controller are being set up
	busRoute := routes.NewBusRoute(busController, router)     // bus routes are initialized
	busRoute.Setup()                                          // bus routes are being setup

	db.DB.AutoMigrate(&models.Bus{}) // migrating Bus model to datbase table

	router.Gin.Run(":" + os.Getenv("APP_PORT"))

	fmt.Println("App running in port: ", os.Getenv("APP_PORT"))
}
