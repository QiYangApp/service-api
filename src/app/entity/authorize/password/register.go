package password

import (
	"github.com/google/uuid"
	"service-api/src/app/entity/http"
)

type RegisteringReq struct {
	http.ReqType `json:"_"`
	ErrMsg       `json:"-"`
	Account      string `form:"account" json:"account" binding:"required" `
}

type RegisteringRsp struct {
	http.RespType `json:"-"`
	Id            string `json:"id"`
	Captcha       string `json:"captcha"`
}

type RegisterCheckReq struct {
	http.ReqType `json:"_"`
	ErrMsg       `json:"-"`
	Account      string `form:"account" json:"account" binding:"required" `
	Email        string `form:"email" json:"email" binding:"required" `
}

type RegisterCheckRsp struct {
	http.RespType `json:"_"`
	State         bool `json:"state"`
}

type RegisteredReq struct {
	http.ReqType
	ErrMsg    `json:"-"`
	CaptchaId string `form:"captchaId" json:"captcha_id" binding:"required" `
	Captcha   string `form:"captcha" json:"captcha" binding:"required" `
	Account   string `form:"account" json:"account" binding:"required" `
	Password  string `form:"password" json:"password" binding:"required" `
}

type RegisteredRsp struct {
	http.RespType `json:"-"`
	Account       string    `json:"account"`
	Avatar        string    `json:"avatar"`
	Nickname      string    `json:"nickname"`
	MemberId      uuid.UUID `json:"member_id"`
	Token         string    `json:"token"`
}
