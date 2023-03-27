package main

import (
	"fmt"
	"test/v2/internal/controllers"
	sql "test/v2/internal/lib"
	"test/v2/internal/middlewares"

	_ "test/v2/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @host      localhost:8080
// @BasePath  /api
func main() {
	pool, err := sql.PgPoolInit()
	if err != nil {
		fmt.Println(err)
	}
	defer pool.Close()

	router := initRouter()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	api.POST("/user/register", controllers.RegisterUser)
	api.POST("/token", controllers.GenerateToken)

	v1 := api.Group("/v1").Use(middlewares.Auth())
	v1.GET("/ping", controllers.Ping)

	return router
}
