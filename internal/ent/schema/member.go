package schema

import (
	"github.com/google/uuid"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.String("email").NotEmpty().MaxLen(128),
		field.String("avatar"),
		field.String("passwd_salt").NotEmpty().MaxLen(120),
		field.String("passwd_hash_algo").NotEmpty().MaxLen(32),
		field.String("passwd").NotEmpty().MaxLen(32),
		field.String("nickname").NotEmpty().MaxLen(32),
		field.Bool("is_restricted").Default(false),
		field.Bool("is_active").Default(false).Comment("true 已激活"),
		field.Time("create_time").Default(time.Now).Immutable(),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return nil
}

func (Member) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").
			Unique(),
	}
}
