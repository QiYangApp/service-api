package password

import (
	"github.com/google/uuid"
	"service-api/src/app/entity/http"
)

type LoginCheckReq struct {
	http.ReqType `json:"_"`
	ErrMsg       `json:"-"`
	Email        string `form:"email" json:"email" binding:"required,email" `
}

type LoginCheckRsp struct {
	http.RespType `json:"_"`
	State         bool `json:"state"`
}

type LoggingReq struct {
	http.ReqType
	ErrMsg `json:"-"`
	Email  string `form:"email" json:"email" binding:"required,email" `
}

type LoggingRsp struct {
	http.RespType `json:"-"`
	Id            string `json:"id"`
	Captcha       string `json:"captcha"`
}

type LoggedReq struct {
	http.ReqType
	ErrMsg    `json:"-"`
	CaptchaId string `form:"captchaId" json:"captcha_id" binding:"required" `
	Captcha   string `form:"captcha" json:"captcha" binding:"required" `
	Account   string `form:"account" json:"account" binding:"required" `
	Password  string `form:"password" json:"password" binding:"required" `
}

type LoggedRsp struct {
	http.RespType `json:"-"`
	Email         string    `json:"Email"`
	Avatar        string    `json:"avatar"`
	Nickname      string    `json:"nickname"`
	MemberId      uuid.UUID `json:"member_id"`
	Token         string    `json:"token"`
}
