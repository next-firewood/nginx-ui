package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func LoggerMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.Path

		c.Next()

		err := c.Writer.Header().Get("err")
		if err != "" {
			log.Printf("Url: %s\tErr: %s", url, err)
		}
	}
}
