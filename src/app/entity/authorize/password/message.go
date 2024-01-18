package password

import (
	"github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"service-api/src/app/entity/http"
	i18n2 "service-api/src/app/enums/i18n"
)

type ErrMsg struct {
}

func (ErrMsg) GetMessage(c *gin.Context) http.ValidatorMessages {
	return http.ValidatorMessages{
		"Email.required":     i18n.MustGetMessage(c, i18n2.EmptyEmail),
		"Email.email":        i18n.MustGetMessage(c, i18n2.ErrorFormatEmail),
		"Captcha.required":   i18n.MustGetMessage(c, i18n2.CaptchaEmptyId),
		"CaptchaId.required": i18n.MustGetMessage(c, i18n2.CaptchaEmptyId),
		"Password.required":  i18n.MustGetMessage(c, i18n2.ErrorPassword),
	}
}
