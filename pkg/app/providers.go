package app

import (
	"app/config"
	"github.com/spf13/viper"
)

type Provider interface {
	Register(r *App)
}

type ConfigProviders struct {
}

func (c *ConfigProviders) Register(r *App) {
	conf := &config.Manage{
		Client: viper.New(),
	}

	config.Instance = conf.ParseFile()
}
