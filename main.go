package main

import (
	"fmt"
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/api/controllers"
	"github.com/IdaDanuartha/atv-backend-app/app/api/middlewares"
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	"github.com/IdaDanuartha/atv-backend-app/app/api/routes"
	"github.com/IdaDanuartha/atv-backend-app/app/api/services"
	"github.com/IdaDanuartha/atv-backend-app/app/config"
)

func init() {
	config.LoadEnv()
}

func main() {
	router := config.NewGinRouter()
	db := config.NewDatabase()

	// User API
	authService := services.NewAuthService()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService, authService)
	userRoute := routes.NewUserRoute(userController, router)

	authMiddleware := middlewares.AuthMiddleware(authService, userService)
	userRoute.Setup(authMiddleware)


	// Master Data API
	entertainmentCategoryRepository := repositories.NewEntertainmentCategoryRepository(db)
	entertainmentCategoryService := services.NewEntertainmentCategoryService(&entertainmentCategoryRepository)
	entertainmentCategoryController := controllers.NewEntertainmentCategoryController(entertainmentCategoryService)
	entertainmentCategoryRoute := routes.NewEntertainmentCategoryRoute(*entertainmentCategoryController, router)
	entertainmentCategoryRoute.Setup(authMiddleware)

	entertainmentPackageRepository := repositories.NewEntertainmentPackageRepository(db)
	entertainmentPackageService := services.NewEntertainmentPackageService(&entertainmentPackageRepository)
	entertainmentPackageController := controllers.NewEntertainmentPackageController(entertainmentPackageService)
	entertainmentPackageRoute := routes.NewEntertainmentPackageRoute(*entertainmentPackageController, router)
	entertainmentPackageRoute.Setup(authMiddleware)
	
	facilityRepository := repositories.NewFacilityRepository(db)
	facilityService := services.NewFacilityService(&facilityRepository)
	facilityController := controllers.NewFacilityController(facilityService)
	facilityRoute := routes.NewFacilityRoute(*facilityController, router)
	facilityRoute.Setup(authMiddleware)

	mandatoryLuggageRepository := repositories.NewMandatoryLuggageRepository(db)
	mandatoryLuggageService := services.NewMandatoryLuggageService(&mandatoryLuggageRepository)
	mandatoryLuggageController := controllers.NewMandatoryLuggageController(mandatoryLuggageService)
	mandatoryLuggageRoute := routes.NewMandatoryLuggageRoute(*mandatoryLuggageController, router)
	mandatoryLuggageRoute.Setup(authMiddleware)

	routeRepository := repositories.NewRouteRepository(db)
	routeService := services.NewRouteService(&routeRepository)
	routeController := controllers.NewRouteController(routeService)
	routeRoute := routes.NewRuteRoute(*routeController, router)
	routeRoute.Setup(authMiddleware)

	// User Management
	instructorRepository := repositories.NewInstructorRepository(db)
    instructorService := services.NewInstructorService(&instructorRepository)
    instructorController := controllers.NewInstructorController(instructorService)
    instructorRoute := routes.NewInstructorRoute(*instructorController, router)
    instructorRoute.Setup(authMiddleware)
	

	router.Gin.Run(":" + os.Getenv("APP_PORT"))

	fmt.Println("App running in port: ", os.Getenv("APP_PORT"))
}
