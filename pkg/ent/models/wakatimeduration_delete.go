// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/wakatimeduration"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WakatimeDurationDelete is the builder for deleting a WakatimeDuration entity.
type WakatimeDurationDelete struct {
	config
	hooks    []Hook
	mutation *WakatimeDurationMutation
}

// Where appends a list predicates to the WakatimeDurationDelete builder.
func (wdd *WakatimeDurationDelete) Where(ps ...predicate.WakatimeDuration) *WakatimeDurationDelete {
	wdd.mutation.Where(ps...)
	return wdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wdd *WakatimeDurationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wdd.sqlExec, wdd.mutation, wdd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wdd *WakatimeDurationDelete) ExecX(ctx context.Context) int {
	n, err := wdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wdd *WakatimeDurationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(wakatimeduration.Table, sqlgraph.NewFieldSpec(wakatimeduration.FieldID, field.TypeInt64))
	if ps := wdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wdd.mutation.done = true
	return affected, err
}

// WakatimeDurationDeleteOne is the builder for deleting a single WakatimeDuration entity.
type WakatimeDurationDeleteOne struct {
	wdd *WakatimeDurationDelete
}

// Where appends a list predicates to the WakatimeDurationDelete builder.
func (wddo *WakatimeDurationDeleteOne) Where(ps ...predicate.WakatimeDuration) *WakatimeDurationDeleteOne {
	wddo.wdd.mutation.Where(ps...)
	return wddo
}

// Exec executes the deletion query.
func (wddo *WakatimeDurationDeleteOne) Exec(ctx context.Context) error {
	n, err := wddo.wdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{wakatimeduration.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wddo *WakatimeDurationDeleteOne) ExecX(ctx context.Context) {
	if err := wddo.Exec(ctx); err != nil {
		panic(err)
	}
}
