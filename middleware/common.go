package middleware

import "github.com/gin-gonic/gin"

// CacheControl 屏蔽客户端缓存
func CacheControl() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "private, no-cache")
	}
}
