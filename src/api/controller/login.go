package controller

import (
	"app/response"
	"app/router"
	"github.com/gin-gonic/gin"
)

// LoginController
// @BasePath(path="/authorize/password/login")
type LoginController struct {
	router.Inject
}

// Check
// @GET(path="check")
func (p *LoginController) Check(c *gin.Context, req any) *gin.Context {

	return response.RSuccess(c, req).ToJson()
}

// Authorizing
// @GET(path="authorizing")
func (p *LoginController) Authorizing(c *gin.Context, req any) *gin.Context {
	return response.RSuccess(c, req).ToJson()
}

// Authorized
// @POST(path="authorized")
func (p *LoginController) Authorized(c *gin.Context, req any) *gin.Context {
	return response.RSuccess(c, req).ToJson()
}
