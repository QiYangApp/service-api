package authorize

import (
	"github.com/gin-gonic/gin"
	"service-api/src/api/controller"
	"service-api/src/app/entity/authorize/request"
	"service-api/src/app/services/authorize"
	"service-api/src/core/helpers/response"
)

// PasswordLoginController
// @BasePath(path="/authorize/password")
type PasswordLoginController struct {
	controller.InjectController
	PasswordLoginService authorize.PasswordLoginService
}

// Check
// @GET(path="check")
func (p *PasswordLoginController) Check(c *gin.Context, req request.PasswordLoginCheckReq) *gin.Context {
	return response.RSuccess(c, p.PasswordLoginService.Check(req)).ToJson()
}

// Authorizing
// @GET(path="authorizing")
func (p *PasswordLoginController) Authorizing(c *gin.Context, req request.PasswordLoggingReq) {

	// todo 待处理账号校验
	//d := t.PasswordLoginService.Authorizing(true)

	//return response.RSuccess[any](c, d).ToJson()
}

// Authorized
// @GET(path="Authorized")
func (p *PasswordLoginController) Authorized(c *gin.Context, req request.PasswordLoggedReq) *gin.Context {
	d := p.PasswordLoginService.Authorized(req)

	return response.RSuccess(c, d).ToJson()
}
