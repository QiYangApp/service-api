package routers

import (
	"framework/router"
	"github.com/gin-gonic/gin"
	"service-api/internal/app/api"
)

type CaptchaRouter struct {
}

func (*CaptchaRouter) Handle(privateRoute, publicRoute *gin.RouterGroup) {
	// 获取验证码
	publicRoute.GET("/captcha/:type", router.Bind(api.Client.CaptchaApi.Index))
	// 验证验证码
	publicRoute.POST("/captcha/:type", router.Bind(api.Client.CaptchaApi.Verify))
}
