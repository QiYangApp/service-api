package config

import (
	"fmt"
	"service-api/src/core/helpers"
	"service-api/src/core/logger"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type ConfigService struct {
	config *Config
}

func (c *ConfigService) initConfig() *ConfigService {
	viper.SetConfigName("config")
	viper.AddConfigPath(helpers.PathInstance.GetCurrentRunPath())
	viper.SetConfigType("toml")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		logger.S().Warnln("check config.toml is exists!")
		logger.S().Panicf("config read fail %v", zap.Error(err))
	}

	// 解析配置项
	if err := viper.Unmarshal(&c.config); err != nil {
		logger.S().Panicf("config parse fail %v", zap.Error(err))
	}

	// 监听配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		logger.S().Info("config file changed: ", zap.String("name", in.Name))
		// 重载配置
		if err := viper.Unmarshal(&c.config); err != nil {
			logger.S().Errorf("config change read fail %v", zap.Error(err))
		}
	})

	return c
}

func (c *ConfigService) GetDomain() string {
	return c.config.Domain
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

func (c ConfigService) StartServiceAddress() string {
	return fmt.Sprintf("%s:%d", c.config.Server.Host, c.config.Server.Port)
}

func (c ConfigService) GetCache() Cache {
	return c.config.Cache
}

func (c ConfigService) IsDebug() bool {
	return c.RunMode() == gin.DebugMode
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
