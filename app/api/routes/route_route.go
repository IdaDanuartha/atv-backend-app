package routes

import (
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/api/controllers"
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/gin-gonic/gin"
)

// RouteRoute -> Route for question module
type RouteRoute struct {
	Controller controllers.RouteController
	Handler    config.GinRouter
}

// NewRouteRoute -> initializes new choice rouets
func NewRouteRoute(
	controller controllers.RouteController,
	handler config.GinRouter,

) RouteRoute {
	return RouteRoute{
		Controller: controller,
		Handler:    handler,
	}
}

// Setup -> setups new choice Routes
func (p RouteRoute) Setup(authMiddleware gin.HandlerFunc) {
	apiPrefix := os.Getenv("APP_PREFIX")

	route := p.Handler.Gin.Group(apiPrefix + "/routes") //Router group

	// Apply AuthMiddleware to the route route group
	// facility.Use(authMiddleware)

	{
		route.GET("/", p.Controller.GetRoutes)
		route.GET("/:id", p.Controller.GetRoute)
		route.POST("/", authMiddleware, p.Controller.AddRoute)
		route.PATCH("/:id", authMiddleware, p.Controller.UpdateRoute)
		route.DELETE("/:id", authMiddleware, p.Controller.DeleteRoute)
	}
}
