package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"frame/util/timeutil"
)

// WakatimeCategory holds the schema definition for the WakatimeCategory entity.
type WakatimeCategory struct {
	ent.Schema
}

// Edges of the WakatimeCategory.
func (WakatimeCategory) Edges() []ent.Edge {
	return nil
}

// Fields of the WakatimeCategory.
func (WakatimeCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("wakatime_id").Comment("wakatime id"),
		field.Int64("user_id"),
		field.String("name").NotEmpty().Default("").Comment("名称"),
		field.Int64("total_seconds").Default(0).Comment("总时长(秒"),
		field.Int64("create_time").GoType(timeutil.TimeStamp(0)).Default(timeutil.TimeStampNow().Int()).Immutable(),
		field.Int64("update_time").GoType(timeutil.TimeStamp(0)).UpdateDefault(timeutil.TimeStampNow),
	}
}

func (WakatimeCategory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("wakatime_id"),
		index.Fields("user_id", "name"),
	}
}

func (WakatimeCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("每日category统计表"),
	}
}
