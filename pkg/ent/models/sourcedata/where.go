// Code generated by ent, DO NOT EDIT.

package sourcedata

import (
	"ent/models/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.SourceData {
	return predicate.SourceData(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.SourceData {
	return predicate.SourceData(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.SourceData {
	return predicate.SourceData(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.SourceData {
	return predicate.SourceData(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.SourceData {
	return predicate.SourceData(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.SourceData {
	return predicate.SourceData(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.SourceData {
	return predicate.SourceData(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldUserID, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldType, v))
}

// SubType applies equality check predicate on the "sub_type" field. It's identical to SubTypeEQ.
func SubType(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldSubType, v))
}

// Info applies equality check predicate on the "info" field. It's identical to InfoEQ.
func Info(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldInfo, v))
}

// Snapshot applies equality check predicate on the "snapshot" field. It's identical to SnapshotEQ.
func Snapshot(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldSnapshot, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldUpdateTime, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.SourceData {
	return predicate.SourceData(sql.FieldLTE(FieldUserID, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.SourceData {
	return predicate.SourceData(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.SourceData {
	return predicate.SourceData(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldContainsFold(FieldType, v))
}

// SubTypeEQ applies the EQ predicate on the "sub_type" field.
func SubTypeEQ(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldSubType, v))
}

// SubTypeNEQ applies the NEQ predicate on the "sub_type" field.
func SubTypeNEQ(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldNEQ(FieldSubType, v))
}

// SubTypeIn applies the In predicate on the "sub_type" field.
func SubTypeIn(vs ...string) predicate.SourceData {
	return predicate.SourceData(sql.FieldIn(FieldSubType, vs...))
}

// SubTypeNotIn applies the NotIn predicate on the "sub_type" field.
func SubTypeNotIn(vs ...string) predicate.SourceData {
	return predicate.SourceData(sql.FieldNotIn(FieldSubType, vs...))
}

// SubTypeGT applies the GT predicate on the "sub_type" field.
func SubTypeGT(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldGT(FieldSubType, v))
}

// SubTypeGTE applies the GTE predicate on the "sub_type" field.
func SubTypeGTE(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldGTE(FieldSubType, v))
}

// SubTypeLT applies the LT predicate on the "sub_type" field.
func SubTypeLT(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldLT(FieldSubType, v))
}

// SubTypeLTE applies the LTE predicate on the "sub_type" field.
func SubTypeLTE(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldLTE(FieldSubType, v))
}

// SubTypeContains applies the Contains predicate on the "sub_type" field.
func SubTypeContains(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldContains(FieldSubType, v))
}

// SubTypeHasPrefix applies the HasPrefix predicate on the "sub_type" field.
func SubTypeHasPrefix(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldHasPrefix(FieldSubType, v))
}

// SubTypeHasSuffix applies the HasSuffix predicate on the "sub_type" field.
func SubTypeHasSuffix(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldHasSuffix(FieldSubType, v))
}

// SubTypeEqualFold applies the EqualFold predicate on the "sub_type" field.
func SubTypeEqualFold(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEqualFold(FieldSubType, v))
}

// SubTypeContainsFold applies the ContainsFold predicate on the "sub_type" field.
func SubTypeContainsFold(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldContainsFold(FieldSubType, v))
}

// InfoEQ applies the EQ predicate on the "info" field.
func InfoEQ(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldInfo, v))
}

// InfoNEQ applies the NEQ predicate on the "info" field.
func InfoNEQ(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldNEQ(FieldInfo, v))
}

// InfoIn applies the In predicate on the "info" field.
func InfoIn(vs ...string) predicate.SourceData {
	return predicate.SourceData(sql.FieldIn(FieldInfo, vs...))
}

// InfoNotIn applies the NotIn predicate on the "info" field.
func InfoNotIn(vs ...string) predicate.SourceData {
	return predicate.SourceData(sql.FieldNotIn(FieldInfo, vs...))
}

// InfoGT applies the GT predicate on the "info" field.
func InfoGT(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldGT(FieldInfo, v))
}

// InfoGTE applies the GTE predicate on the "info" field.
func InfoGTE(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldGTE(FieldInfo, v))
}

