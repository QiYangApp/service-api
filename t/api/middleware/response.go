package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Resposne() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
