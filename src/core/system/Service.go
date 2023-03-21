package system

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"service-api/src/api/middleware"
	"service-api/src/routes"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

	r.Static("./src/resources/assets", "./assets")

	(new(LoggerService)).Handle(r, cfg)

	// 注册路由
	middleware.SetupMiddleware(r)

	// 注册路由
	routes.SetupRoutes(r)
}

func afterStart(r *gin.Engine, cfg ConfigService) {

}

func run(r *gin.Engine, cfg ConfigService) {

	srv := &http.Server{
        Addr:    cfg.startServiceAddress(),
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
