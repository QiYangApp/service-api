package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"frame/util/timeutil"
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
		field.Int64("id"),
		field.Int64("wakatime_id").Comment("wakatime id"),
		field.Int64("user_id"),
		field.String("name").NotEmpty().Default("").Comment("名称"),
		field.Int64("total_seconds").Default(0).Comment("总时长(秒"),
		field.Int64("create_time").GoType(timeutil.TimeStamp(0)).Default(timeutil.TimeStampNow().Int()).Immutable(),
		field.Int64("update_time").GoType(timeutil.TimeStamp(0)).UpdateDefault(timeutil.TimeStampNow),
	}
}

func (WakatimeDependency) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("wakatime_id"),
		index.Fields("user_id", "name"),
		index.Fields("create_time"),
		index.Fields("update_time"),
	}
}

func (WakatimeDependency) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("每日dependency统计表"),
	}
}
