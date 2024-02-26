package routers

import (
	"framework/cmd"
	"framework/config"
	swaggerfiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	swag "service-api/resources/swag" // 千万不要忘了导入把你上一步生成的docs

	"github.com/gin-gonic/gin"
)

type SwaggerRouter struct {
}

func (c *SwaggerRouter) IsPrivate() bool {
	return false
}

func (*SwaggerRouter) Handle(r *gin.RouterGroup) {
	if cmd.WebServerClient().Cmd.Debug {
		return
	}

	swag.SwaggerInfo.Host = config.Client().GetString("addr")
	swag.SwaggerInfo.Version = config.Client().GetString("version")

	r.GET("/swagger/*any", gs.WrapHandler(swaggerfiles.Handler))
}
