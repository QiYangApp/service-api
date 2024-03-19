package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// AccessToken holds the schema definition for the AccessToken entity.
type AccessToken struct {
	ent.Schema
}

// Fields of the AccessToken.
func (AccessToken) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("user_id"),
		field.String("name"),
		field.String("token"),
		field.String("token_hash").Unique(),
		field.String("token_salt"),
		field.String("token_last_eight"),
		field.String("scope"),
		field.String("has_recent_activity"),
		field.String("has_used"),
		field.Time("create_time").Default(time.Now).Immutable(),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now),
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
		index.Fields("update_time"),
	}
}
