package db

import (
	"database/sql"
	"fmt"
	"service-api/src/core/config"
	"service-api/src/core/logger"
	"time"
)

type ConnectManage struct {
	Connect
}

func (p *ConnectManage) connect(t string, cfg config.DatabaseOperation) *sql.DB {
	db, err := sql.Open(t, p.parseUrl(cfg))
	if err != nil {
		logger.S().Fatalf("connecting to the database failed %v", err)
	}

	// 设置数据库连接池相关配置
	db.SetMaxIdleConns(int(cfg.Config.MaxIdleConns))
	db.SetMaxOpenConns(int(cfg.Config.MaxOpenConns))
	db.SetConnMaxLifetime(time.Hour)

	if err = db.Ping(); err != nil {
		logger.S().Fatalf("ping to the database failed %v", err)
	}

	return db
}

func (p ConnectManage) parseUrl(cfg config.DatabaseOperation) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Database,
		cfg.Host,
		cfg.Port,
	)
}
