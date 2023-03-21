package system

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	"log"
	"os"
	"service-api/src/ent"
	"service-api/src/ent/migrate"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type DatabaseService struct {
	service
	cfg Database
}

func (d *DatabaseService) Handle(r *gin.Engine, cfg ConfigService) {
	d.cfg = cfg.GetDatabase()

	d.connect()
}

// 链接
func (d DatabaseService) connect() {
	db := d.pool()

	driver := entsql.OpenDB(dialect.Postgres, db)

	client := ent.NewClient(ent.Driver(driver))

	// 自定义初始化 schema 和 migrate 文件夹位置
	// schemaPath := "app/database/migrate"
	// migratePath := "app/modules/schema"

	err := client.Schema.Create(
		context.Background(),
		schema.WithGlobalUniqueID(true),
		migrate.WithForeignKeys(false),
		migrate.WithDropIndex(true),  // 删除索引
		migrate.WithDropColumn(true), // 删除列
	)

	if err != nil {
		zap.S().Fatalf("failed creating schema resources: %v", err)
	}

	d.write(client)
}

func (d DatabaseService) write(client *ent.Client) {
	f, err := os.Create("./src/database/migrate.sql")
	if err != nil {
		zap.S().Fatalf("create migrate file: %v", err)
	}

	defer f.Close()

	if err := client.Schema.WriteTo(context.Background(), f); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
}

func (d DatabaseService) pool() *sql.DB {
	db, err := sql.Open(d.cfg.Type, d.parseConfig())
	if err != nil {
		zap.S().Fatalf("failed connecting to the database: %v", err)
	}

	// 设置数据库连接池相关配置
	db.SetMaxIdleConns(int(d.cfg.Config.MaxIdleConns))
	db.SetMaxOpenConns(int(d.cfg.Config.MaxOpenConns))
	db.SetConnMaxLifetime(time.Hour)

	return db
}

// 解析配置

func (d DatabaseService) parseConfig() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		d.cfg.Username,
		d.cfg.Password,
		d.cfg.Database,
		d.cfg.Host,
		d.cfg.Port,
	)
}
