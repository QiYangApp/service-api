// Code generated by ent, DO NOT EDIT.

package models

import (
	"ent/models/wakatimeprojectinfo"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// WakatimeProjectInfo is the model entity for the WakatimeProjectInfo schema.
type WakatimeProjectInfo struct {
	config
	// ID of the ent.
	ID           int64 `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WakatimeProjectInfo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case wakatimeprojectinfo.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WakatimeProjectInfo fields.
func (wpi *WakatimeProjectInfo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case wakatimeprojectinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wpi.ID = int64(value.Int64)
		default:
			wpi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WakatimeProjectInfo.
// This includes values selected through modifiers, order, etc.
func (wpi *WakatimeProjectInfo) Value(name string) (ent.Value, error) {
	return wpi.selectValues.Get(name)
}

// Update returns a builder for updating this WakatimeProjectInfo.
// Note that you need to call WakatimeProjectInfo.Unwrap() before calling this method if this WakatimeProjectInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (wpi *WakatimeProjectInfo) Update() *WakatimeProjectInfoUpdateOne {
	return NewWakatimeProjectInfoClient(wpi.config).UpdateOne(wpi)
}

// Unwrap unwraps the WakatimeProjectInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wpi *WakatimeProjectInfo) Unwrap() *WakatimeProjectInfo {
	_tx, ok := wpi.config.driver.(*txDriver)
	if !ok {
		panic("models: WakatimeProjectInfo is not a transactional entity")
	}
	wpi.config.driver = _tx.drv
	return wpi
}

// String implements the fmt.Stringer.
func (wpi *WakatimeProjectInfo) String() string {
	var builder strings.Builder
	builder.WriteString("WakatimeProjectInfo(")
	builder.WriteString(fmt.Sprintf("id=%v", wpi.ID))
	builder.WriteByte(')')
	return builder.String()
}

// WakatimeProjectInfos is a parsable slice of WakatimeProjectInfo.
type WakatimeProjectInfos []*WakatimeProjectInfo
