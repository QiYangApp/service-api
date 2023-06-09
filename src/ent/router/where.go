// Code generated by ent, DO NOT EDIT.

package router

import (
	"service-api/src/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Router {
	return predicate.Router(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Router {
	return predicate.Router(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Router {
	return predicate.Router(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Router {
	return predicate.Router(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Router {
	return predicate.Router(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Router {
	return predicate.Router(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Router {
	return predicate.Router(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Router {
	return predicate.Router(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Router {
	return predicate.Router(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Router {
	return predicate.Router(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Router {
	return predicate.Router(sql.FieldEQ(FieldUpdateTime, v))
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
func State(v string) predicate.Router {
	return predicate.Router(sql.FieldEQ(FieldState, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Router {
	return predicate.Router(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Router {
	return predicate.Router(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Router {
	return predicate.Router(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Router {
	return predicate.Router(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Router {
	return predicate.Router(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Router {
	return predicate.Router(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Router {
	return predicate.Router(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Router {
	return predicate.Router(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Router {
	return predicate.Router(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Router {
	return predicate.Router(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Router {
	return predicate.Router(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Router {
	return predicate.Router(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Router {
	return predicate.Router(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Router {
	return predicate.Router(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Router {
	return predicate.Router(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Router {
	return predicate.Router(sql.FieldLTE(FieldUpdateTime, v))
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
func StateEQ(v string) predicate.Router {
	return predicate.Router(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.Router {
	return predicate.Router(sql.FieldNEQ(FieldState, v))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.Router {
	return predicate.Router(sql.FieldIn(FieldState, vs...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.Router {
	return predicate.Router(sql.FieldNotIn(FieldState, vs...))
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.Router {
	return predicate.Router(sql.FieldGT(FieldState, v))
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.Router {
	return predicate.Router(sql.FieldGTE(FieldState, v))
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.Router {
	return predicate.Router(sql.FieldLT(FieldState, v))
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.Router {
	return predicate.Router(sql.FieldLTE(FieldState, v))
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.Router {
	return predicate.Router(sql.FieldContains(FieldState, v))
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.Router {
	return predicate.Router(sql.FieldHasPrefix(FieldState, v))
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.Router {
	return predicate.Router(sql.FieldHasSuffix(FieldState, v))
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.Router {
	return predicate.Router(sql.FieldEqualFold(FieldState, v))
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.Router {
	return predicate.Router(sql.FieldContainsFold(FieldState, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Router) predicate.Router {
	return predicate.Router(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Router) predicate.Router {
	return predicate.Router(func(s *sql.Selector) {
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
func Not(p predicate.Router) predicate.Router {
	return predicate.Router(func(s *sql.Selector) {
		p(s.Not())
	})
}
