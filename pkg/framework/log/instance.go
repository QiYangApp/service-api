package log

import (
	"framework/config"
	"go.uber.org/zap"
	"sync"
)

var instance *Manage = nil
var once = sync.Once{}

func Instance() *Manage {
	once.Do(func() {
		instance = &Manage{
			Param: &Param{
				Debug:       config.Client().GetBool("DEBUG"),
				MaxBackups:  config.Client().GetInt("LOG.MAX_BACKUPS"),
				MaxAge:      config.Client().GetInt("LOG.MAX_AGE"),
				MaxSize:     config.Client().GetInt("LOG.MAX_SIZE"),
				Compress:    config.Client().GetBool("LOG.COMPRESS"),
				LogFileName: config.Client().GetString("LOG.FILE_NAME"),
				Level:       config.Client().GetString("LOG.LEVEL"),
			},
		}
		instance.Builder()
	})

	return instance
}

func Client() *zap.Logger {
	return Instance().Client
}
