package routes

import (
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/api/controllers"
	"github.com/IdaDanuartha/atv-backend-app/app/config"
)

// MandatoryLuggageRoute -> Route for question module
type MandatoryLuggageRoute struct {
	Controller controllers.MandatoryLuggageController
	Handler    config.GinRouter
}

// NewMandatoryLuggageRoute -> initializes new choice rouets
func NewMandatoryLuggageRoute(
	controller controllers.MandatoryLuggageController,
	handler config.GinRouter,

) MandatoryLuggageRoute {
	return MandatoryLuggageRoute{
		Controller: controller,
		Handler:    handler,
	}
}

// Setup -> setups new choice Routes
func (p MandatoryLuggageRoute) Setup() {
	apiPrefix := os.Getenv("APP_PREFIX")

	bus := p.Handler.Gin.Group(apiPrefix + "/mandatory/luggages") //Router group
	{
		bus.GET("/", p.Controller.GetMandatoryLuggages)
		bus.GET("/:id", p.Controller.GetMandatoryLuggage)
		bus.POST("/", p.Controller.AddMandatoryLuggage)
		bus.PATCH("/:id", p.Controller.UpdateMandatoryLuggage)
		bus.DELETE("/:id", p.Controller.DeleteMandatoryLuggage)
	}
}
