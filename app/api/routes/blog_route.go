package routes

import (
	"os"

	"github.com/IdaDanuartha/atv-backend-app/app/api/controllers"
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/gin-gonic/gin"
)

// BlogRoute -> Route for question module
type BlogRoute struct {
	Controller controllers.BlogController
	Handler    config.GinRouter
}

// NewBlogRoute -> initializes new choice rouets
func NewBlogRoute(
	controller controllers.BlogController,
	handler config.GinRouter,

) BlogRoute {
	return BlogRoute{
		Controller: controller,
		Handler:    handler,
	}
}

// Setup -> setups new choice Routes
func (p BlogRoute) Setup(authMiddleware gin.HandlerFunc) {
	apiPrefix := os.Getenv("APP_PREFIX")

	blog := p.Handler.Gin.Group(apiPrefix + "/blogs") //Router group

	// Apply AuthMiddleware to the blog route group
	// blog.Use(authMiddleware)

	{
		blog.GET("/", p.Controller.GetBlogs)
		blog.GET("/:id", p.Controller.GetBlog)
		blog.POST("/upload/:id", authMiddleware, p.Controller.UploadImage)
		blog.POST("/", authMiddleware, p.Controller.AddBlog)
		blog.PATCH("/:id", authMiddleware, p.Controller.UpdateBlog)
		blog.DELETE("/:id", authMiddleware, p.Controller.DeleteBlog)
	}
}
