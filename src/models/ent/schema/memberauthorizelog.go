package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// MemberAuthorizeLog holds the schema definition for the MemberAuthorizeLog entity.
type MemberAuthorizeLog struct {
	ent.Schema
}

// Fields of the MemberAuthorizeLog.
func (MemberAuthorizeLog) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Comment("UUID of the"),
		field.UUID("member_id", uuid.UUID{}).Default(uuid.New).Comment("member UUID of the"),
		field.String("channel").Comment("登录渠道"),
		field.String("device").NotEmpty().MaxLen(64).Comment("登录设备"),
		field.Text("device_detail").NotEmpty().MaxLen(64).Comment("登录信息"),
		field.String("ipv4").NotEmpty().MaxLen(32),
		field.String("ipv6").NotEmpty().MaxLen(32),
		field.String("snapshot").NotEmpty().MaxLen(254),
	}
}

// Edges of the MemberAuthorizeLog.
func (MemberAuthorizeLog) Edges() []ent.Edge {
	return nil
}

func (MemberAuthorizeLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("member_id", "channel", "device"),
	}
}

func (MemberAuthorizeLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
