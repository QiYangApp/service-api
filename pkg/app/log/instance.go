package log

import (
	"app/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var instance *Manage = nil
var once = sync.Once{}

func Instance() *Manage {
	once.Do(func() {
		instance = &Manage{
			Param: &Param{
				Debug:       config.Client().GetBool("server.debug"),
				MaxBackups:  config.Client().GetInt("log.maxBackups"),
				MaxAge:      config.Client().GetInt("log.maxAge"),
				MaxSize:     config.Client().GetInt("log.maxSize"),
				Compress:    config.Client().GetBool("log.compress"),
				LogFileName: config.Client().GetString("log.fileName") + ".log",
				Level:       zapcore.Level(config.Client().GetInt("log.level")),
			},
		}
	})

	return instance
}

func Client() *zap.Logger {
	return Instance().Client
}