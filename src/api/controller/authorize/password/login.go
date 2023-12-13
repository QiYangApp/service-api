package password

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service-api/src/api/controller"
	password2 "service-api/src/app/entity/authorize/password"
	"service-api/src/app/services/authorize"
	"service-api/src/app/services/authorize/password"
	"service-api/src/core/helpers/response"
	"service-api/src/enums/i18n"
	"service-api/src/errors"
)

// LoginController
// @BasePath(path="/authorize/password/login")
type LoginController struct {
	controller.InjectController
	Service       password.LoginService
	LoggerService authorize.LoggerService
}

// Check
// @GET(path="check")
func (p *LoginController) Check(c *gin.Context, req password2.LoginCheckReq) *gin.Context {

	if req.Email == "" {
		return response.RError(c, errors.WithMes(i18n.EmptyEmail), http.StatusBadRequest, nil).ToJson()
	}

	body, err := p.Service.Check(req)
	if err != nil {
		return response.RError(c, err, http.StatusBadRequest, nil).ToJson()
	}

	return response.RSuccess(c, body).ToJson()
}

// Authorizing
// @GET(path="authorizing")
func (p *LoginController) Authorizing(c *gin.Context, req password2.LoggingReq) *gin.Context {

	if req.Email == "" {
		return response.RError(c, errors.WithMes(i18n.EmptyEmail), http.StatusBadRequest, nil).ToJson()
	}

	body, err := p.Service.Authorizing(req)
	if err != nil {
		return response.RError(c, err, http.StatusBadRequest, nil).ToJson()
	}

	return response.RSuccess(c, body).ToJson()
}

// Authorized
// @GET(path="authorized")
func (p *LoginController) Authorized(c *gin.Context, req password2.LoggedReq) *gin.Context {
	member, err := p.Service.Authorized(req)
	if err != nil {
		return response.RError(c, err, http.StatusBadRequest, nil).ToJson()
	}

	_, err = p.LoggerService.Write(c, member.MemberId, member.Token)
	if err != nil {
		return response.RError(c, err, http.StatusBadRequest, nil).ToJson()
	}

	return response.RSuccess(c, member).ToJson()
}
