package routes

import (
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/api/controllers"
	"github.com/IdaDanuartha/atv-backend-app/app/config"
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
func (p EntertainmentPackageRoute) Setup() {
	apiPrefix := os.Getenv("APP_PREFIX")

	entertainmentPackage := p.Handler.Gin.Group(apiPrefix + "/entertainment/packages") //Router group
	{
		entertainmentPackage.GET("/", p.Controller.GetEntertainmentPackages)
		entertainmentPackage.GET("/:id", p.Controller.GetEntertainmentPackage)
		entertainmentPackage.POST("/", p.Controller.AddEntertainmentPackage)
		entertainmentPackage.PATCH("/:id", p.Controller.UpdateEntertainmentPackage)
		entertainmentPackage.DELETE("/:id", p.Controller.DeleteEntertainmentPackage)
	}
}
