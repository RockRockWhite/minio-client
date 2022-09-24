package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Logger 记录请求中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
