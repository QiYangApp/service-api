package system

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"

	"entgo.io/ent/cmd"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
)

type DatabaseService struct {
	service
}

func (d *DatabaseService) Handle(r *gin.Engine, cfg ConfigService) {
	dbConfig := cfg.GetDatabase()

	d.connect(dbConfig)
}

// 链接
func (d DatabaseService) connect(cfg Database) {
	pool, err := pgxpool.ConnectConfig(context.Background(), d.parseConfig(cfg))
	if err != nil {
		zap.S().Fatalf("Error connecting to database, %s", err)
	}

	driverConfig := &driver.Config{
		DriverName: dialect.Postgres,
		Conn:       pool.Acquire,
		Dialect:    sql.Postgres,
	}

	err = entc.Generate("./src/app/models/schema", &gen.Config{
		NamingStrategy: func(s string) string {
			return "qy_" + s
		},
		Driver: driverConfig,
	})

	if err != nil {
		panic(err)
	}

}

// func (d *DatabaseService) migrate() {
// 	cfg, err := load.Config("./src/app/models/schema")
// 	if err != nil {
// 		zap().S().Fatalf("failed loading config: %v", err)
// 	}
// 	genCfg := &gen.Config{
// 		Schema: cfg,
// 	}
// 	err = entc.Generate("./src/database/migrate", genCfg, entc.TemplateDir("template"))
// 	if err != nil {
// 		zap().S().Fatalf("failed running ent codegen: %v", err)
// 	}
// }

// 解析配置

func (d DatabaseService) parseConfig(cfg Database) *pgxpool.Config {
	// Initialize database connection pool
	config, err := pgxpool.ParseConfig(fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.Database,
	))

	if err != nil {
		zap.S().Fatalf("Error parsing database configuration, %s", err)
	}

	return config
}
