package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Accounts holds the schema definition for the UserAccounts entity.
type Accounts struct {
	ent.Schema
}

// Fields of the UserAccounts.
func (Accounts) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("user_id"),
		field.String("account"),
		field.Uint8("type"),
		field.String("desc"),
		field.Bool("is_activated"),
		field.Bool("is_primary").Default(false),
	}
}

func (Accounts) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("account", "type").Unique(),
	}
}

// Edges of the UserAccounts.
func (Accounts) Edges() []ent.Edge {
	return nil
}
