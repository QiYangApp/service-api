// Code generated by ent, DO NOT EDIT.

package router

import (
	"ent/enums/state"
	"ent/models/predicate"
	"ent/utils/timeutil"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Router {
	return predicate.Router(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Router {
	return predicate.Router(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Router {
	return predicate.Router(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Router {
	return predicate.Router(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Router {
	return predicate.Router(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Router {
	return predicate.Router(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Router {
	return predicate.Router(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Router {
	return predicate.Router(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Router {
	return predicate.Router(sql.FieldLTE(FieldID, id))
}

// RouteName applies equality check predicate on the "route_name" field. It's identical to RouteNameEQ.
func RouteName(v string) predicate.Router {
	return predicate.Router(sql.FieldEQ(FieldRouteName, v))
}

// Route applies equality check predicate on the "route" field. It's identical to RouteEQ.
func Route(v string) predicate.Router {
	return predicate.Router(sql.FieldEQ(FieldRoute, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Router {
	return predicate.Router(sql.FieldEQ(FieldDescription, v))
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v state.SwitchState) predicate.Router {
	vc := int(v)
	return predicate.Router(sql.FieldEQ(FieldState, vc))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v timeutil.TimeStamp) predicate.Router {
	vc := int64(v)
	return predicate.Router(sql.FieldEQ(FieldCreateTime, vc))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v timeutil.TimeStamp) predicate.Router {
	vc := int64(v)
	return predicate.Router(sql.FieldEQ(FieldUpdateTime, vc))
}

// RouteNameEQ applies the EQ predicate on the "route_name" field.
func RouteNameEQ(v string) predicate.Router {
	return predicate.Router(sql.FieldEQ(FieldRouteName, v))
}

// RouteNameNEQ applies the NEQ predicate on the "route_name" field.
func RouteNameNEQ(v string) predicate.Router {
	return predicate.Router(sql.FieldNEQ(FieldRouteName, v))
}

// RouteNameIn applies the In predicate on the "route_name" field.
func RouteNameIn(vs ...string) predicate.Router {
	return predicate.Router(sql.FieldIn(FieldRouteName, vs...))
}

// RouteNameNotIn applies the NotIn predicate on the "route_name" field.
func RouteNameNotIn(vs ...string) predicate.Router {
	return predicate.Router(sql.FieldNotIn(FieldRouteName, vs...))
}

// RouteNameGT applies the GT predicate on the "route_name" field.
func RouteNameGT(v string) predicate.Router {
	return predicate.Router(sql.FieldGT(FieldRouteName, v))
}

// RouteNameGTE applies the GTE predicate on the "route_name" field.
func RouteNameGTE(v string) predicate.Router {
	return predicate.Router(sql.FieldGTE(FieldRouteName, v))
}

// RouteNameLT applies the LT predicate on the "route_name" field.
func RouteNameLT(v string) predicate.Router {
	return predicate.Router(sql.FieldLT(FieldRouteName, v))
}

// RouteNameLTE applies the LTE predicate on the "route_name" field.
func RouteNameLTE(v string) predicate.Router {
	return predicate.Router(sql.FieldLTE(FieldRouteName, v))
}

// RouteNameContains applies the Contains predicate on the "route_name" field.
func RouteNameContains(v string) predicate.Router {
	return predicate.Router(sql.FieldContains(FieldRouteName, v))
}

// RouteNameHasPrefix applies the HasPrefix predicate on the "route_name" field.
func RouteNameHasPrefix(v string) predicate.Router {
	return predicate.Router(sql.FieldHasPrefix(FieldRouteName, v))
}

// RouteNameHasSuffix applies the HasSuffix predicate on the "route_name" field.
func RouteNameHasSuffix(v string) predicate.Router {
	return predicate.Router(sql.FieldHasSuffix(FieldRouteName, v))
}

// RouteNameEqualFold applies the EqualFold predicate on the "route_name" field.
func RouteNameEqualFold(v string) predicate.Router {
	return predicate.Router(sql.FieldEqualFold(FieldRouteName, v))
}

// RouteNameContainsFold applies the ContainsFold predicate on the "route_name" field.
func RouteNameContainsFold(v string) predicate.Router {
	return predicate.Router(sql.FieldContainsFold(FieldRouteName, v))
}

// RouteEQ applies the EQ predicate on the "route" field.
func RouteEQ(v string) predicate.Router {
	return predicate.Router(sql.FieldEQ(FieldRoute, v))
}

// RouteNEQ applies the NEQ predicate on the "route" field.
func RouteNEQ(v string) predicate.Router {
	return predicate.Router(sql.FieldNEQ(FieldRoute, v))
}

// RouteIn applies the In predicate on the "route" field.
func RouteIn(vs ...string) predicate.Router {
	return predicate.Router(sql.FieldIn(FieldRoute, vs...))
}

// RouteNotIn applies the NotIn predicate on the "route" field.
func RouteNotIn(vs ...string) predicate.Router {
	return predicate.Router(sql.FieldNotIn(FieldRoute, vs...))
}

// RouteGT applies the GT predicate on the "route" field.
func RouteGT(v string) predicate.Router {
	return predicate.Router(sql.FieldGT(FieldRoute, v))
}

// RouteGTE applies the GTE predicate on the "route" field.
func RouteGTE(v string) predicate.Router {
	return predicate.Router(sql.FieldGTE(FieldRoute, v))
}

// RouteLT applies the LT predicate on the "route" field.
func RouteLT(v string) predicate.Router {
	return predicate.Router(sql.FieldLT(FieldRoute, v))
}

// RouteLTE applies the LTE predicate on the "route" field.
func RouteLTE(v string) predicate.Router {
	return predicate.Router(sql.FieldLTE(FieldRoute, v))
}

// RouteContains applies the Contains predicate on the "route" field.
func RouteContains(v string) predicate.Router {
	return predicate.Router(sql.FieldContains(FieldRoute, v))
}

// RouteHasPrefix applies the HasPrefix predicate on the "route" field.
func RouteHasPrefix(v string) predicate.Router {
	return predicate.Router(sql.FieldHasPrefix(FieldRoute, v))
}

// RouteHasSuffix applies the HasSuffix predicate on the "route" field.
func RouteHasSuffix(v string) predicate.Router {
	return predicate.Router(sql.FieldHasSuffix(FieldRoute, v))
}

// RouteEqualFold applies the EqualFold predicate on the "route" field.
func RouteEqualFold(v string) predicate.Router {
	return predicate.Router(sql.FieldEqualFold(FieldRoute, v))
}

// RouteContainsFold applies the ContainsFold predicate on the "route" field.
func RouteContainsFold(v string) predicate.Router {
	return predicate.Router(sql.FieldContainsFold(FieldRoute, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Router {
	return predicate.Router(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Router {
	return predicate.Router(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Router {
	return predicate.Router(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Router {
	return predicate.Router(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Router {
	return predicate.Router(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Router {
	return predicate.Router(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Router {
	return predicate.Router(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Router {
	return predicate.Router(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Router {
	return predicate.Router(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Router {
	return predicate.Router(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Router {
	return predicate.Router(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Router {
	return predicate.Router(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Router {
	return predicate.Router(sql.FieldContainsFold(FieldDescription, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v state.SwitchState) predicate.Router {
	vc := int(v)
	return predicate.Router(sql.FieldEQ(FieldState, vc))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v state.SwitchState) predicate.Router {
	vc := int(v)
	return predicate.Router(sql.FieldNEQ(FieldState, vc))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...state.SwitchState) predicate.Router {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.Router(sql.FieldIn(FieldState, v...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...state.SwitchState) predicate.Router {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.Router(sql.FieldNotIn(FieldState, v...))
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v state.SwitchState) predicate.Router {
	vc := int(v)
	return predicate.Router(sql.FieldGT(FieldState, vc))
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v state.SwitchState) predicate.Router {
	vc := int(v)
	return predicate.Router(sql.FieldGTE(FieldState, vc))
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v state.SwitchState) predicate.Router {
	vc := int(v)
	return predicate.Router(sql.FieldLT(FieldState, vc))
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v state.SwitchState) predicate.Router {
	vc := int(v)
	return predicate.Router(sql.FieldLTE(FieldState, vc))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v timeutil.TimeStamp) predicate.Router {
	vc := int64(v)
	return predicate.Router(sql.FieldEQ(FieldCreateTime, vc))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v timeutil.TimeStamp) predicate.Router {
	vc := int64(v)
	return predicate.Router(sql.FieldNEQ(FieldCreateTime, vc))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...timeutil.TimeStamp) predicate.Router {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.Router(sql.FieldIn(FieldCreateTime, v...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...timeutil.TimeStamp) predicate.Router {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.Router(sql.FieldNotIn(FieldCreateTime, v...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v timeutil.TimeStamp) predicate.Router {
	vc := int64(v)
	return predicate.Router(sql.FieldGT(FieldCreateTime, vc))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v timeutil.TimeStamp) predicate.Router {
	vc := int64(v)
	return predicate.Router(sql.FieldGTE(FieldCreateTime, vc))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v timeutil.TimeStamp) predicate.Router {
	vc := int64(v)
	return predicate.Router(sql.FieldLT(FieldCreateTime, vc))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v timeutil.TimeStamp) predicate.Router {
	vc := int64(v)
	return predicate.Router(sql.FieldLTE(FieldCreateTime, vc))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v timeutil.TimeStamp) predicate.Router {
	vc := int64(v)
	return predicate.Router(sql.FieldEQ(FieldUpdateTime, vc))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v timeutil.TimeStamp) predicate.Router {
	vc := int64(v)
	return predicate.Router(sql.FieldNEQ(FieldUpdateTime, vc))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...timeutil.TimeStamp) predicate.Router {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.Router(sql.FieldIn(FieldUpdateTime, v...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...timeutil.TimeStamp) predicate.Router {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.Router(sql.FieldNotIn(FieldUpdateTime, v...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v timeutil.TimeStamp) predicate.Router {
	vc := int64(v)
	return predicate.Router(sql.FieldGT(FieldUpdateTime, vc))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v timeutil.TimeStamp) predicate.Router {
	vc := int64(v)
	return predicate.Router(sql.FieldGTE(FieldUpdateTime, vc))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v timeutil.TimeStamp) predicate.Router {
	vc := int64(v)
	return predicate.Router(sql.FieldLT(FieldUpdateTime, vc))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v timeutil.TimeStamp) predicate.Router {
	vc := int64(v)
	return predicate.Router(sql.FieldLTE(FieldUpdateTime, vc))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Router) predicate.Router {
	return predicate.Router(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Router) predicate.Router {
	return predicate.Router(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Router) predicate.Router {
	return predicate.Router(sql.NotPredicates(p))
}
