package log

import (
	"framework/config"
	"go.uber.org/zap"
	"sync"
)

func init() {
	Instance()
}

var instance *Manage = nil
var once = sync.Once{}
var Client *zap.Logger = Instance().Client

func Instance() *Manage {
	once.Do(func() {
		instance = &Manage{
			Param: &Param{
				Debug:       config.Client.GetBool("DEBUG"),
				MaxBackups:  config.Client.GetInt("LOG.MAX_BACKUPS"),
				MaxAge:      config.Client.GetInt("LOG.MAX_AGE"),
				MaxSize:     config.Client.GetInt("LOG.MAX_SIZE"),
				Compress:    config.Client.GetBool("LOG.COMPRESS"),
				LogFileName: config.Client.GetString("LOG.FILE_NAME"),
				Level:       config.Client.GetString("LOG.LEVEL"),
			},
		}
		instance.Builder()
	})

	return instance
}
