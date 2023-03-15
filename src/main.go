package main

import (
	"github.com/gin-gonic/gin"
	"service-api/src/core/system"
)

func main() {
	var gin = gin.Default()

	system.Start(gin)
}
