package base64Captcha

import (
	"framework/cache"
	"framework/exceptions"
	captchaclient "github.com/mojocn/base64Captcha"
	"service-api/internal/modules/captcha"
	"service-api/resources/lang"
)

// https://github.com/mojocn/base64Captcha/blob/master/store_memory.go

type Opts struct {
	Height   int     `mapstructure:"height"`
	Width    int     `mapstructure:"width"`
	Len      int     `mapstructure:"len"`
	dotCount int     `mapstructure:"dot_count"`
	MaxSkew  float64 `mapstructure:"max_skew"`
}

type Captcha struct {
	Opts  Opts
	Store captcha.Store
}

func (*Captcha) getCacheKey(key, token string) string {
	return captcha.Image.ToString() + "-" + token + "-" + key
}

func (c *Captcha) Generate(token string) (*captcha.Resp, error) {

	driver := captchaclient.NewDriverDigit(
		c.Opts.Height,
		c.Opts.Width,
		c.Opts.Len,
		c.Opts.MaxSkew,
		c.Opts.dotCount,
	)

	cli := captchaclient.NewCaptcha(driver, c.Store)
	id, body, answer, err := cli.Generate()
	if err != nil {
		return nil, exceptions.New(lang.CaptchaErrorGenerateCode)
	}

	// 生成场景存储类型错误
	if err := c.Store.Set(c.getCacheKey(id, token), id); err != nil {
		return nil, exceptions.New(lang.CaptchaErrorGenerateCode)
	}

	return &captcha.Resp{Captcha: body, Key: id, Answer: answer, Type: captcha.Image}, nil
}

func (c *Captcha) Verify(token, key string, answer any, clear bool) bool {
	if cache.Exists(c.getCacheKey(key, token)) {
		return false
	}

	return c.Store.Verify(key, answer.(string), clear)
}

func New(opts Opts, store captcha.Store) captcha.Captcha {
	return &Captcha{Store: store, Opts: opts}
}
