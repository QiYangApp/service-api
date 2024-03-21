// Code generated by ent, DO NOT EDIT.

package sourcedata

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the sourcedata type in the database.
	Label = "source_data"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldSubType holds the string denoting the sub_type field in the database.
	FieldSubType = "sub_type"
	// FieldInfo holds the string denoting the info field in the database.
	FieldInfo = "info"
	// FieldSnapshot holds the string denoting the snapshot field in the database.
	FieldSnapshot = "snapshot"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// Table holds the table name of the sourcedata in the database.
	Table = "source_data"
)

// Columns holds all SQL columns for sourcedata fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldType,
	FieldSubType,
	FieldInfo,
	FieldSnapshot,
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
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType string
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// DefaultSubType holds the default value on creation for the "sub_type" field.
	DefaultSubType string
	// SubTypeValidator is a validator for the "sub_type" field. It is called by the builders before save.
	SubTypeValidator func(string) error
	// DefaultInfo holds the default value on creation for the "info" field.
	DefaultInfo string
	// DefaultSnapshot holds the default value on creation for the "snapshot" field.
	DefaultSnapshot string
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)

// OrderOption defines the ordering options for the SourceData queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// BySubType orders the results by the sub_type field.
func BySubType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubType, opts...).ToFunc()
}

// ByInfo orders the results by the info field.
func ByInfo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInfo, opts...).ToFunc()
}

// BySnapshot orders the results by the snapshot field.
func BySnapshot(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSnapshot, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}