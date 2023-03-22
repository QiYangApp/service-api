package system

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var config Config

type ConfigService struct {
	service
	config Config
}

func (c *ConfigService) Handle(r *gin.Engine, cfg ConfigService) {
	if config.Server.Host == "" {
		c.initConfig()
	}

	c.config = config
}

func (c *ConfigService) initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./src/config")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 解析配置项
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	// 监听配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		zap.S().Info("config file changed: ", in.Name)
		// 重载配置
		if err := viper.Unmarshal(&config); err != nil {
			fmt.Println(err)
		}

		c.config = config
	})
}

func (c ConfigService) GetAllConfig() Config {
	return config
}

func (c ConfigService) GetDatabase() Database {
	return c.config.Database
}

func (c ConfigService) GetService() Server {
	return c.config.Server
}

func (c ConfigService) startServiceAddress() string {
	return fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
}

func (c ConfigService) runMode() string {
	switch config.Server.Env {
	case gin.ReleaseMode:
		return gin.ReleaseMode
	case gin.DebugMode:
		return gin.DebugMode
	case gin.TestMode:
		return gin.TestMode
	}

	return gin.ReleaseMode
}
