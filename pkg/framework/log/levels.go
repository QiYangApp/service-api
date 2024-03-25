package log

import "go.uber.org/zap/zapcore"

type Level string

const (
	DEBUG   Level = "debug"
	INFO          = "info"
	WARN          = "warn"
	ERROR         = "error"
	DPANICE       = "dpanice"
	PANICE        = "panice"
	FATAL         = "fatal"
)

var LevelNames = map[Level]zapcore.Level{
	DEBUG:   zapcore.DebugLevel,
	INFO:    zapcore.InfoLevel,
	WARN:    zapcore.WarnLevel,
	ERROR:   zapcore.ErrorLevel,
	DPANICE: zapcore.DPanicLevel,
	PANICE:  zapcore.PanicLevel,
	FATAL:   zapcore.FatalLevel,
}

func (l Level) Name() zapcore.Level {
	if level, ok := LevelNames[l]; ok {
		return level
	}

	return zapcore.DebugLevel
}
