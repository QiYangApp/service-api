package middlewares

import (
	"framework/config"
	"framework/log"
	"framework/response"
	"framework/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

func Limiter() gin.HandlerFunc {

	if config.Client.GetBool("debug") {
		return func(c *gin.Context) {
			// 中间件逻辑
			c.Next()
		}
	}

	duration := config.Client.GetDuration("max_request_time")
	maxRequests := config.Client.GetInt("max_requests")
	limiter := utils.NewLimiter(rate.Every(duration*time.Second), maxRequests, "request")

	return func(c *gin.Context) {
		// 限流逻辑
		if limiter.Allow() == false {
			// 在请求处理之后记录响应信息
			log.Client.Info(
				"limiter",
				zap.String("url", c.Request.URL.String()),
				zap.Any("headers", c.Request.Header),
				zap.String("method", c.Request.Method),
				zap.Any("body", c.Request.Body),
				zap.Int("status", c.Writer.Status()),
			)

			c.AbortWithStatusJSON(
				http.StatusTooManyRequests,
				response.RFail(
					c,
					"",
					http.StatusTooManyRequests,
					"TooManyRequests",
				).ToStruct(),
			)
		}

		// 继续处理请求
		c.Next()
	}
}
