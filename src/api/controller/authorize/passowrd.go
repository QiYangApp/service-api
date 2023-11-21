package authorize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service-api/src/api/controller"
	"service-api/src/app/entity/authorize/request"
	"service-api/src/app/services/authorize"
	"service-api/src/app/services/authorize/password"
	"service-api/src/core/helpers/response"
	"service-api/src/enums/i18n"
	"service-api/src/errors"
)

// PasswordLoginController
// @BasePath(path="/authorize/password")
type PasswordLoginController struct {
	controller.InjectController
	PasswordLoginService password.LoginService
	LoggerService        authorize.LoggerService
}

// Check
// @GET(path="check")
func (p *PasswordLoginController) Check(c *gin.Context, req request.PasswordLoginCheckReq) *gin.Context {

	if req.Account == "" {
		return response.RError(c, errors.WithMes(i18n.EmptyAccount), http.StatusBadRequest, i18n.EmptyAccount).ToJson()
	}

	body, err := p.PasswordLoginService.Check(req)
	if err != nil {
		return response.RError(c, err, http.StatusBadRequest, i18n.StateError).ToJson()
	}

	return response.RSuccess(c, body).ToJson()
}

// Authorizing
// @GET(path="authorizing")
func (p *PasswordLoginController) Authorizing(c *gin.Context, req request.PasswordLoggingReq) *gin.Context {

	if req.Account == "" {
		return response.RError(c, errors.WithMes(i18n.EmptyAccount), http.StatusBadRequest, i18n.EmptyAccount).ToJson()
	}

	body, err := p.PasswordLoginService.Authorizing(req)
	if err != nil {
		return response.RError(c, err, http.StatusBadRequest, i18n.StateError).ToJson()
	}

	return response.RSuccess(c, body).ToJson()
}

// Authorized
// @GET(path="authorized")
func (p *PasswordLoginController) Authorized(c *gin.Context, req request.PasswordLoggedReq) *gin.Context {
	member, err := p.PasswordLoginService.Authorized(req)
	if err != nil {
		return response.RError(c, err, http.StatusBadRequest, i18n.StateError).ToJson()
	}

	_, err = p.LoggerService.Write(c, member.MemberId)
	if err != nil {
		return response.RError(c, err, http.StatusBadRequest, i18n.StateError).ToJson()
	}

	return response.RSuccess(c, member).ToJson()
}

type PasswordRegisterController struct {
	controller.InjectController
	PasswordRegisterService password.LoginService
	LoggerService           authorize.LoggerService
}
