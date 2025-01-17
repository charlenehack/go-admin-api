package middleware

import (
	"admin-api/common/result"
	"admin-api/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			result.Failed(c, 403, "请求头中的auth为空")
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			result.Failed(c, 405, "请求头中的auth格式有误")
			c.Abort()
			return
		}
		mc, err := jwt.ValidateToken(parts[1])
		if err != nil {
			result.Failed(c, 406, "无效的token，请重新登录")
			c.Abort()
			return
		}
		c.Set("authedUserObj", mc)
		c.Next()
	}
}
