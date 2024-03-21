// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLowerName holds the string denoting the lower_name field in the database.
	FieldLowerName = "lower_name"
	// FieldFullName holds the string denoting the full_name field in the database.
	FieldFullName = "full_name"
	// FieldPasswdSalt holds the string denoting the passwd_salt field in the database.
	FieldPasswdSalt = "passwd_salt"
	// FieldPasswdHashAlgo holds the string denoting the passwd_hash_algo field in the database.
	FieldPasswdHashAlgo = "passwd_hash_algo"
	// FieldPasswd holds the string denoting the passwd field in the database.
	FieldPasswd = "passwd"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldLoginName holds the string denoting the login_name field in the database.
	FieldLoginName = "login_name"
	// FieldLoginSource holds the string denoting the login_source field in the database.
	FieldLoginSource = "login_source"
	// FieldLoginType holds the string denoting the login_type field in the database.
	FieldLoginType = "login_type"
	// FieldIsRestricted holds the string denoting the is_restricted field in the database.
	FieldIsRestricted = "is_restricted"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// FieldProhibitLogin holds the string denoting the prohibit_login field in the database.
	FieldProhibitLogin = "prohibit_login"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldAvatar,
	FieldEmail,
	FieldName,
	FieldLowerName,
	FieldFullName,
	FieldPasswdSalt,
	FieldPasswdHashAlgo,
	FieldPasswd,
	FieldLanguage,
	FieldLoginName,
	FieldLoginSource,
	FieldLoginType,
	FieldIsRestricted,
	FieldIsActive,
	FieldProhibitLogin,
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
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// LowerNameValidator is a validator for the "lower_name" field. It is called by the builders before save.
	LowerNameValidator func(string) error
	// FullNameValidator is a validator for the "full_name" field. It is called by the builders before save.
	FullNameValidator func(string) error
	// PasswdSaltValidator is a validator for the "passwd_salt" field. It is called by the builders before save.
	PasswdSaltValidator func(string) error
	// PasswdHashAlgoValidator is a validator for the "passwd_hash_algo" field. It is called by the builders before save.
	PasswdHashAlgoValidator func(string) error
	// PasswdValidator is a validator for the "passwd" field. It is called by the builders before save.
	PasswdValidator func(string) error
	// LanguageValidator is a validator for the "language" field. It is called by the builders before save.
	LanguageValidator func(string) error
	// LoginNameValidator is a validator for the "login_name" field. It is called by the builders before save.
	LoginNameValidator func(string) error
	// DefaultLoginSource holds the default value on creation for the "login_source" field.
	DefaultLoginSource int64
	// DefaultIsRestricted holds the default value on creation for the "is_restricted" field.
	DefaultIsRestricted bool
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive bool
	// DefaultProhibitLogin holds the default value on creation for the "prohibit_login" field.
	DefaultProhibitLogin bool
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByLowerName orders the results by the lower_name field.
func ByLowerName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLowerName, opts...).ToFunc()
}

// ByFullName orders the results by the full_name field.
func ByFullName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFullName, opts...).ToFunc()
}

// ByPasswdSalt orders the results by the passwd_salt field.
func ByPasswdSalt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPasswdSalt, opts...).ToFunc()
}

// ByPasswdHashAlgo orders the results by the passwd_hash_algo field.
func ByPasswdHashAlgo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPasswdHashAlgo, opts...).ToFunc()
}

// ByPasswd orders the results by the passwd field.
func ByPasswd(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPasswd, opts...).ToFunc()
}

// ByLanguage orders the results by the language field.
func ByLanguage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLanguage, opts...).ToFunc()
}

// ByLoginName orders the results by the login_name field.
func ByLoginName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLoginName, opts...).ToFunc()
}

// ByLoginSource orders the results by the login_source field.
func ByLoginSource(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLoginSource, opts...).ToFunc()
}

// ByLoginType orders the results by the login_type field.
func ByLoginType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLoginType, opts...).ToFunc()
}

// ByIsRestricted orders the results by the is_restricted field.
func ByIsRestricted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsRestricted, opts...).ToFunc()
}

// ByIsActive orders the results by the is_active field.
func ByIsActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsActive, opts...).ToFunc()
}

// ByProhibitLogin orders the results by the prohibit_login field.
func ByProhibitLogin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProhibitLogin, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}