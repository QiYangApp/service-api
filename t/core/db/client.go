package db

import (
	"service-api/src/core/config"
	"service-api/src/models/ent"
	"sync"
)

var Singleton *Instance
var once = sync.Once{}

func Init(cfg config.Database) *Instance {
	once.Do(func() {
		Singleton = &Instance{
			cfg: cfg,
		}

		Singleton.InitClient()
	})

	return Singleton
}

func Client() *ent.Client {
	return Singleton.Client
}
