package middleware

import (
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func SetupMiddleware(r *gin.Engine) {
	r.Use(Limiter(20, time.Minute))
	r.Use(logger())
	r.Use(recovery())
	r.Use(cors())
	r.Use(resposne())
	// 使用 Gzip 中间件
	r.Use(gzip.Gzip(gzip.DefaultCompression))
}
