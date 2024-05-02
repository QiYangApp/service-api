package repo

import (
	"context"
	"ent/models"
	"ent/models/migrate"
	"frame/modules/db"
	"frame/modules/log"
	"frame/util/path"
	"os"
	"path/filepath"
	"service-api/internal/modules/setting"
	"sync"

	"entgo.io/ent/dialect/sql/schema"
)

var once = sync.Once{}

var Client *models.Client = nil

func Init() {
	once.Do(func() {
		conns := db.DB{}

		conns.Init()

		Client = New(
			[]models.Option{
				models.Driver(&db.MultiDriver{R: conns.Read(), W: conns.Write()}),
			},
			setting.AppSetting.Debug,
		).Debug()
	})
}

func New(opts []models.Option, debug bool) *models.Client {
	models.SetDebugState(debug)
	return NewClient(opts)
}

func NewClient(opts []models.Option) *models.Client {
	client := models.NewClient(opts...)

	err := client.Schema.Create(
		context.Background(),
		schema.WithGlobalUniqueID(true),
		migrate.WithForeignKeys(false),
		migrate.WithDropIndex(true),  // 删除索引
		migrate.WithDropColumn(true), // 删除列
	)

	if err != nil {
		log.Sugar().Panicf("failed creating schema resources: %v", err)
	}

	err = write(path.JoinPath(path.RunStoragePath, "migrate.sql"), client)
	if err != nil {
		log.Sugar().Panicf("write sql file fail: %v", err)
	}

	return client
}

func write(path string, client *models.Client) error {
	migrateFile, _ := filepath.Abs(path)
	f, err := os.Create(migrateFile)
	if err != nil {
		return err
	}

	defer f.Close()

	if err := client.Schema.WriteTo(context.Background(), f); err != nil {
		return err
	}

	return nil
}
