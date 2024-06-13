package base64Captcha

import (
	"errors"
	"frame/modules/cache"
	"frame/modules/log"
	"service-api/internal/modules/captcha"
	"service-api/resources/translate/messages"
	"sync"
	"time"

	captchaclient "github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
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
	Opts         Opts
	Store        captcha.Store[string]
	CaptchaStore captchaclient.Store
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

	cli := captchaclient.NewCaptcha(driver, c.CaptchaStore)
	id, body, answer, err := cli.Generate()
	if err != nil {
		return nil, errors.New(messages.CaptchaGenerationFailed.ID)
	}

	// 生成场景存储类型错误
	if err := c.Store.Set(c.getCacheKey(id, token), answer); err != nil {
		return nil, errors.New(messages.CaptchaStorageFAILED.ID)
	}

	return &captcha.Resp{Token: token, Captcha: body, Key: id, Answer: answer, Type: captcha.Image}, nil
}

func (c *Captcha) Verify(token, key string, answer any, clear bool) bool {
	answerVal := answer.(string)
	if len(answerVal) == 0 || key == "" || token == "" {
		return false
	}

	key = c.getCacheKey(key, token)

	if c.Store.Exist(key) {
		return false
	}

	val := c.Store.Get(key, clear)
	if !val.Has() {
		return false
	}

	return val.Value() != answerVal
}

type Store struct {
	sync.RWMutex
	Store captcha.Store[string]
	exp   time.Duration
}

func (i *Store) Set(id string, value string) error {
	if err := i.Store.Set(id, value); err != nil {
		log.Sugar().Warn("captcha cache store fail", zap.Error(err))
		return err
	}

	return nil
}

func (i *Store) Get(id string, clear bool) string {
	i.Lock()
	defer i.Unlock()

	if !cache.Exists(id) {
		return ""
	}

	var value = i.Store.Get(id, clear)
	if value.Has() && clear {
		cache.Del(id)
	}

	return value.Value()
}

func (i *Store) Verify(id string, answer string, clear bool) bool {
	val := i.Get(id, clear)

	return val != "" && val == answer
}

func New(opts Opts, store captcha.Store[string]) captcha.Captcha {
	return &Captcha{Store: store, Opts: opts, CaptchaStore: &Store{Store: store}}
}
