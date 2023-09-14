package middleware

import (
	"go.uber.org/zap"
	"net/http"
	"service-api/src/core/helpers/response"
	loggerInstance "service-api/src/core/logger"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func Limiter(maxRequests int64, duration time.Duration) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(duration), int(maxRequests))

	return func(c *gin.Context) {

		// 在请求处理之后记录响应信息
		loggerInstance.R().Debugw(
			"limiter",
			zap.String("url", c.Request.URL.String()),
			zap.Any("headers", c.Request.Header),
			zap.String("method", c.Request.Method),
			zap.Any("body", c.Request.Body),
			zap.Int("status", c.Writer.Status()),
		)

		// 限流逻辑
		if limiter.Allow() == false {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, response.SingleCustom(
				c,
				"TOO MANY REQUESTS",
				http.StatusTooManyRequests,
				"REQUESTS.TOO_MANY_REQUESTS",
			),
			)
			return
		}

		// 继续处理请求
		c.Next()
	}
}
