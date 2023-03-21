package middleware

import (
	"fmt"
	"net/http"
	"service-api/src/app/helpers"

	"github.com/gin-gonic/gin"
)

func cache() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 自定义输出内容
				errMsg := fmt.Sprintf("panic error: %v", err)

				resp := helpers.NewResponse[string]()
				resp.SetCode(http.StatusInternalServerError)
				resp.SetMessage(errMsg)

				c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Single(""))
			}
		}()
		c.Next()
	}
}