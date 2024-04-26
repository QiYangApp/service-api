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
	Database     string `json:"database,omitempty"`
	MaxOpenConns int    `json:"max_open_conns,omitempty"`
	MaxIdleConns int    `json:"max_idle_conns,omitempty"`
	Host         string `json:"host,omitempty"`
	Port         int    `json:"port,omitempty"`
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
}

type Connect struct {
}

func (p *Connect) Open(t string, cfg Config) *entsql.Driver {
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

func (p *Connect) Connect(t string, cfg Config) *sql.DB {
	db, err := sql.Open(t, p.parseUrl(cfg))
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

func (p Connect) parseUrl(cfg Config) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Database,
		cfg.Host,
		cfg.Port,
	)
}
