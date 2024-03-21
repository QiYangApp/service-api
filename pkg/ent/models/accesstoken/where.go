// Code generated by ent, DO NOT EDIT.

package accesstoken

import (
	"ent/models/predicate"
	"ent/utils/timeutil"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldUserID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldName, v))
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldToken, v))
}

// TokenHash applies equality check predicate on the "token_hash" field. It's identical to TokenHashEQ.
func TokenHash(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldTokenHash, v))
}

// TokenSalt applies equality check predicate on the "token_salt" field. It's identical to TokenSaltEQ.
func TokenSalt(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldTokenSalt, v))
}

// TokenLastEight applies equality check predicate on the "token_last_eight" field. It's identical to TokenLastEightEQ.
func TokenLastEight(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldTokenLastEight, v))
}

// Scope applies equality check predicate on the "scope" field. It's identical to ScopeEQ.
func Scope(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldScope, v))
}

// HasRecentActivity applies equality check predicate on the "has_recent_activity" field. It's identical to HasRecentActivityEQ.
func HasRecentActivity(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldHasRecentActivity, v))
}

// HasUsed applies equality check predicate on the "has_used" field. It's identical to HasUsedEQ.
func HasUsed(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldHasUsed, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v timeutil.TimeStamp) predicate.AccessToken {
	vc := int64(v)
	return predicate.AccessToken(sql.FieldEQ(FieldCreateTime, vc))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v timeutil.TimeStamp) predicate.AccessToken {
	vc := int64(v)
	return predicate.AccessToken(sql.FieldEQ(FieldUpdateTime, vc))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldUserID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContainsFold(FieldName, v))
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldToken, v))
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldToken, v))
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldToken, vs...))
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldToken, vs...))
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldToken, v))
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldToken, v))
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldToken, v))
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldToken, v))
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContains(FieldToken, v))
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasPrefix(FieldToken, v))
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasSuffix(FieldToken, v))
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEqualFold(FieldToken, v))
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContainsFold(FieldToken, v))
}

// TokenHashEQ applies the EQ predicate on the "token_hash" field.
func TokenHashEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldTokenHash, v))
}

// TokenHashNEQ applies the NEQ predicate on the "token_hash" field.
func TokenHashNEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldTokenHash, v))
}

// TokenHashIn applies the In predicate on the "token_hash" field.
func TokenHashIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldTokenHash, vs...))
}

// TokenHashNotIn applies the NotIn predicate on the "token_hash" field.
func TokenHashNotIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldTokenHash, vs...))
}

// TokenHashGT applies the GT predicate on the "token_hash" field.
func TokenHashGT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldTokenHash, v))
}

// TokenHashGTE applies the GTE predicate on the "token_hash" field.
func TokenHashGTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldTokenHash, v))
}

// TokenHashLT applies the LT predicate on the "token_hash" field.
func TokenHashLT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldTokenHash, v))
}

// TokenHashLTE applies the LTE predicate on the "token_hash" field.
func TokenHashLTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldTokenHash, v))
}

// TokenHashContains applies the Contains predicate on the "token_hash" field.
func TokenHashContains(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContains(FieldTokenHash, v))
}

// TokenHashHasPrefix applies the HasPrefix predicate on the "token_hash" field.
func TokenHashHasPrefix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasPrefix(FieldTokenHash, v))
}

// TokenHashHasSuffix applies the HasSuffix predicate on the "token_hash" field.
func TokenHashHasSuffix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasSuffix(FieldTokenHash, v))
}

// TokenHashEqualFold applies the EqualFold predicate on the "token_hash" field.
func TokenHashEqualFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEqualFold(FieldTokenHash, v))
}

// TokenHashContainsFold applies the ContainsFold predicate on the "token_hash" field.
func TokenHashContainsFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContainsFold(FieldTokenHash, v))
}

// TokenSaltEQ applies the EQ predicate on the "token_salt" field.
func TokenSaltEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldTokenSalt, v))
}

// TokenSaltNEQ applies the NEQ predicate on the "token_salt" field.
func TokenSaltNEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldTokenSalt, v))
}

