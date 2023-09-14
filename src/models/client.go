package models

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"service-api/src/models/ent"
	"service-api/src/models/ent/migrate"
)

func NewClient(path string, opts ...ent.Option) (*ent.Client, error) {
	client := ent.NewClient(opts...)

	err := client.Schema.Create(
		context.Background(),
		schema.WithGlobalUniqueID(true),
		migrate.WithForeignKeys(false),
		migrate.WithDropIndex(true),  // 删除索引
		migrate.WithDropColumn(true), // 删除列
	)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed creating schema resources: %v", err))
	}

	schemaDir, _ := filepath.Abs(path + "/models/ent/schema")
	if err := entc.Generate(schemaDir, &gen.Config{}); err != nil {
		return nil, errors.New(fmt.Sprintf("running ent code generate: %v", err))
	}

	err = write(path, client)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("write sql file fail: %v", err))
	}

	return client, nil
}

func write(path string, client *ent.Client) error {
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
