package main

import (
	"fmt"
	"log"
	"net/http"
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
	routeRoute := routes.NewRouteRoute(*routeController, router)
	routeRoute.Setup(authMiddleware)

	entertainmentServiceRepository := repositories.NewEntertainmentServiceRepository(db)
	entertainmentServiceService := services.NewEntertainmentServiceService(&entertainmentServiceRepository)
	entertainmentServiceController := controllers.NewEntertainmentServiceController(entertainmentServiceService)
	entertainmentServiceRoute := routes.NewEntertainmentServiceRoute(*entertainmentServiceController, router)
	entertainmentServiceRoute.Setup(authMiddleware)
	
	blogRepository := repositories.NewBlogRepository(db)
	blogService := services.NewBlogService(&blogRepository)
	blogController := controllers.NewBlogController(blogService)
	blogRoute := routes.NewBlogRoute(*blogController, router)
	blogRoute.Setup(authMiddleware)

	// User Management
	instructorRepository := repositories.NewInstructorRepository(db)
    instructorService := services.NewInstructorService(&instructorRepository, userRepository)
    instructorController := controllers.NewInstructorController(instructorService)
    instructorRoute := routes.NewInstructorRoute(*instructorController, router)
    instructorRoute.Setup(authMiddleware)

	staffRepository := repositories.NewStaffRepository(db)
	staffService := services.NewStaffService(&staffRepository, userRepository)
	staffController := controllers.NewStaffController(staffService)
	staffRoute := routes.NewStaffRoute(*staffController, router)
	staffRoute.Setup(authMiddleware)

	customerRepository := repositories.NewCustomerRepository(db)
	customerService := services.NewCustomerService(&customerRepository)
	customerController := controllers.NewCustomerController(customerService)
	customerRoute := routes.NewCustomerRoute(*customerController, router)
	customerRoute.Setup(authMiddleware)

	// Booking API
	bookingRepository := repositories.NewBookingRepository(db)
    bookingService := services.NewBookingService(&bookingRepository)
    bookingController := controllers.NewBookingController(bookingService)
    bookingRoute := routes.NewBookingRoute(*bookingController, router)
    bookingRoute.Setup(authMiddleware)

	// router.Gin.Run(":" + os.Getenv("APP_PORT"))

	enchancedRouter := enableCORS(jsonContentTypeMiddleware(router.Gin))

	fmt.Println("App running in port: ", os.Getenv("APP_PORT"))
	log.Fatal(http.ListenAndServe(":" + os.Getenv("APP_PORT"), enchancedRouter))
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow any origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Check if the request is for CORS preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass down the request to the next middleware (or final handler)
		next.ServeHTTP(w, r)
	})

}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set JSON Content-Type
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}