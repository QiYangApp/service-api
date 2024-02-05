// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"service-api/src/models/ent/wakatimeprojectduration"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// WakatimeProjectDuration is the model entity for the WakatimeProjectDuration schema.
type WakatimeProjectDuration struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WakatimeProjectDuration) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case wakatimeprojectduration.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WakatimeProjectDuration fields.
func (wpd *WakatimeProjectDuration) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case wakatimeprojectduration.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wpd.ID = int(value.Int64)
		default:
			wpd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WakatimeProjectDuration.
// This includes values selected through modifiers, order, etc.
func (wpd *WakatimeProjectDuration) Value(name string) (ent.Value, error) {
	return wpd.selectValues.Get(name)
}

// Update returns a builder for updating this WakatimeProjectDuration.
// Note that you need to call WakatimeProjectDuration.Unwrap() before calling this method if this WakatimeProjectDuration
// was returned from a transaction, and the transaction was committed or rolled back.
func (wpd *WakatimeProjectDuration) Update() *WakatimeProjectDurationUpdateOne {
	return NewWakatimeProjectDurationClient(wpd.config).UpdateOne(wpd)
}

// Unwrap unwraps the WakatimeProjectDuration entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wpd *WakatimeProjectDuration) Unwrap() *WakatimeProjectDuration {
	_tx, ok := wpd.config.driver.(*txDriver)
	if !ok {
		panic("ent: WakatimeProjectDuration is not a transactional entity")
	}
	wpd.config.driver = _tx.drv
	return wpd
}

// String implements the fmt.Stringer.
func (wpd *WakatimeProjectDuration) String() string {
	var builder strings.Builder
	builder.WriteString("WakatimeProjectDuration(")
	builder.WriteString(fmt.Sprintf("id=%v", wpd.ID))
	builder.WriteByte(')')
	return builder.String()
}

// WakatimeProjectDurations is a parsable slice of WakatimeProjectDuration.
type WakatimeProjectDurations []*WakatimeProjectDuration
