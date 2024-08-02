package captcha

type Resp struct {
	Type    Type
	Key     string `json:"key"`
	Token   string `json:"token"`
	Captcha any    `json:"captcha"`
	Answer  any    `json:"answer"`
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

func (r *Resp) GetAnswer() any {
	return r.Answer
}
