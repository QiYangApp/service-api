package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	loggerInstance "service-api/src/core/logger"
	"time"
)

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 中间件逻辑
		c.Next()

		// 在请求处理之后记录响应信息
		loggerInstance.C(loggerInstance.RequestMode).Info(
			"request",
			zap.String("url", c.Request.URL.String()),
			zap.Any("headers", c.Request.Header),
			zap.String("method", c.Request.Method),
			zap.Any("body", c.Request.Body),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Since(start)),
		)
	}
}
