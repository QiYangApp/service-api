package gocaptcha

import (
	"framework/cache"
	"framework/exceptions"
	"framework/log"
	captchaclient "github.com/wenlng/go-captcha/captcha"
	"service-api/internal/modules/captcha"
	"service-api/resources/lang"
)

type Opts struct {
	Fonts   []string `mapstructure:"fonts"`
	Chars   []string `mapstructure:"chars"`
	Width   int      `mapstructure:"width"`
	Height  int      `mapstructure:"height"`
	Quality int      `mapstructure:"quality"`
	MinLen  int      `mapstructure:"minlen"`
	MaxLen  int      `mapstructure:"maxlen"`
}

type Captcha struct {
	opts  Opts
	store captcha.Store
}

func (*Captcha) getCacheKey(key, token string) string {
	return captcha.TextPoint.ToString() + "-" + token + "-" + key
}

func (c *Captcha) Generate(token string) (*captcha.Resp, error) {
	cli := captchaclient.GetCaptcha()
	cli.SetFont(c.opts.Fonts)
	cli.SetImageSize(captchaclient.Size{Width: c.opts.Width, Height: c.opts.Height})
	cli.SetRangCheckTextLen(captchaclient.RangeVal{Max: c.opts.MaxLen, Min: c.opts.MinLen})

	if err := cli.SetRangChars(c.opts.Chars); err != nil {
		log.Client.Sugar().Warn("CAPTCHA SET RANGE CHARS ERR. ", err)
		return nil, err
	}

	answer, imageBase64, thumbImageBase64, key, err := cli.Generate()
	if err != nil {
		return nil, err
	}

	if err := c.store.Set(c.getCacheKey(token, key), key); err != nil {
		return nil, exceptions.New(lang.CaptchaErrorGenerateCode)
	}

	return &captcha.Resp{
		Token:  token,
		Key:    key,
		Answer: answer,
		Captcha: RespBody{
			Image: imageBase64,
			Thumb: thumbImageBase64,
		},
	}, nil
}

func (c *Captcha) Verify(token, key string, answer any, clear bool) bool {
	if cache.Exists(c.getCacheKey(token, key)) {
		return false
	}

	return c.store.Verify(key, answer, clear) == false
}

func New(opts Opts, store captcha.Store) captcha.Captcha {
	return &Captcha{store: store, opts: opts}
}
