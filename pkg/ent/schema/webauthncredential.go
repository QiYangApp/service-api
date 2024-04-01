package schema

import (
	"ent/utils/timeutil"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// WebAuthnCredential holds the schema definition for the WebAuthnCredential entity.
type WebAuthnCredential struct {
	ent.Schema
}

// Mixins of the WebAuthnCredential.
func (WebAuthnCredential) Mixins() []ent.Mixin {
	return []ent.Mixin{}
}

// Fields of the WebAuthnCredential.
func (WebAuthnCredential) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name"),
		field.String("lower_name"),
		field.Int64("user_id").Unique(),
		field.Bytes("credential_id").MaxLen(1024),
		field.Bytes("public_key"),
		field.String("attestation_type"),
		field.Bytes("AAGUID"),
		field.Uint32("sign_count"),
		field.Bool("clone_warning"),

		// Fields go here.
		field.Int64("create_time").GoType(timeutil.TimeStamp(0)).Default(timeutil.TimeStampNow().Int()).Immutable(),
		field.Int64("update_time").GoType(timeutil.TimeStamp(0)).UpdateDefault(timeutil.TimeStampNow),
	}
}

func (WebAuthnCredential) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("create_time"),
		index.Fields("update_time"),
	}
}
