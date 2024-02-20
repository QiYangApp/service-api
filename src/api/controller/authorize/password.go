package authorize

import (
	"app/response"
	"app/router"
	"github.com/gin-gonic/gin"
	"service-api/src/entity"
	"service-api/src/service/authorize"
)

type Password struct {
	router.Inject
	AuthorizeService authorize.PasswordService
}

// Login
// @GET(path="/authorizing/login")
func (p *Password) Login(c *gin.Context, req entity.LoginRequest) *gin.Context {

	p.AuthorizeService.Login(req)

	return response.RSuccess(c, req).ToJson()
}

// Register
// @GET(path="/authorizing/register")
func (p *Password) Register(c *gin.Context, req any) *gin.Context {
	return response.RSuccess(c, req).ToJson()
}
