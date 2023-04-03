package system

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"service-api/src/api/middleware"
	"service-api/src/app/helpers"
	"service-api/src/core/config"
	"service-api/src/routes"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	if cfg.RunMode() == gin.DebugMode {
		gin.ForceConsoleColor()
	} else {
		gin.DisableConsoleColor()
	}

	gin.SetMode(cfg.RunMode())

	(new(LoggerService)).Handle(r, cfg)

	(new(DatabaseService)).Handle(r, cfg)

	(new(LanguageService)).Handle(r, cfg)

	(new(CacheService)).Handle(r, cfg)

	// 注册路由
	middleware.SetupMiddleware(r)

	// 注册路由
	routes.SetupRoutes(r)

}

func afterStart(r *gin.Engine, cfg *config.ConfigService) {

}

func run(r *gin.Engine, cfg *config.ConfigService) {

	r.Static("/resources/assets", helpers.NewPathMange().JoinCurrentRunPath("resources/assets"))

	srv := &http.Server{
		Addr:    cfg.StartServiceAddress(),
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	zap.S().Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.S().Fatal("Server Shutdown:", err)
	}
	zap.S().Info("Server exiting")
}
