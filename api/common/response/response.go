package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"rui/common/errorx"
	"strings"
)

func Response(c *gin.Context, resp interface{}, err error) {
	if err != nil {
		var codeErr *errorx.CodeError
		if errors.As(err, &codeErr) {
			c.JSON(http.StatusOK, gin.H{
				"Code": codeErr.Code,
				"Msg":  codeErr.Msg,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Code": errorx.DefaultCode,
				"Msg":  errorx.ErrMap[errorx.DefaultCode].Error(),
			})

			c.Writer.Header().Add("err", strings.TrimSpace(err.Error()))
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Code": errorx.Success,
			"Msg":  "OK",
			"Data": resp,
		})
	}
}
