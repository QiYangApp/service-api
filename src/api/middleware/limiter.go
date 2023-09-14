package middleware

import (
	"net/http"
	"service-api/src/core/helpers/response"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func Limiter(maxRequests int64, duration time.Duration) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(duration), int(maxRequests))

	return func(c *gin.Context) {
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
