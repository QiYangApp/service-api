package captcha

import (
	"github.com/mojocn/base64Captcha"
	"service-api/src/core/cache"
	"service-api/src/enums/i18n"
	"service-api/src/errors"
	"sync"
	"time"
)

// https://github.com/mojocn/base64Captcha/blob/master/store_memory.go

type ImageParam struct {
	Height, Width, Len, dotCount int
	MaxSkew                      float64
}

type ImageResp struct {
	Id      string `json:"id"`
	Captcha string `json:"captcha"`
	Code    string `json:"code"`
}

type ImageStore struct {
	sync.RWMutex
	exp time.Duration
}

func (i *ImageStore) Set(id string, value string) error {
	if !cache.SetEx(id, value, i.exp) {
		return errors.WithMes(i18n.CaptchaErrorStoreCode)
	}

	return nil
}

func (i *ImageStore) Get(id string, clear bool) (value string) {
	i.Lock()
	defer i.Unlock()

	if !cache.Exists(id) {
		return
	}

	err := cache.Get(id, value)
	if err == nil && clear {
		cache.Del(id)
	}

	return
}

func (i *ImageStore) Verify(id string, answer string, clear bool) bool {
	key := i.Get(id, clear)

	return key != "" && key == answer
}

type Image struct {
	store *ImageStore
}

func (i *Image) Generate(p *ImageParam) (*ImageResp, error) {

	if p == nil {
		p = &ImageParam{
			70,
			130,
			4,
			100,
			0.8,
		}
	}

	driver := base64Captcha.NewDriverDigit(
		p.Height,
		p.Width,
		p.Len,
		p.MaxSkew,
		p.dotCount,
	)

	captcha := base64Captcha.NewCaptcha(driver, i.store)
	id, body, err := captcha.Generate()
	if err != nil {
		return nil, errors.WithMes(i18n.CaptchaErrorGenerateCode)
	}

	return &ImageResp{Captcha: body, Id: id}, nil
}

func (i *Image) Check(id, value string) bool {
	return i.store.Verify(id, value, true)
}

func NewImage() *Image {
	return &Image{
		store: &ImageStore{exp: 2 * time.Minute},
	}
}
