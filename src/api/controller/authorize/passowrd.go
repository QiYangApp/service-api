package authorize

import (
	"github.com/gin-gonic/gin"
	"service-api/src/api/controller"
	"service-api/src/app/services/authorize"
	"service-api/src/core/helpers/response"
)

type PasswordLoginController struct {
	*controller.InjectController
	PasswordLoginService authorize.PasswordLoginService
}

// Check
// @GET(path="check")
func (*PasswordLoginController) Check(c *gin.Context) {
	//return c
}

func (t *PasswordLoginController) Authorizing(c *gin.Context) {

	// todo 待处理账号校验
	//d := t.PasswordLoginService.Authorizing(true)

	//return response.RSuccess[any](c, d).ToJson()
}

func (t *PasswordLoginController) Authorized(c *gin.Context) *gin.Context {
	d := t.PasswordLoginService.Authorizing("ssss")

	return response.RSuccess(c, d).ToJson()
}
