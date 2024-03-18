package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// WakatimeDependency holds the schema definition for the WakatimeDependency entity.
type WakatimeDependency struct {
	ent.Schema
}

func (WakatimeDependency) Edges() []ent.Edge {
	return nil
}

func (WakatimeDependency) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Comment(""),
		field.UUID("wakatime_id", uuid.UUID{}).Comment("wakatime id"),
		field.UUID("user_id", uuid.UUID{}).Comment("会员id"),
		field.String("name").NotEmpty().Default("").Comment("名称"),
		field.Int64("total_seconds").Default(0).Comment("总时长(秒"),
	}
}

func (WakatimeDependency) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("wakatime_id"),
		index.Fields("user_id", "name"),
	}
}

func (WakatimeDependency) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (WakatimeDependency) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("每日dependency统计表"),
	}
}