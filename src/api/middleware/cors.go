package middleware

import (
	corsService "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func cors() gin.HandlerFunc {
	return corsService.New(corsService.Config{
		AllowOrigins:     []string{"http://localhost"},                        // 允许跨域请求的源地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许跨域请求的 HTTP 方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 允许跨域请求的 HTTP Header
		ExposeHeaders:    []string{"Content-Length"},                          // 暴露的 HTTP Header
		AllowCredentials: true,                                                // 是否允许跨域请求携带 Cookie 等认证信息
		MaxAge:           12 * time.Hour,                                      // 缓存的最大时间
	})
}
