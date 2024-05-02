package main

import (
	"frame"
	"frame/cmd"
	"service-api/providers"
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

// go run -mod=mod entgo.io/ent/cmd/ent new --target "./internal/ent/schema" --template "./internal/ent/schema/template/entinit.tmpl"
//
//go:generate go run -mod=mod  github.com/swaggo/swag/cmd/swag init --output  ./resources/swag
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate "./internal/ent/schema" --template "./internal/ent/schema/template/gen"
func main() {

	client := cmd.WebCli()
	client.Providers = append(client.Providers,
		providers.SettingRegister,
		providers.CronRegister,
		providers.DatabaseRegister,
		providers.RouterRegister,
	)

	frame.NewApp()
}
