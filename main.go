package main

import (
	"framework"
	"service-api/api/providers"
)

//go:generate go run ./pkg/framework/script/router/main.go
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate "./ent/schema"
func main() {
	context := framework.Client()
	context.Providers(
		&providers.Cron{},
		&providers.Database{},
		&providers.Router{},
	)

	context.Run(&framework.Cmd{})
}
