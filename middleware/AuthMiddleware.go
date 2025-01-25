package middleware

import (
	"easydemo/common"
	"easydemo/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// gin的中间件返回一个gin.HandlerFunc类型
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从header中获取authorization header
		tokenString := c.GetHeader("Authorization")

		// 验证正确的token格式: 非空并且以Bearer开头
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}

		// 格式正确,提取token的有效部分
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Token错误"})
			c.Abort()
			return
		}

		// 验证通过后获取claims中的UserId
		userId := claims.UserId
		db := common.GetDB()
		var user model.User
		db.First(&user, userId)

		// 用户不存在
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}

		// 验证通过后,将user信息写入上下文,这样用户才能够查询信息
		c.Set("user", user)
	}
}
