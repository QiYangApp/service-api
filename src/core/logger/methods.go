package logger

import "go.uber.org/zap/zapcore"

type LoggerMode string

var singleton = make(map[LoggerMode]*Instance, 0)
var defaultOutputLevel = zapcore.WarnLevel

const (
	DefaultMode LoggerMode = "default"
	RequestMode            = "request"
	SystemMode             = "system"
)

func R() Instance {
	return NewSingletonLogger(LoggerCoreParam{
		Mode:        RequestMode,
		OutputLevel: defaultOutputLevel,
	})
}

func D() Instance {
	return NewSingletonLogger(LoggerCoreParam{
		Mode:        DefaultMode,
		OutputLevel: defaultOutputLevel,
	})
}

func S() Instance {
	return NewSingletonLogger(LoggerCoreParam{
		Mode:        SystemMode,
		OutputLevel: defaultOutputLevel,
	})
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
