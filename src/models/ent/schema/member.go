package schema

import (
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

// Annotations of the User.
func (Member) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		schema.Comment("member"),
	}
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Comment("UUID of the"),
		field.String("account").Comment("login account"),
		field.String("email").NotEmpty().MaxLen(128),
		field.String("avatar"),
		field.String("password_sing").NotEmpty().MaxLen(64),
		field.String("password").NotEmpty().MaxLen(32),
		field.String("mobile").NotEmpty().MaxLen(32),
		field.String("nickname").NotEmpty().MaxLen(32),
		field.String("state").NotEmpty().MaxLen(32),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return nil
}

func (Member) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("mobile", "mobile").
			Unique(),
	}
}

func (Member) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
