package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerMode string

var singleton = make(map[LoggerMode]*Instance, 0)
var defaultOutputLevel = zapcore.WarnLevel

const (
	DefaultMode  LoggerMode = "default"
	RequestMode             = "request"
	SystemMode              = "system"
	RecoveryMode            = "recovery"
)

func Recovery() *zap.SugaredLogger {
	log := NewSingletonLogger(LoggerCoreParam{
		Mode:        RequestMode,
		OutputLevel: defaultOutputLevel,
	})

	return log.Logger().Sugar()
}

func R() *zap.SugaredLogger {
	log := NewSingletonLogger(LoggerCoreParam{
		Mode:        RequestMode,
		OutputLevel: defaultOutputLevel,
	})

	return log.Logger().Sugar()
}

func D() *zap.SugaredLogger {
	log := NewSingletonLogger(LoggerCoreParam{
		Mode:        DefaultMode,
		OutputLevel: defaultOutputLevel,
	})

	return log.Logger().Sugar()
}

func S() *zap.SugaredLogger {
	log := NewSingletonLogger(LoggerCoreParam{
		Mode:        SystemMode,
		OutputLevel: defaultOutputLevel,
	})

	return log.Logger().Sugar()
}

func C(mode LoggerMode) Instance {
	return NewSingletonLogger(LoggerCoreParam{
		Mode:        mode,
		OutputLevel: defaultOutputLevel,
	})
}

func NewSingletonLogger(param LoggerCoreParam) Instance {
	if logger, ok := singleton[param.Mode]; ok {
		return *logger
	}

	singleton[param.Mode] = NewLogger(param)

	return *singleton[param.Mode]
}

func NewLogger(param LoggerCoreParam) *Instance {

	logger := LoggerCoreBuilder{
		LoggerCoreParam: param,
	}

	return &Instance{
		logger: logger.Builder(),
	}
}
