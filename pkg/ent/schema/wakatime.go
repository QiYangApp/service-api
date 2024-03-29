package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
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
		field.Time("create_time").Default(time.Now).Immutable(),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now),
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
	}
}
