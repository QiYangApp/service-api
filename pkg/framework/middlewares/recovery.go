package middlewares

import (
	"fmt"
	"framework/config"
	"framework/log"
	"framework/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Recovery() gin.HandlerFunc {
	if config.Client.GetBool("debug") {
		return gin.Recovery()
	}

	var custom = func(c *gin.Context, err any) {
		defer func() {
			// 自定义输出内容
			errMsg := fmt.Sprintf("%v", err)

			log.Client.Sugar().Errorf("url: %s, err: %s", c.Request.URL, errMsg)

			response.RFail(
				c,
				err,
				http.StatusInternalServerError,
				errMsg,
			)
		}()
		c.Next()
	}

	return gin.CustomRecovery(custom)

}
