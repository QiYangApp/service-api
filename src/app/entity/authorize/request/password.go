package request

import (
	"github.com/google/uuid"
	"service-api/src/app/entity/http"
)

type PasswordLoginCheckReq struct {
	http.ReqType `json:"_"`
	Account      string `form:"account"`
}

type PasswordLoginCheckRsp struct {
	http.RespType `json:"_"`
	State         bool `json:"state"`
}

type PasswordLoggingReq struct {
	http.ReqType
	Account string `form:"account"`
}

type PasswordLoggingRsp struct {
	http.RespType `json:"-"`
	Id            string `json:"id"`
	Captcha       string `json:"captcha"`
}

type PasswordLoggedReq struct {
	http.ReqType
	CaptchaId string `form:"captchaId"`
	Captcha   string `form:"captcha"`
	Account   string `form:"account"`
	Password  string `form:"password"`
}

type PasswordLoggedRsp struct {
	http.RespType `json:"-"`
	Account       string    `json:"account,omitempty"`
	Avatar        string    `json:"avatar,omitempty"`
	Nickname      string    `json:"nickname,omitempty"`
	MemberId      uuid.UUID `json:"member_id,omitempty"`
	Token         string    `json:"token,omitempty"`
}
