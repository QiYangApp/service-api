package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"frame/util/timeutil"
	"github.com/google/uuid"
)

// PermissionRelatedRouter holds the schema definition for the PermissionRelatedRouter entity.
type PermissionRelatedRouter struct {
	ent.Schema
}

// Fields of the PermissionRelatedRouter.
func (PermissionRelatedRouter) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.UUID("router_id", uuid.UUID{}).Comment("路由id"),
		field.UUID("permission_group_id", uuid.UUID{}).Comment("权限分组"),
		field.Int64("create_time").GoType(timeutil.TimeStamp(0)).Default(timeutil.TimeStampNow().Int()).Immutable(),
		field.Int64("update_time").GoType(timeutil.TimeStamp(0)).UpdateDefault(timeutil.TimeStampNow),
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
