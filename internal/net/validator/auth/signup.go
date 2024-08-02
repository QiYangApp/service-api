package auth

import (
	"frame/modules/router"
	"frame/modules/translate"
	"github.com/gin-gonic/gin"
	"service-api/resources/translate/messages"
)

type SignUpForm struct {
	UserName string `form:"userName" json:"userName" binding:"required,max=40"`
	Email    string `form:"email" json:"email" binding:"required,email,max=254"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=255"`
	Retype   string `form:"retype" json:"retype" binding:"required,eqfield=Password"`
}

func (SignUpForm) GetMessage(c *gin.Context) router.ValidatorMessages {
	return router.ValidatorMessages{
		"userName.required": translate.MustGetMessage(c, messages.UserNameEmpty),
		"userName.max":      translate.MustGetMessage(c, messages.UserNameMaxLength),
		"email.required":    translate.MustGetMessage(c, messages.UserEmailEmpty),
		"email.max":         translate.MustGetMessage(c, messages.UserEmailMaxLength),
		"email.email":       translate.MustGetMessage(c, messages.UserEmailFormatInvalid),
		"password.required": translate.MustGetMessage(c, messages.UserPasswordEmpty),
		"password.min":      translate.MustGetMessage(c, messages.UserPasswordCheckFailed),
		"password.max":      translate.MustGetMessage(c, messages.UserPasswordCheckFailed),
		"retype.required":   translate.MustGetMessage(c, messages.UserPasswordEmpty),
		"retype.eqfield":    translate.MustGetMessage(c, messages.UserPasswordCheckFailed),
	}
}
