// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/userrelatedrole"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserRelatedRoleDelete is the builder for deleting a UserRelatedRole entity.
type UserRelatedRoleDelete struct {
	config
	hooks    []Hook
	mutation *UserRelatedRoleMutation
}

// Where appends a list predicates to the UserRelatedRoleDelete builder.
func (urrd *UserRelatedRoleDelete) Where(ps ...predicate.UserRelatedRole) *UserRelatedRoleDelete {
	urrd.mutation.Where(ps...)
	return urrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (urrd *UserRelatedRoleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, urrd.sqlExec, urrd.mutation, urrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (urrd *UserRelatedRoleDelete) ExecX(ctx context.Context) int {
	n, err := urrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (urrd *UserRelatedRoleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userrelatedrole.Table, sqlgraph.NewFieldSpec(userrelatedrole.FieldID, field.TypeInt64))
	if ps := urrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, urrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	urrd.mutation.done = true
	return affected, err
}

// UserRelatedRoleDeleteOne is the builder for deleting a single UserRelatedRole entity.
type UserRelatedRoleDeleteOne struct {
	urrd *UserRelatedRoleDelete
}

// Where appends a list predicates to the UserRelatedRoleDelete builder.
func (urrdo *UserRelatedRoleDeleteOne) Where(ps ...predicate.UserRelatedRole) *UserRelatedRoleDeleteOne {
	urrdo.urrd.mutation.Where(ps...)
	return urrdo
}

// Exec executes the deletion query.
func (urrdo *UserRelatedRoleDeleteOne) Exec(ctx context.Context) error {
	n, err := urrdo.urrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userrelatedrole.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (urrdo *UserRelatedRoleDeleteOne) ExecX(ctx context.Context) {
	if err := urrdo.Exec(ctx); err != nil {
		panic(err)
	}
}