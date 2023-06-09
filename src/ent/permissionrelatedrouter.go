// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"service-api/src/ent/permissionrelatedrouter"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// PermissionRelatedRouter is the model entity for the PermissionRelatedRouter schema.
type PermissionRelatedRouter struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// 路由id
	RouterID uuid.UUID `json:"router_id,omitempty"`
	// 权限分组
	PermissionGroupID uuid.UUID `json:"permission_group_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PermissionRelatedRouter) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case permissionrelatedrouter.FieldCreateTime, permissionrelatedrouter.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case permissionrelatedrouter.FieldID, permissionrelatedrouter.FieldRouterID, permissionrelatedrouter.FieldPermissionGroupID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PermissionRelatedRouter", columns[i])
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
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				prr.ID = *value
			}
		case permissionrelatedrouter.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				prr.CreateTime = value.Time
			}
		case permissionrelatedrouter.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				prr.UpdateTime = value.Time
			}
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
		}
	}
	return nil
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
		panic("ent: PermissionRelatedRouter is not a transactional entity")
	}
	prr.config.driver = _tx.drv
	return prr
}

// String implements the fmt.Stringer.
func (prr *PermissionRelatedRouter) String() string {
	var builder strings.Builder
	builder.WriteString("PermissionRelatedRouter(")
	builder.WriteString(fmt.Sprintf("id=%v, ", prr.ID))
	builder.WriteString("create_time=")
	builder.WriteString(prr.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(prr.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("router_id=")
	builder.WriteString(fmt.Sprintf("%v", prr.RouterID))
	builder.WriteString(", ")
	builder.WriteString("permission_group_id=")
	builder.WriteString(fmt.Sprintf("%v", prr.PermissionGroupID))
	builder.WriteByte(')')
	return builder.String()
}

// PermissionRelatedRouters is a parsable slice of PermissionRelatedRouter.
type PermissionRelatedRouters []*PermissionRelatedRouter
