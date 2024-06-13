package schema

import (
	"ent/enums/state"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"frame/util/timeutil"
)

// UserRole holds the schema definition for the MemberRole entity.
type UserRole struct {
	ent.Schema
}

// Fields of the MemberRole.
func (UserRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("role_name").MaxLen(64).NotEmpty().Comment("规则名称"),
		field.Int("state").GoType(state.SwitchState(0)).Default(1),
		field.Int64("create_time").GoType(timeutil.TimeStamp(0)).Default(timeutil.TimeStampNow().Int()).Immutable(),
		field.Int64("update_time").GoType(timeutil.TimeStamp(0)).UpdateDefault(timeutil.TimeStampNow),
	}
}

// Edges of the MemberRole.
func (UserRole) Edges() []ent.Edge {
	return nil
}

func (UserRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("state", "role_name").
			Unique(),
		index.Fields("create_time"),
		index.Fields("update_time"),
	}
}
