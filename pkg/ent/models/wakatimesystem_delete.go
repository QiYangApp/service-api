// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/wakatimesystem"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WakatimeSystemDelete is the builder for deleting a WakatimeSystem entity.
type WakatimeSystemDelete struct {
	config
	hooks    []Hook
	mutation *WakatimeSystemMutation
}

// Where appends a list predicates to the WakatimeSystemDelete builder.
func (wsd *WakatimeSystemDelete) Where(ps ...predicate.WakatimeSystem) *WakatimeSystemDelete {
	wsd.mutation.Where(ps...)
	return wsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wsd *WakatimeSystemDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wsd.sqlExec, wsd.mutation, wsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wsd *WakatimeSystemDelete) ExecX(ctx context.Context) int {
	n, err := wsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wsd *WakatimeSystemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(wakatimesystem.Table, sqlgraph.NewFieldSpec(wakatimesystem.FieldID, field.TypeInt64))
	if ps := wsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wsd.mutation.done = true
	return affected, err
}

// WakatimeSystemDeleteOne is the builder for deleting a single WakatimeSystem entity.
type WakatimeSystemDeleteOne struct {
	wsd *WakatimeSystemDelete
}

// Where appends a list predicates to the WakatimeSystemDelete builder.
func (wsdo *WakatimeSystemDeleteOne) Where(ps ...predicate.WakatimeSystem) *WakatimeSystemDeleteOne {
	wsdo.wsd.mutation.Where(ps...)
	return wsdo
}

// Exec executes the deletion query.
func (wsdo *WakatimeSystemDeleteOne) Exec(ctx context.Context) error {
	n, err := wsdo.wsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{wakatimesystem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wsdo *WakatimeSystemDeleteOne) ExecX(ctx context.Context) {
	if err := wsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
