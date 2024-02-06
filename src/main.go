package main

import (
	"app"
	"service-api/src/providers"
)

//go:generate go run ./pkg/app/script/router/main.go
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate "./src/app/models/ent/schema"
func main() {
	context := app.New()
	context.PackageName = "service-api/src"
	context.Providers(
		&providers.Cron{},
		&providers.Database{},
	)

	context.Run(&app.Cmd{})
}
