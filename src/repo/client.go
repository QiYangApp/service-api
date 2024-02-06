package repo

import (
	"app/db"
	"app/log"
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"os"
	"path/filepath"
	"service-api/src/repo/models"
	"service-api/src/repo/models/migrate"
	"sync"
)

var client *models.Client

var once = sync.Once{}

func New() *models.Client {
	once.Do(func() {
		conns := db.DB{}

		conns.Init()

		client = NewClient([]models.Option{
			models.Driver(&db.MultiDriver{R: conns.Read(), W: conns.Write()}),
		})
	})

	return client
}

func Client() *models.Client {
	return New()
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
		log.Client().Sugar().Panicf("failed creating schema resources: %v", err)
		return nil
	}

	//schemaDir, _ := filepath.Abs(path + "/models/ent/schema")
	//if err := entc.Generate(schemaDir, &gen.Config{}); err != nil {
	//	return nil, errors.New(fmt.Sprintf("running ent code generate: %v", err))
	//}
	//
	//err = write(path, client)
	//if err != nil {
	//	return nil, errors.New(fmt.Sprintf("write sql file fail: %v", err))
	//}

	return client
}

func write(path string, client *models.Client) error {
	migrateFile, _ := filepath.Abs(path + "/models/ent/migrate/migrate.sql")
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
