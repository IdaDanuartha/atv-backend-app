package routes

import (
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/api/controllers"
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/gin-gonic/gin"
)

// EntertainmentServiceRoute -> Route for question module
type EntertainmentServiceRoute struct {
	Controller controllers.EntertainmentServiceController
	Handler    config.GinRouter
}

// NewEntertainmentServiceRoute -> initializes new choice rouets
func NewEntertainmentServiceRoute(
	controller controllers.EntertainmentServiceController,
	handler config.GinRouter,

) EntertainmentServiceRoute {
	return EntertainmentServiceRoute{
		Controller: controller,
		Handler:    handler,
	}
}

// Setup -> setups new choice Routes
func (p EntertainmentServiceRoute) Setup(authMiddleware gin.HandlerFunc) {
	apiPrefix := os.Getenv("APP_PREFIX")

	entertainmentService := p.Handler.Gin.Group(apiPrefix + "/entertainment/services") //Router group

	// Apply AuthMiddleware to the entertainmentService route group
	// entertainmentService.Use(authMiddleware)

	{
		entertainmentService.GET("/", p.Controller.GetEntertainmentServices)
		entertainmentService.GET("/:id", p.Controller.GetEntertainmentService)
		entertainmentService.POST("/upload/:id", authMiddleware, p.Controller.UploadAvatar)
		entertainmentService.POST("/", authMiddleware, p.Controller.AddEntertainmentService)
		entertainmentService.PATCH("/:id", authMiddleware, p.Controller.UpdateEntertainmentService)
		entertainmentService.DELETE("/:id", authMiddleware, p.Controller.DeleteEntertainmentService)
	}
}
