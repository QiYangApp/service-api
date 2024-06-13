package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"frame/util/timeutil"
)

// SourceData holds the schema definition for the SourceData entity.
type SourceData struct {
	ent.Schema
}

// Fields of the SourceData.
func (SourceData) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id"),
		field.String("type").MaxLen(32).Default("").NotEmpty().Comment("类型"),
		field.String("sub_type").MaxLen(32).Default("").Comment("子类型"),
		field.String("info").Default("").Comment("信息"),
		field.Text("snapshot").Default("").Comment("数据"),
		field.Int64("create_time").GoType(timeutil.TimeStamp(0)).Default(timeutil.TimeStampNow().Int()).Immutable(),
		field.Int64("update_time").GoType(timeutil.TimeStamp(0)).UpdateDefault(timeutil.TimeStampNow),
	}
}

// Edges of the SourceData.
func (SourceData) Edges() []ent.Edge {
	return nil
}

func (SourceData) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("type", "sub_type"),
		index.Fields("create_time"),
		index.Fields("update_time"),
	}
}
