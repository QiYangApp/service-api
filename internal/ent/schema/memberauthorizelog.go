package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"time"
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
		field.String("token").NotEmpty().Comment("授权token").MaxLen(254),
		field.String("channel").MaxLen(64).Comment("登录渠道"),
		field.String("device").MaxLen(64).Comment("登录设备"),
		field.Text("device_detail").MaxLen(64).Comment("登录信息"),
		field.String("client_ip").NotEmpty().MaxLen(32),
		field.String("remote_ip").NotEmpty().MaxLen(32),
		field.String("snapshot").MaxLen(254),
		field.Time("create_time").Default(time.Now).Immutable(),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now),
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
