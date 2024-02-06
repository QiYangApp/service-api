package main

import (
	"app"
	"service-api/src/providers"
)

// go run -mod=mod entgo.io/ent/cmd/ent  generate  .\ent\schema\  --target .\src\repo\models
//
//go:generate go run ./src/script/ioc/main.go
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate "./src/app/models/ent/schema"
func main() {
	context := app.New()
	context.PackageName = "service-api/src"
	context.Providers(
		&providers.Cron{},
		&providers.Database{},
		&providers.Router{},
	)

	context.Run(&app.Cmd{})
}
