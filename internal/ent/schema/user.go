package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the Member entity.
type User struct {
	ent.Schema
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		schema.Comment("member"),
	}
}

// Fields of the Member.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("email").NotEmpty().MaxLen(128),
		field.String("avatar"),
		field.String("passwd_salt").NotEmpty().MaxLen(120),
		field.String("passwd_hash_algo").NotEmpty().MaxLen(32),
		field.String("passwd").NotEmpty().MaxLen(32),
		field.String("nickname").NotEmpty().MaxLen(32),
		field.String("language").NotEmpty().MaxLen(32),
		field.String("login_name").NotEmpty(),
		field.Int("login_source").Default(0),
		field.Int("login_type").Default(0),
		field.Bool("is_restricted").Default(false),
		field.Bool("is_active").Default(false).Comment("true is activated"),
		field.Bool("prohibit_login").Default(false).Comment("is web login"),
		field.Time("create_time").Default(time.Now).Immutable(),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Member.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").
			Unique(),
	}
}