package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// MemberRelatedRole holds the schema definition for the MemberRelatedRole entity.
type MemberRelatedRole struct {
	ent.Schema
}

// Fields of the MemberRelatedRole.
func (MemberRelatedRole) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("member_id", uuid.UUID{}).Comment("会员id"),
		field.UUID("role_id", uuid.UUID{}).Comment("角色"),
	}
}

// Edges of the MemberRelatedRole.
func (MemberRelatedRole) Edges() []ent.Edge {
	return nil
}

func (MemberRelatedRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("member_id", "role_id").
			Unique(),
	}
}

func (MemberRelatedRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
