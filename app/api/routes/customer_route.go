package routes

import (
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/api/controllers"
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/gin-gonic/gin"
)

// CustomerRoute -> Route for question module
type CustomerRoute struct {
	Controller controllers.CustomerController
	Handler    config.GinRouter
}

// NewCustomerRoute -> initializes new choice rouets
func NewCustomerRoute(
	controller controllers.CustomerController,
	handler config.GinRouter,

) CustomerRoute {
	return CustomerRoute{
		Controller: controller,
		Handler:    handler,
	}
}

// Setup -> setups new choice Routes
func (p CustomerRoute) Setup(authMiddleware gin.HandlerFunc) {
	apiPrefix := os.Getenv("APP_PREFIX")

	customer := p.Handler.Gin.Group(apiPrefix + "/customers") //Router group

	// Apply AuthMiddleware to the customer route group
	customer.Use(authMiddleware)

	{
		customer.GET("/", p.Controller.GetCustomers)
		customer.GET("/:id", p.Controller.GetCustomer)
	}
}
