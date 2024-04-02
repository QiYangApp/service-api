package setting

import (
	"framework/log"
	"github.com/spf13/viper"
	"time"
)

var AuthSetting = &struct {
	TwoFA struct {
		Expires time.Duration `mapstructure:"expires"`
	}
}{}

func loadAuthSetting(viper *viper.Viper) {
	log.Client.Debug("start load auth setting")
	if err := viper.Unmarshal(AuthSetting); err != nil {
		log.Client.Error("load auth setting")
	}

	AuthSetting.TwoFA.Expires *= time.Minute
}
