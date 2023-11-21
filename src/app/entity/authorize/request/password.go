package request

import (
	"github.com/google/uuid"
	"service-api/src/app/entity/http"
)

type PasswordLoginCheckReq struct {
	http.ReqType `json:"_"`
	Account      string `form:"account" json:"account" binding:"required" `
}

type PasswordLoginCheckRsp struct {
	http.RespType `json:"_"`
	State         bool `json:"state"`
}

type PasswordLoggingReq struct {
	http.ReqType
	Account string `form:"account" json:"account" binding:"required" `
}

type PasswordLoggingRsp struct {
	http.RespType `json:"-"`
	Id            string `json:"id"`
	Captcha       string `json:"captcha"`
}

type PasswordLoggedReq struct {
	http.ReqType
	CaptchaId string `form:"captchaId" json:"captcha_id" binding:"required" `
	Captcha   string `form:"captcha" json:"captcha" binding:"required" `
	Account   string `form:"account" json:"account" binding:"required" `
	Password  string `form:"password" json:"password" binding:"required" `
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
