package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"service-api/enums"
)

// MemberRole holds the schema definition for the MemberRole entity.
type MemberRole struct {
	ent.Schema
}

// Fields of the MemberRole.
func (MemberRole) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("role_name").MaxLen(64).NotEmpty().Comment("规则名称"),
		field.String("state").Default(enums.ON).NotEmpty().MaxLen(32).Comment("状态 on 开启 off 关闭"),
	}
}

// Edges of the MemberRole.
func (MemberRole) Edges() []ent.Edge {
	return nil
}

func (MemberRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("state", "role_name").
			Unique(),
	}
}

func (MemberRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
