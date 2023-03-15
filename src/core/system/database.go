package system

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/postgres"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema/field"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"go.uber.org/zap"
	"time"
)

type DatabaseService struct {
	service
}

func (d *Database) Handle(r *gin.Engine, cfg ConfigService) {
	dbConfig := cfg.GetDatabase()

	path := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Database, dbConfig.Password)
	client, err := ent.Open(cfg.config.Database.Type, path)
	if err != nil {
		zap.S().Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
}

func (d *DatabaseService) DbPool(cfg Database, pgConnect *pg.DB) {

	// 创建数据库链接池
	pool := pg.NewConnPool(&pg.ConnPoolConfig{
		MaxSize: dbCfg.MaxConn,
		MinSize: dbCfg.MaxConn / 2,
	})
}

func (d *DatabaseService) DbConnect(cfg Database, gin *gin.Engine) *pg.DB {
	pg := pg.Connect(&pg.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		User:         cfg.Username,
		Password:     cfg.Password,
		Database:     cfg.Database,
		MaxRetries:   5,
		MinIdleConns: int(cfg.Config.MaxOpenConns) / 2,
		MaxConnAge:   time.Hour * 4,
		PoolSize:     int(cfg.Config.MaxOpenConns),
		IdleTimeout:  time.Minute * 2,
	})

	defer pg.Close()

	return pg
}