// InfoLT applies the LT predicate on the "info" field.
func InfoLT(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldLT(FieldInfo, v))
}

// InfoLTE applies the LTE predicate on the "info" field.
func InfoLTE(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldLTE(FieldInfo, v))
}

// InfoContains applies the Contains predicate on the "info" field.
func InfoContains(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldContains(FieldInfo, v))
}

// InfoHasPrefix applies the HasPrefix predicate on the "info" field.
func InfoHasPrefix(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldHasPrefix(FieldInfo, v))
}

// InfoHasSuffix applies the HasSuffix predicate on the "info" field.
func InfoHasSuffix(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldHasSuffix(FieldInfo, v))
}

// InfoEqualFold applies the EqualFold predicate on the "info" field.
func InfoEqualFold(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEqualFold(FieldInfo, v))
}

// InfoContainsFold applies the ContainsFold predicate on the "info" field.
func InfoContainsFold(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldContainsFold(FieldInfo, v))
}

// SnapshotEQ applies the EQ predicate on the "snapshot" field.
func SnapshotEQ(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldSnapshot, v))
}

// SnapshotNEQ applies the NEQ predicate on the "snapshot" field.
func SnapshotNEQ(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldNEQ(FieldSnapshot, v))
}

// SnapshotIn applies the In predicate on the "snapshot" field.
func SnapshotIn(vs ...string) predicate.SourceData {
	return predicate.SourceData(sql.FieldIn(FieldSnapshot, vs...))
}

// SnapshotNotIn applies the NotIn predicate on the "snapshot" field.
func SnapshotNotIn(vs ...string) predicate.SourceData {
	return predicate.SourceData(sql.FieldNotIn(FieldSnapshot, vs...))
}

// SnapshotGT applies the GT predicate on the "snapshot" field.
func SnapshotGT(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldGT(FieldSnapshot, v))
}

// SnapshotGTE applies the GTE predicate on the "snapshot" field.
func SnapshotGTE(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldGTE(FieldSnapshot, v))
}

// SnapshotLT applies the LT predicate on the "snapshot" field.
func SnapshotLT(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldLT(FieldSnapshot, v))
}

// SnapshotLTE applies the LTE predicate on the "snapshot" field.
func SnapshotLTE(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldLTE(FieldSnapshot, v))
}

// SnapshotContains applies the Contains predicate on the "snapshot" field.
func SnapshotContains(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldContains(FieldSnapshot, v))
}

// SnapshotHasPrefix applies the HasPrefix predicate on the "snapshot" field.
func SnapshotHasPrefix(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldHasPrefix(FieldSnapshot, v))
}

// SnapshotHasSuffix applies the HasSuffix predicate on the "snapshot" field.
func SnapshotHasSuffix(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldHasSuffix(FieldSnapshot, v))
}

// SnapshotEqualFold applies the EqualFold predicate on the "snapshot" field.
func SnapshotEqualFold(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldEqualFold(FieldSnapshot, v))
}

// SnapshotContainsFold applies the ContainsFold predicate on the "snapshot" field.
func SnapshotContainsFold(v string) predicate.SourceData {
	return predicate.SourceData(sql.FieldContainsFold(FieldSnapshot, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.SourceData {
	return predicate.SourceData(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.SourceData {
	return predicate.SourceData(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.SourceData {
	return predicate.SourceData(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.SourceData {
	return predicate.SourceData(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.SourceData {
	return predicate.SourceData(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.SourceData {
	return predicate.SourceData(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.SourceData {
	return predicate.SourceData(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.SourceData {
	return predicate.SourceData(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.SourceData {
	return predicate.SourceData(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.SourceData {
	return predicate.SourceData(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.SourceData {
	return predicate.SourceData(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.SourceData {
	return predicate.SourceData(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.SourceData {
	return predicate.SourceData(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.SourceData {
	return predicate.SourceData(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.SourceData {
	return predicate.SourceData(sql.FieldLTE(FieldUpdateTime, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SourceData) predicate.SourceData {
	return predicate.SourceData(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SourceData) predicate.SourceData {
	return predicate.SourceData(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SourceData) predicate.SourceData {
	return predicate.SourceData(sql.NotPredicates(p))
}
