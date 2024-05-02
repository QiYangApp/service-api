package routers

import (
	"service-api/internal/modules/setting"
	swag "service-api/resources/swag" // 千万不要忘了导入把你上一步生成的docs

	swaggerfiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type SwaggerRouter struct {
}

func (*SwaggerRouter) Handle(r *gin.RouterGroup) {
	if setting.AppSetting.Debug {
		return
	}

	swag.SwaggerInfo.Title = setting.AppSetting.Name
	swag.SwaggerInfo.Host = setting.AppSetting.Addr
	swag.SwaggerInfo.Version = setting.AppSetting.Version

	r.GET("/swagger/*any", gs.WrapHandler(swaggerfiles.Handler))
}
