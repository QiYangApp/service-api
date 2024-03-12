package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// UserRelatedRole holds the schema definition for the MemberRelatedRole entity.
type UserRelatedRole struct {
	ent.Schema
}

// Fields of the MemberRelatedRole.
func (UserRelatedRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.UUID("user_id", uuid.UUID{}).Comment("会员id"),
		field.UUID("role_id", uuid.UUID{}).Comment("角色"),
		field.Time("create_time").Default(time.Now).Immutable(),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the MemberRelatedRole.
func (UserRelatedRole) Edges() []ent.Edge {
	return nil
}

func (UserRelatedRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "role_id").
			Unique(),
	}
}
