package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"frame/util/timeutil"
)

// Accounts holds the schema definition for the UserAccounts entity.
type Accounts struct {
	ent.Schema
}

// Fields of the UserAccounts.
func (Accounts) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id"),
		field.String("account"),
		field.Uint8("type"),
		field.String("desc"),
		field.Bool("is_private").Default(true),
		field.Bool("is_activated"),
		field.Bool("is_primary").Default(false),
		field.Int64("create_time").GoType(timeutil.TimeStamp(0)).Default(timeutil.TimeStampNow().Int()).Immutable(),
		field.Int64("update_time").GoType(timeutil.TimeStamp(0)).UpdateDefault(timeutil.TimeStampNow),
	}
}

func (Accounts) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("account", "type").Unique(),
		index.Fields("create_time"),
		index.Fields("update_time"),
	}
}

// Edges of the UserAccounts.
func (Accounts) Edges() []ent.Edge {
	return nil
}
