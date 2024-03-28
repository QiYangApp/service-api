package captcha

import "framework/utils/optional"

type Type string

const (
	Image     Type = "Image"
	TextPoint Type = "TextPoint"
)

func (t Type) ToString() string {
	return string(t)
}

type Store[T any] interface {
	Exist(key string) bool

	Del(key string) bool

	Get(key string, clear bool) optional.Option[T]

	Set(key string, value any) error
}

type Captcha interface {
	Generate(token string) (*Resp, error)
	Verify(token, key string, answer any, clear bool) bool
}