// TokenSaltIn applies the In predicate on the "token_salt" field.
func TokenSaltIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldTokenSalt, vs...))
}

// TokenSaltNotIn applies the NotIn predicate on the "token_salt" field.
func TokenSaltNotIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldTokenSalt, vs...))
}

// TokenSaltGT applies the GT predicate on the "token_salt" field.
func TokenSaltGT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldTokenSalt, v))
}

// TokenSaltGTE applies the GTE predicate on the "token_salt" field.
func TokenSaltGTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldTokenSalt, v))
}

// TokenSaltLT applies the LT predicate on the "token_salt" field.
func TokenSaltLT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldTokenSalt, v))
}

// TokenSaltLTE applies the LTE predicate on the "token_salt" field.
func TokenSaltLTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldTokenSalt, v))
}

// TokenSaltContains applies the Contains predicate on the "token_salt" field.
func TokenSaltContains(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContains(FieldTokenSalt, v))
}

// TokenSaltHasPrefix applies the HasPrefix predicate on the "token_salt" field.
func TokenSaltHasPrefix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasPrefix(FieldTokenSalt, v))
}

// TokenSaltHasSuffix applies the HasSuffix predicate on the "token_salt" field.
func TokenSaltHasSuffix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasSuffix(FieldTokenSalt, v))
}

// TokenSaltEqualFold applies the EqualFold predicate on the "token_salt" field.
func TokenSaltEqualFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEqualFold(FieldTokenSalt, v))
}

// TokenSaltContainsFold applies the ContainsFold predicate on the "token_salt" field.
func TokenSaltContainsFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContainsFold(FieldTokenSalt, v))
}

// TokenLastEightEQ applies the EQ predicate on the "token_last_eight" field.
func TokenLastEightEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldTokenLastEight, v))
}

// TokenLastEightNEQ applies the NEQ predicate on the "token_last_eight" field.
func TokenLastEightNEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldTokenLastEight, v))
}

// TokenLastEightIn applies the In predicate on the "token_last_eight" field.
func TokenLastEightIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldTokenLastEight, vs...))
}

// TokenLastEightNotIn applies the NotIn predicate on the "token_last_eight" field.
func TokenLastEightNotIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldTokenLastEight, vs...))
}

// TokenLastEightGT applies the GT predicate on the "token_last_eight" field.
func TokenLastEightGT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldTokenLastEight, v))
}

// TokenLastEightGTE applies the GTE predicate on the "token_last_eight" field.
func TokenLastEightGTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldTokenLastEight, v))
}

// TokenLastEightLT applies the LT predicate on the "token_last_eight" field.
func TokenLastEightLT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldTokenLastEight, v))
}

// TokenLastEightLTE applies the LTE predicate on the "token_last_eight" field.
func TokenLastEightLTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldTokenLastEight, v))
}

// TokenLastEightContains applies the Contains predicate on the "token_last_eight" field.
func TokenLastEightContains(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContains(FieldTokenLastEight, v))
}

// TokenLastEightHasPrefix applies the HasPrefix predicate on the "token_last_eight" field.
func TokenLastEightHasPrefix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasPrefix(FieldTokenLastEight, v))
}

// TokenLastEightHasSuffix applies the HasSuffix predicate on the "token_last_eight" field.
func TokenLastEightHasSuffix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasSuffix(FieldTokenLastEight, v))
}

// TokenLastEightEqualFold applies the EqualFold predicate on the "token_last_eight" field.
func TokenLastEightEqualFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEqualFold(FieldTokenLastEight, v))
}

// TokenLastEightContainsFold applies the ContainsFold predicate on the "token_last_eight" field.
func TokenLastEightContainsFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContainsFold(FieldTokenLastEight, v))
}

// ScopeEQ applies the EQ predicate on the "scope" field.
func ScopeEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldScope, v))
}

// ScopeNEQ applies the NEQ predicate on the "scope" field.
func ScopeNEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldScope, v))
}

// ScopeIn applies the In predicate on the "scope" field.
func ScopeIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldScope, vs...))
}

// ScopeNotIn applies the NotIn predicate on the "scope" field.
func ScopeNotIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldScope, vs...))
}

