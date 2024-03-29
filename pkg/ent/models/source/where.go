// Code generated by ent, DO NOT EDIT.

package source

import (
	"ent/models/predicate"
	"ent/types/auth"
	"ent/utils/timeutil"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Source {
	return predicate.Source(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Source {
	return predicate.Source(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Source {
	return predicate.Source(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Source {
	return predicate.Source(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Source {
	return predicate.Source(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Source {
	return predicate.Source(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Source {
	return predicate.Source(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Source {
	return predicate.Source(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Source {
	return predicate.Source(sql.FieldLTE(FieldID, id))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v auth.Type) predicate.Source {
	vc := int(v)
	return predicate.Source(sql.FieldEQ(FieldType, vc))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Source {
	return predicate.Source(sql.FieldEQ(FieldName, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.Source {
	return predicate.Source(sql.FieldEQ(FieldIsActive, v))
}

// IsSyncEnabled applies equality check predicate on the "is_sync_enabled" field. It's identical to IsSyncEnabledEQ.
func IsSyncEnabled(v bool) predicate.Source {
	return predicate.Source(sql.FieldEQ(FieldIsSyncEnabled, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v timeutil.TimeStamp) predicate.Source {
	vc := int64(v)
	return predicate.Source(sql.FieldEQ(FieldCreateTime, vc))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v timeutil.TimeStamp) predicate.Source {
	vc := int64(v)
	return predicate.Source(sql.FieldEQ(FieldUpdateTime, vc))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v auth.Type) predicate.Source {
	vc := int(v)
	return predicate.Source(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v auth.Type) predicate.Source {
	vc := int(v)
	return predicate.Source(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...auth.Type) predicate.Source {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.Source(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...auth.Type) predicate.Source {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.Source(sql.FieldNotIn(FieldType, v...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v auth.Type) predicate.Source {
	vc := int(v)
	return predicate.Source(sql.FieldGT(FieldType, vc))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v auth.Type) predicate.Source {
	vc := int(v)
	return predicate.Source(sql.FieldGTE(FieldType, vc))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v auth.Type) predicate.Source {
	vc := int(v)
	return predicate.Source(sql.FieldLT(FieldType, vc))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v auth.Type) predicate.Source {
	vc := int(v)
	return predicate.Source(sql.FieldLTE(FieldType, vc))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Source {
	return predicate.Source(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Source {
	return predicate.Source(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Source {
	return predicate.Source(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Source {
	return predicate.Source(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Source {
	return predicate.Source(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Source {
	return predicate.Source(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Source {
	return predicate.Source(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Source {
	return predicate.Source(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Source {
	return predicate.Source(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Source {
	return predicate.Source(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Source {
	return predicate.Source(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Source {
	return predicate.Source(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Source {
	return predicate.Source(sql.FieldContainsFold(FieldName, v))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.Source {
	return predicate.Source(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.Source {
	return predicate.Source(sql.FieldNEQ(FieldIsActive, v))
}

// IsSyncEnabledEQ applies the EQ predicate on the "is_sync_enabled" field.
func IsSyncEnabledEQ(v bool) predicate.Source {
	return predicate.Source(sql.FieldEQ(FieldIsSyncEnabled, v))
}

// IsSyncEnabledNEQ applies the NEQ predicate on the "is_sync_enabled" field.
func IsSyncEnabledNEQ(v bool) predicate.Source {
	return predicate.Source(sql.FieldNEQ(FieldIsSyncEnabled, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v timeutil.TimeStamp) predicate.Source {
	vc := int64(v)
	return predicate.Source(sql.FieldEQ(FieldCreateTime, vc))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v timeutil.TimeStamp) predicate.Source {
	vc := int64(v)
	return predicate.Source(sql.FieldNEQ(FieldCreateTime, vc))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...timeutil.TimeStamp) predicate.Source {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.Source(sql.FieldIn(FieldCreateTime, v...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...timeutil.TimeStamp) predicate.Source {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.Source(sql.FieldNotIn(FieldCreateTime, v...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v timeutil.TimeStamp) predicate.Source {
	vc := int64(v)
	return predicate.Source(sql.FieldGT(FieldCreateTime, vc))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v timeutil.TimeStamp) predicate.Source {
	vc := int64(v)
	return predicate.Source(sql.FieldGTE(FieldCreateTime, vc))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v timeutil.TimeStamp) predicate.Source {
	vc := int64(v)
	return predicate.Source(sql.FieldLT(FieldCreateTime, vc))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v timeutil.TimeStamp) predicate.Source {
	vc := int64(v)
	return predicate.Source(sql.FieldLTE(FieldCreateTime, vc))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v timeutil.TimeStamp) predicate.Source {
	vc := int64(v)
	return predicate.Source(sql.FieldEQ(FieldUpdateTime, vc))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v timeutil.TimeStamp) predicate.Source {
	vc := int64(v)
	return predicate.Source(sql.FieldNEQ(FieldUpdateTime, vc))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...timeutil.TimeStamp) predicate.Source {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.Source(sql.FieldIn(FieldUpdateTime, v...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...timeutil.TimeStamp) predicate.Source {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.Source(sql.FieldNotIn(FieldUpdateTime, v...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v timeutil.TimeStamp) predicate.Source {
	vc := int64(v)
	return predicate.Source(sql.FieldGT(FieldUpdateTime, vc))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v timeutil.TimeStamp) predicate.Source {
	vc := int64(v)
	return predicate.Source(sql.FieldGTE(FieldUpdateTime, vc))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v timeutil.TimeStamp) predicate.Source {
	vc := int64(v)
	return predicate.Source(sql.FieldLT(FieldUpdateTime, vc))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v timeutil.TimeStamp) predicate.Source {
	vc := int64(v)
	return predicate.Source(sql.FieldLTE(FieldUpdateTime, vc))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Source) predicate.Source {
	return predicate.Source(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Source) predicate.Source {
	return predicate.Source(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Source) predicate.Source {
	return predicate.Source(sql.NotPredicates(p))
}
