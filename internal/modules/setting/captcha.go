package setting

import (
	"framework/cache"
	"framework/log"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	captchaclient "github.com/wenlng/go-captcha/captcha"
	"service-api/internal/modules/captcha"
	"service-api/internal/modules/captcha/source/base64Captcha"
	"service-api/internal/modules/captcha/source/gocaptcha"
	"time"
)

type CaptchaFeature string

func (t CaptchaFeature) ToString() string {
	return string(t)
}

const (
	CaptchaFeatureSignIn CaptchaFeature = "sign_in"
	CaptchaFeatureSignUp CaptchaFeature = "sign_up"
)

var CaptchaSetting = &struct {
	Enable  bool                    `mapstructure:"enable"`
	Store   string                  `mapstructure:"store"`
	Type    captcha.Type            `mapstructure:"type"`
	EXPIRES time.Duration           `mapstructure:"expires"`
	Drivers map[captcha.Type]any    `mapstructure:"drivers"`
	Feature map[CaptchaFeature]bool `mapstructure:"feature"`
}{
	Enable: false,
}

func loadCaptchaSetting(viper *viper.Viper) {
	if err := viper.Unmarshal(CaptchaSetting); err != nil {
		log.Client.Error("load service setting")
	}
}

func GetCaptchaClient() (captcha.Captcha, error) {

	var driver cache.Drive
	var err error
	if driver, err = cache.Instance().Register(CaptchaSetting.Store); err != nil {
		log.Client.Sugar().Errorf("captcha store driver init fail, err: %v", err)
	}

	var setting any
	var ok bool
	if setting, ok = CaptchaSetting.Drivers[CaptchaSetting.Type]; !ok {
		log.Client.Sugar().Errorf("captcha setting type not exists, err: %v", err)
		return nil, err
	}

	var exp = CaptchaSetting.EXPIRES * time.Second

	switch CaptchaSetting.Type {
	case captcha.Image:
		var opts base64Captcha.Opts
		if err := mapstructure.Decode(setting, &opts); err != nil {
			log.Client.Sugar().Errorf("captcha config init fail, err: %v", err)
			return nil, err
		}

		store := cache.NewOperation[string](driver)
		return base64Captcha.New(opts, captcha.NewStoreValue(store, exp)), nil
	case captcha.TextPoint:
		var opts gocaptcha.Opts
		if err := mapstructure.Decode(setting, &opts); err != nil {
			log.Client.Sugar().Errorf("captcha config init fail, err: %v", err)
			return nil, err
		}

		store := cache.NewOperation[map[int]captchaclient.CharDot](driver)
		return gocaptcha.New(opts, captcha.NewStoreValue(store, exp)), nil
	}

	return nil, nil
}

// CheckCaptchaFeatureEnable 校验限定功能的验证是否启用
func CheckCaptchaFeatureEnable(t CaptchaFeature) bool {
	if st, ok := CaptchaSetting.Feature[t]; ok {
		return st
	}

	return false
}
