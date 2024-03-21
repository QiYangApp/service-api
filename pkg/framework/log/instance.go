package log

import (
	"framework/config"
	"go.uber.org/zap"
	"sync"
)

var instance *Manage = nil
var once = sync.Once{}
var Client *zap.Logger = nil

func init() {
	Instance()
}

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

		Client = instance.Client
	})

	return instance
}
