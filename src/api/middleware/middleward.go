package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

func SetupMiddleware(r *gin.Engine) {
	r.Use(Limiter(20, time.Minute))
	r.Use(logger())
	r.Use(recovery())
	r.Use(cors())
	r.Use(resposne())
}
