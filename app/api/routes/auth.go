package routes

import (
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/api/controllers"
	"github.com/IdaDanuartha/atv-backend-app/app/config"
)

// AuthRoute -> Route for question module
type AuthRoute struct {
	Controller controllers.AuthController
	Handler    config.GinRouter
}

// NewAuthRoute -> initializes new choice rouets
func NewAuthRoute(
	controller controllers.AuthController,
	handler config.GinRouter,

) AuthRoute {
	return AuthRoute{
		Controller: controller,
		Handler:    handler,
	}
}

// Setup -> setups new choice Routes
func (p AuthRoute) Setup() {
	apiPrefix := os.Getenv("APP_PREFIX")

	auth := p.Handler.Gin.Group(apiPrefix + "/auth") //Router group
	{
		auth.POST("/register", p.Controller.RegisterUser)
		// auth.POST("/login", p.Controller.Login)
	}
}
