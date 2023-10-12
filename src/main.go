package main

import (
	"github.com/gin-gonic/gin"
	_ "service-api/src/api/controller"
	"service-api/src/core/system"
)

//go:generate ioc "./src/api/controller"
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate "./src/models/ent/schema"
func main() {
	system.Start(gin.New())
}
