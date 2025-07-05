package public

import (
	"github.com/gin-gonic/gin"
	"rui/internal/handler/public"
	"rui/internal/svc"
)

func Router(router *gin.RouterGroup, svcCtx *svc.ServiceContext) {
	group := router.Group("/public")

	group.POST("/login", public.LoginHandler(svcCtx))
}
