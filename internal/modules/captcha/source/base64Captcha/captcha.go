package base64Captcha

import (
	"fmt"
	"framework/cache"
	"framework/exceptions"
	captchaclient "github.com/mojocn/base64Captcha"
	"service-api/internal/modules/captcha"
	"service-api/resources/lang"
	"time"
)

// https://github.com/mojocn/base64Captcha/blob/master/store_memory.go

type Opts struct {
	Height, Width, Len, dotCount int
	MaxSkew                      float64
	exp                          time.Duration
}

type Captcha struct {
	opts  Opts
	store *Store
}

func (c *Captcha) Generate(token string) (captcha.Resp[string, string], error) {
	return c.GenerateCustom(token)
}

func (c *Captcha) GenerateCustom(token string) (captcha.Resp[string, string], error) {
	driver := captchaclient.NewDriverDigit(
		c.opts.Height,
		c.opts.Width,
		c.opts.Len,
		c.opts.MaxSkew,
		c.opts.dotCount,
	)

	cli := captchaclient.NewCaptcha(driver, c.store)
	id, body, answer, err := cli.Generate()
	if err != nil {
		return nil, exceptions.New(lang.CaptchaErrorGenerateCode)
	}

	// 生成场景存储类型错误
	if cache.SetEx(fmt.Sprintf("captcha-%s-%s", token, id), id, c.opts.exp) == false {
		return nil, exceptions.New(lang.CaptchaErrorGenerateCode)
	}

	return &Resp{Captcha: body, Key: id, Answer: answer}, nil
}

func (c *Captcha) Verify(token, key, answer string, clear bool) bool {
	if cache.Exists(fmt.Sprintf("captcha-%s-%s", token, key)) {
		return false
	}

	return c.store.Verify(key, answer, clear) == false
}

func New() *Captcha {
	return Custom(
		Opts{
			70,
			130,
			4,
			100,
			0.8,
			1 * time.Minute,
		},
		NewStore(time.Duration(10)),
	)
}

func Custom(opts Opts, store *Store) *Captcha {
	return &Captcha{store: store, opts: opts}
}