// ScopeGT applies the GT predicate on the "scope" field.
func ScopeGT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldScope, v))
}

// ScopeGTE applies the GTE predicate on the "scope" field.
func ScopeGTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldScope, v))
}

// ScopeLT applies the LT predicate on the "scope" field.
func ScopeLT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldScope, v))
}

// ScopeLTE applies the LTE predicate on the "scope" field.
func ScopeLTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldScope, v))
}

// ScopeContains applies the Contains predicate on the "scope" field.
func ScopeContains(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContains(FieldScope, v))
}

// ScopeHasPrefix applies the HasPrefix predicate on the "scope" field.
func ScopeHasPrefix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasPrefix(FieldScope, v))
}

// ScopeHasSuffix applies the HasSuffix predicate on the "scope" field.
func ScopeHasSuffix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasSuffix(FieldScope, v))
}

// ScopeEqualFold applies the EqualFold predicate on the "scope" field.
func ScopeEqualFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEqualFold(FieldScope, v))
}

// ScopeContainsFold applies the ContainsFold predicate on the "scope" field.
func ScopeContainsFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContainsFold(FieldScope, v))
}

// HasRecentActivityEQ applies the EQ predicate on the "has_recent_activity" field.
func HasRecentActivityEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldHasRecentActivity, v))
}

// HasRecentActivityNEQ applies the NEQ predicate on the "has_recent_activity" field.
func HasRecentActivityNEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldHasRecentActivity, v))
}

// HasRecentActivityIn applies the In predicate on the "has_recent_activity" field.
func HasRecentActivityIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldHasRecentActivity, vs...))
}

// HasRecentActivityNotIn applies the NotIn predicate on the "has_recent_activity" field.
func HasRecentActivityNotIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldHasRecentActivity, vs...))
}

// HasRecentActivityGT applies the GT predicate on the "has_recent_activity" field.
func HasRecentActivityGT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldHasRecentActivity, v))
}

// HasRecentActivityGTE applies the GTE predicate on the "has_recent_activity" field.
func HasRecentActivityGTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldHasRecentActivity, v))
}

// HasRecentActivityLT applies the LT predicate on the "has_recent_activity" field.
func HasRecentActivityLT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldHasRecentActivity, v))
}

// HasRecentActivityLTE applies the LTE predicate on the "has_recent_activity" field.
func HasRecentActivityLTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldHasRecentActivity, v))
}

// HasRecentActivityContains applies the Contains predicate on the "has_recent_activity" field.
func HasRecentActivityContains(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContains(FieldHasRecentActivity, v))
}

// HasRecentActivityHasPrefix applies the HasPrefix predicate on the "has_recent_activity" field.
func HasRecentActivityHasPrefix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasPrefix(FieldHasRecentActivity, v))
}

// HasRecentActivityHasSuffix applies the HasSuffix predicate on the "has_recent_activity" field.
func HasRecentActivityHasSuffix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasSuffix(FieldHasRecentActivity, v))
}

// HasRecentActivityEqualFold applies the EqualFold predicate on the "has_recent_activity" field.
func HasRecentActivityEqualFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEqualFold(FieldHasRecentActivity, v))
}

// HasRecentActivityContainsFold applies the ContainsFold predicate on the "has_recent_activity" field.
func HasRecentActivityContainsFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContainsFold(FieldHasRecentActivity, v))
}

// HasUsedEQ applies the EQ predicate on the "has_used" field.
func HasUsedEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEQ(FieldHasUsed, v))
}

// HasUsedNEQ applies the NEQ predicate on the "has_used" field.
func HasUsedNEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNEQ(FieldHasUsed, v))
}

// HasUsedIn applies the In predicate on the "has_used" field.
func HasUsedIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldIn(FieldHasUsed, vs...))
}

// HasUsedNotIn applies the NotIn predicate on the "has_used" field.
func HasUsedNotIn(vs ...string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldNotIn(FieldHasUsed, vs...))
}

// HasUsedGT applies the GT predicate on the "has_used" field.
func HasUsedGT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGT(FieldHasUsed, v))
}

