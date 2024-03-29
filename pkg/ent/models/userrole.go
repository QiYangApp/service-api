// Code generated by ent, DO NOT EDIT.

package models

import (
	"ent/enums/state"
	"ent/models/userrole"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserRole is the model entity for the UserRole schema.
type UserRole struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// 规则名称
	RoleName string `json:"role_name,omitempty"`
	// State holds the value of the "state" field.
	State state.SwitchState `json:"state,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime   time.Time `json:"update_time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserRole) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userrole.FieldID, userrole.FieldState:
			values[i] = new(sql.NullInt64)
		case userrole.FieldRoleName:
			values[i] = new(sql.NullString)
		case userrole.FieldCreateTime, userrole.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserRole fields.
func (ur *UserRole) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userrole.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ur.ID = int64(value.Int64)
		case userrole.FieldRoleName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role_name", values[i])
			} else if value.Valid {
				ur.RoleName = value.String
			}
		case userrole.FieldState:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				ur.State = state.SwitchState(value.Int64)
			}
		case userrole.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ur.CreateTime = value.Time
			}
		case userrole.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ur.UpdateTime = value.Time
			}
		default:
			ur.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserRole.
// This includes values selected through modifiers, order, etc.
func (ur *UserRole) Value(name string) (ent.Value, error) {
	return ur.selectValues.Get(name)
}

// Update returns a builder for updating this UserRole.
// Note that you need to call UserRole.Unwrap() before calling this method if this UserRole
// was returned from a transaction, and the transaction was committed or rolled back.
func (ur *UserRole) Update() *UserRoleUpdateOne {
	return NewUserRoleClient(ur.config).UpdateOne(ur)
}

// Unwrap unwraps the UserRole entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ur *UserRole) Unwrap() *UserRole {
	_tx, ok := ur.config.driver.(*txDriver)
	if !ok {
		panic("models: UserRole is not a transactional entity")
	}
	ur.config.driver = _tx.drv
	return ur
}

// String implements the fmt.Stringer.
func (ur *UserRole) String() string {
	var builder strings.Builder
	builder.WriteString("UserRole(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ur.ID))
	builder.WriteString("role_name=")
	builder.WriteString(ur.RoleName)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", ur.State))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(ur.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(ur.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UserRoles is a parsable slice of UserRole.
type UserRoles []*UserRole
