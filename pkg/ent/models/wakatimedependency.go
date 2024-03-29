// Code generated by ent, DO NOT EDIT.

package models

import (
	"ent/models/wakatimedependency"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// 每日dependency统计表
type WakatimeDependency struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// wakatime id
	WakatimeID uuid.UUID `json:"wakatime_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 总时长(秒
	TotalSeconds int64 `json:"total_seconds,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WakatimeDependency) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case wakatimedependency.FieldID, wakatimedependency.FieldUserID, wakatimedependency.FieldTotalSeconds:
			values[i] = new(sql.NullInt64)
		case wakatimedependency.FieldName:
			values[i] = new(sql.NullString)
		case wakatimedependency.FieldCreateTime, wakatimedependency.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case wakatimedependency.FieldWakatimeID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WakatimeDependency fields.
func (wd *WakatimeDependency) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case wakatimedependency.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wd.ID = int64(value.Int64)
		case wakatimedependency.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				wd.CreateTime = value.Time
			}
		case wakatimedependency.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				wd.UpdateTime = value.Time
			}
		case wakatimedependency.FieldWakatimeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field wakatime_id", values[i])
			} else if value != nil {
				wd.WakatimeID = *value
			}
		case wakatimedependency.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				wd.UserID = value.Int64
			}
		case wakatimedependency.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				wd.Name = value.String
			}
		case wakatimedependency.FieldTotalSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_seconds", values[i])
			} else if value.Valid {
				wd.TotalSeconds = value.Int64
			}
		default:
			wd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WakatimeDependency.
// This includes values selected through modifiers, order, etc.
func (wd *WakatimeDependency) Value(name string) (ent.Value, error) {
	return wd.selectValues.Get(name)
}

// Update returns a builder for updating this WakatimeDependency.
// Note that you need to call WakatimeDependency.Unwrap() before calling this method if this WakatimeDependency
// was returned from a transaction, and the transaction was committed or rolled back.
func (wd *WakatimeDependency) Update() *WakatimeDependencyUpdateOne {
	return NewWakatimeDependencyClient(wd.config).UpdateOne(wd)
}

// Unwrap unwraps the WakatimeDependency entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wd *WakatimeDependency) Unwrap() *WakatimeDependency {
	_tx, ok := wd.config.driver.(*txDriver)
	if !ok {
		panic("models: WakatimeDependency is not a transactional entity")
	}
	wd.config.driver = _tx.drv
	return wd
}

// String implements the fmt.Stringer.
func (wd *WakatimeDependency) String() string {
	var builder strings.Builder
	builder.WriteString("WakatimeDependency(")
	builder.WriteString(fmt.Sprintf("id=%v, ", wd.ID))
	builder.WriteString("create_time=")
	builder.WriteString(wd.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(wd.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("wakatime_id=")
	builder.WriteString(fmt.Sprintf("%v", wd.WakatimeID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", wd.UserID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(wd.Name)
	builder.WriteString(", ")
	builder.WriteString("total_seconds=")
	builder.WriteString(fmt.Sprintf("%v", wd.TotalSeconds))
	builder.WriteByte(')')
	return builder.String()
}

// WakatimeDependencies is a parsable slice of WakatimeDependency.
type WakatimeDependencies []*WakatimeDependency
