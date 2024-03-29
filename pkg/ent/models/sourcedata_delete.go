// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/sourcedata"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SourceDataDelete is the builder for deleting a SourceData entity.
type SourceDataDelete struct {
	config
	hooks    []Hook
	mutation *SourceDataMutation
}

// Where appends a list predicates to the SourceDataDelete builder.
func (sdd *SourceDataDelete) Where(ps ...predicate.SourceData) *SourceDataDelete {
	sdd.mutation.Where(ps...)
	return sdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sdd *SourceDataDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sdd.sqlExec, sdd.mutation, sdd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sdd *SourceDataDelete) ExecX(ctx context.Context) int {
	n, err := sdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sdd *SourceDataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sourcedata.Table, sqlgraph.NewFieldSpec(sourcedata.FieldID, field.TypeInt64))
	if ps := sdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sdd.mutation.done = true
	return affected, err
}

// SourceDataDeleteOne is the builder for deleting a single SourceData entity.
type SourceDataDeleteOne struct {
	sdd *SourceDataDelete
}

// Where appends a list predicates to the SourceDataDelete builder.
func (sddo *SourceDataDeleteOne) Where(ps ...predicate.SourceData) *SourceDataDeleteOne {
	sddo.sdd.mutation.Where(ps...)
	return sddo
}

// Exec executes the deletion query.
func (sddo *SourceDataDeleteOne) Exec(ctx context.Context) error {
	n, err := sddo.sdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sourcedata.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sddo *SourceDataDeleteOne) ExecX(ctx context.Context) {
	if err := sddo.Exec(ctx); err != nil {
		panic(err)
	}
}
