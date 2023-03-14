package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"service-api/src/app/helpers"
	"time"
)

func Limiter(maxRequests int64, duration time.Duration) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(duration), int(maxRequests))

	return func(c *gin.Context) {
		// 限流逻辑
		if limiter.Allow() == false {
			response := helpers.NewResponse[string]()
			response.SetCode(http.StatusTooManyRequests)
			response.SetMessage("too many requests")
			response.SetState(helpers.Fail)

			c.AbortWithStatusJSON(http.StatusTooManyRequests, response.Single(""))
			return
		}

		// 继续处理请求
		c.Next()
	}
}
