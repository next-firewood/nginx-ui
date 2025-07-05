package user

import (
	"github.com/gin-gonic/gin"
	"rui/internal/handler/user"
	"rui/internal/svc"
)

func Router(router *gin.RouterGroup, svcCtx *svc.ServiceContext) {
	group := router.Group("/user")

	group.GET("/detail", user.UserDetailHandler(svcCtx))
}
