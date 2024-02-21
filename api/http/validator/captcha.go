package validator

import (
	"framework/router"
	"github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
)

type CaptchaForm struct {
	Type  string `uri:"type" binding:"required" `
	Token string `form:"token" binding:"required"`
}

type ErrMsg struct {
}

func (ErrMsg) GetMessage(c *gin.Context) router.ValidatorMessages {
	return router.ValidatorMessages{
		"Email.required":     i18n.MustGetMessage(c, i18n.EmptyEmail),
		"Email.email":        i18n.MustGetMessage(c, i18n2.ErrorFormatEmail),
		"Captcha.required":   i18n.MustGetMessage(c, i18n2.CaptchaEmptyId),
		"CaptchaId.required": i18n.MustGetMessage(c, i18n2.CaptchaEmptyId),
		"Password.required":  i18n.MustGetMessage(c, i18n2.ErrorPassword),
	}
}

type CaptchaResponse struct {
	Id      string `json:"id"`
	Captcha string `json:"captcha"`
}
