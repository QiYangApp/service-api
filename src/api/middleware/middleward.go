package middleware

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"time"
)

func SetupMiddleware(r *gin.Engine) {
	r.Use(Limiter(20, time.Minute))
	r.Use(Logger())
	r.Use(Recovery())
	r.Use(Cors())
	r.Use(Resposne())
	r.Use(CSRF())

	//使用 Gzip
	r.Use(gzip.Gzip(gzip.DefaultCompression))
}
