// Code generated by ent, DO NOT EDIT.

package permissiongroup

import (
	"ent/enums/state"
	"frame/util/timeutil"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the permissiongroup type in the database.
	Label = "permission_group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPermissionName holds the string denoting the permission_name field in the database.
	FieldPermissionName = "permission_name"
	// FieldIoc holds the string denoting the ioc field in the database.
	FieldIoc = "ioc"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldLeft holds the string denoting the left field in the database.
	FieldLeft = "left"
	// FieldRight holds the string denoting the right field in the database.
	FieldRight = "right"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// Table holds the table name of the permissiongroup in the database.
	Table = "permission_groups"
)

// Columns holds all SQL columns for permissiongroup fields.
var Columns = []string{
	FieldID,
	FieldPermissionName,
	FieldIoc,
	FieldSort,
	FieldLeft,
	FieldRight,
	FieldState,
	FieldCreateTime,
	FieldUpdateTime,
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
	// DefaultPermissionName holds the default value on creation for the "permission_name" field.
	DefaultPermissionName string
	// PermissionNameValidator is a validator for the "permission_name" field. It is called by the builders before save.
	PermissionNameValidator func(string) error
	// DefaultIoc holds the default value on creation for the "ioc" field.
	DefaultIoc string
	// IocValidator is a validator for the "ioc" field. It is called by the builders before save.
	IocValidator func(string) error
	// DefaultSort holds the default value on creation for the "sort" field.
	DefaultSort int32
	// DefaultLeft holds the default value on creation for the "left" field.
	DefaultLeft int32
	// DefaultRight holds the default value on creation for the "right" field.
	DefaultRight int32
	// DefaultState holds the default value on creation for the "state" field.
	DefaultState state.SwitchState
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime timeutil.TimeStamp
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() timeutil.TimeStamp
)

// OrderOption defines the ordering options for the PermissionGroup queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPermissionName orders the results by the permission_name field.
func ByPermissionName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPermissionName, opts...).ToFunc()
}

// ByIoc orders the results by the ioc field.
func ByIoc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIoc, opts...).ToFunc()
}

// BySort orders the results by the sort field.
func BySort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSort, opts...).ToFunc()
}

// ByLeft orders the results by the left field.
func ByLeft(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLeft, opts...).ToFunc()
}

// ByRight orders the results by the right field.
func ByRight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRight, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}
