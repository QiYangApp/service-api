package password

import (
	"github.com/google/uuid"
	"service-api/src/app/entity/http"
)

type RegisteringReq struct {
	http.ReqType `json:"_"`
	ErrMsg       `json:"-"`
	Email        string `form:"email" json:"email" binding:"required,email" `
}

type RegisteringRsp struct {
	http.RespType `json:"-"`
	Id            string `json:"id"`
	Captcha       string `json:"captcha"`
}

type RegisteredReq struct {
	http.ReqType
	ErrMsg         `json:"-"`
	CaptchaId      string `form:"captchaId" json:"captcha_id" binding:"required" `
	Captcha        string `form:"captcha" json:"captcha" binding:"required" `
	Email          string `form:"email" json:"email" binding:"required,email" `
	Nickname       string `form:"nickname" json:"nickname" binding:"required,min=4,max=10" `
	Password       string `form:"password" json:"password" binding:"required,max=16,min=8" `
	RepeatPassword string `form:"repeat_password" json:"repeat_password" binding:"required,eqcsfield=Password" `
}

type RegisteredRsp struct {
	http.RespType `json:"-"`
	Email         string    `json:"email"`
	Avatar        string    `json:"avatar"`
	Nickname      string    `json:"nickname"`
	MemberId      uuid.UUID `json:"member_id"`
	Token         string    `json:"token"`
}
