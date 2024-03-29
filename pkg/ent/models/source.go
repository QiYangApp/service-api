// Code generated by ent, DO NOT EDIT.

package models

import (
	"encoding/json"
	"ent/models/source"
	"ent/types/auth"
	"ent/utils/timeutil"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Source is the model entity for the Source schema.
type Source struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type auth.Type `json:"type,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// IsSyncEnabled holds the value of the "is_sync_enabled" field.
	IsSyncEnabled bool `json:"is_sync_enabled,omitempty"`
	// Cfg holds the value of the "cfg" field.
	Cfg auth.Config[interface{}] `json:"cfg,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime timeutil.TimeStamp `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime   timeutil.TimeStamp `json:"update_time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Source) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case source.FieldCfg:
			values[i] = new([]byte)
		case source.FieldIsActive, source.FieldIsSyncEnabled:
			values[i] = new(sql.NullBool)
		case source.FieldID, source.FieldType, source.FieldCreateTime, source.FieldUpdateTime:
			values[i] = new(sql.NullInt64)
		case source.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Source fields.
func (s *Source) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case source.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int64(value.Int64)
		case source.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				s.Type = auth.Type(value.Int64)
			}
		case source.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case source.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				s.IsActive = value.Bool
			}
		case source.FieldIsSyncEnabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_sync_enabled", values[i])
			} else if value.Valid {
				s.IsSyncEnabled = value.Bool
			}
		case source.FieldCfg:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field cfg", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.Cfg); err != nil {
					return fmt.Errorf("unmarshal field cfg: %w", err)
				}
			}
		case source.FieldCreateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				s.CreateTime = timeutil.TimeStamp(value.Int64)
			}
		case source.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				s.UpdateTime = timeutil.TimeStamp(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Source.
// This includes values selected through modifiers, order, etc.
func (s *Source) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Source.
// Note that you need to call Source.Unwrap() before calling this method if this Source
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Source) Update() *SourceUpdateOne {
	return NewSourceClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Source entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Source) Unwrap() *Source {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("models: Source is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Source) String() string {
	var builder strings.Builder
	builder.WriteString("Source(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", s.Type))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", s.IsActive))
	builder.WriteString(", ")
	builder.WriteString("is_sync_enabled=")
	builder.WriteString(fmt.Sprintf("%v", s.IsSyncEnabled))
	builder.WriteString(", ")
	builder.WriteString("cfg=")
	builder.WriteString(fmt.Sprintf("%v", s.Cfg))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(fmt.Sprintf("%v", s.CreateTime))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(fmt.Sprintf("%v", s.UpdateTime))
	builder.WriteByte(')')
	return builder.String()
}

// Sources is a parsable slice of Source.
type Sources []*Source
