package main

import (
	"service-api/src/core/system"

	"github.com/gin-gonic/gin"
)

//go:generate ioc "./src/api/controller"
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate "./src/models/ent/schema"
func main() {
	system.Start(gin.New())
}
