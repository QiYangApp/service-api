package db

import (
	"framework/config"
	"framework/log"
	"encoding/json"
	entsql "entgo.io/ent/dialect/sql"
)

type DB struct {
	Cfg    ConfigConnsMany
	Driver string
}

func (i *DB) Init() {
	var driver = config.Client().GetString("database.driver")
	if driver == "" {
		log.Client().Panic("database config type error, name: " + driver)
	}

	conns := config.Client().Get("conns").([]interface{})
	var conn map[string]any
	for _, t := range conns {
		conn = t.(map[string]any)
		if name, ok := conn["driver"].(string); ok && name == driver {
			break
		}
	}

	if len(conn) == 0 {
		log.Client().Panic("database config conns database empty")
	}

	var cfg ConfigConnsMany
	var err error
	if _, ok := conn["driver"]; ok {
		// 使用encoding/json包进行转换
		jsonData, _ := json.Marshal(conn)
		err = json.Unmarshal(jsonData, &cfg)
	}

	if err != nil {
		log.Client().Sugar().Panic(err)
	}

	i.Driver = driver
	i.Cfg = cfg
}

func (i *DB) Read() *entsql.Driver {
	return new(Connect).Open(i.Driver, i.Cfg.Read)
}

func (i *DB) Write() *entsql.Driver {
	return new(Connect).Open(i.Driver, i.Cfg.Write)
}
