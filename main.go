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

	authService := services.NewAuthService()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService, authService)
	userRoute := routes.NewUserRoute(userController, router)
	userRoute.Setup()

	authMiddleware := middlewares.AuthMiddleware(authService, userService)

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
	

	router.Gin.Run(":" + os.Getenv("APP_PORT"))

	fmt.Println("App running in port: ", os.Getenv("APP_PORT"))
}
