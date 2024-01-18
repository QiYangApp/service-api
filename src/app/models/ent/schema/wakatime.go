package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Wakatime holds the schema definition for the Wakatime entity.
type Wakatime struct {
	ent.Schema
}

// Fields of the Wakatime.
func (Wakatime) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Comment(""),
		field.UUID("member_id", uuid.UUID{}).Comment("会员id"),
		field.String("key").NotEmpty().Comment("密钥"),
		field.String("api").NotEmpty().Comment("地址"),
		field.String("state").NotEmpty().Comment("状态"),
	}
}

// Edges of the Wakatime.
func (Wakatime) Edges() []ent.Edge {
	return nil
}

func (Wakatime) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("member_id"),
		index.Fields("key"),
	}
}

func (Wakatime) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
