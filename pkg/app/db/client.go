package db

import (
	"app/config"
	"app/log"
	"encoding/json"
	"entgo.io/ent/dialect"
)

type DB struct {
	Cfg      ConfigConnsMany
	Classify string
}

func (i *DB) Init() {
	var t = config.Client().GetString("database.driver")
	if t == "" {
		log.Client().Panic("database config type error")
	}

	database := config.Client().GetStringMap("database.conns")
	if len(database) == 0 {
		log.Client().Panic("database config conns database empty")
	}

	var cfg ConfigConnsMany
	var err error
	if temp, ok := database[t]; ok {
		// 使用encoding/json包进行转换
		jsonData, _ := json.Marshal(temp)
		err = json.Unmarshal(jsonData, &cfg)
	}

	if err != nil {
		log.Client().Sugar().Panic(err)
	}

	i.Classify = t
	i.Cfg = cfg
}

func (i *DB) Read() dialect.Driver {
	return new(Connect).Open(i.Classify, i.Cfg.Read)
}

func (i *DB) Write() dialect.Driver {
	return new(Connect).Open(i.Classify, i.Cfg.Write)
}
