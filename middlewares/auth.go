package middlewares

import (
	"blog/config"
	"blog/pkg/jwt"
	"blog/response"
	"strings"

	"github.com/gin-gonic/gin"
)

// 登录状态检查
func LoginAuthCheck(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		response.ResponseError(c, response.CodeNeedLogin)
		c.Abort()
		return
	}

	parts := strings.SplitN(token, " ", 2)
	if len(parts) < 2 || parts[0] != "Bearer" {
		response.ResponseError(c, response.CodeInvalidToken)
		c.Abort()
		return
	}

	userClaim, err := jwt.AnalyzeToken(parts[1])
	if err != nil {
		response.ResponseError(c, response.CodeInvalidToken)
		c.Abort()
		return
	}

	c.Set(config.UserClaimKey, userClaim)
	c.Next()
}
