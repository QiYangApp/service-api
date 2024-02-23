package framework

import (
	"errors"
	"framework/config"
	"framework/log"
	"framework/middlewares"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

var singleton *App
var once = sync.Once{}

type Cmd struct {
	RunMode string
	Debug   bool
}

type App struct {
	PackageName string
	Cmd         *Cmd
	Engine      *gin.Engine
	Middlewares []middlewares.Middleware
	providers   []Provider
}

func (t *App) Middleware(middlewares ...middlewares.Middleware) *App {
	t.Middlewares = append(t.Middlewares, middlewares...)

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

	t.SetMode(t.Cmd.RunMode)

	_ = t.Engine.SetTrustedProxies(nil)

	addr := config.Client().GetString("addr")

	log.Client().Sugar().Infof("Server addr %s", addr)

	srv := &http.Server{
		Addr:                         addr,
		Handler:                      t.Engine,
		DisableGeneralOptionsHandler: true,
		ReadTimeout:                  time.Duration(config.Client().GetInt("read_timeout")) * time.Second,
		WriteTimeout:                 time.Duration(config.Client().GetInt("write_timeout")) * time.Second,
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

func Client() *App {

	once.Do(func() {
		singleton = &App{
			Engine: gin.New(),
			providers: []Provider{
				&ConfigProviders{},
				&CacheProviders{},
				&CronProviders{},
			},
			Middlewares: []middlewares.Middleware{
				middlewares.Recovery,
				middlewares.I18nLocal,
				middlewares.I18nUrl,
				middlewares.Limiter,
				middlewares.Option,
				middlewares.Secure,
				middlewares.Logger,
			},
		}
	})

	return singleton
}
