package router

import (
	"blog/controller"
	middlewares "blog/middleware"

	_ "blog/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func Init(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(middlewares.GinLogger(), middlewares.GinRecovery(false), middlewares.Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/register", controller.RegisterHandler)

	return r
}
