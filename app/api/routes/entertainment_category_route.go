package routes

import (
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/api/controllers"
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/gin-gonic/gin"
)

// EntertainmentCategoryRoute -> Route for question module
type EntertainmentCategoryRoute struct {
	Controller controllers.EntertainmentCategoryController
	Handler    config.GinRouter
}

// NewEntertainmentCategoryRoute -> initializes new choice rouets
func NewEntertainmentCategoryRoute(
	controller controllers.EntertainmentCategoryController,
	handler config.GinRouter,

) EntertainmentCategoryRoute {
	return EntertainmentCategoryRoute{
		Controller: controller,
		Handler:    handler,
	}
}

// Setup -> setups new choice Routes
func (p EntertainmentCategoryRoute) Setup(authMiddleware gin.HandlerFunc)  {
	apiPrefix := os.Getenv("APP_PREFIX")

	entertainmentCategory := p.Handler.Gin.Group(apiPrefix + "/entertainment/categories") //Router group
	
	// Apply AuthMiddleware to the entertainmentCategory route group
	// entertainmentCategory.Use(authMiddleware)
	
	{
		entertainmentCategory.GET("/", p.Controller.GetEntertainmentCategories)
		entertainmentCategory.GET("/:id", p.Controller.GetEntertainmentCategory)
		entertainmentCategory.POST("/", authMiddleware, p.Controller.AddEntertainmentCategory)
		entertainmentCategory.PATCH("/:id", authMiddleware, p.Controller.UpdateEntertainmentCategory)
		entertainmentCategory.DELETE("/:id", authMiddleware, p.Controller.DeleteEntertainmentCategory)
	}
}
