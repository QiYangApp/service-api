package authorize

import (
	"framework/response"
	"github.com/gin-gonic/gin"
	"service-api/entity"
	"service-api/services/authorize"
)

type PasswordSingUpApi struct {
	AuthorizeService authorize.PasswordSignUpService
}

func (p *PasswordSingUpApi) Login(c *gin.Context, req entity.LoginRequest) *gin.Context {

	p.AuthorizeService.Authorized(req)

	return response.RSuccess(c, req).ToJson()
}

func (p *PasswordSingUpApi) Register(c *gin.Context, req any) *gin.Context {
	return response.RSuccess(c, req).ToJson()
}

type PasswordSingInApi struct {
	AuthorizeService authorize.PasswordSignUpService
}
