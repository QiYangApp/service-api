// Code generated by ent, DO NOT EDIT.

package twofactor

import (
	"ent/utils/timeutil"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the twofactor type in the database.
	Label = "two_factor"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldSecret holds the string denoting the secret field in the database.
	FieldSecret = "secret"
	// FieldScratchSalt holds the string denoting the scratch_salt field in the database.
	FieldScratchSalt = "scratch_salt"
	// FieldScratchHash holds the string denoting the scratch_hash field in the database.
	FieldScratchHash = "scratch_hash"
	// FieldLastUsedPasscode holds the string denoting the last_used_passcode field in the database.
	FieldLastUsedPasscode = "last_used_passcode"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// Table holds the table name of the twofactor in the database.
	Table = "two_factors"
)

// Columns holds all SQL columns for twofactor fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldSecret,
	FieldScratchSalt,
	FieldScratchHash,
	FieldLastUsedPasscode,
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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime timeutil.TimeStamp
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() timeutil.TimeStamp
)

// OrderOption defines the ordering options for the TwoFactor queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// BySecret orders the results by the secret field.
func BySecret(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSecret, opts...).ToFunc()
}

// ByScratchSalt orders the results by the scratch_salt field.
func ByScratchSalt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScratchSalt, opts...).ToFunc()
}

// ByScratchHash orders the results by the scratch_hash field.
func ByScratchHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScratchHash, opts...).ToFunc()
}

// ByLastUsedPasscode orders the results by the last_used_passcode field.
func ByLastUsedPasscode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastUsedPasscode, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}