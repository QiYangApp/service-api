package system

import (
	"github.com/gin-gonic/gin"
	"service-api/src/api/middleware"
	"service-api/src/core/config"
	"service-api/src/core/helpers"
)

type service interface {
	Handle(r *gin.Engine, cfg *config.ConfigService)
}

func Start(r *gin.Engine) {
	beforeStart(r, config.Instance)
	run(r, config.Instance)
	afterStart(r, config.Instance)
}

func beforeStart(r *gin.Engine, cfg *config.ConfigService) {

	// 日志颜色
	if cfg.IsDebug() {
		gin.ForceConsoleColor()
	} else {
		gin.DisableConsoleColor()
	}

	gin.SetMode(cfg.RunMode())

	(new(LoggerService)).Handle(r, cfg)

	(new(DatabaseService)).Handle(r, cfg)

	(new(LanguageService)).Handle(r, cfg)

	(new(CacheService)).Handle(r, cfg)

	(new(CronService)).Handle(r, cfg)

	// 注册路由
	middleware.SetupMiddleware(r)

	(new(InjectService)).Handle(r, cfg)
}

func afterStart(r *gin.Engine, cfg *config.ConfigService) {

}

func run(r *gin.Engine, cfg *config.ConfigService) {
	r.SetTrustedProxies(nil)
	r.Static("/resources/assets", helpers.PathInstance.JoinCurrentRunPath("resources/assets"))

	if cfg.IsDebug() {
		s := DevelopmentStartMode{
			startMode{
				Gin:  r,
				Path: cfg.StartServiceAddress(),
			},
		}

		s.Start()
	} else {
		s := ProductionStartMode{
			startMode{
				Gin:  r,
				Path: cfg.StartServiceAddress(),
			},
		}

		s.Start()
	}

}
