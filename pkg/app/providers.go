package app

import (
	"app/config"
)

type Provider interface {
	Register(r *App)
}

type ConfigProviders struct {
}

func (c *ConfigProviders) Register(r *App) {
	config.Instance().ParseFile()
}
