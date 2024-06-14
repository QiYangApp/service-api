package schema

import (
	"ent/types/auth"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	timeutil "frame/util/timeutil"
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
		field.Int64("id"),
		field.String("avatar"),
		field.String("email").NotEmpty().MaxLen(128),
		field.String("name").NotEmpty().MaxLen(32),
		field.String("lower_name").NotEmpty().MaxLen(32),
		field.String("full_name").NotEmpty().MaxLen(32),
		field.String("passwd_salt").Default("argon2").NotEmpty().MaxLen(120),
		field.String("passwd_hash_algo").NotEmpty().MaxLen(32),
		field.String("passwd").NotEmpty().MaxLen(32),
		field.String("language").NotEmpty().MaxLen(32),
		field.String("theme").NotEmpty().MaxLen(32),
		field.String("login_name").NotEmpty(),
		field.Int64("login_source").Default(0),
		field.Int("login_type").GoType(auth.Type(0)),
		field.Bool("is_restricted").Default(false),
		field.Bool("is_active").Default(false).Comment("true is activated"),
		field.Bool("prohibit_login").Default(false).Comment("is web login"),
		field.Int64("create_time").GoType(timeutil.TimeStamp(0)).Default(timeutil.TimeStampNow().Int()).Immutable(),
		field.Int64("update_time").GoType(timeutil.TimeStamp(0)).UpdateDefault(timeutil.TimeStampNow),
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
		index.Fields("create_time"),
		index.Fields("update_time"),
	}
}
