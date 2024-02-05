// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"service-api/src/models/ent/memberauthorizelog"
	"service-api/src/models/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberAuthorizeLogDelete is the builder for deleting a MemberAuthorizeLog entity.
type MemberAuthorizeLogDelete struct {
	config
	hooks    []Hook
	mutation *MemberAuthorizeLogMutation
}

// Where appends a list predicates to the MemberAuthorizeLogDelete builder.
func (mald *MemberAuthorizeLogDelete) Where(ps ...predicate.MemberAuthorizeLog) *MemberAuthorizeLogDelete {
	mald.mutation.Where(ps...)
	return mald
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mald *MemberAuthorizeLogDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, mald.sqlExec, mald.mutation, mald.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mald *MemberAuthorizeLogDelete) ExecX(ctx context.Context) int {
	n, err := mald.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mald *MemberAuthorizeLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(memberauthorizelog.Table, sqlgraph.NewFieldSpec(memberauthorizelog.FieldID, field.TypeUUID))
	if ps := mald.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mald.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mald.mutation.done = true
	return affected, err
}

// MemberAuthorizeLogDeleteOne is the builder for deleting a single MemberAuthorizeLog entity.
type MemberAuthorizeLogDeleteOne struct {
	mald *MemberAuthorizeLogDelete
}

// Where appends a list predicates to the MemberAuthorizeLogDelete builder.
func (maldo *MemberAuthorizeLogDeleteOne) Where(ps ...predicate.MemberAuthorizeLog) *MemberAuthorizeLogDeleteOne {
	maldo.mald.mutation.Where(ps...)
	return maldo
}

// Exec executes the deletion query.
func (maldo *MemberAuthorizeLogDeleteOne) Exec(ctx context.Context) error {
	n, err := maldo.mald.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{memberauthorizelog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (maldo *MemberAuthorizeLogDeleteOne) ExecX(ctx context.Context) {
	if err := maldo.Exec(ctx); err != nil {
		panic(err)
	}
}
