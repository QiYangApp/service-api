// Code generated by ent, DO NOT EDIT.

package memberrolerelatedpermission

import (
	"ent/models/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldLTE(FieldID, id))
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v uuid.UUID) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldEQ(FieldRoleID, v))
}

// PermissionGroupID applies equality check predicate on the "permission_group_id" field. It's identical to PermissionGroupIDEQ.
func PermissionGroupID(v uuid.UUID) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldEQ(FieldPermissionGroupID, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldEQ(FieldUpdateTime, v))
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v uuid.UUID) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldEQ(FieldRoleID, v))
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v uuid.UUID) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldNEQ(FieldRoleID, v))
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...uuid.UUID) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldIn(FieldRoleID, vs...))
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...uuid.UUID) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldNotIn(FieldRoleID, vs...))
}

// RoleIDGT applies the GT predicate on the "role_id" field.
func RoleIDGT(v uuid.UUID) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldGT(FieldRoleID, v))
}

// RoleIDGTE applies the GTE predicate on the "role_id" field.
func RoleIDGTE(v uuid.UUID) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldGTE(FieldRoleID, v))
}

// RoleIDLT applies the LT predicate on the "role_id" field.
func RoleIDLT(v uuid.UUID) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldLT(FieldRoleID, v))
}

// RoleIDLTE applies the LTE predicate on the "role_id" field.
func RoleIDLTE(v uuid.UUID) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldLTE(FieldRoleID, v))
}

// PermissionGroupIDEQ applies the EQ predicate on the "permission_group_id" field.
func PermissionGroupIDEQ(v uuid.UUID) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldEQ(FieldPermissionGroupID, v))
}

// PermissionGroupIDNEQ applies the NEQ predicate on the "permission_group_id" field.
func PermissionGroupIDNEQ(v uuid.UUID) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldNEQ(FieldPermissionGroupID, v))
}

// PermissionGroupIDIn applies the In predicate on the "permission_group_id" field.
func PermissionGroupIDIn(vs ...uuid.UUID) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldIn(FieldPermissionGroupID, vs...))
}

// PermissionGroupIDNotIn applies the NotIn predicate on the "permission_group_id" field.
func PermissionGroupIDNotIn(vs ...uuid.UUID) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldNotIn(FieldPermissionGroupID, vs...))
}

// PermissionGroupIDGT applies the GT predicate on the "permission_group_id" field.
func PermissionGroupIDGT(v uuid.UUID) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldGT(FieldPermissionGroupID, v))
}

// PermissionGroupIDGTE applies the GTE predicate on the "permission_group_id" field.
func PermissionGroupIDGTE(v uuid.UUID) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldGTE(FieldPermissionGroupID, v))
}

// PermissionGroupIDLT applies the LT predicate on the "permission_group_id" field.
func PermissionGroupIDLT(v uuid.UUID) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldLT(FieldPermissionGroupID, v))
}

// PermissionGroupIDLTE applies the LTE predicate on the "permission_group_id" field.
func PermissionGroupIDLTE(v uuid.UUID) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldLTE(FieldPermissionGroupID, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.FieldLTE(FieldUpdateTime, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MemberRoleRelatedPermission) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MemberRoleRelatedPermission) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MemberRoleRelatedPermission) predicate.MemberRoleRelatedPermission {
	return predicate.MemberRoleRelatedPermission(sql.NotPredicates(p))
}
