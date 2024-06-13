// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/wakatimeproject"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WakatimeProjectUpdate is the builder for updating WakatimeProject entities.
type WakatimeProjectUpdate struct {
	config
	hooks    []Hook
	mutation *WakatimeProjectMutation
}

// Where appends a list predicates to the WakatimeProjectUpdate builder.
func (wpu *WakatimeProjectUpdate) Where(ps ...predicate.WakatimeProject) *WakatimeProjectUpdate {
	wpu.mutation.Where(ps...)
	return wpu
}

// Mutation returns the WakatimeProjectMutation object of the builder.
func (wpu *WakatimeProjectUpdate) Mutation() *WakatimeProjectMutation {
	return wpu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wpu *WakatimeProjectUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, wpu.sqlSave, wpu.mutation, wpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wpu *WakatimeProjectUpdate) SaveX(ctx context.Context) int {
	affected, err := wpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wpu *WakatimeProjectUpdate) Exec(ctx context.Context) error {
	_, err := wpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wpu *WakatimeProjectUpdate) ExecX(ctx context.Context) {
	if err := wpu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wpu *WakatimeProjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(wakatimeproject.Table, wakatimeproject.Columns, sqlgraph.NewFieldSpec(wakatimeproject.FieldID, field.TypeInt64))
	if ps := wpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wakatimeproject.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wpu.mutation.done = true
	return n, nil
}

// WakatimeProjectUpdateOne is the builder for updating a single WakatimeProject entity.
type WakatimeProjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WakatimeProjectMutation
}

// Mutation returns the WakatimeProjectMutation object of the builder.
func (wpuo *WakatimeProjectUpdateOne) Mutation() *WakatimeProjectMutation {
	return wpuo.mutation
}

// Where appends a list predicates to the WakatimeProjectUpdate builder.
func (wpuo *WakatimeProjectUpdateOne) Where(ps ...predicate.WakatimeProject) *WakatimeProjectUpdateOne {
	wpuo.mutation.Where(ps...)
	return wpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wpuo *WakatimeProjectUpdateOne) Select(field string, fields ...string) *WakatimeProjectUpdateOne {
	wpuo.fields = append([]string{field}, fields...)
	return wpuo
}

// Save executes the query and returns the updated WakatimeProject entity.
func (wpuo *WakatimeProjectUpdateOne) Save(ctx context.Context) (*WakatimeProject, error) {
	return withHooks(ctx, wpuo.sqlSave, wpuo.mutation, wpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wpuo *WakatimeProjectUpdateOne) SaveX(ctx context.Context) *WakatimeProject {
	node, err := wpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wpuo *WakatimeProjectUpdateOne) Exec(ctx context.Context) error {
	_, err := wpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wpuo *WakatimeProjectUpdateOne) ExecX(ctx context.Context) {
	if err := wpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wpuo *WakatimeProjectUpdateOne) sqlSave(ctx context.Context) (_node *WakatimeProject, err error) {
	_spec := sqlgraph.NewUpdateSpec(wakatimeproject.Table, wakatimeproject.Columns, sqlgraph.NewFieldSpec(wakatimeproject.FieldID, field.TypeInt64))
	id, ok := wpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "WakatimeProject.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wakatimeproject.FieldID)
		for _, f := range fields {
			if !wakatimeproject.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != wakatimeproject.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &WakatimeProject{config: wpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wakatimeproject.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wpuo.mutation.done = true
	return _node, nil
}
