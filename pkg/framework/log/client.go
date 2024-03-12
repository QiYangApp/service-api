package log

import (
	"framework/utils"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Param struct {
	Level       string
	MaxSize     int
	MaxAge      int
	MaxBackups  int
	Debug       bool
	Compress    bool
	LogFileName string
}

type Manage struct {
	*Param
	Client *zap.Logger
}

func (b *Manage) Builder() *Manage {

	// 设置初始化字段
	field := zap.Fields(zap.String("time", time.Now().Format("2006-01-02 15:04:05")), zap.Any("level", b.Level))

	// 开启开发模式，堆栈跟踪, 文件和行号
	if b.Debug {
		b.Client = zap.New(b.core(), zap.AddCaller(), zap.Development(), field)
	} else {
		b.Client = zap.New(b.core(), field)
	}

	return b
}

func (b *Manage) core() zapcore.Core {

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
	atomicLevel.SetLevel(b.GetLevel(b.Level))

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(b.hook()...),
		atomicLevel,
	)

	return core
}
func (b *Manage) GetLevel(level string) zapcore.Level {

	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "DPanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	}

	return zap.DebugLevel
}

func (b *Manage) hook() []zapcore.WriteSyncer {
	hooks := []zapcore.WriteSyncer{
		zapcore.AddSync(
			&lumberjack.Logger{
				Filename:   utils.Path.Join(utils.Path.LogPath, b.LogFileName), // 日志文件路径
				MaxSize:    b.MaxSize,                                          // 每个日志文件保存的大小 单位:M
				MaxAge:     b.MaxAge,                                           // 文件最多保存多少天
				MaxBackups: b.MaxBackups,                                       // 日志文件最多保存多少个备份
				Compress:   b.Compress,                                         // 是否压缩
			},
		),
	}

	// 如果是开发环境，同时在控制台上也输出
	if b.Debug {
		hooks = append(hooks, zapcore.AddSync(os.Stdout))
	}

	return hooks
}
