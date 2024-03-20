// Code generated by ent, DO NOT EDIT.

package models

import (
	"ent/models/memberrolerelatedpermission"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// MemberRoleRelatedPermission is the model entity for the MemberRoleRelatedPermission schema.
type MemberRoleRelatedPermission struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 规则id
	RoleID uuid.UUID `json:"role_id,omitempty"`
	// 权限分组id
	PermissionGroupID uuid.UUID `json:"permission_group_id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime   time.Time `json:"update_time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MemberRoleRelatedPermission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case memberrolerelatedpermission.FieldID:
			values[i] = new(sql.NullInt64)
		case memberrolerelatedpermission.FieldCreateTime, memberrolerelatedpermission.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case memberrolerelatedpermission.FieldRoleID, memberrolerelatedpermission.FieldPermissionGroupID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MemberRoleRelatedPermission fields.
func (mrrp *MemberRoleRelatedPermission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case memberrolerelatedpermission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mrrp.ID = int(value.Int64)
		case memberrolerelatedpermission.FieldRoleID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value != nil {
				mrrp.RoleID = *value
			}
		case memberrolerelatedpermission.FieldPermissionGroupID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field permission_group_id", values[i])
			} else if value != nil {
				mrrp.PermissionGroupID = *value
			}
		case memberrolerelatedpermission.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				mrrp.CreateTime = value.Time
			}
		case memberrolerelatedpermission.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				mrrp.UpdateTime = value.Time
			}
		default:
			mrrp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MemberRoleRelatedPermission.
// This includes values selected through modifiers, order, etc.
func (mrrp *MemberRoleRelatedPermission) Value(name string) (ent.Value, error) {
	return mrrp.selectValues.Get(name)
}

// Update returns a builder for updating this MemberRoleRelatedPermission.
// Note that you need to call MemberRoleRelatedPermission.Unwrap() before calling this method if this MemberRoleRelatedPermission
// was returned from a transaction, and the transaction was committed or rolled back.
func (mrrp *MemberRoleRelatedPermission) Update() *MemberRoleRelatedPermissionUpdateOne {
	return NewMemberRoleRelatedPermissionClient(mrrp.config).UpdateOne(mrrp)
}

// Unwrap unwraps the MemberRoleRelatedPermission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mrrp *MemberRoleRelatedPermission) Unwrap() *MemberRoleRelatedPermission {
	_tx, ok := mrrp.config.driver.(*txDriver)
	if !ok {
		panic("models: MemberRoleRelatedPermission is not a transactional entity")
	}
	mrrp.config.driver = _tx.drv
	return mrrp
}

// String implements the fmt.Stringer.
func (mrrp *MemberRoleRelatedPermission) String() string {
	var builder strings.Builder
	builder.WriteString("MemberRoleRelatedPermission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mrrp.ID))
	builder.WriteString("role_id=")
	builder.WriteString(fmt.Sprintf("%v", mrrp.RoleID))
	builder.WriteString(", ")
	builder.WriteString("permission_group_id=")
	builder.WriteString(fmt.Sprintf("%v", mrrp.PermissionGroupID))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(mrrp.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(mrrp.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// MemberRoleRelatedPermissions is a parsable slice of MemberRoleRelatedPermission.
type MemberRoleRelatedPermissions []*MemberRoleRelatedPermission
