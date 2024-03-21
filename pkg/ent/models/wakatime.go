// Code generated by ent, DO NOT EDIT.

package models

import (
	"ent/models/wakatime"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Wakatime is the model entity for the Wakatime schema.
type Wakatime struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// 密钥
	Key string `json:"key,omitempty"`
	// 地址
	API string `json:"api,omitempty"`
	// 状态
	State string `json:"state,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime   time.Time `json:"update_time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Wakatime) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case wakatime.FieldID, wakatime.FieldUserID:
			values[i] = new(sql.NullInt64)
		case wakatime.FieldKey, wakatime.FieldAPI, wakatime.FieldState:
			values[i] = new(sql.NullString)
		case wakatime.FieldCreateTime, wakatime.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Wakatime fields.
func (w *Wakatime) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case wakatime.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int64(value.Int64)
		case wakatime.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				w.UserID = value.Int64
			}
		case wakatime.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				w.Key = value.String
			}
		case wakatime.FieldAPI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field api", values[i])
			} else if value.Valid {
				w.API = value.String
			}
		case wakatime.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				w.State = value.String
			}
		case wakatime.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				w.CreateTime = value.Time
			}
		case wakatime.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				w.UpdateTime = value.Time
			}
		default:
			w.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Wakatime.
// This includes values selected through modifiers, order, etc.
func (w *Wakatime) Value(name string) (ent.Value, error) {
	return w.selectValues.Get(name)
}

// Update returns a builder for updating this Wakatime.
// Note that you need to call Wakatime.Unwrap() before calling this method if this Wakatime
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Wakatime) Update() *WakatimeUpdateOne {
	return NewWakatimeClient(w.config).UpdateOne(w)
}

// Unwrap unwraps the Wakatime entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Wakatime) Unwrap() *Wakatime {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("models: Wakatime is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Wakatime) String() string {
	var builder strings.Builder
	builder.WriteString("Wakatime(")
	builder.WriteString(fmt.Sprintf("id=%v, ", w.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", w.UserID))
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(w.Key)
	builder.WriteString(", ")
	builder.WriteString("api=")
	builder.WriteString(w.API)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(w.State)
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(w.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(w.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Wakatimes is a parsable slice of Wakatime.
type Wakatimes []*Wakatime