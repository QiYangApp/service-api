package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	RunMode     string
	Engine      *gin.Engine
	middlewares []gin.HandlerFunc
	providers   []Provider
	container   *Container
	events      map[string]*Event
}

func (t *App) Events(events ...*Event) *App {
	for _, event := range events {
		if t.events[event.Key] != nil && t.events[event.Key].Share == false {
			continue
		}

		t.events[event.Key] = event
	}

	return t
}

func (t *App) Middlewares(middlewares ...gin.HandlerFunc) *App {
	t.middlewares = append(t.middlewares, middlewares...)

	return t
}

func (t *App) Providers(providers ...Provider) *App {
	t.providers = append(t.providers, providers...)

	return t
}

func (t *App) Container(container *Container) *App {
	t.container = container

	return t
}

func (t *App) Run() {
	gin.SetMode(t.RunMode)

	_ = t.Engine.SetTrustedProxies(nil)

	srv := &http.Server{
		Addr:    d.Path,
		Handler: d.Gin,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
