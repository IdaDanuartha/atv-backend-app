package routes

import (
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/api/controllers"
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/gin-gonic/gin"
)

// StaffRoute -> Route for question module
type StaffRoute struct {
	Controller controllers.StaffController
	Handler    config.GinRouter
}

// NewStaffRoute -> initializes new choice rouets
func NewStaffRoute(
	controller controllers.StaffController,
	handler config.GinRouter,

) StaffRoute {
	return StaffRoute{
		Controller: controller,
		Handler:    handler,
	}
}

// Setup -> setups new choice Routes
func (p StaffRoute) Setup(authMiddleware gin.HandlerFunc) {
	apiPrefix := os.Getenv("APP_PREFIX")

	staff := p.Handler.Gin.Group(apiPrefix + "/staffs") //Router group

	// Apply AuthMiddleware to the staff route group
	staff.Use(authMiddleware)

	{
		staff.GET("/", p.Controller.GetStaffs)
		staff.GET("/:id", p.Controller.GetStaff)
		staff.POST("/", p.Controller.AddStaff)
		staff.PATCH("/:id", p.Controller.UpdateStaff)
		staff.DELETE("/:id", p.Controller.DeleteStaff)
	}
}
