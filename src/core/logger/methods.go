package logger

import "go.uber.org/zap/zapcore"

type LoggerMode string

var singleton = make(map[LoggerMode]*Instance, 0)

const (
	RequestMode LoggerMode = "request"
	SystemMode             = "system"
)

func S(mode *string) Instance {
	if mode == nil {
		m := SystemMode
		mode = &mode
	}

	return NewSingletonLogger(LoggerCoreParam{
		Mode:        *mode,
		OutputLevel: zapcore.WarnLevel,
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
