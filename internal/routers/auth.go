package routers

import (
	"framework/router"
	"github.com/gin-gonic/gin"
	api "service-api/internal/app/api"
)

type AuthRouter struct {
}

func (c *AuthRouter) IsPrivate() bool {
	return false
}

func (*AuthRouter) Handle(r *gin.RouterGroup) {
	// 获取前置数据
	r.GET("/v1/auth/:type", router.Bind(api.Client.AuthApi.PreAuthorizationData))
	// 前置校验
	r.GET("/v1/auth/:type/verify", router.Bind(api.Client.AuthApi.PreAuthorizationVerify))
	// 授权进行中
	r.POST("/v1/auth/:type/authorizing", router.Bind(api.Client.CaptchaApi.Verify))
	// 授权完成
	r.POST("/v1/auth/:type/authorized", router.Bind(api.Client.CaptchaApi.Verify))
}
