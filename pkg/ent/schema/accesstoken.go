package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"frame/util/timeutil"
)

// AccessToken holds the schema definition for the AccessToken entity.
type AccessToken struct {
	ent.Schema
}

// Fields of the AccessToken.
func (AccessToken) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id"),
		field.String("name"),
		field.String("token"),
		field.String("token_hash").Unique(),
		field.String("token_salt"),
		field.String("token_last_eight"),
		field.String("scope"),
		field.String("has_recent_activity"),
		field.String("has_used"),
		field.Int64("create_time").GoType(timeutil.TimeStamp(0)).Default(timeutil.TimeStampNow().Int()).Immutable(),
	}
}

// Edges of the AccessToken.
func (AccessToken) Edges() []ent.Edge {
	return nil
}

func (AccessToken) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("token_last_eight"),
		index.Fields("create_time"),
	}
}
