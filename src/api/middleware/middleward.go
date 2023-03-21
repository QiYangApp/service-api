package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"github.com/gin-contrib/gzip"
)

func SetupMiddleware(r *gin.Engine) {
	r.Use(Limiter(20, time.Minute))
	r.Use(logger())
	r.Use(recovery())
	r.Use(cors())
	r.Use(resposne())

	r.Use(gzip.Gzip(gzip.DefaultCompression))
}
