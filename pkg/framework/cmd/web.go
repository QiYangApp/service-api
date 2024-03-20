package cmd

// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

import (
	"context"
	"errors"
	"framework/config"
	"framework/log"
	"framework/middlewares"
	"framework/providers"
	"framework/utils"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

var webSingleton *WebServer
var webOnce = sync.Once{}

var Web = cli.Command{
	Name:  "web",
	Usage: "Start web server",
	Description: `web server is the only thing you need to run,
and it takes care of all the other things for you`,
	Action: runWeb,
	Flags: []cli.Flag{
		stringFlag("addr, a", "localhost:3000", "Temporary port number to prevent conflict"),
		stringFlag("config, c", "config.toml", "Custom configuration file path"),
	},
}

type WebCmd struct {
	RunMode string
	Debug   bool
	Addr    string
	Config  string
}

type WebServer struct {
	Cmd         *WebCmd
	Engine      *gin.Engine
	Middlewares []middlewares.Middleware
	Providers   []providers.Provider
	Commands    []cli.Command
}

func (t *WebServer) SetMode(mod string) {
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

func (t *WebServer) Run(cmd *WebCmd) {
	t.Cmd = cmd

	for _, provider := range t.Providers {
		provider.Register()
	}

	t.SetMode(t.Cmd.RunMode)

	_ = t.Engine.SetTrustedProxies(nil)

	addr := config.Client.GetString("ADDR")

	log.Client.Sugar().Infof("Server addr %s", addr)

	srv := &http.Server{
		Addr:                         addr,
		Handler:                      t.Engine,
		DisableGeneralOptionsHandler: true,
		ReadTimeout:                  time.Duration(config.Client.GetInt("READ_TIMEOUT")) * time.Second,
		WriteTimeout:                 time.Duration(config.Client.GetInt("WRITE_TIMEOUT")) * time.Second,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Client.Sugar().Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Client.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Client.Warn("Shutdown Server ...", zap.Error(err))
	}
	log.Client.Info("Server exiting")
}

func WebServerClient() *WebServer {
	webOnce.Do(func() {
		webSingleton = &WebServer{
			Engine: gin.New(),
			Providers: []providers.Provider{
				&providers.CacheProviders{},
				&providers.CronProviders{},
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

	return webSingleton
}

func runWeb(c *cli.Context) {

	path := utils.Path.RootPath
	if c.IsSet("config") {
		path = c.GlobalString("config")
	}

	config.Instance().ParseFile(path)

	runCmd := &WebCmd{
		Debug:   config.Client.GetBool("debug"),
		RunMode: config.Client.GetString("run_mode"),
		Addr:    config.Client.GetString("addr"),
	}

	if c.IsSet("addr") {
		runCmd.Addr = c.GlobalString("addr")
	}

	WebServerClient().Run(runCmd)
}
