package routes

import (
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/api/controllers"
	"github.com/IdaDanuartha/atv-backend-app/app/config"
)

//EntertainmentCategoryRoute -> Route for question module
type EntertainmentCategoryRoute struct {
    Controller controllers.EntertainmentCategoryController
    Handler    config.GinRouter
}

//NewEntertainmentCategoryRoute -> initializes new choice rouets
func NewEntertainmentCategoryRoute(
    controller controllers.EntertainmentCategoryController,
    handler config.GinRouter,

) EntertainmentCategoryRoute {
    return EntertainmentCategoryRoute{
        Controller: controller,
        Handler:    handler,
    }
}

//Setup -> setups new choice Routes
func (p EntertainmentCategoryRoute) Setup() {
    apiPrefix := os.Getenv("APP_PREFIX")

    bus := p.Handler.Gin.Group(apiPrefix + "/entertainment/categories") //Router group
    {
        bus.GET("/", p.Controller.GetEntertainmentCategories)
        bus.GET("/:id", p.Controller.GetEntertainmentCategory)
        bus.POST("/", p.Controller.AddEntertainmentCategory)
        bus.PATCH("/:id", p.Controller.UpdateEntertainmentCategory)
        bus.DELETE("/:id", p.Controller.DeleteEntertainmentCategory)
    }
}