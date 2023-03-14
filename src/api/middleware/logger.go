package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 在请求处理之前记录请求信息
		zap.L().Info(
			"request started",
			zap.String("method", c.Request.Method),
			zap.String("url", c.Request.URL.String()),
			zap.Any("headers", c.Request.Header),
			zap.Any("body", c.Request.Body),
		)
		// 中间件逻辑
		c.Next()

		// 在请求处理之后记录响应信息
		zap.L().Info(
			"request finished",
			zap.String("method", c.Request.Method),
			zap.String("url", c.Request.URL.String()),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Since(start)),
		)
	}
}
