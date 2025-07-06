package public

import (
	"github.com/gin-gonic/gin"
	"rui/common/response"
	public "rui/internal/logic/public"
	"rui/internal/svc"
)

// InitStatusHandler 初始化状态
// @Summary 初始化状态
// @Tags public
// @Success 200 {object} InitStatusRes
// @Router /api/public/init/status [get]
func InitStatusHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		logic := public.NewInitStatusLogic(c.Request.Context(), svcCtx)

		resp, err := logic.InitStatusLogic()
		response.Response(c, resp, err)
	}
}
