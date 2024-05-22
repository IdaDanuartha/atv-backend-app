package routes

import (
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/api/controllers"
	"github.com/IdaDanuartha/atv-backend-app/app/config"
)

// FacilityRoute -> Route for question module
type FacilityRoute struct {
	Controller controllers.FacilityController
	Handler    config.GinRouter
}

// NewFacilityRoute -> initializes new choice rouets
func NewFacilityRoute(
	controller controllers.FacilityController,
	handler config.GinRouter,

) FacilityRoute {
	return FacilityRoute{
		Controller: controller,
		Handler:    handler,
	}
}

// Setup -> setups new choice Routes
func (p FacilityRoute) Setup() {
	apiPrefix := os.Getenv("APP_PREFIX")

	bus := p.Handler.Gin.Group(apiPrefix + "/facilities") //Router group
	{
		bus.GET("/", p.Controller.GetFacilities)
		bus.GET("/:id", p.Controller.GetFacility)
		bus.POST("/", p.Controller.AddFacility)
		bus.PATCH("/:id", p.Controller.UpdateFacility)
		bus.DELETE("/:id", p.Controller.DeleteFacility)
	}
}
