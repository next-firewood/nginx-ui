package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strings"
	"time"
)

type AuthConf struct {
	SecretKey  string
	Expiration time.Duration
}

func AuthMiddleware(secretKey string, skipPaths ...string) gin.HandlerFunc {
	skipPathMap := make(map[string]bool)
	for _, path := range skipPaths {
		skipPathMap[path] = true
	}

	return func(c *gin.Context) {
		// 检查是否为白名单路径
		if skipPathMap[c.Request.URL.Path] {
			c.Next()
			return
		}
		// 从请求头获取令牌
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			unauthorized(c, http.StatusUnauthorized, "missing auth header")
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			unauthorized(c, http.StatusUnauthorized, "invalid auth header format")
			return
		}

		tokenStr := parts[1]

		// 解析和验证令牌
		parsedToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}

			return secretKey, nil
		})

		if err != nil || !parsedToken.Valid {
			unauthorized(c, http.StatusUnauthorized, "invalid token")
			return
		}

		// 将用户信息存储到上下文
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
			c.Set("uc", claims)
		}

		c.Next()
	}
}

// 未授权响应
func unauthorized(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

// 生成 JWT 令牌
func (a *AuthConf) GenerateToken(secretKey []byte, claims jwt.MapClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	// 设置过期时间
	if _, ok := claims["exp"]; !ok {
		claims["exp"] = time.Now().Add(a.Expiration).Unix()
	}

	return token.SignedString(secretKey)
}
