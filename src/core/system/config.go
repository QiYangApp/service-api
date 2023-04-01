package system

import (
	"fmt"
	"service-api/src/app/helpers"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var ConfigInstance *ConfigService

type ConfigService struct {
	config *Config
}

func (c *ConfigService) initConfig() *ConfigService {
	viper.SetConfigName("config")
	viper.AddConfigPath(helpers.NewPathMange().JoinCurrentRunPath("config"))
	viper.SetConfigType("toml")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 解析配置项
	if err := viper.Unmarshal(&c.config); err != nil {
		panic(err)
	}

	// 监听配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		zap.S().Info("config file changed: ", in.Name)
		// 重载配置
		if err := viper.Unmarshal(&c.config); err != nil {
			fmt.Println(err)
		}
	})

	return c
}

func (c ConfigService) GetAllConfig() *Config {
	return c.config
}

func (c ConfigService) GetDatabase() Database {
	return c.config.Database
}

func (c ConfigService) GetService() Server {
	return c.config.Server
}

func (c ConfigService) GetToken() Token {
	return c.config.Token
}

func (c ConfigService) startServiceAddress() string {
	return fmt.Sprintf("%s:%d", c.config.Server.Host, c.config.Server.Port)
}

func (c ConfigService) GetCache() Cache {
	return c.config.Cache
}

func (c ConfigService) RunMode() string {
	switch c.config.Server.Env {
	case gin.ReleaseMode:
		return gin.ReleaseMode
	case gin.DebugMode:
		return gin.DebugMode
	case gin.TestMode:
		return gin.TestMode
	}

	return gin.ReleaseMode
}
