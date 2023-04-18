package system

import (
	"fmt"
	"os"
	"service-api/src/core/config"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerService struct {
	service
}

func (l *LoggerService) Handle(r *gin.Engine, cfg *config.ConfigService) {

	hook := lumberjack.Logger{
		Filename:   fmt.Sprintf("./storage/logs/%s/%s.log", time.Now().Format("2006-01-02"), "system"), // 日志文件路径
		MaxSize:    128,                                                                                // 每个日志文件保存的大小 单位:M
		MaxAge:     7,                                                                                  // 文件最多保存多少天
		MaxBackups: 30,                                                                                 // 日志文件最多保存多少个备份
		Compress:   true,                                                                               // 是否压缩
	}

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
	atomicLevel.SetLevel(zap.DebugLevel)
	var writes = []zapcore.WriteSyncer{zapcore.AddSync(&hook)}
	// 如果是开发环境，同时在控制台上也输出
	if cfg.GetService().Env == gin.DebugMode {
		writes = append(writes, zapcore.AddSync(os.Stdout))
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writes...),
		atomicLevel,
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()

	// 设置初始化字段
	field := zap.Fields(zap.String("appName", "default"))

	// 构造日志
	ZapLogger := zap.New(core, caller, development, field)
	ZapLogger.Info("log 初始化成功")

	zap.ReplaceGlobals(ZapLogger)
}
