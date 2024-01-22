package middleware

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"service-api/src/core/helpers/response"
	"service-api/src/core/logger"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	var custom = func(c *gin.Context, err any) {
		// 自定义输出内容
		errMsg := fmt.Sprintf("error: %v", err)

		logger.Recovery().Errorln("url: %s, err: %v", c.Request.URL, zap.Error(err.(error)))

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