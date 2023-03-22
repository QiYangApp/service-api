package main

import (
	"github.com/gin-gonic/gin"
	"service-api/src/core/system"
)

func main() {
	system.Start(gin.Default())
}
