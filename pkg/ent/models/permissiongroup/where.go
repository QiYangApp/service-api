// Code generated by ent, DO NOT EDIT.

package permissiongroup

import (
	"ent/enums/state"
	"ent/models/predicate"
	"ent/utils/timeutil"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLTE(FieldID, id))
}

// PermissionName applies equality check predicate on the "permission_name" field. It's identical to PermissionNameEQ.
func PermissionName(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldPermissionName, v))
}

// Ioc applies equality check predicate on the "ioc" field. It's identical to IocEQ.
func Ioc(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldIoc, v))
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldSort, v))
}

// Left applies equality check predicate on the "left" field. It's identical to LeftEQ.
func Left(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldLeft, v))
}

// Right applies equality check predicate on the "right" field. It's identical to RightEQ.
func Right(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldRight, v))
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v state.SwitchState) predicate.PermissionGroup {
	vc := int(v)
	return predicate.PermissionGroup(sql.FieldEQ(FieldState, vc))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v timeutil.TimeStamp) predicate.PermissionGroup {
	vc := int64(v)
	return predicate.PermissionGroup(sql.FieldEQ(FieldCreateTime, vc))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v timeutil.TimeStamp) predicate.PermissionGroup {
	vc := int64(v)
	return predicate.PermissionGroup(sql.FieldEQ(FieldUpdateTime, vc))
}

// PermissionNameEQ applies the EQ predicate on the "permission_name" field.
func PermissionNameEQ(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldPermissionName, v))
}

// PermissionNameNEQ applies the NEQ predicate on the "permission_name" field.
func PermissionNameNEQ(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNEQ(FieldPermissionName, v))
}

// PermissionNameIn applies the In predicate on the "permission_name" field.
func PermissionNameIn(vs ...string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldIn(FieldPermissionName, vs...))
}

// PermissionNameNotIn applies the NotIn predicate on the "permission_name" field.
func PermissionNameNotIn(vs ...string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNotIn(FieldPermissionName, vs...))
}

// PermissionNameGT applies the GT predicate on the "permission_name" field.
func PermissionNameGT(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGT(FieldPermissionName, v))
}

// PermissionNameGTE applies the GTE predicate on the "permission_name" field.
func PermissionNameGTE(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGTE(FieldPermissionName, v))
}

// PermissionNameLT applies the LT predicate on the "permission_name" field.
func PermissionNameLT(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLT(FieldPermissionName, v))
}

// PermissionNameLTE applies the LTE predicate on the "permission_name" field.
func PermissionNameLTE(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLTE(FieldPermissionName, v))
}

// PermissionNameContains applies the Contains predicate on the "permission_name" field.
func PermissionNameContains(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldContains(FieldPermissionName, v))
}

// PermissionNameHasPrefix applies the HasPrefix predicate on the "permission_name" field.
func PermissionNameHasPrefix(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldHasPrefix(FieldPermissionName, v))
}

// PermissionNameHasSuffix applies the HasSuffix predicate on the "permission_name" field.
func PermissionNameHasSuffix(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldHasSuffix(FieldPermissionName, v))
}

// PermissionNameEqualFold applies the EqualFold predicate on the "permission_name" field.
func PermissionNameEqualFold(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEqualFold(FieldPermissionName, v))
}

// PermissionNameContainsFold applies the ContainsFold predicate on the "permission_name" field.
func PermissionNameContainsFold(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldContainsFold(FieldPermissionName, v))
}

// IocEQ applies the EQ predicate on the "ioc" field.
func IocEQ(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldIoc, v))
}

// IocNEQ applies the NEQ predicate on the "ioc" field.
func IocNEQ(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNEQ(FieldIoc, v))
}

// IocIn applies the In predicate on the "ioc" field.
func IocIn(vs ...string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldIn(FieldIoc, vs...))
}

// IocNotIn applies the NotIn predicate on the "ioc" field.
func IocNotIn(vs ...string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNotIn(FieldIoc, vs...))
}

// IocGT applies the GT predicate on the "ioc" field.
func IocGT(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGT(FieldIoc, v))
}

// IocGTE applies the GTE predicate on the "ioc" field.
func IocGTE(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGTE(FieldIoc, v))
}

// IocLT applies the LT predicate on the "ioc" field.
func IocLT(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLT(FieldIoc, v))
}

// IocLTE applies the LTE predicate on the "ioc" field.
func IocLTE(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLTE(FieldIoc, v))
}

// IocContains applies the Contains predicate on the "ioc" field.
func IocContains(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldContains(FieldIoc, v))
}

// IocHasPrefix applies the HasPrefix predicate on the "ioc" field.
func IocHasPrefix(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldHasPrefix(FieldIoc, v))
}

// IocHasSuffix applies the HasSuffix predicate on the "ioc" field.
func IocHasSuffix(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldHasSuffix(FieldIoc, v))
}

// IocEqualFold applies the EqualFold predicate on the "ioc" field.
func IocEqualFold(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEqualFold(FieldIoc, v))
}

// IocContainsFold applies the ContainsFold predicate on the "ioc" field.
func IocContainsFold(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldContainsFold(FieldIoc, v))
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldSort, v))
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNEQ(FieldSort, v))
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldIn(FieldSort, vs...))
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNotIn(FieldSort, vs...))
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGT(FieldSort, v))
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGTE(FieldSort, v))
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLT(FieldSort, v))
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLTE(FieldSort, v))
}

// LeftEQ applies the EQ predicate on the "left" field.
func LeftEQ(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldLeft, v))
}

// LeftNEQ applies the NEQ predicate on the "left" field.
func LeftNEQ(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNEQ(FieldLeft, v))
}

