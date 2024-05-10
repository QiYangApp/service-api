// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"errors"
	"frame/conf"
	"frame/modules/log"
	"frame/modules/middlewares"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

var webCli *WebServer
var webOnce = sync.Once{}

var WebCommand = &cli.Command{
	Name:  "web",
	Usage: "Start web server",
	Description: `web server is the only thing you need to run,
and it takes care of all the other things for you`,
	Action: runWeb,
}

type WebPath struct {
	Root     string
	App      string
	Storage  string
	Resource string
	Public   string
	Log      string
	I18      string
}

type WebServer struct {
	Path        *WebPath
	Engine      *gin.Engine
	Serve       *http.Server
	Middlewares []middlewares.Middleware
	Providers   []func()
	Commands    []cli.Command
}

func (w *WebServer) SetMode(mod string) {
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

func (w *WebServer) Run(cmd *cli.Context) {
	for _, provider := range w.Providers {
		provider()
	}

	w.SetMode(conf.RunMode())

	_ = w.Engine.SetTrustedProxies(nil)

	addr := conf.Client().GetString("addr")
	log.Sugar().Infof("Server addr %s", addr)
	w.Serve = &http.Server{
		Addr:                         addr,
		Handler:                      w.Engine,
		DisableGeneralOptionsHandler: true,
		ReadTimeout:                  time.Duration(conf.Client().GetInt("read_timeout")) * time.Second,
		WriteTimeout:                 time.Duration(conf.Client().GetInt("write_timeout")) * time.Second,
	}

	go func() {
		// 服务连接
		if err := w.Serve.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Sugar().Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Sugar().Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := w.Serve.Shutdown(ctx); err != nil {
		log.Sugar().Warn("Shutdown Server ...", zap.Error(err))
	}
	log.Sugar().Info("Server exiting")

	defer log.Client().Sync()

}

func runWeb(c *cli.Context) error {
	WebCli().Run(c)

	return nil
}

func WebCli() *WebServer {
	webOnce.Do(func() {
		webCli = &WebServer{
			Engine: gin.New(),
			Path: &WebPath{
				Root:     "",
				App:      "",
				Storage:  "",
				Resource: "",
				Public:   "",
				Log:      "",
				I18:      "",
			},
			Middlewares: []middlewares.Middleware{
				middlewares.Recovery,
				middlewares.I18nLocal,
				middlewares.I18nUrl,
				middlewares.Session,
				middlewares.Limiter,
				middlewares.Option,
				middlewares.Secure,
				middlewares.Logger,
			},
		}

	})

	return webCli
}
