// Code generated by ent, DO NOT EDIT.

package models

import (
	"ent/models/wakatimecategory"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// 每日category统计表
type WakatimeCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// wakatime id
	WakatimeID uuid.UUID `json:"wakatime_id,omitempty"`
	// 会员id
	UserID uuid.UUID `json:"user_id,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 总时长(秒
	TotalSeconds int64 `json:"total_seconds,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WakatimeCategory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case wakatimecategory.FieldID, wakatimecategory.FieldTotalSeconds:
			values[i] = new(sql.NullInt64)
		case wakatimecategory.FieldName:
			values[i] = new(sql.NullString)
		case wakatimecategory.FieldCreateTime, wakatimecategory.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case wakatimecategory.FieldWakatimeID, wakatimecategory.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WakatimeCategory fields.
func (wc *WakatimeCategory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case wakatimecategory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wc.ID = int64(value.Int64)
		case wakatimecategory.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				wc.CreateTime = value.Time
			}
		case wakatimecategory.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				wc.UpdateTime = value.Time
			}
		case wakatimecategory.FieldWakatimeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field wakatime_id", values[i])
			} else if value != nil {
				wc.WakatimeID = *value
			}
		case wakatimecategory.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				wc.UserID = *value
			}
		case wakatimecategory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				wc.Name = value.String
			}
		case wakatimecategory.FieldTotalSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_seconds", values[i])
			} else if value.Valid {
				wc.TotalSeconds = value.Int64
			}
		default:
			wc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WakatimeCategory.
// This includes values selected through modifiers, order, etc.
func (wc *WakatimeCategory) Value(name string) (ent.Value, error) {
	return wc.selectValues.Get(name)
}

// Update returns a builder for updating this WakatimeCategory.
// Note that you need to call WakatimeCategory.Unwrap() before calling this method if this WakatimeCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (wc *WakatimeCategory) Update() *WakatimeCategoryUpdateOne {
	return NewWakatimeCategoryClient(wc.config).UpdateOne(wc)
}

// Unwrap unwraps the WakatimeCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wc *WakatimeCategory) Unwrap() *WakatimeCategory {
	_tx, ok := wc.config.driver.(*txDriver)
	if !ok {
		panic("models: WakatimeCategory is not a transactional entity")
	}
	wc.config.driver = _tx.drv
	return wc
}

// String implements the fmt.Stringer.
func (wc *WakatimeCategory) String() string {
	var builder strings.Builder
	builder.WriteString("WakatimeCategory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", wc.ID))
	builder.WriteString("create_time=")
	builder.WriteString(wc.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(wc.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("wakatime_id=")
	builder.WriteString(fmt.Sprintf("%v", wc.WakatimeID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", wc.UserID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(wc.Name)
	builder.WriteString(", ")
	builder.WriteString("total_seconds=")
	builder.WriteString(fmt.Sprintf("%v", wc.TotalSeconds))
	builder.WriteByte(')')
	return builder.String()
}

// WakatimeCategories is a parsable slice of WakatimeCategory.
type WakatimeCategories []*WakatimeCategory
