// Code generated by ent, DO NOT EDIT.

package models

import (
	"ent/models/permissionrelatedrouter"
	"ent/utils/timeutil"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// PermissionRelatedRouter is the model entity for the PermissionRelatedRouter schema.
type PermissionRelatedRouter struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// 路由id
	RouterID uuid.UUID `json:"router_id,omitempty"`
	// 权限分组
	PermissionGroupID uuid.UUID `json:"permission_group_id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime timeutil.TimeStamp `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime   timeutil.TimeStamp `json:"update_time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PermissionRelatedRouter) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case permissionrelatedrouter.FieldID, permissionrelatedrouter.FieldCreateTime, permissionrelatedrouter.FieldUpdateTime:
			values[i] = new(sql.NullInt64)
		case permissionrelatedrouter.FieldRouterID, permissionrelatedrouter.FieldPermissionGroupID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PermissionRelatedRouter fields.
func (prr *PermissionRelatedRouter) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case permissionrelatedrouter.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			prr.ID = int64(value.Int64)
		case permissionrelatedrouter.FieldRouterID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field router_id", values[i])
			} else if value != nil {
				prr.RouterID = *value
			}
		case permissionrelatedrouter.FieldPermissionGroupID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field permission_group_id", values[i])
			} else if value != nil {
				prr.PermissionGroupID = *value
			}
		case permissionrelatedrouter.FieldCreateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				prr.CreateTime = timeutil.TimeStamp(value.Int64)
			}
		case permissionrelatedrouter.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				prr.UpdateTime = timeutil.TimeStamp(value.Int64)
			}
		default:
			prr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PermissionRelatedRouter.
// This includes values selected through modifiers, order, etc.
func (prr *PermissionRelatedRouter) Value(name string) (ent.Value, error) {
	return prr.selectValues.Get(name)
}

// Update returns a builder for updating this PermissionRelatedRouter.
// Note that you need to call PermissionRelatedRouter.Unwrap() before calling this method if this PermissionRelatedRouter
// was returned from a transaction, and the transaction was committed or rolled back.
func (prr *PermissionRelatedRouter) Update() *PermissionRelatedRouterUpdateOne {
	return NewPermissionRelatedRouterClient(prr.config).UpdateOne(prr)
}

// Unwrap unwraps the PermissionRelatedRouter entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (prr *PermissionRelatedRouter) Unwrap() *PermissionRelatedRouter {
	_tx, ok := prr.config.driver.(*txDriver)
	if !ok {
		panic("models: PermissionRelatedRouter is not a transactional entity")
	}
	prr.config.driver = _tx.drv
	return prr
}

// String implements the fmt.Stringer.
func (prr *PermissionRelatedRouter) String() string {
	var builder strings.Builder
	builder.WriteString("PermissionRelatedRouter(")
	builder.WriteString(fmt.Sprintf("id=%v, ", prr.ID))
	builder.WriteString("router_id=")
	builder.WriteString(fmt.Sprintf("%v", prr.RouterID))
	builder.WriteString(", ")
	builder.WriteString("permission_group_id=")
	builder.WriteString(fmt.Sprintf("%v", prr.PermissionGroupID))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(fmt.Sprintf("%v", prr.CreateTime))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(fmt.Sprintf("%v", prr.UpdateTime))
	builder.WriteByte(')')
	return builder.String()
}

// PermissionRelatedRouters is a parsable slice of PermissionRelatedRouter.
type PermissionRelatedRouters []*PermissionRelatedRouter