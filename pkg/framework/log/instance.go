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
				Debug:       config.Client().GetBool("debug"),
				MaxBackups:  config.Client().GetInt("log.max_backups"),
				MaxAge:      config.Client().GetInt("log.max_age"),
				MaxSize:     config.Client().GetInt("log.max_size"),
				Compress:    config.Client().GetBool("log.compress"),
				LogFileName: config.Client().GetString("log.file_name") + ".log",
				Level:       config.Client().GetString("log.level"),
			},
		}
		instance.Builder()
	})

	return instance
}

func Client() *zap.Logger {
	return Instance().Client
}
