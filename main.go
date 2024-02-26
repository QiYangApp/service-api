package main

import (
	"framework"
	"framework/cmd"
	"service-api/internal/app/providers"
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

//go:generate  go run -mod=mod  github.com/swaggo/swag/cmd/swag init --output  ./resources/swag
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate "./internal/ent/schema"
func main() {
	client := cmd.WebServerClient()

	client.Providers = append(client.Providers,
		&providers.Cron{},
		&providers.Database{},
		&providers.Router{},
	)

	framework.NewApp()
}
