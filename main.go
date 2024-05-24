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
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

func init() {
	config.LoadEnv()
}

func main() {
	router := config.NewGinRouter()
	db := config.NewDatabase()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService()
	userController := controllers.NewUserController(userService, authService)
	userRoute := routes.NewUserRoute(userController, router)
	userRoute.Setup()
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.Admin{})
	db.DB.AutoMigrate(&models.Staff{})
	db.DB.AutoMigrate(&models.Instructor{})
	db.DB.AutoMigrate(&models.Customer{})

	authMiddleware := middlewares.AuthMiddleware(authService, userService)
	entertainmentCategoryRepository := repositories.NewEntertainmentCategoryRepository(db)
	entertainmentCategoryService := services.NewEntertainmentCategoryService(&entertainmentCategoryRepository)
	entertainmentCategoryController := controllers.NewEntertainmentCategoryController(entertainmentCategoryService)
	entertainmentCategoryRoute := routes.NewEntertainmentCategoryRoute(*entertainmentCategoryController, router)
	entertainmentCategoryRoute.Setup(authMiddleware)
	db.DB.AutoMigrate(&models.EntertainmentCategory{})

	router.Gin.Run(":" + os.Getenv("APP_PORT"))

	fmt.Println("App running in port: ", os.Getenv("APP_PORT"))
}
