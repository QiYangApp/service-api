package schema

import "entgo.io/ent"

// WakatimeProject holds the schema definition for the WakatimeProject entity.
type WakatimeProject struct {
	ent.Schema
}

// Fields of the WakatimeProject.
func (WakatimeProject) Fields() []ent.Field {
	return nil
}

// Edges of the WakatimeProject.
func (WakatimeProject) Edges() []ent.Edge {
	return nil
}
