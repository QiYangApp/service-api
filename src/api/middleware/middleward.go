package middleware

import (
	"time"

	"github.com/gin-contrib/gzip"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupMiddleware(r *gin.Engine) {
	r.Use(Limiter(20, time.Minute))
	// r.Use(logger())
	r.Use(recovery())
	r.Use(cors())
	r.Use(resposne())
	// 使用 Gzip 中间件
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))

}
