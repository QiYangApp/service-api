package routers

import (
	"framework/router"
	"github.com/gin-gonic/gin"
	"service-api/api/http/controller"
)

type CaptchaRouter struct {
}

func (c *CaptchaRouter) IsPrivate() bool {
	return false
}

func (*CaptchaRouter) Handle(r *gin.RouterGroup) {
	// 获取验证码
	r.GET("/captcha/", router.Bind(controller.Client.CaptchaApi.Index))
	// 验证验证码
	r.GET("/captcha/:code", router.Bind(controller.Client.CaptchaApi.Verify))
}
