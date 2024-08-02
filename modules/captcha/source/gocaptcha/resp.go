package gocaptcha

import (
	captchaclient "github.com/wenlng/go-captcha/captcha"
)

type Resp struct {
	Key     string                        `json:"key"`
	Token   string                        `json:"token"`
	Captcha RespBody                      `json:"captcha"`
	Answer  map[int]captchaclient.CharDot `json:"answer"`
}

type RespBody struct {
	Image string `json:"image"`
	Thumb string `json:"thumb"`
}

func (r *Resp) GetToken() string {
	return r.Token
}

func (r *Resp) GetKey() string {
	return r.Key
}

func (r *Resp) GetBody() any {
	return r.Captcha
}

func (r *Resp) GetAnswer() map[int]captchaclient.CharDot {
	return r.Answer
}
