// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/wakatimeprojectinfo"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WakatimeProjectInfoDelete is the builder for deleting a WakatimeProjectInfo entity.
type WakatimeProjectInfoDelete struct {
	config
	hooks    []Hook
	mutation *WakatimeProjectInfoMutation
}

// Where appends a list predicates to the WakatimeProjectInfoDelete builder.
func (wpid *WakatimeProjectInfoDelete) Where(ps ...predicate.WakatimeProjectInfo) *WakatimeProjectInfoDelete {
	wpid.mutation.Where(ps...)
	return wpid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wpid *WakatimeProjectInfoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wpid.sqlExec, wpid.mutation, wpid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wpid *WakatimeProjectInfoDelete) ExecX(ctx context.Context) int {
	n, err := wpid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wpid *WakatimeProjectInfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(wakatimeprojectinfo.Table, sqlgraph.NewFieldSpec(wakatimeprojectinfo.FieldID, field.TypeInt64))
	if ps := wpid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wpid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wpid.mutation.done = true
	return affected, err
}

// WakatimeProjectInfoDeleteOne is the builder for deleting a single WakatimeProjectInfo entity.
type WakatimeProjectInfoDeleteOne struct {
	wpid *WakatimeProjectInfoDelete
}

// Where appends a list predicates to the WakatimeProjectInfoDelete builder.
func (wpido *WakatimeProjectInfoDeleteOne) Where(ps ...predicate.WakatimeProjectInfo) *WakatimeProjectInfoDeleteOne {
	wpido.wpid.mutation.Where(ps...)
	return wpido
}

// Exec executes the deletion query.
func (wpido *WakatimeProjectInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := wpido.wpid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{wakatimeprojectinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wpido *WakatimeProjectInfoDeleteOne) ExecX(ctx context.Context) {
	if err := wpido.Exec(ctx); err != nil {
		panic(err)
	}
}
