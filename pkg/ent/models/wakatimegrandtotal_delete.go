// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/wakatimegrandtotal"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WakatimeGrandTotalDelete is the builder for deleting a WakatimeGrandTotal entity.
type WakatimeGrandTotalDelete struct {
	config
	hooks    []Hook
	mutation *WakatimeGrandTotalMutation
}

// Where appends a list predicates to the WakatimeGrandTotalDelete builder.
func (wgtd *WakatimeGrandTotalDelete) Where(ps ...predicate.WakatimeGrandTotal) *WakatimeGrandTotalDelete {
	wgtd.mutation.Where(ps...)
	return wgtd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wgtd *WakatimeGrandTotalDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wgtd.sqlExec, wgtd.mutation, wgtd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wgtd *WakatimeGrandTotalDelete) ExecX(ctx context.Context) int {
	n, err := wgtd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wgtd *WakatimeGrandTotalDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(wakatimegrandtotal.Table, sqlgraph.NewFieldSpec(wakatimegrandtotal.FieldID, field.TypeInt))
	if ps := wgtd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wgtd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wgtd.mutation.done = true
	return affected, err
}

// WakatimeGrandTotalDeleteOne is the builder for deleting a single WakatimeGrandTotal entity.
type WakatimeGrandTotalDeleteOne struct {
	wgtd *WakatimeGrandTotalDelete
}

// Where appends a list predicates to the WakatimeGrandTotalDelete builder.
func (wgtdo *WakatimeGrandTotalDeleteOne) Where(ps ...predicate.WakatimeGrandTotal) *WakatimeGrandTotalDeleteOne {
	wgtdo.wgtd.mutation.Where(ps...)
	return wgtdo
}

// Exec executes the deletion query.
func (wgtdo *WakatimeGrandTotalDeleteOne) Exec(ctx context.Context) error {
	n, err := wgtdo.wgtd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{wakatimegrandtotal.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wgtdo *WakatimeGrandTotalDeleteOne) ExecX(ctx context.Context) {
	if err := wgtdo.Exec(ctx); err != nil {
		panic(err)
	}
}
