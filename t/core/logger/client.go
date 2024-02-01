package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var singleton = make(map[string]*Instance, 0)
var defaultOutputLevel = zapcore.WarnLevel

const (
	DefaultMode  string = "default"
	RequestMode         = "request"
	SystemMode          = "system"
	RecoveryMode        = "recovery"
)

func Recovery() *zap.SugaredLogger {
	log := NewSingletonLogger(CoreParam{
		Mode:        RequestMode,
		OutputLevel: defaultOutputLevel,
	})

	return log.Logger().Sugar()
}

func R() *zap.SugaredLogger {
	log := NewSingletonLogger(CoreParam{
		Mode:        RequestMode,
		OutputLevel: defaultOutputLevel,
	})

	return log.Logger().Sugar()
}

func D() *zap.SugaredLogger {
	log := NewSingletonLogger(CoreParam{
		Mode:        DefaultMode,
		OutputLevel: defaultOutputLevel,
	})

	return log.Logger().Sugar()
}

func S() *zap.SugaredLogger {
	log := NewSingletonLogger(CoreParam{
		Mode:        SystemMode,
		OutputLevel: defaultOutputLevel,
	})

	return log.Logger().Sugar()
}

func C(mode string) Instance {
	return NewSingletonLogger(CoreParam{
		Mode:        mode,
		OutputLevel: defaultOutputLevel,
	})
}

func NewSingletonLogger(param CoreParam) Instance {
	if logger, ok := singleton[param.Mode]; ok {
		return *logger
	}

	singleton[param.Mode] = NewLogger(param)

	return *singleton[param.Mode]
}

func NewLogger(param CoreParam) *Instance {

	logger := CoreBuilder{
		CoreParam: param,
	}

	return &Instance{
		logger: logger.Builder(),
	}
}
