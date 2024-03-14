package schema

import (
	"service-api/internal/modules/timeutil"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.Time("create_time").GoType(timeutil.TimeStamp(0)).Default(time.Now).Immutable(),
		field.Time("update_time").GoType(timeutil.TimeStamp(0)).Default(time.Now).UpdateDefault(time.Now),
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
