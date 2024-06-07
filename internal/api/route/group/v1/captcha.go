package v1

import (
	"frame/modules/router"

	"github.com/gin-gonic/gin"
	"service-api/internal/api/v1/captcha"
)

type CaptchaRouter struct {
}

func (*CaptchaRouter) Handle(r *gin.RouterGroup) {
	router.GetBind(r, "v1/captcha/:type", captcha.Index)
	router.PostBind(r, "v1/captcha/:type", captcha.Verify)
}
