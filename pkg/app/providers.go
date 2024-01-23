package app

import (
	"app/config"
	"app/cron"
)

type Provider interface {
	Register(app *App)
}

type ConfigProviders struct {
}

func (c *ConfigProviders) Register(app *App) {
	config.Instance().ParseFile()
}

type CronProviders struct {
}

func (c *CronProviders) Register(app *App) {
	cron.Instance().Init()
}
