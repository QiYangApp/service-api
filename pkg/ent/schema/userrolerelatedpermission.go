package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"frame/util/timeutil"
	"github.com/google/uuid"
)

// MemberRoleRelatedPermission holds the schema definition for the MemberRoleRelatedPermission entity.
type MemberRoleRelatedPermission struct {
	ent.Schema
}

// Fields of the MemberRoleRelatedPermission.
func (MemberRoleRelatedPermission) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.UUID("role_id", uuid.UUID{}).Comment("规则id"),
		field.UUID("permission_group_id", uuid.UUID{}).Comment("权限分组id"),
		field.Int64("create_time").GoType(timeutil.TimeStamp(0)).Default(timeutil.TimeStampNow().Int()).Immutable(),
		field.Int64("update_time").GoType(timeutil.TimeStamp(0)).UpdateDefault(timeutil.TimeStampNow),
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
		index.Fields("create_time"),
		index.Fields("update_time"),
	}
}
