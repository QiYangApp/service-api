package main

import (
	"service-api/src/core/system"

	"github.com/gin-gonic/gin"
)

func main() {
	system.Start(gin.Default())
}
