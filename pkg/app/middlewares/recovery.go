package middlewares

import (
	"app/config"
	"app/log"
	"app/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func Recovery() gin.HandlerFunc {
	if config.Client().GetBool("server.debug") {
		return gin.Recovery()
	}

	var custom = func(c *gin.Context, err any) {
		// 自定义输出内容
		errMsg := fmt.Sprintf("error: %v", err)

		log.Client().Sugar().Error("url: %s, exceptions: %v", c.Request.URL, zap.Error(err.(error)))

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			response.RFail(
				c,
				err,
				http.StatusInternalServerError,
				errMsg,
			).ToStruct(),
		)
	}

	return gin.CustomRecovery(custom)

}
