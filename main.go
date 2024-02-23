package main

import (
	"framework"
	"service-api/api/providers"
)

// @title QiYang
// @version 1.0
// @description QiYangApiService
// @termsOfService http://swagger.io/terms/

// @contact.name QiYang
// @contact.url https://github.com/QiYangApp/service-api
// @contact.email notice@myadream.com

// @license.name BSD 3-Clause "New" or "Revised" License
// @license.url https://github.com/QiYangApp/service-api/blob/main/LICENSE

// @BasePath /api
// @securityDefinitions.basic BasicAuth

//go:generate  go run -mod=mod  github.com/swaggo/swag/cmd/swag init
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
