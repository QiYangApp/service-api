package validator

import (
	"frame/modules/router"
	"frame/modules/translate"
	"service-api/internal/modules/setting"
	"service-api/resources/translate/messages"

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
	Key    string                 `form:"key" json:"key"`
	Type   setting.CaptchaFeature `form:"type" json:"type"`
	Token  string                 `form:"token" json:"token"`
	Id     string                 `form:"id" json:"id"`
	Answer string                 `form:"answer" json:"answer"`
}

type Msg struct {
}

func (Msg) GetMessage(c *gin.Context) router.ValidatorMessages {
	return router.ValidatorMessages{
		"Token.required":  translate.MustGetMessage(c, messages.CaptchaInputEmpty.ID),
		"Id.required":     translate.MustGetMessage(c, messages.CaptchaIdEmpty.ID),
		"Answer.required": translate.MustGetMessage(c, messages.CaptchaIdEmpty.ID),
	}
}
