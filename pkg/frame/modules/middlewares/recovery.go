package middlewares

import (
	"fmt"
	"frame/conf"
	"frame/modules/log"
	"frame/modules/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Recovery() gin.HandlerFunc {
	if conf.IsDebug() {
		return gin.Recovery()
	}

	var custom = func(c *gin.Context, err any) {
		defer func() {
			// 自定义输出内容
			errMsg := fmt.Sprintf("%v", err)

			log.Sugar().Errorf("url: %s, err: %s", c.Request.URL, errMsg)

			resp.Fail(
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
