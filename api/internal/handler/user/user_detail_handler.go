package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rui/common/response"
	"rui/internal/api"
	"rui/internal/logic/user"
	"rui/internal/svc"
)

// UserDetailHandler 获取所有配置文件
// @Summary 获取配置列表
// @Tags user
// @Param UuidForm query api.UuidForm true "uuid"
// @Success 200 {object} UserDetailResp
// @Router /api/user/detail [get]
func UserDetailHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &api.UuidForm{}

		if err := c.ShouldBindQuery(req); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		}

		logic := user.NewUserDetailLogic(c.Request.Context(), svcCtx)

		resp, err := logic.UserDetail(req)
		response.Response(c, resp, err)
	}
}
