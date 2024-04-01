package schema

import (
	"ent/utils/timeutil"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// TwoFactor holds the schema definition for the TwoFactor entity.
type TwoFactor struct {
	ent.Schema
}

// Mixins of the TwoFactor.
func (TwoFactor) Mixins() []ent.Mixin {
	return []ent.Mixin{}
}

// Fields of the TwoFactor.
func (TwoFactor) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id").Unique(),
		field.String("secret"),
		field.String("scratch_salt"),
		field.String("scratch_hash"),
		field.String("last_used_passcode"),

		// Fields go here.
		field.Int64("create_time").GoType(timeutil.TimeStamp(0)).Default(timeutil.TimeStampNow().Int()).Immutable(),
		field.Int64("update_time").GoType(timeutil.TimeStamp(0)).UpdateDefault(timeutil.TimeStampNow),
	}
}

func (TwoFactor) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("create_time"),
		index.Fields("update_time"),
	}
}
