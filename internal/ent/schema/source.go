package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"service-api/internal/modules/timeutil"
)

// Source holds the schema definition for the Source entity.
type Source struct {
	ent.Schema
}

// Fields of the Source.
func (Source) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("type"),
		field.String("name"),
		field.Bool("is_active"),
		field.Bool("is_sync_enabled"),
		field.Text("cfg"),
		field.Int64("create_time").GoType(timeutil.TimeStamp(0)).Default(timeutil.TimeStampNow().Int()).Immutable(),
		field.Int64("update_time").GoType(timeutil.TimeStamp(0)).Default(timeutil.TimeStampNow().Int()).UpdateDefault(timeutil.TimeStampNow()),
	}
}

// Edges of the Source.
func (Source) Edges() []ent.Edge {
	return nil
}

func (Source) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("is_active", "is_sync_enabled"),
		index.Fields("create_time"),
		index.Fields("update_time"),
		index.Fields("name").Unique(),
	}
}
