package password

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service-api/src/api/controller"
	entity "service-api/src/app/entity/authorize/password"
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
	Service       password.RegisterService
	LoggerService authorize.LoggerService
}

// Authorizing
// @POST(path="authorizing")
func (p *RegisterController) Authorizing(c *gin.Context, req entity.RegisteringReq) *gin.Context {

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
// @POST(path="authorized")
func (p *RegisterController) Authorized(c *gin.Context, req entity.RegisteredReq) *gin.Context {
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