// HasUsedGTE applies the GTE predicate on the "has_used" field.
func HasUsedGTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldGTE(FieldHasUsed, v))
}

// HasUsedLT applies the LT predicate on the "has_used" field.
func HasUsedLT(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLT(FieldHasUsed, v))
}

// HasUsedLTE applies the LTE predicate on the "has_used" field.
func HasUsedLTE(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldLTE(FieldHasUsed, v))
}

// HasUsedContains applies the Contains predicate on the "has_used" field.
func HasUsedContains(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContains(FieldHasUsed, v))
}

// HasUsedHasPrefix applies the HasPrefix predicate on the "has_used" field.
func HasUsedHasPrefix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasPrefix(FieldHasUsed, v))
}

// HasUsedHasSuffix applies the HasSuffix predicate on the "has_used" field.
func HasUsedHasSuffix(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldHasSuffix(FieldHasUsed, v))
}

// HasUsedEqualFold applies the EqualFold predicate on the "has_used" field.
func HasUsedEqualFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldEqualFold(FieldHasUsed, v))
}

// HasUsedContainsFold applies the ContainsFold predicate on the "has_used" field.
func HasUsedContainsFold(v string) predicate.AccessToken {
	return predicate.AccessToken(sql.FieldContainsFold(FieldHasUsed, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v timeutil.TimeStamp) predicate.AccessToken {
	vc := int64(v)
	return predicate.AccessToken(sql.FieldEQ(FieldCreateTime, vc))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v timeutil.TimeStamp) predicate.AccessToken {
	vc := int64(v)
	return predicate.AccessToken(sql.FieldNEQ(FieldCreateTime, vc))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...timeutil.TimeStamp) predicate.AccessToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.AccessToken(sql.FieldIn(FieldCreateTime, v...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...timeutil.TimeStamp) predicate.AccessToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.AccessToken(sql.FieldNotIn(FieldCreateTime, v...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v timeutil.TimeStamp) predicate.AccessToken {
	vc := int64(v)
	return predicate.AccessToken(sql.FieldGT(FieldCreateTime, vc))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v timeutil.TimeStamp) predicate.AccessToken {
	vc := int64(v)
	return predicate.AccessToken(sql.FieldGTE(FieldCreateTime, vc))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v timeutil.TimeStamp) predicate.AccessToken {
	vc := int64(v)
	return predicate.AccessToken(sql.FieldLT(FieldCreateTime, vc))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v timeutil.TimeStamp) predicate.AccessToken {
	vc := int64(v)
	return predicate.AccessToken(sql.FieldLTE(FieldCreateTime, vc))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v timeutil.TimeStamp) predicate.AccessToken {
	vc := int64(v)
	return predicate.AccessToken(sql.FieldEQ(FieldUpdateTime, vc))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v timeutil.TimeStamp) predicate.AccessToken {
	vc := int64(v)
	return predicate.AccessToken(sql.FieldNEQ(FieldUpdateTime, vc))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...timeutil.TimeStamp) predicate.AccessToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.AccessToken(sql.FieldIn(FieldUpdateTime, v...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...timeutil.TimeStamp) predicate.AccessToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.AccessToken(sql.FieldNotIn(FieldUpdateTime, v...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v timeutil.TimeStamp) predicate.AccessToken {
	vc := int64(v)
	return predicate.AccessToken(sql.FieldGT(FieldUpdateTime, vc))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v timeutil.TimeStamp) predicate.AccessToken {
	vc := int64(v)
	return predicate.AccessToken(sql.FieldGTE(FieldUpdateTime, vc))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v timeutil.TimeStamp) predicate.AccessToken {
	vc := int64(v)
	return predicate.AccessToken(sql.FieldLT(FieldUpdateTime, vc))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v timeutil.TimeStamp) predicate.AccessToken {
	vc := int64(v)
	return predicate.AccessToken(sql.FieldLTE(FieldUpdateTime, vc))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AccessToken) predicate.AccessToken {
	return predicate.AccessToken(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AccessToken) predicate.AccessToken {
	return predicate.AccessToken(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AccessToken) predicate.AccessToken {
	return predicate.AccessToken(sql.NotPredicates(p))
}
