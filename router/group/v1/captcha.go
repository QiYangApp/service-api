package v1

import (
	"frame/modules/router"
	"service-api/internal/net/api/v1/captcha"

	"github.com/gin-gonic/gin"
)

type CaptchaRouter struct {
}

func (*CaptchaRouter) Handle(r *gin.RouterGroup) {
	router.GetBind(r, "v1/captcha/:type", captcha.Index)
}
