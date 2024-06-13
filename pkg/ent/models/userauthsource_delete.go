// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/userauthsource"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserAuthSourceDelete is the builder for deleting a UserAuthSource entity.
type UserAuthSourceDelete struct {
	config
	hooks    []Hook
	mutation *UserAuthSourceMutation
}

// Where appends a list predicates to the UserAuthSourceDelete builder.
func (uasd *UserAuthSourceDelete) Where(ps ...predicate.UserAuthSource) *UserAuthSourceDelete {
	uasd.mutation.Where(ps...)
	return uasd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uasd *UserAuthSourceDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, uasd.sqlExec, uasd.mutation, uasd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (uasd *UserAuthSourceDelete) ExecX(ctx context.Context) int {
	n, err := uasd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uasd *UserAuthSourceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userauthsource.Table, sqlgraph.NewFieldSpec(userauthsource.FieldID, field.TypeInt64))
	if ps := uasd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, uasd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	uasd.mutation.done = true
	return affected, err
}

// UserAuthSourceDeleteOne is the builder for deleting a single UserAuthSource entity.
type UserAuthSourceDeleteOne struct {
	uasd *UserAuthSourceDelete
}

// Where appends a list predicates to the UserAuthSourceDelete builder.
func (uasdo *UserAuthSourceDeleteOne) Where(ps ...predicate.UserAuthSource) *UserAuthSourceDeleteOne {
	uasdo.uasd.mutation.Where(ps...)
	return uasdo
}

// Exec executes the deletion query.
func (uasdo *UserAuthSourceDeleteOne) Exec(ctx context.Context) error {
	n, err := uasdo.uasd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userauthsource.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uasdo *UserAuthSourceDeleteOne) ExecX(ctx context.Context) {
	if err := uasdo.Exec(ctx); err != nil {
		panic(err)
	}
}
