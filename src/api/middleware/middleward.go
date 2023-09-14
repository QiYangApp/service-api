package middleware

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"time"
)

func SetupMiddleware(r *gin.Engine) {
	r.Use(Recovery())

	r.Use(Resposne())
	r.Use(Limiter(20, time.Minute))
	r.Use(Secure())
	r.Use(Cors())
	r.Use(CSRF())
	r.Use(Logger())

	//使用 Gzip
	r.Use(gzip.Gzip(gzip.DefaultCompression))
}
