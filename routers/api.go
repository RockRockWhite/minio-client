package routers

import (
	"github.com/RockRockWhite/minio-client/controllers"
	"github.com/RockRockWhite/minio-client/middlewares"
	"github.com/gin-gonic/gin"
)

func InitApiRouter() *gin.Engine {
	// 初始化Controllers
	router := gin.Default()

	// 配置中间件
	router.Use(middlewares.Logger())

	token := router.Group("/tokens")
	{
		token.POST("", controllers.CreateToken)
	}

	buckets := router.Group("/buckets")
	{
		buckets.GET("/:filename", controllers.GetFile)
		buckets.POST("/:filename", middlewares.AccessToken, controllers.UploadObject)
	}

	return router
}
