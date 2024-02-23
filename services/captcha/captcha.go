package captcha

import (
	"fmt"
	"framework/cache"
	"framework/exceptions"
	"github.com/mojocn/base64Captcha"
	"service-api/resources/lang"
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
	Answer  string `json:"answer"`
}

type ImageStore struct {
	sync.RWMutex
	exp time.Duration
}

func (i *ImageStore) Set(id string, value string) error {
	if !cache.SetEx(id, value, i.exp) {
		return exceptions.New(lang.CaptchaErrorStoreCode)
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

func (i *Image) Generate(scenes string, p *ImageParam) (*ImageResp, error) {

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
	id, body, answer, err := captcha.Generate()
	if err != nil {
		return nil, exceptions.New(lang.CaptchaErrorGenerateCode)
	}

	// 生成场景存储类型错误
	if cache.SetEx(fmt.Sprintf("captcha-%s-%s", scenes, id), id, 2*time.Minute) == false {
		return nil, exceptions.New(lang.CaptchaErrorGenerateCode)
	}

	return &ImageResp{Captcha: body, Id: id, Answer: answer}, nil
}

func (i *Image) Verify(scenes, id, answer string) bool {
	if cache.Exists(fmt.Sprintf("captcha-%s-%s", scenes, id)) {
		return false
	}

	return i.store.Verify(id, answer, false) == false
}

func NewImage() *Image {
	return &Image{
		store: &ImageStore{exp: 2 * time.Minute},
	}
}
