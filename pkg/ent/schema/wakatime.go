package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"frame/util/timeutil"
)

// Wakatime holds the schema definition for the Wakatime entity.
type Wakatime struct {
	ent.Schema
}

// Fields of the Wakatime.
func (Wakatime) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id"),
		field.String("key").NotEmpty().Comment("密钥"),
		field.String("api").NotEmpty().Comment("地址"),
		field.String("state").NotEmpty().Comment("状态"),
		field.Int64("create_time").GoType(timeutil.TimeStamp(0)).Default(timeutil.TimeStampNow().Int()).Immutable(),
		field.Int64("update_time").GoType(timeutil.TimeStamp(0)).UpdateDefault(timeutil.TimeStampNow),
	}
}

// Edges of the Wakatime.
func (Wakatime) Edges() []ent.Edge {
	return nil
}

func (Wakatime) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("key"),
		index.Fields("create_time"),
		index.Fields("update_time"),
	}
}
