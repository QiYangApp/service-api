package app

import (
	"github.com/spf13/viper"
	"github.com/imdario/mergo"
)

type ConfigSet struct {
	Timezone string
	App      struct {
		Name   string
		Domain string
		Sing   string
	}
	Server struct {
		Path         string
		Port         uint32
		Debug        bool
		RunMode      string
		MaxMemory    uint64
		WriteTimeout uint32
		ReadTimeout  uint32
	}
	Database struct {
	}
	Cache struct {
	}
}

type ConfigManage[C ConfigSet] struct {
	Path   string
	Config map[string]interface{}
}

func (c *ConfigManage[C]) ParseStruct() *ConfigManage[C] {

	viper.SetConfigFile(c.Path)

	//viper.SetConfigName("config")
	//viper.AddConfigPath(helpers.PathInstance.GetCurrentRunPath())
	viper.SetConfigType("toml")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		//logger.S().Warnln("check config.toml is exists!")
		//logger.S().Panicf("config read fail %v", zap.Error(err))
	}

	// 解析配置项
	if err := viper.Unmarshal(&c.Config); err != nil {
		//logger.S().Panicf("config parse fail %v", zap.Error(err))
	}

	// 监听配置文件
	// 不启用监听，考虑后续会使用线上acm或者其他远程配置中心
	//viper.WatchConfig()
	//viper.OnConfigChange(func(in fsnotify.Event) {
	//	//logger.S().Info("config file changed: ", zap.String("name", in.Name))
	//	// 重载配置
	//	if err := viper.Unmarshal(&c.Config); err != nil {
	//		//logger.S().Errorf("config change read fail %v", zap.Error(err))
	//	}
	//})

	return c
}

func (c *ConfigManage[C]) merge(config *ConfigSet) {
	c.Config =
}
