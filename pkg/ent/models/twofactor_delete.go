// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/twofactor"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TwoFactorDelete is the builder for deleting a TwoFactor entity.
type TwoFactorDelete struct {
	config
	hooks    []Hook
	mutation *TwoFactorMutation
}

// Where appends a list predicates to the TwoFactorDelete builder.
func (tfd *TwoFactorDelete) Where(ps ...predicate.TwoFactor) *TwoFactorDelete {
	tfd.mutation.Where(ps...)
	return tfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tfd *TwoFactorDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tfd.sqlExec, tfd.mutation, tfd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tfd *TwoFactorDelete) ExecX(ctx context.Context) int {
	n, err := tfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tfd *TwoFactorDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(twofactor.Table, sqlgraph.NewFieldSpec(twofactor.FieldID, field.TypeInt64))
	if ps := tfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tfd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tfd.mutation.done = true
	return affected, err
}

// TwoFactorDeleteOne is the builder for deleting a single TwoFactor entity.
type TwoFactorDeleteOne struct {
	tfd *TwoFactorDelete
}

// Where appends a list predicates to the TwoFactorDelete builder.
func (tfdo *TwoFactorDeleteOne) Where(ps ...predicate.TwoFactor) *TwoFactorDeleteOne {
	tfdo.tfd.mutation.Where(ps...)
	return tfdo
}

// Exec executes the deletion query.
func (tfdo *TwoFactorDeleteOne) Exec(ctx context.Context) error {
	n, err := tfdo.tfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{twofactor.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tfdo *TwoFactorDeleteOne) ExecX(ctx context.Context) {
	if err := tfdo.Exec(ctx); err != nil {
		panic(err)
	}
}
