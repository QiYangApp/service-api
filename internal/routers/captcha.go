package routers

import (
	"frame/modules/router"

	"github.com/gin-gonic/gin"
	"service-api/internal/app/api/v1/captcha"
)

type CaptchaRouter struct {
}

func (*CaptchaRouter) Handle(r *gin.RouterGroup) {
	// 获取验证码
	r.GET("/captcha/:type", router.Bind(captcha.Index))
	// 验证验证码
	r.POST("/captcha/:type", router.Bind(captcha.Verify))
}
