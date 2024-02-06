package app

import (
	"app/config"
	"app/log"
	"app/middlewares"
	"app/router"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type Cmd struct {
	RunMode string
	Debug   bool
}

type App struct {
	PackageName string
	Cmd         *Cmd
	Engine      *gin.Engine
	middlewares []middlewares.Middleware
	providers   []Provider
	Scan        *router.Scan
}

func (t *App) Middlewares(middlewares ...middlewares.Middleware) *App {
	t.middlewares = append(t.middlewares, middlewares...)

	return t
}

func (t *App) Providers(providers ...Provider) *App {
	t.providers = append(t.providers, providers...)

	return t
}

func (t *App) SetMode(mod string) {
	ginMode := ""
	switch mod {
	case "debug":
		ginMode = gin.DebugMode
	case "release":
		ginMode = gin.ReleaseMode
	case "test":
		ginMode = gin.TestMode
	default:
		panic("gin mode unknown: " + mod + " (available mode: debug release test)")
	}
	gin.SetMode(ginMode)
}

func (t *App) Run(cmd *Cmd) {
	t.Cmd = cmd

	for _, provider := range t.providers {
		provider.Register(t)
	}

	t.Cmd.Debug = config.Client().GetBool(`debug`)
	t.Cmd.RunMode = config.Client().GetString(`runMode`)
	t.SetMode(t.Cmd.RunMode)

	funcs := []gin.HandlerFunc{}
	for _, middleware := range t.middlewares {
		funcs = append(funcs, middleware())
	}

	t.Engine.Use(funcs...)

	_ = t.Engine.SetTrustedProxies(nil)

	router.Apply(t.Engine, t.Scan, true)

	srv := &http.Server{
		Addr:                         config.Client().GetString("addr") + ":" + config.Client().GetString("port"),
		Handler:                      t.Engine,
		DisableGeneralOptionsHandler: true,
		ReadTimeout:                  time.Duration(config.Client().GetInt("readTimeout")) * time.Second,
		WriteTimeout:                 time.Duration(config.Client().GetInt("writeTimeout")) * time.Second,
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
			&CronProviders{},
		},
		middlewares: []middlewares.Middleware{
			middlewares.Recovery,
			middlewares.I18nLocal,
			middlewares.I18nUrl,
			middlewares.Limiter,
			middlewares.Option,
			middlewares.Secure,
			middlewares.Logger,
		},
	}
}
