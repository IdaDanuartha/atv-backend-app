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

	bus := p.Handler.Gin.Group(apiPrefix + "/entertainment/packages") //Router group
	{
		bus.GET("/", p.Controller.GetEntertainmentPackages)
		bus.GET("/:id", p.Controller.GetEntertainmentPackage)
		bus.POST("/", p.Controller.AddEntertainmentPackage)
		bus.PATCH("/:id", p.Controller.UpdateEntertainmentPackage)
		bus.DELETE("/:id", p.Controller.DeleteEntertainmentPackage)
	}
}
