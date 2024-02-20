package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"service-api/enums"
)

// Router holds the schema definition for the Router entity.
type Router struct {
	ent.Schema
}

// Fields of the Router.
func (Router) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Comment(""),
		field.String("route_name").MaxLen(32).Default("").NotEmpty().Comment("路由名称"),
		field.String("route").MaxLen(254).Default("").NotEmpty().Comment("路由"),
		field.String("description").MaxLen(254).Default("").NotEmpty().Comment("描述"),
		field.String("state").MaxLen(32).Default(enums.ON).NotEmpty().Comment("状态 开启 on 关闭 off"),
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
	}
}

func (Router) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
