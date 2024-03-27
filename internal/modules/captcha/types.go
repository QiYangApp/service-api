package captcha

type Resp[B any, A any] interface {
	GetKey() string

	GetToken() string

	GetBody() B

	GetAnswer() A
}

type Captcha[B any, A any] interface {
	Generate(token string) (Resp[B, A], error)
	GenerateCustom(token string) (Resp[B, A], error)
	Verify(token, key string, answer A, clear bool) bool
}
