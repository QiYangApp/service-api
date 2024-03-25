package log

import (
	"framework/config"
	"go.uber.org/zap"
	"sync"
)

var instance *Manage = nil
var once = sync.Once{}
var Client *zap.Logger = nil

type ConfigType struct {
	DEBUG bool
}

func Instance() *Manage {
	once.Do(func() {
		var param = &Param{
			Debug: config.Client.GetBool("DEBUG"),
		}

		if err := config.Client.Sub("log").Unmarshal(param); err != nil {
			zap.S().Panic("log init fail, err: ", err)
		}

		instance = &Manage{
			Param: param,
		}
		instance.Builder()

		Client = instance.Client
	})

	return instance
}
