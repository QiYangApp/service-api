package log

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Param struct {
	Level      Level `mapstructure:"level"`
	MaxSize    int   `mapstructure:"max_size"`
	MaxAge     int   `mapstructure:"max_age"`
	MaxBackups int   `mapstructure:"max_backups"`
	Debug      bool
	Compress   bool   `mapstructure:"compress"`
	FileName   string `mapstructure:"file_name"`
}

type Manage struct {
	*Param
	Client *zap.Logger
}

func (b *Manage) Builder() *Manage {

	// 设置初始化字段
	//field := zap.Fields(zap.String("time", time.Now().Format("2006-01-02 15:04:05")), zap.Any("level", b.Level))

	// 开启开发模式，堆栈跟踪, 文件和行号
	if b.Debug {
		b.Client = zap.New(b.core(), zap.AddCaller(), zap.Development())
	} else {
		b.Client = zap.New(b.core())
	}

	return b
}

func (b *Manage) core() zapcore.Core {
	encodelevel := zapcore.LowercaseLevelEncoder
	if b.Debug {
		encodelevel = zapcore.LowercaseColorLevelEncoder
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "file",
		StacktraceKey: "trace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   encodelevel,
		// EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString("[" + t.Format("2006-01-02 15:04::05.000") + "]")
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(b.Level.Name())

	encoder := zapcore.NewJSONEncoder(encoderConfig)
	if b.Debug {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(b.hook()...),
		atomicLevel,
	)

	return core
}

func (b *Manage) hook() []zapcore.WriteSyncer {
	hooks := []zapcore.WriteSyncer{
		zapcore.AddSync(
			&lumberjack.Logger{
				Filename:   b.FileName,   // 日志文件路径
				MaxSize:    b.MaxSize,    // 每个日志文件保存的大小 单位:M
				MaxAge:     b.MaxAge,     // 文件最多保存多少天
				MaxBackups: b.MaxBackups, // 日志文件最多保存多少个备份
				Compress:   b.Compress,   // 是否压缩
			},
		),
	}

	// 如果是开发环境，同时在控制台上也输出
	if b.Debug {
		hooks = append(hooks, zapcore.AddSync(os.Stdout))
	}

	return hooks
}
