// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/wakatimeheartbeat"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WakatimeHeartBeatDelete is the builder for deleting a WakatimeHeartBeat entity.
type WakatimeHeartBeatDelete struct {
	config
	hooks    []Hook
	mutation *WakatimeHeartBeatMutation
}

// Where appends a list predicates to the WakatimeHeartBeatDelete builder.
func (whbd *WakatimeHeartBeatDelete) Where(ps ...predicate.WakatimeHeartBeat) *WakatimeHeartBeatDelete {
	whbd.mutation.Where(ps...)
	return whbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (whbd *WakatimeHeartBeatDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, whbd.sqlExec, whbd.mutation, whbd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (whbd *WakatimeHeartBeatDelete) ExecX(ctx context.Context) int {
	n, err := whbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (whbd *WakatimeHeartBeatDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(wakatimeheartbeat.Table, sqlgraph.NewFieldSpec(wakatimeheartbeat.FieldID, field.TypeInt64))
	if ps := whbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, whbd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	whbd.mutation.done = true
	return affected, err
}

// WakatimeHeartBeatDeleteOne is the builder for deleting a single WakatimeHeartBeat entity.
type WakatimeHeartBeatDeleteOne struct {
	whbd *WakatimeHeartBeatDelete
}

// Where appends a list predicates to the WakatimeHeartBeatDelete builder.
func (whbdo *WakatimeHeartBeatDeleteOne) Where(ps ...predicate.WakatimeHeartBeat) *WakatimeHeartBeatDeleteOne {
	whbdo.whbd.mutation.Where(ps...)
	return whbdo
}

// Exec executes the deletion query.
func (whbdo *WakatimeHeartBeatDeleteOne) Exec(ctx context.Context) error {
	n, err := whbdo.whbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{wakatimeheartbeat.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (whbdo *WakatimeHeartBeatDeleteOne) ExecX(ctx context.Context) {
	if err := whbdo.Exec(ctx); err != nil {
		panic(err)
	}
}
