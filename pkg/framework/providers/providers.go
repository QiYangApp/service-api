package providers

import (
	"framework/cache"
	"framework/config"
	"framework/cron"
	"framework/log"
	"go.uber.org/zap"
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
	driver := config.Client.GetString("CACHE.DRIVER")
	if err := cache.Client.SetDefaultDrive(driver); err != nil {
		zap.S().Fatal(err)
	}
}

type LogProviders struct {
}

func (c *LogProviders) Register() {
	log.Instance()
}
