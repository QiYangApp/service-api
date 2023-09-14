package db

import (
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"go.uber.org/zap"
	"service-api/src/core/config"
	"service-api/src/core/helpers"
	"service-api/src/core/logger"
	"service-api/src/models"
	"service-api/src/models/ent"
)

type Instance struct {
	cfg    config.Database
	Client *ent.Client
}

func (i *Instance) InitClient() {
	var err error
	i.Client, err = models.NewClient(
		helpers.PathInstance.GetCurrentRunPath(),
		ent.Driver(&multiDriver{r: i.read(), w: i.write()}),
	)

	if err != nil {
		logger.S().Panicln(zap.Error(err))
	}
}

func (i *Instance) read() *entsql.Driver {
	return i.open(i.cfg.Read)
}

func (i *Instance) write() *entsql.Driver {
	return i.open(i.cfg.Write)
}

func (i *Instance) open(cfg config.DatabaseOperation) *entsql.Driver {
	return entsql.OpenDB(i.dbType(i.cfg.Type), i.connect(cfg))
}

func (i *Instance) connect(cfg config.DatabaseOperation) *sql.DB {
	return new(ConnectManage).connect(i.cfg.Type, cfg)
}

func (i *Instance) dbType(t string) string {
	switch t {
	case "pgx":
		return dialect.Postgres
	case "mysql":
		return dialect.MySQL
	case "sqlite":
		return dialect.SQLite
	default:
		return dialect.Postgres
	}
}
