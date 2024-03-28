package db

import (
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"framework/log"
	_ "github.com/jackc/pgx/v5/stdlib"
	"time"
)

type Config struct {
	Type  string                     `json:"type,omitempty"`
	Conns map[string]ConfigConnsMany `json:"conns,omitempty"`
}

type ConfigConnsSingle struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type ConfigConnsMany struct {
	Database     string            `json:"database,omitempty"`
	MaxOpenConns int               `json:"max_open_conns,omitempty"`
	MaxIdleConns int               `json:"max_idle_conns,omitempty"`
	Read         ConfigConnsSingle `json:"read"`
	Write        ConfigConnsSingle `json:"write"`
}

type Connect struct {
}

func (p *Connect) Open(t string, cfg ConfigConnsMany, singcfg ConfigConnsSingle) *entsql.Driver {
	return entsql.OpenDB(p.dbType(t), p.Connect(t, cfg, singcfg))
}

func (i *Connect) dbType(t string) string {
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

func (p *Connect) Connect(t string, cfg ConfigConnsMany, singcfg ConfigConnsSingle) *sql.DB {
	db, err := sql.Open(t, p.parseUrl(cfg, singcfg))
	if err != nil {
		log.Client.Sugar().Fatalf("connecting to the database failed %v", err)
	}

	// 设置数据库连接池相关配置
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(time.Minute)

	if err = db.Ping(); err != nil {
		log.Client.Sugar().Fatalf("ping to the database failed %v", err)
	}

	return db
}

func (p Connect) parseUrl(cfg ConfigConnsMany, singcfg ConfigConnsSingle) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		singcfg.Username,
		singcfg.Password,
		cfg.Database,
		singcfg.Host,
		singcfg.Port,
	)
}
