package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Response(c *gin.Context, resp interface{}, err error) {
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Code": 500,
			"Msg":  "发生未知错误",
		})

		c.Writer.Header().Add("err", strings.TrimSpace(err.Error()))
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Code": 0,
			"Msg":  "OK",
			"Data": resp,
		})
	}
}
