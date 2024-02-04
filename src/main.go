package main

import (
	"app"
	"service-api/src/api/providers"
)

//go:generate go run ./src/script/ioc/main.go
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate "./src/app/models/ent/schema"
func main() {
	//system.Start(gin.New())
	context := app.New()
	context.Providers(
		&providers.Cron{},
	)

	context.Run(&app.Cmd{})
}
