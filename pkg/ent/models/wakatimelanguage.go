// Code generated by ent, DO NOT EDIT.

package models

import (
	"ent/models/wakatimelanguage"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// WakatimeLanguage is the model entity for the WakatimeLanguage schema.
type WakatimeLanguage struct {
	config
	// ID of the ent.
	ID           int64 `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WakatimeLanguage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case wakatimelanguage.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WakatimeLanguage fields.
func (wl *WakatimeLanguage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case wakatimelanguage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wl.ID = int64(value.Int64)
		default:
			wl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WakatimeLanguage.
// This includes values selected through modifiers, order, etc.
func (wl *WakatimeLanguage) Value(name string) (ent.Value, error) {
	return wl.selectValues.Get(name)
}

// Update returns a builder for updating this WakatimeLanguage.
// Note that you need to call WakatimeLanguage.Unwrap() before calling this method if this WakatimeLanguage
// was returned from a transaction, and the transaction was committed or rolled back.
func (wl *WakatimeLanguage) Update() *WakatimeLanguageUpdateOne {
	return NewWakatimeLanguageClient(wl.config).UpdateOne(wl)
}

// Unwrap unwraps the WakatimeLanguage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wl *WakatimeLanguage) Unwrap() *WakatimeLanguage {
	_tx, ok := wl.config.driver.(*txDriver)
	if !ok {
		panic("models: WakatimeLanguage is not a transactional entity")
	}
	wl.config.driver = _tx.drv
	return wl
}

// String implements the fmt.Stringer.
func (wl *WakatimeLanguage) String() string {
	var builder strings.Builder
	builder.WriteString("WakatimeLanguage(")
	builder.WriteString(fmt.Sprintf("id=%v", wl.ID))
	builder.WriteByte(')')
	return builder.String()
}

// WakatimeLanguages is a parsable slice of WakatimeLanguage.
type WakatimeLanguages []*WakatimeLanguage
