package setting

import (
	"framework/cache"
	"framework/log"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"service-api/internal/modules/captcha"
	"service-api/internal/modules/captcha/source/base64Captcha"
	"service-api/internal/modules/captcha/source/gocaptcha"
	"time"
)

var CaptchaSetting = struct {
	Enable  bool           `mapstructure:"enable"`
	Store   string         `mapstructure:"store"`
	Type    captcha.Type   `mapstructure:"type"`
	EXPIRES time.Duration  `mapstructure:"expires"`
	Drivers map[string]any `mapstructure:"drivers"`
}{
	Enable: false,
}

func loadCaptchaSetting(viper *viper.Viper) {
	if err := viper.Unmarshal(CaptchaSetting); err != nil {
		log.Client.Error("load service setting")
	}
}

func GetCaptchaClient() (captcha.Captcha, error) {
	store, err := cache.Instance().Register(CaptchaSetting.Store)
	if err != nil {
		log.Client.Sugar().Errorf("captcha store driver init fail, err: %v", err)
		return nil, err
	}

	var setting any
	var ok bool
	if setting, ok = CaptchaSetting.Drivers[CaptchaSetting.Type.ToString()]; !ok {
		log.Client.Sugar().Errorf("captcha setting type not exists, err: %v", err)
		return nil, err
	}

	switch CaptchaSetting.Type {
	case captcha.Image:
		var opts base64Captcha.Opts
		if err := mapstructure.Decode(setting, &opts); err != nil {
			log.Client.Sugar().Errorf("captcha config init fail, err: %v", err)
			return nil, err
		}

		return base64Captcha.New(opts, captcha.NewStoreValue(store, CaptchaSetting.EXPIRES)), nil
	case captcha.TextPoint:
		var opts gocaptcha.Opts
		if err := mapstructure.Decode(setting, &opts); err != nil {
			log.Client.Sugar().Errorf("captcha config init fail, err: %v", err)
			return nil, err
		}

		return gocaptcha.New(opts, captcha.NewStoreValue(store, CaptchaSetting.EXPIRES)), nil
	}

	return nil, nil
}
