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

    entertainmentCategory := p.Handler.Gin.Group(apiPrefix + "/entertainment/categories") //Router group
    {
        entertainmentCategory.GET("/", p.Controller.GetEntertainmentCategories)
        entertainmentCategory.GET("/:id", p.Controller.GetEntertainmentCategory)
        entertainmentCategory.POST("/", p.Controller.AddEntertainmentCategory)
        entertainmentCategory.PATCH("/:id", p.Controller.UpdateEntertainmentCategory)
        entertainmentCategory.DELETE("/:id", p.Controller.DeleteEntertainmentCategory)
    }
}