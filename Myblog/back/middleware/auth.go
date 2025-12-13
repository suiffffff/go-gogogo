package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 简单的 Cookie 认证
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Cookie
		userID, err := c.Cookie("user_id")
		if err != nil {
			// 如果没登录，就把 user_id 设为 0，或者直接跳转去登录页
			// 这里我们设为 0，方便前端判断是否显示“登录”按钮
			c.Set("user_id", 0)
			c.Next()
			return
		}
		// 如果登录了，把 ID 存到上下文，方便后续使用
		c.Set("user_id", userID)
		c.Next()
	}
}

// ForceLogin 强制登录中间件 (用于发布文章、评论等接口)
func ForceLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("user_id")
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}
