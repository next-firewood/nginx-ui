package public

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rui/common/response"
	publicApi "rui/internal/api/public"
	public "rui/internal/logic/public"
	"rui/internal/svc"
)

// LoginHandler 登录
// @Summary 登录
// @Tags public
// @Param LoginReq body LoginReq true "LoginReq"
// @Success 200 {object} LoginResp
// @Router /api/public/login [post]
func LoginHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &publicApi.LoginReq{}

		if err := c.ShouldBindBodyWithJSON(req); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		}

		logic := public.NewLoginLogic(c.Request.Context(), svcCtx)

		resp, err := logic.Login(req)
		response.Response(c, resp, err)
	}
}
