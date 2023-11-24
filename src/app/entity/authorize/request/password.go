package request

import (
	"github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"service-api/src/app/entity/http"
	msg "service-api/src/enums/i18n"
)

type PasswordReqErrMsg struct {
}

func (PasswordReqErrMsg) GetMessage(c *gin.Context) http.ValidatorMessages {
	return http.ValidatorMessages{
		"Account.required": i18n.MustGetMessage(c, msg.EmptyAccount),
		"Email.required":   i18n.MustGetMessage(c, msg.EmptyEmail),
	}
}

type PasswordLoginCheckReq struct {
	http.ReqType      `json:"_"`
	PasswordReqErrMsg `json:"-"`
	Account           string `form:"account" json:"account" binding:"required" `
}

type PasswordLoginCheckRsp struct {
	http.RespType `json:"_"`
	State         bool `json:"state"`
}

type PasswordLoggingReq struct {
	http.ReqType
	PasswordReqErrMsg `json:"-"`
	Account           string `form:"account" json:"account" binding:"required" `
}

type PasswordLoggingRsp struct {
	http.RespType `json:"-"`
	Id            string `json:"id"`
	Captcha       string `json:"captcha"`
}

type PasswordLoggedReq struct {
	http.ReqType
	PasswordReqErrMsg `json:"-"`
	CaptchaId         string `form:"captchaId" json:"captcha_id" binding:"required" `
	Captcha           string `form:"captcha" json:"captcha" binding:"required" `
	Account           string `form:"account" json:"account" binding:"required" `
	Password          string `form:"password" json:"password" binding:"required" `
}

type PasswordLoggedRsp struct {
	http.RespType `json:"-"`
	Account       string    `json:"account"`
	Avatar        string    `json:"avatar"`
	Nickname      string    `json:"nickname"`
	MemberId      uuid.UUID `json:"member_id"`
	Token         string    `json:"token"`
}

type PasswordRegisterCheckReq struct {
	http.ReqType `json:"_"`
	Account      string `form:"account" json:"account" binding:"required" `
	Email        string `form:"email" json:"email" binding:"required" `
}

type PasswordRegisterCheckRsp struct {
	http.RespType `json:"_"`
	State         bool `json:"state"`
}
