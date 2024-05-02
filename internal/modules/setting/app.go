package setting

import (
	"frame/modules/log"

	"github.com/spf13/viper"
)

var AppSetting = &struct {
	Name           string `mapstructure:"name"`
	RunMode        string `mapstructure:"run_mode"`
	RunUser        string `mapstructure:"run_user"`
	Debug          bool   `mapstructure:"debug"`
	Domain         string `mapstructure:"domain"`
	Addr           string `mapstructure:"addr"`
	Version        string `mapstructure:"version"`
	WriteTimeout   int    `mapstructure:"write_timeout"`
	ReadTimeout    int    `mapstructure:"read_timeout"`
	MaxRequests    int    `mapstructure:"max_requests"`
	MaxRequestTime int    `mapstructure:"max_request_time"`
	TimeZones      string `mapstructure:"time_zones"`
}{}

func loadApp(viper *viper.Viper) {
	if err := viper.Unmarshal(AppSetting); err != nil {
		log.Sugar().Errorf("app config parser fail, ", err)
	}
}
