package routes

import (
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/api/controllers"
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/gin-gonic/gin"
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
func (p MandatoryLuggageRoute) Setup(authMiddleware gin.HandlerFunc) {
	apiPrefix := os.Getenv("APP_PREFIX")

	mandatoryLuggage := p.Handler.Gin.Group(apiPrefix + "/mandatory/luggages") //Router group

	// Apply AuthMiddleware to the mandatoryLuggage route group
	// mandatoryLuggage.Use(authMiddleware)

	{
		mandatoryLuggage.GET("/", p.Controller.GetMandatoryLuggages)
		mandatoryLuggage.GET("/:id", p.Controller.GetMandatoryLuggage)
		mandatoryLuggage.POST("/", authMiddleware, p.Controller.AddMandatoryLuggage)
		mandatoryLuggage.PATCH("/:id", authMiddleware, p.Controller.UpdateMandatoryLuggage)
		mandatoryLuggage.DELETE("/:id", authMiddleware, p.Controller.DeleteMandatoryLuggage)
	}
}
