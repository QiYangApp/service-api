package base64Captcha

import (
	"fmt"
	"framework/cache"
	"framework/exceptions"
	captcha "github.com/mojocn/base64Captcha"
	"service-api/resources/lang"
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

	driver := captcha.NewDriverDigit(
		p.Height,
		p.Width,
		p.Len,
		p.MaxSkew,
		p.dotCount,
	)

	captcha := captcha.NewCaptcha(driver, i.store)
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

func (i *Image) Verify(scenes, id, answer string, clear bool) bool {
	if cache.Exists(fmt.Sprintf("captcha-%s-%s", scenes, id)) {
		return false
	}

	return i.store.Verify(id, answer, clear) == false
}

func NewImage() *Image {
	return &Image{
		store: &ImageStore{exp: 2 * time.Minute},
	}
}
