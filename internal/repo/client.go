package repo

import (
	"context"
	"ent/models"
	"ent/models/migrate"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"errors"
	"fmt"
	"framework/config"
	"framework/db"
	"framework/log"
	"framework/utils"
	"os"
	"path/filepath"
	"sync"
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
			config.Client.GetBool("debug"),
			config.Client.GetBool("database.schema_init"),
		)
	})
}

func New(opts []models.Option, debug, schemaInit bool) *models.Client {
	models.SetDebugState(debug)
	return NewClient(opts, debug, schemaInit)
}

func NewClient(opts []models.Option, debug, schemaInit bool) *models.Client {
	client := models.NewClient(opts...)

	err := client.Schema.Create(
		context.Background(),
		schema.WithGlobalUniqueID(true),
		migrate.WithForeignKeys(false),
		migrate.WithDropIndex(true),  // 删除索引
		migrate.WithDropColumn(true), // 删除列
	)

	if err != nil {
t 		log.Client.Sugar().Panicf("failed creating schema resources: %v", err)
	}

	if schemaInit {
		if err := schemaInitRun(); err != nil {
			log.Client.Sugar().Panicf("failed init schema: %v", err)
		}
	}

	err = write(utils.Path.Join(utils.Path.StoragePath, "migrate.sql"), client)
	if err != nil {
		log.Client.Sugar().Panicf("write sql file fail: %v", err)
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

func schemaInitRun() error {
	ex, err := entgql.NewExtension()
	if err != nil {
		log.Client.Sugar().Fatalf("creating entgql extension: %v", err)
	}

	schemaDir := utils.Path.Join(utils.Path.RootPath, "/pkg/ent/schema")
	targetDir := utils.Path.Join(utils.Path.RootPath, "/pkg/ent/models")
	templateDir := utils.Path.Join(utils.Path.RootPath, "/pkg/ent/template/gen")

	opts := []entc.Option{
		entc.TemplateDir(templateDir),
		entc.Extensions(ex),
	}

	if err := entc.Generate(schemaDir, &gen.Config{
		Schema:    schemaDir,
		Target:    targetDir,
		Templates: entgql.AllTemplates,
	}, opts...); err != nil {
		return errors.New(fmt.Sprintf("running ent code generate: %v", err))
	}

	return nil
}
