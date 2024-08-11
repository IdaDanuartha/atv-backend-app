package routes

import (
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/api/controllers"
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/gin-gonic/gin"
)

// BookingRoute -> Route for question module
type BookingRoute struct {
	Controller controllers.BookingController
	Handler    config.GinRouter
}

// NewBookingRoute -> initializes new choice rouets
func NewBookingRoute(
	controller controllers.BookingController,
	handler config.GinRouter,

) BookingRoute {
	return BookingRoute{
		Controller: controller,
		Handler:    handler,
	}
}

// Setup -> setups new choice Routes
func (p BookingRoute) Setup(authMiddleware gin.HandlerFunc) {
	apiPrefix := os.Getenv("APP_PREFIX")

	booking := p.Handler.Gin.Group(apiPrefix + "/bookings") //Router group

	// Apply AuthMiddleware to the booking route group
	booking.Use(authMiddleware)

	{
		booking.GET("/", p.Controller.GetBookings)
		booking.GET("/:id", p.Controller.GetBooking)
		booking.GET("/export/excel", p.Controller.ExportToExcel)
		booking.POST("/", p.Controller.AddBooking)
		// booking.PATCH("/:id", p.Controller.UpdateBooking)
		booking.DELETE("/:id", p.Controller.DeleteBooking)
	}
}
