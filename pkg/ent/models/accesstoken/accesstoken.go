// Code generated by ent, DO NOT EDIT.

package accesstoken

import (
	"ent/utils/timeutil"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the accesstoken type in the database.
	Label = "access_token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldTokenHash holds the string denoting the token_hash field in the database.
	FieldTokenHash = "token_hash"
	// FieldTokenSalt holds the string denoting the token_salt field in the database.
	FieldTokenSalt = "token_salt"
	// FieldTokenLastEight holds the string denoting the token_last_eight field in the database.
	FieldTokenLastEight = "token_last_eight"
	// FieldScope holds the string denoting the scope field in the database.
	FieldScope = "scope"
	// FieldHasRecentActivity holds the string denoting the has_recent_activity field in the database.
	FieldHasRecentActivity = "has_recent_activity"
	// FieldHasUsed holds the string denoting the has_used field in the database.
	FieldHasUsed = "has_used"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// Table holds the table name of the accesstoken in the database.
	Table = "access_tokens"
)

// Columns holds all SQL columns for accesstoken fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldName,
	FieldToken,
	FieldTokenHash,
	FieldTokenSalt,
	FieldTokenLastEight,
	FieldScope,
	FieldHasRecentActivity,
	FieldHasUsed,
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

// OrderOption defines the ordering options for the AccessToken queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByToken orders the results by the token field.
func ByToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToken, opts...).ToFunc()
}

// ByTokenHash orders the results by the token_hash field.
func ByTokenHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTokenHash, opts...).ToFunc()
}

// ByTokenSalt orders the results by the token_salt field.
func ByTokenSalt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTokenSalt, opts...).ToFunc()
}

// ByTokenLastEight orders the results by the token_last_eight field.
func ByTokenLastEight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTokenLastEight, opts...).ToFunc()
}

// ByScope orders the results by the scope field.
func ByScope(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScope, opts...).ToFunc()
}

// ByHasRecentActivity orders the results by the has_recent_activity field.
func ByHasRecentActivity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasRecentActivity, opts...).ToFunc()
}

// ByHasUsed orders the results by the has_used field.
func ByHasUsed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasUsed, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}
