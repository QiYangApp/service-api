package system

import (
	"service-api/src/core/config"
	"service-api/src/core/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LoggerService struct {
	service
}

func (l *LoggerService) Handle(r *gin.Engine, cfg *config.ConfigService) {

	level := zap.WarnLevel
	if cfg.RunMode() == gin.DebugMode {
		level = zap.DebugLevel
	}

	instance := logger.NewSingletonLogger(logger.LoggerCoreParam{
		OutputLevel: level,
		Mode:        logger.SystemMode,
	})

	logger.NewSingletonLogger(logger.LoggerCoreParam{
		OutputLevel: level,
		Mode:        logger.DefaultMode,
	})

	logger.NewSingletonLogger(logger.LoggerCoreParam{
		OutputLevel: level,
		Mode:        logger.RequestMode,
	})

	zap.ReplaceGlobals(instance.Logger())
}
