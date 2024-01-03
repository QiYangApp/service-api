package schema

import "entgo.io/ent"

// WakatimeEntity holds the schema definition for the WakatimeEntity entity.
type WakatimeEntity struct {
	ent.Schema
}

// Fields of the WakatimeEntity.
func (WakatimeEntity) Fields() []ent.Field {
	return nil
}

// Edges of the WakatimeEntity.
func (WakatimeEntity) Edges() []ent.Edge {
	return nil
}
