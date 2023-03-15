package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	// 定义路由
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})
}
