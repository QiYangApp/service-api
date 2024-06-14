package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"frame/util/timeutil"
)

// UserAuthSource holds the schema definition for the UserAuthSource entity.
type UserAuthSource struct {
	ent.Schema
}

// Fields of the MemberAuthorizeLog.
func (UserAuthSource) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id"),
		field.String("token").NotEmpty().Comment("授权token").MaxLen(254),
		field.String("token_salt"),
		field.String("token_last_eight"),
		field.String("channel").MaxLen(64).Comment("登录渠道"),
		field.String("device").MaxLen(64).Comment("登录设备"),
		field.Text("device_detail").MaxLen(64).Comment("登录信息"),
		field.String("client_ip").NotEmpty().MaxLen(32),
		field.String("remote_ip").NotEmpty().MaxLen(32),
		field.String("snapshot").MaxLen(254),
		field.String("login_name").NotEmpty(),
		field.Int("login_source").Default(0),
		field.Int("login_type").Default(0),
		field.Int64("create_time").GoType(timeutil.TimeStamp(0)).Default(timeutil.TimeStampNow().Int()).Immutable(),
		field.Int64("update_time").GoType(timeutil.TimeStamp(0)).UpdateDefault(timeutil.TimeStampNow),
	}
}

// Edges of the MemberAuthorizeLog.
func (UserAuthSource) Edges() []ent.Edge {
	return nil
}

func (UserAuthSource) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "channel", "device"),
		index.Fields("create_time"),
		index.Fields("update_time"),
	}
}
