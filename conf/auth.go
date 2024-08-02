package conf

import (
	"frame/modules/log"
	"time"

	"github.com/spf13/viper"
)

var AuthSetting = &struct {
	TwoFA struct {
		Expires time.Duration `mapstructure:"expires"`
	}
}{}

func loadAuthSetting(viper *viper.Viper) {
	if err := viper.Unmarshal(AuthSetting); err != nil {
		log.Sugar().Error("load auth conf")
	}

	AuthSetting.TwoFA.Expires *= time.Minute
}
