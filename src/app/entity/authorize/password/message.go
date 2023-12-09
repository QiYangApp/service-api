package password

import (
	"github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"service-api/src/app/entity/http"
	msg "service-api/src/enums/i18n"
)

type ErrMsg struct {
}

func (ErrMsg) GetMessage(c *gin.Context) http.ValidatorMessages {
	return http.ValidatorMessages{
		"Email.required": i18n.MustGetMessage(c, msg.EmptyEmail),
		"Email.email":    i18n.MustGetMessage(c, msg.ErrorFormatEmail),
	}
}
