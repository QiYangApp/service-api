package schema

import (
	"time"

	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("account"),
		field.String("email").Unique(),
		field.String("avatar"),
		field.String("mobile"),
		field.String("nickname").Optional().Nillable(),
		field.String("state"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return nil
}
