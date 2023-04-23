package main

import (
	_ "test/v2/docs"
	"test/v2/internal/adapter"
	"test/v2/internal/adapter/sqlite"
	"test/v2/internal/middlewares"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           REST API base on Clean Arch.
// @version         1.0
// @host      localhost:8080
// @BasePath  /api
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	sqlite.ConnDB()
	router := setupRouter()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	adapter.LoadLoginRouter(api)

	v1 := api.Group("/v1")
	v1.Use(middlewares.Auth())
	adapter.LoadUserRouter(v1)

	return router
}
