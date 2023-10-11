package authorize

import (
	"github.com/gin-gonic/gin"
	"service-api/src/api/controller"
	"service-api/src/app/entity/http"
	"service-api/src/app/services/authorize"
	"service-api/src/core/helpers/response"
)

type PasswordLoginController[P http.VerifyType] struct {
	AuthorizedType[P]
	controller.AbstractController
}

func (PasswordLoginController[P]) CheckAccountExists(c *gin.Context, p P) *gin.Context {
	return c
}

func (PasswordLoginController[P]) Authorizing(c *gin.Context, p P) *gin.Context {

	// todo 待处理账号校验
	var a = authorize.PasswordLoginService{}
	d := a.Authorizing(true)

	return response.RSuccess[any](c, d).ToJson()
}

func (PasswordLoginController[P]) Authorized(c *gin.Context, p P) *gin.Context {
	return response.RSuccess[bool](c, true).ToJson()
}
