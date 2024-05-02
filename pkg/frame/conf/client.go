package conf

import (
	"go.uber.org/zap"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var client = viper.New()

var once = sync.Once{}

func Client() *viper.Viper {
	once.Do(func() {
		dir, _ := os.Getwd()
		if err := ParseLocalFile(client, dir); err != nil {
			zap.S().Panic("conf read error, %v", err)
		}
	})

	return client
}

func ParseLocalFile(cli *viper.Viper, dir string) error {
	cli.SetConfigType("yaml")
	cli.SetConfigName("config")
	cli.AddConfigPath(dir)

	// 读取配置文件
	if err := cli.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
