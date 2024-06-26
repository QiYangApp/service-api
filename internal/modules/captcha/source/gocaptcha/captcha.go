package gocaptcha

import (
	"errors"
	"fmt"
	"frame/modules/cache"
	"frame/modules/log"

	"service-api/internal/modules/captcha"
	lang "service-api/resources/i18n"
	"strconv"

	captchaclient "github.com/wenlng/go-captcha/captcha"
)

type Opts struct {
	Fonts   []string `mapstructure:"fonts"`
	Chars   []string `mapstructure:"chars"`
	Asset   []string `mapstructure:"asset"`
	Width   int      `mapstructure:"width"`
	Height  int      `mapstructure:"height"`
	Quality int      `mapstructure:"quality"`
	MinLen  int      `mapstructure:"min_len"`
	MaxLen  int      `mapstructure:"max_len"`
}

type Captcha struct {
	opts  Opts
	store captcha.Store[map[int]captchaclient.CharDot]
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
		log.Sugar().Warn("CAPTCHA SET RANGE CHARS ERR. ", err)
		return nil, err
	}

	answer, imageBase64, thumbImageBase64, key, err := cli.Generate()
	if err != nil {
		return nil, err
	}

	if err := c.store.Set(c.getCacheKey(token, key), answer); err != nil {
		return nil, errors.New(lang.CaptchaErrorGenerateCode)
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
	dots := answer.(map[int]captchaclient.CharDot)
	if len(dots) == 0 || key == "" || token == "" {
		return false
	}

	cachekey := c.getCacheKey(token, key)
	if cache.Exists(cachekey) {
		return false
	}

	dct := c.store.Get(cachekey, clear)
	if !dct.Has() || len(dots)*2 != len(dct.Value()) {
		return false
	}

	check := false
	for i, dot := range dct.Value() {
		j := i * 2
		k := i*2 + 1
		sx, _ := strconv.ParseFloat(fmt.Sprintf("%v", dots[j]), 64)
		sy, _ := strconv.ParseFloat(fmt.Sprintf("%v", dots[k]), 64)

		// 检测点位置
		// chkRet = captcha.CheckPointDist(int64(sx), int64(sy), int64(dot.Dx), int64(dot.Dy), int64(dot.Width), int64(dot.Height))

		// 校验点的位置,在原有的区域上添加额外边距进行扩张计算区域,不推荐设置过大的padding
		// 例如：文本的宽和高为30，校验范围x为10-40，y为15-45，此时扩充5像素后校验范围宽和高为40，则校验范围x为5-45，位置y为10-50
		check = captchaclient.CheckPointDistWithPadding(int64(sx), int64(sy), int64(dot.Dx), int64(dot.Dy), int64(dot.Width), int64(dot.Height), 5)
		if !check {
			break
		}
	}

	return check
}

func New(opts Opts, store captcha.Store[map[int]captchaclient.CharDot]) captcha.Captcha {
	return &Captcha{store: store, opts: opts}
}
