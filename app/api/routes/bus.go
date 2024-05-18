package routes

import (
	"github.com/IdaDanuartha/teman-bus-backend-app/app/api/controllers"
	"github.com/IdaDanuartha/teman-bus-backend-app/app/config"
)

//BusRoute -> Route for question module
type BusRoute struct {
    Controller controllers.BusController
    Handler    config.GinRouter
}

//NewBusRoute -> initializes new choice rouets
func NewBusRoute(
    controller controllers.BusController,
    handler config.GinRouter,

) BusRoute {
    return BusRoute{
        Controller: controller,
        Handler:    handler,
    }
}

//Setup -> setups new choice Routes
func (p BusRoute) Setup() {
    bus := p.Handler.Gin.Group("/bus") //Router group
    {
        bus.GET("/", p.Controller.GetBuses)
        bus.GET("/:id", p.Controller.GetBus)
        bus.POST("/", p.Controller.AddBus)
        bus.PATCH("/:id", p.Controller.UpdateBus)
        bus.DELETE("/:id", p.Controller.DeleteBus)
    }
}