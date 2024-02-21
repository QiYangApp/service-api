package config

import (
	"framework/helpers"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"sync"
)

var singleton *Manage
var once = sync.Once{}

type Methods func(config *viper.Viper) *Manage

type Manage struct {
	Client *viper.Viper
}

func (c *Manage) ParseFile() *Manage {

	c.Client.SetConfigType("toml")
	c.Client.SetConfigName("config")
	c.Client.AddConfigPath(helpers.Path.RootPath)

	// 读取配置文件
	if err := c.Client.ReadInConfig(); err != nil {
		zap.S().Error(err)
	}

	return c
}

func (c *Manage) ParserRemote(m Methods) *Manage {
	return m(c.Client)
}

func Instance() *Manage {

	once.Do(func() {
		singleton = &Manage{
			Client: viper.New(),
		}
	})

	return singleton
}

func Client() *viper.Viper {
	return Instance().Client
}
