package controller

import (
	"github.com/gin-gonic/gin"
	"service-api/src/app/services/authorize"
)

type TestController struct {
	*InjectController
	PasswordLoginService authorize.PasswordLoginService
}

// Test
// @GET(path="test")
func (t *TestController) Test(c *gin.Context) {
	c.JSON(200, "test")
}
