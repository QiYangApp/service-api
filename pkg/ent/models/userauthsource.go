// Code generated by ent, DO NOT EDIT.

package models

import (
	"ent/models/userauthsource"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserAuthSource is the model entity for the UserAuthSource schema.
type UserAuthSource struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// 授权token
	Token string `json:"token,omitempty"`
	// 登录渠道
	Channel string `json:"channel,omitempty"`
	// 登录设备
	Device string `json:"device,omitempty"`
	// 登录信息
	DeviceDetail string `json:"device_detail,omitempty"`
	// ClientIP holds the value of the "client_ip" field.
	ClientIP string `json:"client_ip,omitempty"`
	// RemoteIP holds the value of the "remote_ip" field.
	RemoteIP string `json:"remote_ip,omitempty"`
	// Snapshot holds the value of the "snapshot" field.
	Snapshot string `json:"snapshot,omitempty"`
	// LoginName holds the value of the "login_name" field.
	LoginName string `json:"login_name,omitempty"`
	// LoginSource holds the value of the "login_source" field.
	LoginSource int `json:"login_source,omitempty"`
	// LoginType holds the value of the "login_type" field.
	LoginType int `json:"login_type,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime   time.Time `json:"update_time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserAuthSource) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userauthsource.FieldID, userauthsource.FieldUserID, userauthsource.FieldLoginSource, userauthsource.FieldLoginType:
			values[i] = new(sql.NullInt64)
		case userauthsource.FieldToken, userauthsource.FieldChannel, userauthsource.FieldDevice, userauthsource.FieldDeviceDetail, userauthsource.FieldClientIP, userauthsource.FieldRemoteIP, userauthsource.FieldSnapshot, userauthsource.FieldLoginName:
			values[i] = new(sql.NullString)
		case userauthsource.FieldCreateTime, userauthsource.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserAuthSource fields.
func (uas *UserAuthSource) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userauthsource.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uas.ID = int64(value.Int64)
		case userauthsource.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				uas.UserID = value.Int64
			}
		case userauthsource.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				uas.Token = value.String
			}
		case userauthsource.FieldChannel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field channel", values[i])
			} else if value.Valid {
				uas.Channel = value.String
			}
		case userauthsource.FieldDevice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device", values[i])
			} else if value.Valid {
				uas.Device = value.String
			}
		case userauthsource.FieldDeviceDetail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device_detail", values[i])
			} else if value.Valid {
				uas.DeviceDetail = value.String
			}
		case userauthsource.FieldClientIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_ip", values[i])
			} else if value.Valid {
				uas.ClientIP = value.String
			}
		case userauthsource.FieldRemoteIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remote_ip", values[i])
			} else if value.Valid {
				uas.RemoteIP = value.String
			}
		case userauthsource.FieldSnapshot:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field snapshot", values[i])
			} else if value.Valid {
				uas.Snapshot = value.String
			}
		case userauthsource.FieldLoginName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field login_name", values[i])
			} else if value.Valid {
				uas.LoginName = value.String
			}
		case userauthsource.FieldLoginSource:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field login_source", values[i])
			} else if value.Valid {
				uas.LoginSource = int(value.Int64)
			}
		case userauthsource.FieldLoginType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field login_type", values[i])
			} else if value.Valid {
				uas.LoginType = int(value.Int64)
			}
		case userauthsource.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				uas.CreateTime = value.Time
			}
		case userauthsource.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				uas.UpdateTime = value.Time
			}
		default:
			uas.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserAuthSource.
// This includes values selected through modifiers, order, etc.
func (uas *UserAuthSource) Value(name string) (ent.Value, error) {
	return uas.selectValues.Get(name)
}

// Update returns a builder for updating this UserAuthSource.
// Note that you need to call UserAuthSource.Unwrap() before calling this method if this UserAuthSource
// was returned from a transaction, and the transaction was committed or rolled back.
func (uas *UserAuthSource) Update() *UserAuthSourceUpdateOne {
	return NewUserAuthSourceClient(uas.config).UpdateOne(uas)
}

// Unwrap unwraps the UserAuthSource entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uas *UserAuthSource) Unwrap() *UserAuthSource {
	_tx, ok := uas.config.driver.(*txDriver)
	if !ok {
		panic("models: UserAuthSource is not a transactional entity")
	}
	uas.config.driver = _tx.drv
	return uas
}

// String implements the fmt.Stringer.
func (uas *UserAuthSource) String() string {
	var builder strings.Builder
	builder.WriteString("UserAuthSource(")
	builder.WriteString(fmt.Sprintf("id=%v, ", uas.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", uas.UserID))
	builder.WriteString(", ")
	builder.WriteString("token=")
	builder.WriteString(uas.Token)
	builder.WriteString(", ")
	builder.WriteString("channel=")
	builder.WriteString(uas.Channel)
	builder.WriteString(", ")
	builder.WriteString("device=")
	builder.WriteString(uas.Device)
	builder.WriteString(", ")
	builder.WriteString("device_detail=")
	builder.WriteString(uas.DeviceDetail)
	builder.WriteString(", ")
	builder.WriteString("client_ip=")
	builder.WriteString(uas.ClientIP)
	builder.WriteString(", ")
	builder.WriteString("remote_ip=")
	builder.WriteString(uas.RemoteIP)
	builder.WriteString(", ")
	builder.WriteString("snapshot=")
	builder.WriteString(uas.Snapshot)
	builder.WriteString(", ")
	builder.WriteString("login_name=")
	builder.WriteString(uas.LoginName)
	builder.WriteString(", ")
	builder.WriteString("login_source=")
	builder.WriteString(fmt.Sprintf("%v", uas.LoginSource))
	builder.WriteString(", ")
	builder.WriteString("login_type=")
	builder.WriteString(fmt.Sprintf("%v", uas.LoginType))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(uas.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(uas.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UserAuthSources is a parsable slice of UserAuthSource.
type UserAuthSources []*UserAuthSource
