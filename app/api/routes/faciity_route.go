package routes

import (
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/api/controllers"
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/gin-gonic/gin"
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
func (p FacilityRoute) Setup(authMiddleware gin.HandlerFunc) {
	apiPrefix := os.Getenv("APP_PREFIX")

	facility := p.Handler.Gin.Group(apiPrefix + "/facilities") //Router group

	// Apply AuthMiddleware to the facility route group
	// facility.Use(authMiddleware)

	{
		facility.GET("/", p.Controller.GetFacilities)
		facility.GET("/:id", p.Controller.GetFacility)
		facility.POST("/", authMiddleware, p.Controller.AddFacility)
		facility.PATCH("/:id", authMiddleware, p.Controller.UpdateFacility)
		facility.DELETE("/:id", authMiddleware, p.Controller.DeleteFacility)
	}
}
