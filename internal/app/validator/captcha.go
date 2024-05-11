package validator

import (
	"frame/modules/router"
	"service-api/internal/modules/setting"
	lang "service-api/resources/i18n"

	"github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
)

type CaptchaRequest struct {
	Msg
	Token string `form:"token" binding:"required"`
	Type  string `uri:"type"`
}

type CaptchaResponse struct {
	Id      string `json:"id"`
	Captcha any    `json:"captcha"`
	Token   string `json:"token"`
}

type CaptchaVerifyRequest struct {
	Msg
	Key    string                 `uri:"key"`
	Type   setting.CaptchaFeature `uri:"type"`
	Token  string                 `form:"token" binding:"required"`
	Id     string                 `form:"id" binding:"required"`
	Answer string                 `form:"answer" binding:"required"`
}

type Msg struct {
}

func (Msg) GetMessage(c *gin.Context) router.ValidatorMessages {
	return router.ValidatorMessages{
		"Token.required":  i18n.MustGetMessage(c, lang.CaptchaEmpty),
		"Id.required":     i18n.MustGetMessage(c, lang.CaptchaEmptyId),
		"Answer.required": i18n.MustGetMessage(c, lang.CaptchaEmptyId),
	}
}
