// Code generated by ent, DO NOT EDIT.

package accounts

import (
	"ent/utils/timeutil"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the accounts type in the database.
	Label = "accounts"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldAccount holds the string denoting the account field in the database.
	FieldAccount = "account"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// FieldIsPrivate holds the string denoting the is_private field in the database.
	FieldIsPrivate = "is_private"
	// FieldIsActivated holds the string denoting the is_activated field in the database.
	FieldIsActivated = "is_activated"
	// FieldIsPrimary holds the string denoting the is_primary field in the database.
	FieldIsPrimary = "is_primary"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// Table holds the table name of the accounts in the database.
	Table = "accounts"
)

// Columns holds all SQL columns for accounts fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldAccount,
	FieldType,
	FieldDesc,
	FieldIsPrivate,
	FieldIsActivated,
	FieldIsPrimary,
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
	// DefaultIsPrivate holds the default value on creation for the "is_private" field.
	DefaultIsPrivate bool
	// DefaultIsPrimary holds the default value on creation for the "is_primary" field.
	DefaultIsPrimary bool
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime timeutil.TimeStamp
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() timeutil.TimeStamp
)

// OrderOption defines the ordering options for the Accounts queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByAccount orders the results by the account field.
func ByAccount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccount, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByDesc orders the results by the desc field.
func ByDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesc, opts...).ToFunc()
}

// ByIsPrivate orders the results by the is_private field.
func ByIsPrivate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsPrivate, opts...).ToFunc()
}

// ByIsActivated orders the results by the is_activated field.
func ByIsActivated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsActivated, opts...).ToFunc()
}

// ByIsPrimary orders the results by the is_primary field.
func ByIsPrimary(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsPrimary, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}