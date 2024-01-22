package app

import (
	"app/config"
	"app/log"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Cmd struct {
	RunMode string
	Debug   bool
}

type App struct {
	Cmd         *Cmd
	Engine      *gin.Engine
	middlewares []gin.HandlerFunc
	providers   []Provider
	container   *Container
}

func (t *App) Middlewares(middlewares ...gin.HandlerFunc) *App {
	t.middlewares = append(t.middlewares, middlewares...)

	return t
}

func (t *App) Providers(providers ...Provider) *App {
	t.providers = append(t.providers, providers...)

	return t
}

func (t *App) Run(cmd *Cmd) {
	t.Cmd = cmd

	for _, provider := range t.providers {
		provider.Register(t)
	}

	t.Cmd.Debug = config.Client().GetBool("server.debug")
	t.Cmd.RunMode = config.Client().GetString("server.runMode")

	t.Engine.Use(t.middlewares...)

	gin.SetMode(t.Cmd.RunMode)

	_ = t.Engine.SetTrustedProxies(nil)

	srv := &http.Server{
		Addr:                         config.Client().GetString("server.addr") + ":" + config.Client().GetString("server.prot"),
		Handler:                      t.Engine,
		DisableGeneralOptionsHandler: true,
		ReadTimeout:                  time.Duration(config.Client().GetInt("server.config.readTimeout")),
		WriteTimeout:                 time.Duration(config.Client().GetInt("server.config.writeTimeout")),
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Client().Sugar().Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Client().Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Client().Warn("Shutdown Server ...", zap.Error(err))
	}
	log.Client().Info("Server exiting")
}

func New() *App {
	return &App{
		Engine: gin.New(),
		providers: []Provider{
			&ConfigProviders{},
		},
	}
}
