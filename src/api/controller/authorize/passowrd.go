package authorize

import (
	"github.com/gin-gonic/gin"
	"service-api/src/api/controller"
	"service-api/src/app/entity/http"
	"service-api/src/core/helpers/response"
	"service-api/src/core/logger"
)

type PasswordController[P http.VerifyType] struct {
	controller.AbstractController
}

func (PasswordController[P]) Authorizing(c *gin.Context, p P) *gin.Context {
	logger.S().Info(p)

	data := response.RSuccess[bool](c, true)

	c.JSON(data.GetCode(), data)

	return c
}

func (PasswordController[P]) Authorized(c *gin.Context, p P) *gin.Context {
	logger.S().Info(p)

	data := response.RSuccess[P](c, p)

	c.JSON(data.GetCode(), data)

	return c
}
