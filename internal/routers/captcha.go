package routers

import (
	"framework/router"
	"github.com/gin-gonic/gin"
	"service-api/internal/app/api"
)

type CaptchaRouter struct {
}

func (c *CaptchaRouter) IsPrivate() bool {
	return false
}

func (*CaptchaRouter) Handle(r *gin.RouterGroup) {
	// 获取验证码
	r.GET("/captcha/:type", router.Bind(api.Client.CaptchaApi.Index))
	// 验证验证码
	r.POST("/captcha/:type", router.Bind(api.Client.CaptchaApi.Verify))
}
