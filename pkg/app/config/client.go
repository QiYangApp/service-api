package config

import (
	"app/helpers"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"sync"
)

var Instance *Manage
var once = sync.Once{}

type Methods func(config *viper.Viper) *Manage

type Manage struct {
	Client *viper.Viper
}

func (c *Manage) ParseFile() *Manage {

	c.Client.SetConfigType("toml")
	viper.AddConfigPath(helpers.Path.RootPath)
	viper.AddConfigPath("config")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		zap.S().Error(err)
	}

	return c
}

func (c *Manage) ParserRemote(m Methods) *Manage {
	return m(c.Client)
}

func () *Manage {


	return instance
}

func Client() *viper.Viper {
	return Instance().Client
}
