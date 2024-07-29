package routes

import (
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/api/controllers"
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/gin-gonic/gin"
)

// InstructorRoute -> Route for question module
type InstructorRoute struct {
	Controller controllers.InstructorController
	Handler    config.GinRouter
}

// NewInstructorRoute -> initializes new choice rouets
func NewInstructorRoute(
	controller controllers.InstructorController,
	handler config.GinRouter,

) InstructorRoute {
	return InstructorRoute{
		Controller: controller,
		Handler:    handler,
	}
}

// Setup -> setups new choice Routes
func (p InstructorRoute) Setup(authMiddleware gin.HandlerFunc) {
	apiPrefix := os.Getenv("APP_PREFIX")

	instructor := p.Handler.Gin.Group(apiPrefix + "/instructors") //Router group

	// Apply AuthMiddleware to the instructor route group
	// instructor.Use(authMiddleware)

	{
		instructor.GET("/", p.Controller.GetInstructors)
		instructor.GET("/:id", p.Controller.GetInstructor)
		instructor.POST("/", authMiddleware, p.Controller.AddInstructor)
		instructor.PATCH("/:id", authMiddleware, p.Controller.UpdateInstructor)
		instructor.DELETE("/:id", authMiddleware, p.Controller.DeleteInstructor)
	}
}
