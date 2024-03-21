// Code generated by ent, DO NOT EDIT.

package wakatimecategory

import (
	"ent/models/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldEQ(FieldUpdateTime, v))
}

// WakatimeID applies equality check predicate on the "wakatime_id" field. It's identical to WakatimeIDEQ.
func WakatimeID(v uuid.UUID) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldEQ(FieldWakatimeID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldEQ(FieldUserID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldEQ(FieldName, v))
}

// TotalSeconds applies equality check predicate on the "total_seconds" field. It's identical to TotalSecondsEQ.
func TotalSeconds(v int64) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldEQ(FieldTotalSeconds, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldLTE(FieldUpdateTime, v))
}

// WakatimeIDEQ applies the EQ predicate on the "wakatime_id" field.
func WakatimeIDEQ(v uuid.UUID) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldEQ(FieldWakatimeID, v))
}

// WakatimeIDNEQ applies the NEQ predicate on the "wakatime_id" field.
func WakatimeIDNEQ(v uuid.UUID) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldNEQ(FieldWakatimeID, v))
}

// WakatimeIDIn applies the In predicate on the "wakatime_id" field.
func WakatimeIDIn(vs ...uuid.UUID) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldIn(FieldWakatimeID, vs...))
}

// WakatimeIDNotIn applies the NotIn predicate on the "wakatime_id" field.
func WakatimeIDNotIn(vs ...uuid.UUID) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldNotIn(FieldWakatimeID, vs...))
}

// WakatimeIDGT applies the GT predicate on the "wakatime_id" field.
func WakatimeIDGT(v uuid.UUID) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldGT(FieldWakatimeID, v))
}

// WakatimeIDGTE applies the GTE predicate on the "wakatime_id" field.
func WakatimeIDGTE(v uuid.UUID) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldGTE(FieldWakatimeID, v))
}

// WakatimeIDLT applies the LT predicate on the "wakatime_id" field.
func WakatimeIDLT(v uuid.UUID) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldLT(FieldWakatimeID, v))
}

// WakatimeIDLTE applies the LTE predicate on the "wakatime_id" field.
func WakatimeIDLTE(v uuid.UUID) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldLTE(FieldWakatimeID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldLTE(FieldUserID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldContainsFold(FieldName, v))
}

// TotalSecondsEQ applies the EQ predicate on the "total_seconds" field.
func TotalSecondsEQ(v int64) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldEQ(FieldTotalSeconds, v))
}

// TotalSecondsNEQ applies the NEQ predicate on the "total_seconds" field.
func TotalSecondsNEQ(v int64) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldNEQ(FieldTotalSeconds, v))
}

// TotalSecondsIn applies the In predicate on the "total_seconds" field.
func TotalSecondsIn(vs ...int64) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldIn(FieldTotalSeconds, vs...))
}

// TotalSecondsNotIn applies the NotIn predicate on the "total_seconds" field.
func TotalSecondsNotIn(vs ...int64) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldNotIn(FieldTotalSeconds, vs...))
}

// TotalSecondsGT applies the GT predicate on the "total_seconds" field.
func TotalSecondsGT(v int64) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldGT(FieldTotalSeconds, v))
}

// TotalSecondsGTE applies the GTE predicate on the "total_seconds" field.
func TotalSecondsGTE(v int64) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldGTE(FieldTotalSeconds, v))
}

// TotalSecondsLT applies the LT predicate on the "total_seconds" field.
func TotalSecondsLT(v int64) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldLT(FieldTotalSeconds, v))
}

// TotalSecondsLTE applies the LTE predicate on the "total_seconds" field.
func TotalSecondsLTE(v int64) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.FieldLTE(FieldTotalSeconds, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WakatimeCategory) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WakatimeCategory) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WakatimeCategory) predicate.WakatimeCategory {
	return predicate.WakatimeCategory(sql.NotPredicates(p))
}
