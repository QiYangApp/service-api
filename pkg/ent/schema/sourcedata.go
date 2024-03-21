package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// SourceData holds the schema definition for the SourceData entity.
type SourceData struct {
	ent.Schema
}

// Fields of the SourceData.
func (SourceData) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.UUID("user_id", uuid.UUID{}).Comment("会员id"),
		field.String("type").MaxLen(32).Default("").NotEmpty().Comment("类型"),
		field.String("sub_type").MaxLen(32).Default("").Comment("子类型"),
		field.String("info").Default("").Comment("信息"),
		field.Text("snapshot").Default("").Comment("数据"),
		field.Time("create_time").Default(time.Now).Immutable(),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the SourceData.
func (SourceData) Edges() []ent.Edge {
	return nil
}

func (SourceData) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("type", "sub_type"),
	}
}
