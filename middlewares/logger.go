package middlewares

import (
	"github.com/RockRockWhite/minio-client/utils"
	"github.com/gin-gonic/gin"
	"time"
)

// Logger 记录请求中间件
func Logger(c *gin.Context) {
	startTime := time.Now()
	// 处理请求
	c.Next()
	// 计算处理请求消耗时间
	latency := time.Since(startTime)
	// 记录日志
	utils.GetLogger().Printf("[request]\tstatus:%3d\tcost:%3v\taddr:%15s\tmethod:%s\turl:%s", c.Writer.Status(), latency, c.ClientIP(), c.Request.Method, c.Request.RequestURI)
}
