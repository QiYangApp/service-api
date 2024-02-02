package middlewares

import (
	"app/config"
	"app/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func Logger() gin.HandlerFunc {
	if config.Client().GetBool("server.debug") {
		return gin.Logger()
	}

	return func(c *gin.Context) {
		start := time.Now()
		// 中间件逻辑
		c.Next()

		// 在请求处理之后记录响应信息
		log.Client().Sugar().Infow(
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
