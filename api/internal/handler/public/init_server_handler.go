package public

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rui/common/response"
	publicApi "rui/internal/api/public"
	public "rui/internal/logic/public"
	"rui/internal/svc"
)

// InitServerHandler 初始化服务
// @Summary 初始化服务
// @Tags public
// @Param InitServerReq body InitServerReq true "初始化服务请求体"
// @Router /api/public/init/server [post]
func InitServerHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &publicApi.InitServerReq{}

		if err := c.ShouldBindBodyWithJSON(req); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		}

		logic := public.NewInitServerLogic(c.Request.Context(), svcCtx)

		err := logic.InitServerLogic(req)
		response.Response(c, nil, err)
	}
}
