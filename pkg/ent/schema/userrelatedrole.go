package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"frame/util/timeutil"
	"github.com/google/uuid"
)

// UserRelatedRole holds the schema definition for the MemberRelatedRole entity.
type UserRelatedRole struct {
	ent.Schema
}

// Fields of the MemberRelatedRole.
func (UserRelatedRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id"),
		field.UUID("role_id", uuid.UUID{}).Comment("角色"),
		field.Int64("create_time").GoType(timeutil.TimeStamp(0)).Default(timeutil.TimeStampNow().Int()).Immutable(),
		field.Int64("update_time").GoType(timeutil.TimeStamp(0)).UpdateDefault(timeutil.TimeStampNow),
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
		index.Fields("create_time"),
		index.Fields("update_time"),
	}
}
