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

// RegisterController
// @BasePath(path="/authorize/password/register")
type RegisterController struct {
	controller.InjectController
	Service       password.LoginService
	LoggerService authorize.LoggerService
}

// Check
// @GET(path="check")
func (p *RegisterController) Check(c *gin.Context, req password2.LoginCheckReq) *gin.Context {

	if req.Account == "" {
		return response.RError(c, errors.WithMes(i18n.EmptyAccount), http.StatusBadRequest, nil).ToJson()
	}

	body, err := p.Service.Check(req)
	if err != nil {
		return response.RError(c, err, http.StatusBadRequest, nil).ToJson()
	}

	return response.RSuccess(c, body).ToJson()
}

// Authorizing
// @GET(path="authorizing")
func (p *RegisterController) Authorizing(c *gin.Context, req password2.LoggingReq) *gin.Context {

	if req.Account == "" {
		return response.RError(c, errors.WithMes(i18n.EmptyAccount), http.StatusBadRequest, nil).ToJson()
	}

	body, err := p.Service.Authorizing(req)
	if err != nil {
		return response.RError(c, err, http.StatusBadRequest, nil).ToJson()
	}

	return response.RSuccess(c, body).ToJson()
}

// Authorized
// @GET(path="authorized")
func (p *RegisterController) Authorized(c *gin.Context, req password2.LoggedReq) *gin.Context {
	member, err := p.Service.Authorized(req)
	if err != nil {
		return response.RError(c, err, http.StatusBadRequest, nil).ToJson()
	}

	_, err = p.LoggerService.Write(c, member.MemberId)
	if err != nil {
		return response.RError(c, err, http.StatusBadRequest, nil).ToJson()
	}

	return response.RSuccess(c, member).ToJson()
}
