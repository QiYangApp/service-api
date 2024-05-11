package group

import (
	"frame/modules/router"

	"github.com/gin-gonic/gin"
	"service-api/internal/api/v1/captcha"
)

type CaptchaRouter struct {
}

func (*CaptchaRouter) Handle(r *gin.RouterGroup) {
	router.GetBind(r, PrefixV1+"/captcha/:type", captcha.Index)
	router.PostBind(r, PrefixV1+"/captcha/:type", captcha.Verify)
}
