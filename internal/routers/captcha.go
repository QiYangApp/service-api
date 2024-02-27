package routers

import (
	"framework/router"
	"github.com/gin-gonic/gin"
	v1 "service-api/internal/app/http/v1"
)

type CaptchaRouter struct {
}

func (c *CaptchaRouter) IsPrivate() bool {
	return false
}

func (*CaptchaRouter) Handle(r *gin.RouterGroup) {
	// 获取验证码
	r.GET("/captcha/:type", router.Bind(v1.Client.CaptchaApi.Index))
	// 验证验证码
	r.POST("/captcha/:type", router.Bind(v1.Client.CaptchaApi.Verify))
}
