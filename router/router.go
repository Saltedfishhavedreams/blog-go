package router

import (
	"blog/controller"
	middlewares "blog/middlewares"

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
	r.POST("/login", controller.LoginHandler)

	auth := r.Group("/", middlewares.LoginAuthCheck)

	auth.GET("/refresh_token", controller.RefreshTokenHandler)
	auth.GET("/role", controller.GetRoleListHandler)
	auth.GET("/role/:role_id", controller.GetRoleHandler)

	return r
}
