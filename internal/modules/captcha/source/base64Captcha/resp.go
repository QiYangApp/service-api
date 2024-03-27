package base64Captcha

type Resp struct {
	Key     string `json:"key"`
	Token   string `json:"token"`
	Captcha string `json:"captcha"`
	Answer  string `json:"answer"`
}

func (r *Resp) GetToken() string {
	return r.Token
}

func (r *Resp) GetKey() string {
	return r.Key
}

func (r *Resp) GetBody() string {
	return r.Captcha
}

func (r *Resp) GetAnswer() string {
	return r.Captcha
}
