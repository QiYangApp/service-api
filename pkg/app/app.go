package app

import (
	"app/config"
	"app/log"
	"app/middlewares"
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
	Cmd         *Cmd
	Engine      *gin.Engine
	middlewares []middlewares.Middleware
	providers   []Provider
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

	t.Cmd.Debug = config.Client().GetBool(`server.debug`)
	t.Cmd.RunMode = config.Client().GetString(`server.runMode`)
	t.SetMode(t.Cmd.RunMode)

	for _, middleware := range t.middlewares {
		t.Engine.Use(middleware())
	}

	_ = t.Engine.SetTrustedProxies(nil)

	srv := &http.Server{
		Addr:                         config.Client().GetString("server.addr") + ":" + config.Client().GetString("server.port"),
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
			&CronProviders{},
			&RouterProviders{},
		},
		middlewares: []middlewares.Middleware{
			middlewares.I18nLocal,
			middlewares.I18nUrl,
			middlewares.Recovery,
			middlewares.Limiter,
			middlewares.Option,
			middlewares.Secure,
			middlewares.DefaultLogger,
			middlewares.Logger,
		},
	}
}
