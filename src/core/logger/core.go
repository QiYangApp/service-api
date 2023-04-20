package logger

import (
	"fmt"
	"os"
	"service-api/src/core/helpers"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerCoreParam struct {
	Mode        LoggerMode
	OutputLevel zapcore.Level
}

type LoggerCoreBuilder struct {
	LoggerCoreParam
	logger *zap.Logger
}

func (b *LoggerCoreBuilder) Logger() *zap.Logger {
	return b.logger
}

func (b *LoggerCoreBuilder) setOutPutLevel(level zapcore.Level) *LoggerCoreBuilder {
	b.OutputLevel = level

	return b
}

func (b *LoggerCoreBuilder) Builder() *LoggerCoreBuilder {

	// 设置初始化字段
	field := zap.Fields(zap.String("name", string(b.Mode)), zap.Any("level", b.OutputLevel))

	// 开启开发模式，堆栈跟踪, 文件和行号
	if b.OutputLevel == zap.DebugLevel {
		b.logger = zap.New(b.core(), zap.AddCaller(), zap.Development(), field)
	} else {
		b.logger = zap.New(b.core(), field)
	}

	return b
}

func (b *LoggerCoreBuilder) core() zapcore.Core {

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "file",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder, //zapcore.LowercaseLevelEncoder,
		// EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(b.OutputLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(b.hook()...),
		atomicLevel,
	)

	return core
}

func (b *LoggerCoreBuilder) hook() []zapcore.WriteSyncer {
	hooks := []zapcore.WriteSyncer{
		zapcore.AddSync(
			&lumberjack.Logger{
				Filename:   helpers.PathInstance.JoinCurrentRunPath(fmt.Sprintf("storage/logs/%s/%s.log", time.Now().Format("2006-01-02"), b.Mode)), // 日志文件路径
				MaxSize:    128,                                                                                                                     // 每个日志文件保存的大小 单位:M
				MaxAge:     7,                                                                                                                       // 文件最多保存多少天
				MaxBackups: 30,                                                                                                                      // 日志文件最多保存多少个备份
				Compress:   true,                                                                                                                    // 是否压缩
			},
		),
	}

	// 如果是开发环境，同时在控制台上也输出
	if b.OutputLevel == zap.DebugLevel {
		hooks = append(hooks, zapcore.AddSync(os.Stdout))
	}

	return hooks
}
