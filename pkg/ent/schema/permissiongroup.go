package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"service-api/internal/enums"
	"time"
)

// PermissionGroup holds the schema definition for the PermissionGroup entity.
type PermissionGroup struct {
	ent.Schema
}

// Fields of the PermissionGroup.
func (PermissionGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("permission_name").MaxLen(32).Default("").NotEmpty().Comment("权限名称"),
		field.String("ioc").Default("").MaxLen(254).NotEmpty().Comment(""),
		field.Int32("sort").Default(0).Comment("排序"),
		field.Int32("left").Default(0).Comment(""),
		field.Int32("right").Default(0).Comment(""),
		field.String("state").Default(enums.ON).NotEmpty().Comment("状态 开启 on 关闭 off"),
		field.Time("create_time").Default(time.Now).Immutable(),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now),
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
