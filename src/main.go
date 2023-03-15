package main

import (
	"github.com/gin-gonic/gin"
	"service-api/src/app/helpers"
)

func main() {
	resposne := helpers.NewResponse[map[string]string]()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, resposne.Single(map[string]string{"memssage": "a"}))
	})
	r.Run(":8000") // 监听并在 0.0.0.0:8080 上启动服务
}
