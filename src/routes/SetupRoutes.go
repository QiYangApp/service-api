package routes

import (
	"service-api/src/app/helpers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// 定义路由
	router.GET("/", func(c *gin.Context) {
		res := helpers.NewResponse[string]()
		c.JSON(200, res.Single(""))
	})
}
