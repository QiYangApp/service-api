package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"service-api/src/models/enums"
)

// PermissionGroup holds the schema definition for the PermissionGroup entity.
type PermissionGroup struct {
	ent.Schema
}

// Fields of the PermissionGroup.
func (PermissionGroup) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Comment(""),
		field.String("permission_name").MaxLen(32).Default("").NotEmpty().Comment("权限名称"),
		field.String("ioc").Default("").MaxLen(254).NotEmpty().Comment(""),
		field.Int32("sort").Default(0).Comment("排序"),
		field.Int32("left").Default(0).Comment(""),
		field.Int32("right").Default(0).Comment(""),
		field.String("state").Default(enums.ON).NotEmpty().Comment("状态 开启 on 关闭 off"),
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

func (PermissionGroup) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
