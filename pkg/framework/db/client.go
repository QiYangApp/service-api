package db

import (
	entsql "entgo.io/ent/dialect/sql"
	"framework/config"
	"framework/log"
)

type DB struct {
	Cfg    ConfigConnsMany
	Driver string
}

func (i *DB) Init() {
	var driver = config.Client.GetString("database.driver")
	if driver == "" {
		log.Client.Panic("database config type error, name: " + driver)
	}

	var cfg ConfigConnsMany
	if err := config.Client.Sub("conns." + driver).Unmarshal(&cfg); err != nil {
		log.Client.Sugar().Panicf("database config read fail, name: %s, err: %v", driver, err)
	}

	i.Driver = driver
	i.Cfg = cfg
}

func (i *DB) Read() *entsql.Driver {
	return new(Connect).Open(i.Driver, i.Cfg, i.Cfg.Read)
}

func (i *DB) Write() *entsql.Driver {
	return new(Connect).Open(i.Driver, i.Cfg, i.Cfg.Write)
}
