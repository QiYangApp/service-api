package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// MemberRoleRelatedPermission holds the schema definition for the MemberRoleRelatedPermission entity.
type MemberRoleRelatedPermission struct {
	ent.Schema
}

// Fields of the MemberRoleRelatedPermission.
func (MemberRoleRelatedPermission) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.UUID("role_id", uuid.UUID{}).Comment("规则id"),
		field.UUID("permission_group_id", uuid.UUID{}).Comment("权限分组id"),
		field.Time("create_time").Default(time.Now).Immutable(),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the MemberRoleRelatedPermission.
func (MemberRoleRelatedPermission) Edges() []ent.Edge {
	return nil
}

func (MemberRoleRelatedPermission) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("role_id", "permission_group_id").
			Unique(),
	}
}
