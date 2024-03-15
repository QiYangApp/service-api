package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"service-api/internal/modules/timeutil"
)

// Test holds the schema definition for the Test entity.
type Test struct {
	ent.Schema
}

// Mixins of the Test.
func (Test) Mixins() []ent.Mixin {
	return []ent.Mixin{}
}

// Fields of the Test.
func (Test) Fields() []ent.Field {
	return []ent.Field{
		// Fields go here.
		field.Int64("create_time").GoType(timeutil.TimeStamp(0)).Default(timeutil.TimeStampNow().Int()).Immutable(),
		field.Int64("update_time").GoType(timeutil.TimeStamp(0)).Default(timeutil.TimeStampNow().Int()).UpdateDefault(timeutil.TimeStampNow()),
	}
}

func (Test) Indexes() []ent.Index {
	return []ent.Index{}
}