// LeftIn applies the In predicate on the "left" field.
func LeftIn(vs ...int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldIn(FieldLeft, vs...))
}

// LeftNotIn applies the NotIn predicate on the "left" field.
func LeftNotIn(vs ...int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNotIn(FieldLeft, vs...))
}

// LeftGT applies the GT predicate on the "left" field.
func LeftGT(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGT(FieldLeft, v))
}

// LeftGTE applies the GTE predicate on the "left" field.
func LeftGTE(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGTE(FieldLeft, v))
}

// LeftLT applies the LT predicate on the "left" field.
func LeftLT(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLT(FieldLeft, v))
}

// LeftLTE applies the LTE predicate on the "left" field.
func LeftLTE(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLTE(FieldLeft, v))
}

// RightEQ applies the EQ predicate on the "right" field.
func RightEQ(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldRight, v))
}

// RightNEQ applies the NEQ predicate on the "right" field.
func RightNEQ(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNEQ(FieldRight, v))
}

// RightIn applies the In predicate on the "right" field.
func RightIn(vs ...int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldIn(FieldRight, vs...))
}

// RightNotIn applies the NotIn predicate on the "right" field.
func RightNotIn(vs ...int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNotIn(FieldRight, vs...))
}

// RightGT applies the GT predicate on the "right" field.
func RightGT(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGT(FieldRight, v))
}

// RightGTE applies the GTE predicate on the "right" field.
func RightGTE(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGTE(FieldRight, v))
}

// RightLT applies the LT predicate on the "right" field.
func RightLT(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLT(FieldRight, v))
}

// RightLTE applies the LTE predicate on the "right" field.
func RightLTE(v int32) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLTE(FieldRight, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v state.SwitchState) predicate.PermissionGroup {
	vc := int(v)
	return predicate.PermissionGroup(sql.FieldEQ(FieldState, vc))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v state.SwitchState) predicate.PermissionGroup {
	vc := int(v)
	return predicate.PermissionGroup(sql.FieldNEQ(FieldState, vc))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...state.SwitchState) predicate.PermissionGroup {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.PermissionGroup(sql.FieldIn(FieldState, v...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...state.SwitchState) predicate.PermissionGroup {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.PermissionGroup(sql.FieldNotIn(FieldState, v...))
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v state.SwitchState) predicate.PermissionGroup {
	vc := int(v)
	return predicate.PermissionGroup(sql.FieldGT(FieldState, vc))
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v state.SwitchState) predicate.PermissionGroup {
	vc := int(v)
	return predicate.PermissionGroup(sql.FieldGTE(FieldState, vc))
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v state.SwitchState) predicate.PermissionGroup {
	vc := int(v)
	return predicate.PermissionGroup(sql.FieldLT(FieldState, vc))
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v state.SwitchState) predicate.PermissionGroup {
	vc := int(v)
	return predicate.PermissionGroup(sql.FieldLTE(FieldState, vc))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v timeutil.TimeStamp) predicate.PermissionGroup {
	vc := int64(v)
	return predicate.PermissionGroup(sql.FieldEQ(FieldCreateTime, vc))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v timeutil.TimeStamp) predicate.PermissionGroup {
	vc := int64(v)
	return predicate.PermissionGroup(sql.FieldNEQ(FieldCreateTime, vc))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...timeutil.TimeStamp) predicate.PermissionGroup {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.PermissionGroup(sql.FieldIn(FieldCreateTime, v...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...timeutil.TimeStamp) predicate.PermissionGroup {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.PermissionGroup(sql.FieldNotIn(FieldCreateTime, v...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v timeutil.TimeStamp) predicate.PermissionGroup {
	vc := int64(v)
	return predicate.PermissionGroup(sql.FieldGT(FieldCreateTime, vc))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v timeutil.TimeStamp) predicate.PermissionGroup {
	vc := int64(v)
	return predicate.PermissionGroup(sql.FieldGTE(FieldCreateTime, vc))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v timeutil.TimeStamp) predicate.PermissionGroup {
	vc := int64(v)
	return predicate.PermissionGroup(sql.FieldLT(FieldCreateTime, vc))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v timeutil.TimeStamp) predicate.PermissionGroup {
	vc := int64(v)
	return predicate.PermissionGroup(sql.FieldLTE(FieldCreateTime, vc))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v timeutil.TimeStamp) predicate.PermissionGroup {
	vc := int64(v)
	return predicate.PermissionGroup(sql.FieldEQ(FieldUpdateTime, vc))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v timeutil.TimeStamp) predicate.PermissionGroup {
	vc := int64(v)
	return predicate.PermissionGroup(sql.FieldNEQ(FieldUpdateTime, vc))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...timeutil.TimeStamp) predicate.PermissionGroup {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.PermissionGroup(sql.FieldIn(FieldUpdateTime, v...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...timeutil.TimeStamp) predicate.PermissionGroup {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.PermissionGroup(sql.FieldNotIn(FieldUpdateTime, v...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v timeutil.TimeStamp) predicate.PermissionGroup {
	vc := int64(v)
	return predicate.PermissionGroup(sql.FieldGT(FieldUpdateTime, vc))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v timeutil.TimeStamp) predicate.PermissionGroup {
	vc := int64(v)
	return predicate.PermissionGroup(sql.FieldGTE(FieldUpdateTime, vc))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v timeutil.TimeStamp) predicate.PermissionGroup {
	vc := int64(v)
	return predicate.PermissionGroup(sql.FieldLT(FieldUpdateTime, vc))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v timeutil.TimeStamp) predicate.PermissionGroup {
	vc := int64(v)
	return predicate.PermissionGroup(sql.FieldLTE(FieldUpdateTime, vc))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PermissionGroup) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PermissionGroup) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PermissionGroup) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.NotPredicates(p))
}