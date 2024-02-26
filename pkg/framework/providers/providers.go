package providers

import (
	"framework/cache"
	"framework/config"
	"framework/cron"
)

type Provider interface {
	Register()
}

type CronProviders struct {
}

func (c *CronProviders) Register() {
	cron.Instance().Init()
}

type CacheProviders struct {
}

func (c *CacheProviders) Register() {
	cache.NewInstance(config.Client().GetString("CACHE.DRIVER"))
}
