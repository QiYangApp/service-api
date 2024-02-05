// Code generated by ent, DO NOT EDIT.

package wakatimeheartbeat

import (
	"service-api/src/models/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.WakatimeHeartBeat {
	return predicate.WakatimeHeartBeat(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.WakatimeHeartBeat {
	return predicate.WakatimeHeartBeat(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.WakatimeHeartBeat {
	return predicate.WakatimeHeartBeat(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.WakatimeHeartBeat {
	return predicate.WakatimeHeartBeat(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.WakatimeHeartBeat {
	return predicate.WakatimeHeartBeat(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.WakatimeHeartBeat {
	return predicate.WakatimeHeartBeat(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.WakatimeHeartBeat {
	return predicate.WakatimeHeartBeat(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.WakatimeHeartBeat {
	return predicate.WakatimeHeartBeat(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.WakatimeHeartBeat {
	return predicate.WakatimeHeartBeat(sql.FieldLTE(FieldID, id))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WakatimeHeartBeat) predicate.WakatimeHeartBeat {
	return predicate.WakatimeHeartBeat(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WakatimeHeartBeat) predicate.WakatimeHeartBeat {
	return predicate.WakatimeHeartBeat(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WakatimeHeartBeat) predicate.WakatimeHeartBeat {
	return predicate.WakatimeHeartBeat(sql.NotPredicates(p))
}
