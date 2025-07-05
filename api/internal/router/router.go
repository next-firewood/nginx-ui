package router

import (
	"github.com/gin-gonic/gin"
	"rui/internal/middleware"
	"rui/internal/router/public"
	"rui/internal/router/user"
	"rui/internal/svc"
)

func InitRouter(router *gin.Engine, svcCtx *svc.ServiceContext) {
	globalRouter := router.Group("/api")
	globalRouter.Use(middleware.AuthMiddleware(svcCtx.Config.Auth.SecretKey, SkipPaths...))
	globalRouter.Use(middleware.LoggerMiddle()) // 中间件

	user.Router(globalRouter, svcCtx)
	public.Router(globalRouter, svcCtx)
}
