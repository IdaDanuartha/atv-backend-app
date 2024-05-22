package config

import (
    // "net/http"

    "github.com/gin-gonic/gin"
)

//GinRouter -> Gin Router
type GinRouter struct {
    Gin *gin.Engine
}

//NewGinRouter all the routes are defined here
func NewGinRouter() GinRouter {

    httpRouter := gin.Default()

    httpRouter.ForwardedByClientIP = true
    httpRouter.SetTrustedProxies([]string{"127.0.0.1"})

    // httpRouter.GET("/", func(c *gin.Context) {
    //     c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
    // })
    return GinRouter{
        Gin: httpRouter,
    }

}