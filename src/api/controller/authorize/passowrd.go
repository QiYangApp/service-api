package authorize

import (
	"github.com/gin-gonic/gin"
	"service-api/src/app/entity/http"
	"service-api/src/app/services/authorize"
	"service-api/src/core/helpers/response"
	"service-api/src/core/inject"
)

// PasswordLoginController
// @BasePath(path="/api/authorize/password")
type PasswordLoginController struct {
	inject.Controller
	PasswordLoginService authorize.PasswordLoginService
}

// Check
// @GET(path="check")
func (*PasswordLoginController) Check(c *gin.Context, p http.VerifyType) *gin.Context {
	return c
}

func (t *PasswordLoginController) Authorizing(c *gin.Context, p http.VerifyType) *gin.Context {

	// todo 待处理账号校验
	d := t.PasswordLoginService.Authorizing(true)

	return response.RSuccess[any](c, d).ToJson()
}

func (t *PasswordLoginController) Authorized(c *gin.Context, p http.VerifyType) *gin.Context {
	d := t.PasswordLoginService.Authorizing("ssss")

	return response.RSuccess(c, d).ToJson()
}
