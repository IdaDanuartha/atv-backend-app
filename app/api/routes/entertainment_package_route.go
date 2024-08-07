package routes

import (
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/api/controllers"
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/gin-gonic/gin"
)

// EntertainmentPackageRoute -> Route for question module
type EntertainmentPackageRoute struct {
	Controller controllers.EntertainmentPackageController
	Handler    config.GinRouter
}

// NewEntertainmentPackageRoute -> initializes new choice rouets
func NewEntertainmentPackageRoute(
	controller controllers.EntertainmentPackageController,
	handler config.GinRouter,

) EntertainmentPackageRoute {
	return EntertainmentPackageRoute{
		Controller: controller,
		Handler:    handler,
	}
}

// Setup -> setups new choice Routes
func (p EntertainmentPackageRoute) Setup(authMiddleware gin.HandlerFunc) {
	apiPrefix := os.Getenv("APP_PREFIX")

	entertainmentPackage := p.Handler.Gin.Group(apiPrefix + "/entertainment/packages") //Router group

	// Apply AuthMiddleware to the entertainmentPackage route group
	// entertainmentPackage.Use(authMiddleware)

	{
		entertainmentPackage.GET("/", p.Controller.GetEntertainmentPackages)
		entertainmentPackage.GET("/:id", p.Controller.GetEntertainmentPackage)
		entertainmentPackage.POST("/upload/:id", authMiddleware, p.Controller.UploadImage)
		entertainmentPackage.POST("/", authMiddleware, p.Controller.AddEntertainmentPackage)
		entertainmentPackage.PATCH("/:id", authMiddleware, p.Controller.UpdateEntertainmentPackage)
		entertainmentPackage.DELETE("/:id", authMiddleware, p.Controller.DeleteEntertainmentPackage)
	}
}
