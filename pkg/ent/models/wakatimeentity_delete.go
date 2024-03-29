// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/wakatimeentity"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WakatimeEntityDelete is the builder for deleting a WakatimeEntity entity.
type WakatimeEntityDelete struct {
	config
	hooks    []Hook
	mutation *WakatimeEntityMutation
}

// Where appends a list predicates to the WakatimeEntityDelete builder.
func (wed *WakatimeEntityDelete) Where(ps ...predicate.WakatimeEntity) *WakatimeEntityDelete {
	wed.mutation.Where(ps...)
	return wed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wed *WakatimeEntityDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wed.sqlExec, wed.mutation, wed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wed *WakatimeEntityDelete) ExecX(ctx context.Context) int {
	n, err := wed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wed *WakatimeEntityDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(wakatimeentity.Table, sqlgraph.NewFieldSpec(wakatimeentity.FieldID, field.TypeInt64))
	if ps := wed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wed.mutation.done = true
	return affected, err
}

// WakatimeEntityDeleteOne is the builder for deleting a single WakatimeEntity entity.
type WakatimeEntityDeleteOne struct {
	wed *WakatimeEntityDelete
}

// Where appends a list predicates to the WakatimeEntityDelete builder.
func (wedo *WakatimeEntityDeleteOne) Where(ps ...predicate.WakatimeEntity) *WakatimeEntityDeleteOne {
	wedo.wed.mutation.Where(ps...)
	return wedo
}

// Exec executes the deletion query.
func (wedo *WakatimeEntityDeleteOne) Exec(ctx context.Context) error {
	n, err := wedo.wed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{wakatimeentity.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wedo *WakatimeEntityDeleteOne) ExecX(ctx context.Context) {
	if err := wedo.Exec(ctx); err != nil {
		panic(err)
	}
}
