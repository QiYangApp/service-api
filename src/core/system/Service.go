package system

import (
	"github.com/gin-gonic/gin"
	"service-api/src/api/middleware"
	"service-api/src/routes"
)

type service interface {
	Handle(r *gin.Engine, cfg ConfigService)
}

func Start(r *gin.Engine) {
	cfg := ConfigService{}
	cfg.Handle(r, ConfigService{})

	beforeStart(r, cfg)
	run(r, cfg)
	afterStart(r, cfg)
}

func beforeStart(r *gin.Engine, cfg ConfigService) {
	gin.SetMode(cfg.runMode())

	(new(LoggerService)).Handle(r, cfg)

	// 注册路由
	middleware.SetupMiddleware(r)

	// 注册路由
	routes.SetupRoutes(r)
}

func afterStart(r *gin.Engine, cfg ConfigService) {

}

func run(r *gin.Engine, cfg ConfigService) {

	r.Run(cfg.startServiceAddress())
}
