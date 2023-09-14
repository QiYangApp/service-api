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
				errMsg := fmt.Sprintf("error: %v", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, response.SingleCustom(c, errMsg, http.StatusInternalServerError, "STATE.FAIL"))
			}
		}()
		c.Next()
	}
}
