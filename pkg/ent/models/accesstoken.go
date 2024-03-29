// Code generated by ent, DO NOT EDIT.

package models

import (
	"ent/models/accesstoken"
	"ent/utils/timeutil"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// AccessToken is the model entity for the AccessToken schema.
type AccessToken struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"token,omitempty"`
	// TokenHash holds the value of the "token_hash" field.
	TokenHash string `json:"token_hash,omitempty"`
	// TokenSalt holds the value of the "token_salt" field.
	TokenSalt string `json:"token_salt,omitempty"`
	// TokenLastEight holds the value of the "token_last_eight" field.
	TokenLastEight string `json:"token_last_eight,omitempty"`
	// Scope holds the value of the "scope" field.
	Scope string `json:"scope,omitempty"`
	// HasRecentActivity holds the value of the "has_recent_activity" field.
	HasRecentActivity string `json:"has_recent_activity,omitempty"`
	// HasUsed holds the value of the "has_used" field.
	HasUsed string `json:"has_used,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime timeutil.TimeStamp `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime   timeutil.TimeStamp `json:"update_time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AccessToken) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case accesstoken.FieldID, accesstoken.FieldUserID, accesstoken.FieldCreateTime, accesstoken.FieldUpdateTime:
			values[i] = new(sql.NullInt64)
		case accesstoken.FieldName, accesstoken.FieldToken, accesstoken.FieldTokenHash, accesstoken.FieldTokenSalt, accesstoken.FieldTokenLastEight, accesstoken.FieldScope, accesstoken.FieldHasRecentActivity, accesstoken.FieldHasUsed:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AccessToken fields.
func (at *AccessToken) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case accesstoken.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			at.ID = int64(value.Int64)
		case accesstoken.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				at.UserID = value.Int64
			}
		case accesstoken.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				at.Name = value.String
			}
		case accesstoken.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				at.Token = value.String
			}
		case accesstoken.FieldTokenHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token_hash", values[i])
			} else if value.Valid {
				at.TokenHash = value.String
			}
		case accesstoken.FieldTokenSalt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token_salt", values[i])
			} else if value.Valid {
				at.TokenSalt = value.String
			}
		case accesstoken.FieldTokenLastEight:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token_last_eight", values[i])
			} else if value.Valid {
				at.TokenLastEight = value.String
			}
		case accesstoken.FieldScope:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field scope", values[i])
			} else if value.Valid {
				at.Scope = value.String
			}
		case accesstoken.FieldHasRecentActivity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field has_recent_activity", values[i])
			} else if value.Valid {
				at.HasRecentActivity = value.String
			}
		case accesstoken.FieldHasUsed:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field has_used", values[i])
			} else if value.Valid {
				at.HasUsed = value.String
			}
		case accesstoken.FieldCreateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				at.CreateTime = timeutil.TimeStamp(value.Int64)
			}
		case accesstoken.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				at.UpdateTime = timeutil.TimeStamp(value.Int64)
			}
		default:
			at.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AccessToken.
// This includes values selected through modifiers, order, etc.
func (at *AccessToken) Value(name string) (ent.Value, error) {
	return at.selectValues.Get(name)
}

// Update returns a builder for updating this AccessToken.
// Note that you need to call AccessToken.Unwrap() before calling this method if this AccessToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (at *AccessToken) Update() *AccessTokenUpdateOne {
	return NewAccessTokenClient(at.config).UpdateOne(at)
}

// Unwrap unwraps the AccessToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (at *AccessToken) Unwrap() *AccessToken {
	_tx, ok := at.config.driver.(*txDriver)
	if !ok {
		panic("models: AccessToken is not a transactional entity")
	}
	at.config.driver = _tx.drv
	return at
}

// String implements the fmt.Stringer.
func (at *AccessToken) String() string {
	var builder strings.Builder
	builder.WriteString("AccessToken(")
	builder.WriteString(fmt.Sprintf("id=%v, ", at.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", at.UserID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(at.Name)
	builder.WriteString(", ")
	builder.WriteString("token=")
	builder.WriteString(at.Token)
	builder.WriteString(", ")
	builder.WriteString("token_hash=")
	builder.WriteString(at.TokenHash)
	builder.WriteString(", ")
	builder.WriteString("token_salt=")
	builder.WriteString(at.TokenSalt)
	builder.WriteString(", ")
	builder.WriteString("token_last_eight=")
	builder.WriteString(at.TokenLastEight)
	builder.WriteString(", ")
	builder.WriteString("scope=")
	builder.WriteString(at.Scope)
	builder.WriteString(", ")
	builder.WriteString("has_recent_activity=")
	builder.WriteString(at.HasRecentActivity)
	builder.WriteString(", ")
	builder.WriteString("has_used=")
	builder.WriteString(at.HasUsed)
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(fmt.Sprintf("%v", at.CreateTime))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(fmt.Sprintf("%v", at.UpdateTime))
	builder.WriteByte(')')
	return builder.String()
}

// AccessTokens is a parsable slice of AccessToken.
type AccessTokens []*AccessToken
