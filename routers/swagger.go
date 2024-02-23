package routers

import (
	"framework"
	"framework/config"
	swaggerfiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	docs "service-api/docs" // 千万不要忘了导入把你上一步生成的docs

	"github.com/gin-gonic/gin"
)

type SwaggerRouter struct {
}

func (c *SwaggerRouter) IsPrivate() bool {
	return false
}

func (*SwaggerRouter) Handle(r *gin.RouterGroup) {
	if !framework.Client().Cmd.Debug {
	}

	docs.SwaggerInfo.Host = config.Client().GetString("addr")
	docs.SwaggerInfo.Version = config.Client().GetString("version")

	r.GET("/swagger/*any", gs.WrapHandler(swaggerfiles.Handler))
}
