package config

import (
	// "net/http"

	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//GinRouter -> Gin Router
type GinRouter struct {
    Gin *gin.Engine
}

//NewGinRouter all the routes are defined here
func NewGinRouter() GinRouter {

    httpRouter := gin.Default()
    // Apply CORS middleware
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"*"}
    // config.AllowHeaders = []string{"Content-Type", "Content-Disposition"} 
    httpRouter.Use(cors.Default())

	cookieStore := cookie.NewStore([]byte(os.Getenv("SECRET_KEY")))
	httpRouter.Use(sessions.Sessions("atv_system", cookieStore))

    httpRouter.ForwardedByClientIP = true
    httpRouter.SetTrustedProxies([]string{"127.0.0.1"})

    httpRouter.Static("/uploads", "./uploads")

    // httpRouter.GET("/", func(c *gin.Context) {
    //     c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
    // })
    return GinRouter{
        Gin: httpRouter,
    }

}