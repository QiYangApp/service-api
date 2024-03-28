package captcha

type Type string

const (
	Image     Type = "Image"
	TextPoint Type = "TextPoint"
)

func (t Type) ToString() string {
	return string(t)
}

type Store interface {
	Get(key string, clear bool) string

	Set(id string, value string) error

	Verify(id string, answer string, clear bool) bool
}

type Captcha interface {
	Generate(token string) (*Resp, error)
	Verify(token, key string, answer any, clear bool) bool
}
