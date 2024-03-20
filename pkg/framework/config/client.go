package config

import (
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

func (c *Manage) ParseFile(path string) *Manage {

	c.Client.SetConfigType("toml")
	c.Client.SetConfigName("config")
	c.Client.AddConfigPath(path)

	// 读取配置文件
	if err := c.Client.ReadInConfig(); err != nil {
		zap.S().Panic(err)
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

var Client *viper.Viper = Instance().Client
