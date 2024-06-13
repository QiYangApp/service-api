// Code generated by ent, DO NOT EDIT.

package models

import (
	"ent/models/wakatimeentity"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// WakatimeEntity is the model entity for the WakatimeEntity schema.
type WakatimeEntity struct {
	config
	// ID of the ent.
	ID           int64 `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WakatimeEntity) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case wakatimeentity.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WakatimeEntity fields.
func (we *WakatimeEntity) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case wakatimeentity.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			we.ID = int64(value.Int64)
		default:
			we.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WakatimeEntity.
// This includes values selected through modifiers, order, etc.
func (we *WakatimeEntity) Value(name string) (ent.Value, error) {
	return we.selectValues.Get(name)
}

// Update returns a builder for updating this WakatimeEntity.
// Note that you need to call WakatimeEntity.Unwrap() before calling this method if this WakatimeEntity
// was returned from a transaction, and the transaction was committed or rolled back.
func (we *WakatimeEntity) Update() *WakatimeEntityUpdateOne {
	return NewWakatimeEntityClient(we.config).UpdateOne(we)
}

// Unwrap unwraps the WakatimeEntity entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (we *WakatimeEntity) Unwrap() *WakatimeEntity {
	_tx, ok := we.config.driver.(*txDriver)
	if !ok {
		panic("models: WakatimeEntity is not a transactional entity")
	}
	we.config.driver = _tx.drv
	return we
}

// String implements the fmt.Stringer.
func (we *WakatimeEntity) String() string {
	var builder strings.Builder
	builder.WriteString("WakatimeEntity(")
	builder.WriteString(fmt.Sprintf("id=%v", we.ID))
	builder.WriteByte(')')
	return builder.String()
}

// WakatimeEntities is a parsable slice of WakatimeEntity.
type WakatimeEntities []*WakatimeEntity
