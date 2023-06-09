// Code generated by ent, DO NOT EDIT.

package router

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the router type in the database.
	Label = "router"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldRouteName holds the string denoting the route_name field in the database.
	FieldRouteName = "route_name"
	// FieldRoute holds the string denoting the route field in the database.
	FieldRoute = "route"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// Table holds the table name of the router in the database.
	Table = "routers"
)

// Columns holds all SQL columns for router fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldRouteName,
	FieldRoute,
	FieldDescription,
	FieldState,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultRouteName holds the default value on creation for the "route_name" field.
	DefaultRouteName string
	// RouteNameValidator is a validator for the "route_name" field. It is called by the builders before save.
	RouteNameValidator func(string) error
	// DefaultRoute holds the default value on creation for the "route" field.
	DefaultRoute string
	// RouteValidator is a validator for the "route" field. It is called by the builders before save.
	RouteValidator func(string) error
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultState holds the default value on creation for the "state" field.
	DefaultState string
	// StateValidator is a validator for the "state" field. It is called by the builders before save.
	StateValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
