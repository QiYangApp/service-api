package middleware

import (
	"fmt"
	"net/http"
	"service-api/src/core/helpers/response"

	"github.com/gin-gonic/gin"
)

func recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 自定义输出内容
				errMsg := fmt.Sprintf("panic error: %v", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, response.SingleCustom(errMsg, http.StatusInternalServerError, "FAIL"))
			}
		}()
		c.Next()
	}
}
