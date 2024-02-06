package db

import (
	"app/log"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"time"
)

type Config struct {
	Type  string                     `json:"type,omitempty"`
	Conns map[string]ConfigConnsMany `json:"conns,omitempty"`
}

type ConfigConnsSingle struct {
	MaxOpenConns int    `json:"MaxOpenConns,omitempty"`
	MaxIdleConns int    `json:"MaxIdleConns,omitempty"`
	Host         string `json:"host,omitempty"`
	Port         int    `json:"port,omitempty"`
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	Database     string `json:"database,omitempty"`
}

type ConfigConnsMany struct {
	Read  ConfigConnsSingle `json:"read"`
	Write ConfigConnsSingle `json:"write"`
}

type Connect struct {
}

func (p *Connect) Open(t string, cfg ConfigConnsSingle) *entsql.Driver {
	return entsql.OpenDB(p.dbType(t), p.Connect(t, cfg))
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

func (p *Connect) Connect(t string, cfg ConfigConnsSingle) *sql.DB {
	db, err := sql.Open(t, p.parseUrl(cfg))
	if err != nil {
		log.Client().Sugar().Fatalf("connecting to the database failed %v", err)
	}

	// 设置数据库连接池相关配置
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(time.Minute)

	if err = db.Ping(); err != nil {
		log.Client().Sugar().Fatalf("ping to the database failed %v", err)
	}

	return db
}

func (p Connect) parseUrl(cfg ConfigConnsSingle) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Database,
		cfg.Host,
		cfg.Port,
	)
}
