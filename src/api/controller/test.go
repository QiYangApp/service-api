package controller

import (
	"github.com/gin-gonic/gin"
)

type TestController struct {
	AbstractController
}

// Test
// @GET(path="test")
func (t *TestController) Test(c *gin.Context) {
	c.JSON(200, "test")
}
