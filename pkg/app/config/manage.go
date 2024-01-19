package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Limit interface {
	Config | interface{}
}

type Config struct {
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
	Config struct {
		Type string
	}
}

type Client[C Limit] struct {
	Config C
}

func (c *Client[C]) Parse(path, tp string) *Client[C] {

	viper.SetConfigFile(path)

	//viper.SetConfigName("config")
	//viper.AddConfigPath(helpers.PathInstance.GetCurrentRunPath())
	viper.SetConfigType(tp)

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		//logger.S().Warnln("check config.toml is exists!")
		//logger.S().Panicf("config read fail %v", zap.Error(err))
	}

	// 解析配置项
	var config = &Config{}
	if err := viper.Unmarshal(config); err != nil {
		//logger.S().Panicf("config parse fail %v", zap.Error(err))
	}

	// 监听配置文件
	// 不启用监听，考虑后续会使用线上acm或者其他远程配置中心
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		//logger.S().Info("config file changed: ", zap.String("name", in.Name))
		// 重载配置
		if err := viper.Unmarshal(config); err != nil {
			//logger.S().Errorf("config change read fail %v", zap.Error(err))
		}
	})

	c.Merge(config)

	return c
}

func (c *Client[C]) Merge(config C) *Client[C] {
	c.Config = config

	return c
}
