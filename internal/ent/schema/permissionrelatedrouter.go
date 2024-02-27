package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"time"
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
		field.Time("create_time").Default(time.Now).Immutable(),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now),
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
