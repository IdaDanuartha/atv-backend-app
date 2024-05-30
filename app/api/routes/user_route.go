package routes

import (
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/api/controllers"
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/gin-gonic/gin"
)

// UserRoute -> Route for question module
type UserRoute struct {
	Controller controllers.UserController
	Handler    config.GinRouter
}

// NewUserRoute -> initializes new choice rouets
func NewUserRoute(
	controller *controllers.UserController,
	handler config.GinRouter,

) UserRoute {
	return UserRoute{
		Controller: *controller,
		Handler:    handler,
	}
}

// Setup -> setups new choice Routes
func (p UserRoute) Setup(authMiddleware gin.HandlerFunc) {
	apiPrefix := os.Getenv("APP_PREFIX")

	auth := p.Handler.Gin.Group(apiPrefix + "/auth") //Router group
	{
		auth.GET("/user/me", authMiddleware, p.Controller.FetchUser)
		auth.POST("/register", p.Controller.RegisterUser)
		auth.POST("/login", p.Controller.Login)
		auth.POST("/upload-avatar", authMiddleware, p.Controller.UploadAvatar)
		auth.PATCH("/update", authMiddleware, p.Controller.UpdateProfile)
	}
}
