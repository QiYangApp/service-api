package main

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
	"log"
)

func main() {
	//ex, err := entgql.NewExtension()
	//if err != nil {
	//	log.Fatalf("creating entgql extension: %v", err)
	//}

	schemaDir := "./schema"
	targetDir := "./models"
	templateDir := "./template/gen"

	log.Default().Println(schemaDir)

	opts := []entc.Option{
		entc.TemplateDir(templateDir),
		//entc.Extensions(ex),
	}

	if err := entc.Generate(schemaDir, &gen.Config{
		Package:   "ent/models",
		IDType:    &field.TypeInfo{Type: field.TypeInt64},
		Schema:    schemaDir,
		Target:    targetDir,
		Templates: entgql.AllTemplates,
	}, opts...); err != nil {
		log.Fatalf("running ent code generate: %v", err)
	}

}
