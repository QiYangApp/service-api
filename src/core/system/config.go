package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
		Env  string `mapstructure:"env"`
	} `mapstructure:"server"`
	Database struct {
		Type     string `mapstructure:"type"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
	} `mapstructure:"database"`
}

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
}

func (c ConfigService) GetAllConfig() Config {
	return config
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
