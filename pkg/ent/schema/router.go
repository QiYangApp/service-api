package schema

import (
	"ent/enums/state"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"frame/util/timeutil"
)

// Router holds the schema definition for the Router entity.
type Router struct {
	ent.Schema
}

// Fields of the Router.
func (Router) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("route_name").MaxLen(32).Default("").NotEmpty().Comment("路由名称"),
		field.String("route").MaxLen(254).Default("").NotEmpty().Comment("路由"),
		field.String("description").MaxLen(254).Default("").NotEmpty().Comment("描述"),
		field.Int("state").GoType(state.SwitchState(0)).Default(1),
		field.Int64("create_time").GoType(timeutil.TimeStamp(0)).Default(timeutil.TimeStampNow().Int()).Immutable(),
		field.Int64("update_time").GoType(timeutil.TimeStamp(0)).UpdateDefault(timeutil.TimeStampNow),
	}
}

// Edges of the Router.
func (Router) Edges() []ent.Edge {
	return nil
}

func (Router) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("state", "route_name"),
		index.Fields("state", "route"),
		index.Fields("create_time"),
		index.Fields("update_time"),
	}
}
