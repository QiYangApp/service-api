package system

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"service-api/src/core/logger"
	"time"
)

type StartModeMethods interface {
	Start()
}

type startMode struct {
	StartModeMethods
	Gin  *gin.Engine
	Path string
}

type DevelopmentStartMode struct {
	startMode
}

func (d DevelopmentStartMode) Start() {
	err := d.Gin.Run(d.Path)
	if err != nil {
		logger.S().Fatalln("development service start failed: %v", err)
	}
}

type ProductionStartMode struct {
	startMode
}

func (d ProductionStartMode) Start() {

	_ = d.Gin.SetTrustedProxies(nil)

	srv := &http.Server{
		Addr:    d.Path,
		Handler: d.Gin,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.S().Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logger.S().Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.S().Fatal("Server Shutdown:", zap.Error(err))
	}
	logger.S().Info("Server exiting")
}
