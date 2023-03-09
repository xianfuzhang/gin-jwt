package main

import (
	"test/v2/controllers"
	"test/v2/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	api.POST("/user/register", controllers.RegisterUser)
	// api.POST("/token", controllers.GenerateToken)

	v1 := api.Group("/v1").Use(middlewares.Auth())
	v1.GET("/ping", controllers.Ping)

	return router
}
