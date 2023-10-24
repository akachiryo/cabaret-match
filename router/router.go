package router

import (
	"api/controllers"
	middleware "api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	api := router.Group("/api")
	v1 := api.Group("/v1")
	hosts := v1.Group("/hosts")
	hosts.POST("", controllers.RegisterHost)
	hosts.POST("/login", controllers.LoginHost)

    hosts.Use(middleware.AuthMiddleware) // middlewareを設定
    {
        hosts.GET("", controllers.GetHost)
    }

	return router
}
