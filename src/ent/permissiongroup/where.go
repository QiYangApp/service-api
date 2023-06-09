// Code generated by ent, DO NOT EDIT.

package permissiongroup

import (
	"service-api/src/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldUpdateTime, v))
}

// PermissionName applies equality check predicate on the "permission_name" field. It's identical to PermissionNameEQ.
func PermissionName(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldPermissionName, v))
}

// Icon applies equality check predicate on the "icon" field. It's identical to IconEQ.
func Icon(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldIcon, v))
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
func State(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldState, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLTE(FieldUpdateTime, v))
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

// IconEQ applies the EQ predicate on the "icon" field.
func IconEQ(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldIcon, v))
}

// IconNEQ applies the NEQ predicate on the "icon" field.
func IconNEQ(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNEQ(FieldIcon, v))
}

// IconIn applies the In predicate on the "icon" field.
func IconIn(vs ...string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldIn(FieldIcon, vs...))
}

// IconNotIn applies the NotIn predicate on the "icon" field.
func IconNotIn(vs ...string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNotIn(FieldIcon, vs...))
}

// IconGT applies the GT predicate on the "icon" field.
func IconGT(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGT(FieldIcon, v))
}

// IconGTE applies the GTE predicate on the "icon" field.
func IconGTE(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGTE(FieldIcon, v))
}

// IconLT applies the LT predicate on the "icon" field.
func IconLT(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLT(FieldIcon, v))
}

// IconLTE applies the LTE predicate on the "icon" field.
func IconLTE(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLTE(FieldIcon, v))
}

// IconContains applies the Contains predicate on the "icon" field.
func IconContains(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldContains(FieldIcon, v))
}

// IconHasPrefix applies the HasPrefix predicate on the "icon" field.
func IconHasPrefix(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldHasPrefix(FieldIcon, v))
}

// IconHasSuffix applies the HasSuffix predicate on the "icon" field.
func IconHasSuffix(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldHasSuffix(FieldIcon, v))
}

// IconEqualFold applies the EqualFold predicate on the "icon" field.
func IconEqualFold(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEqualFold(FieldIcon, v))
}

// IconContainsFold applies the ContainsFold predicate on the "icon" field.
func IconContainsFold(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldContainsFold(FieldIcon, v))
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
func StateEQ(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNEQ(FieldState, v))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldIn(FieldState, vs...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldNotIn(FieldState, vs...))
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGT(FieldState, v))
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldGTE(FieldState, v))
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLT(FieldState, v))
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldLTE(FieldState, v))
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldContains(FieldState, v))
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldHasPrefix(FieldState, v))
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldHasSuffix(FieldState, v))
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldEqualFold(FieldState, v))
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.PermissionGroup {
	return predicate.PermissionGroup(sql.FieldContainsFold(FieldState, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PermissionGroup) predicate.PermissionGroup {
	return predicate.PermissionGroup(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PermissionGroup) predicate.PermissionGroup {
	return predicate.PermissionGroup(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PermissionGroup) predicate.PermissionGroup {
	return predicate.PermissionGroup(func(s *sql.Selector) {
		p(s.Not())
	})
}
