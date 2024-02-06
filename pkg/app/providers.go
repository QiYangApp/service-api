package app

import (
	"app/cache"
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

type CacheProviders struct {
}

func (c *CacheProviders) Register(app *App) {
	cache.NewInstance(config.Client().GetString("cache.driver"))
}

//type RouterProviders struct {
//}
//
//func (c *RouterProviders) Register(app *App) {
//	scan := &router.Scan{
//		Apis:       make(map[string]*router.MethodInfo),
//		BasePaths:  make(map[string]string),
//		Names:      make(map[string]router.Names),
//		Controller: make([]router.Inject, 0),
//	}
//
//	scan.RecursionPkgDir(helpers.Path.ControllerPath, app.PackageName+"/api/controller", "controller")
//	app.Scan = scan
//
//	router.GenerateApi(scan, helpers.Path.RootPath)
//}
