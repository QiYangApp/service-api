package ent

import (
	"context"
	"ent/models"
	"ent/models/migrate"
	"entgo.io/ent/dialect/sql/schema"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var Client *models.Client = nil

var once = sync.Once{}

func New(opts []models.Option, migratePath string, debug bool) *models.Client {
	once.Do(func() {
		models.SetDebugState(debug)
		Client = NewClient(opts, migratePath, debug)
	})

	return Client
}

func NewClient(opts []models.Option, migratePath string, debug bool) *models.Client {
	client := models.NewClient(opts...)

	err := client.Schema.Create(
		context.Background(),
		schema.WithGlobalUniqueID(true),
		migrate.WithForeignKeys(false),
		migrate.WithDropIndex(true),  // 删除索引
		migrate.WithDropColumn(true), // 删除列
	)

	if err != nil {
		log.Panicf("failed creating schema resources: %v", err)
		return nil
	}
	//
	//if err := schemaInitRun(debug); err != nil {
	//
	//}

	err = write(migratePath, client)
	if err != nil {
		log.Panicf("write sql file fail: %v", err)
		return nil
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

//func schemaInitRun(debug bool) error {
//	schemaDir, _ := filepath.Abs("/ent/schema")
//	if err := entc.Generate(schemaDir, &gen.Config{}); err != nil {
//		return errors.New(fmt.Sprintf("running ent code generate: %v", err))
//	}
//
//	return nil
//}
