package routers

import (
	"framework/config"
	swaggerfiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"service-api/internal/modules/setting"
	swag "service-api/resources/swag" // 千万不要忘了导入把你上一步生成的docs

	"github.com/gin-gonic/gin"
)

type SwaggerRouter struct {
}

func (*SwaggerRouter) Handle(private, public *gin.RouterGroup) {
	if setting.AppSetting.Debug {
		return
	}

	swag.SwaggerInfo.Host = config.Client.GetString("addr")
	swag.SwaggerInfo.Version = config.Client.GetString("version")

	public.GET("/swagger/*any", gs.WrapHandler(swaggerfiles.Handler))
}
