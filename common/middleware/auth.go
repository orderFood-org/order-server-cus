package middleware

import (
	"github.com/gin-gonic/gin"
)

// Auth 认证中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取令牌
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "未授权访问"})
			c.Abort()
			return
		}

		// 验证令牌
		// 通常这里会调用token.Service来验证令牌
		// 现在只是简单示例，实际实现需要详细的令牌验证逻辑

		// 如果验证通过，设置用户信息到上下文
		c.Set("userID", uint(1)) // 示例用户ID
		c.Next()
	}
}
