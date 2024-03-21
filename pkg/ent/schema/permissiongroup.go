package schema

import (
	"ent/enums/state"
	"ent/utils/timeutil"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PermissionGroup holds the schema definition for the PermissionGroup entity.
type PermissionGroup struct {
	ent.Schema
}

// Fields of the PermissionGroup.
func (PermissionGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("permission_name").MaxLen(32).Default("").NotEmpty().Comment("权限名称"),
		field.String("ioc").Default("").MaxLen(254).NotEmpty().Comment(""),
		field.Int32("sort").Default(0).Comment("排序"),
		field.Int32("left").Default(0).Comment(""),
		field.Int32("right").Default(0).Comment(""),
		field.Int("state").GoType(state.SwitchState(0)).Default(1),
		field.Int64("create_time").GoType(timeutil.TimeStamp(0)).Default(timeutil.TimeStampNow().Int()).Immutable(),
		field.Int64("update_time").GoType(timeutil.TimeStamp(0)).UpdateDefault(timeutil.TimeStampNow),
	}
}

// Edges of the PermissionGroup.
func (PermissionGroup) Edges() []ent.Edge {
	return nil
}

func (PermissionGroup) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("left", "right"),
		index.Fields("state", "sort", "permission_name"),
	}
}
