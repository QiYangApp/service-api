package setting

import (
	"frame/conf"
	"github.com/spf13/viper"
)

func GetConnConf(name string) *viper.Viper {
	return conf.Client().Sub("conns." + name)
}
