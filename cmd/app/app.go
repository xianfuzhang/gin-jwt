package main

import (
	// "test/v2/internal/controllers"

	_ "test/v2/docs"
	"test/v2/internal/adapter"
	"test/v2/internal/adapter/sqlite"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           GIN JWT API
// @version         1.0
// @host      localhost:8080
// @BasePath  /api
func main() {
	sqlite.ConnDB()
	router := setupRouter()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	adapter.LoadUserRouter(api)
	// api.POST("/user/register", controllers.RegisterUser)
	// api.POST("/token", controllers.GenerateToken)

	// v1 := api.Group("/v1").Use(middlewares.Auth())

	return router
}
