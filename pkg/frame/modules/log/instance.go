package log

import (
	"frame/conf"
	"frame/util/path"
	"sync"

	"go.uber.org/zap"
)

var instance *Manage = nil
var once = sync.Once{}

func Client() *zap.Logger {
	return Instance().Client
}

func Sugar() *zap.SugaredLogger {
	return Instance().Client.Sugar()
}

type ConfigType struct {
	DEBUG bool
}

func Instance() *Manage {
	once.Do(func() {

		instance = &Manage{
			Param: builderParam(),
		}
		instance.Builder()
	})

	return instance
}

func builderParam() *Param {
	var param = &Param{
		Debug: conf.IsDebug(),
	}

	if err := conf.Client().Sub("log").Unmarshal(param); err != nil {
		zap.S().Panic("log init fail, err: ", err)
	}

	if param.FileName == "" {
		param.FileName = "record"
	}

	param.FileName = path.JoinPath(path.RunStoragePath, "logs", param.FileName)

	return param
}
