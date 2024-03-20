package schema

import (
	"ent/enums/state"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// UserRole holds the schema definition for the MemberRole entity.
type UserRole struct {
	ent.Schema
}

// Fields of the MemberRole.
func (UserRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("role_name").MaxLen(64).NotEmpty().Comment("规则名称"),
		field.Int("state").GoType(state.SwitchState(0)).Default(1),
		field.Time("create_time").Default(time.Now).Immutable(),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now),
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
	}
}
