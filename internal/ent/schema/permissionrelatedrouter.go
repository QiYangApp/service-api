package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// PermissionRelatedRouter holds the schema definition for the PermissionRelatedRouter entity.
type PermissionRelatedRouter struct {
	ent.Schema
}

// Fields of the PermissionRelatedRouter.
func (PermissionRelatedRouter) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Comment(""),
		field.UUID("router_id", uuid.UUID{}).Comment("路由id"),
		field.UUID("permission_group_id", uuid.UUID{}).Comment("权限分组"),
	}
}

// Edges of the PermissionRelatedRouter.
func (PermissionRelatedRouter) Edges() []ent.Edge {
	return nil
}

func (PermissionRelatedRouter) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("permission_group_id", "router_id"),
	}
}

func (PermissionRelatedRouter) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
