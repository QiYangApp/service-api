package main

import (
	"github.com/gin-gonic/gin"
	"service-api/src/core/system"
)

//go:generate go run ./src/script/ioc/main.go
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate "./src/app/models/ent/schema"
func main() {
	system.Start(gin.New())
}
