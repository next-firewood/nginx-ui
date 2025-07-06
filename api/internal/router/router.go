package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"rui/internal/middleware"
	"rui/internal/router/public"
	"rui/internal/router/user"
	"rui/internal/svc"
)

func InitRouter(router *gin.Engine, svcCtx *svc.ServiceContext) {
	if svcCtx.Config.Mode == "dev" {
		// Swagger 接口
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	globalRouter := router.Group("/api")
	globalRouter.Use(middleware.AuthMiddleware(svcCtx.Config.Auth.SecretKey, SkipPaths...))
	globalRouter.Use(middleware.LoggerMiddle()) // 中间件

	user.Router(globalRouter, svcCtx)
	public.Router(globalRouter, svcCtx)
}
