// Code generated by ent, DO NOT EDIT.

package wakatimeprojectinfo

import (
	"ent/models/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.WakatimeProjectInfo {
	return predicate.WakatimeProjectInfo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.WakatimeProjectInfo {
	return predicate.WakatimeProjectInfo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.WakatimeProjectInfo {
	return predicate.WakatimeProjectInfo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.WakatimeProjectInfo {
	return predicate.WakatimeProjectInfo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.WakatimeProjectInfo {
	return predicate.WakatimeProjectInfo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.WakatimeProjectInfo {
	return predicate.WakatimeProjectInfo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.WakatimeProjectInfo {
	return predicate.WakatimeProjectInfo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.WakatimeProjectInfo {
	return predicate.WakatimeProjectInfo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.WakatimeProjectInfo {
	return predicate.WakatimeProjectInfo(sql.FieldLTE(FieldID, id))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WakatimeProjectInfo) predicate.WakatimeProjectInfo {
	return predicate.WakatimeProjectInfo(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WakatimeProjectInfo) predicate.WakatimeProjectInfo {
	return predicate.WakatimeProjectInfo(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WakatimeProjectInfo) predicate.WakatimeProjectInfo {
	return predicate.WakatimeProjectInfo(sql.NotPredicates(p))
}
