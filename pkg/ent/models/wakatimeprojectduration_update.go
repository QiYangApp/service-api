// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/wakatimeprojectduration"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WakatimeProjectDurationUpdate is the builder for updating WakatimeProjectDuration entities.
type WakatimeProjectDurationUpdate struct {
	config
	hooks    []Hook
	mutation *WakatimeProjectDurationMutation
}

// Where appends a list predicates to the WakatimeProjectDurationUpdate builder.
func (wpdu *WakatimeProjectDurationUpdate) Where(ps ...predicate.WakatimeProjectDuration) *WakatimeProjectDurationUpdate {
	wpdu.mutation.Where(ps...)
	return wpdu
}

// Mutation returns the WakatimeProjectDurationMutation object of the builder.
func (wpdu *WakatimeProjectDurationUpdate) Mutation() *WakatimeProjectDurationMutation {
	return wpdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wpdu *WakatimeProjectDurationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, wpdu.sqlSave, wpdu.mutation, wpdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wpdu *WakatimeProjectDurationUpdate) SaveX(ctx context.Context) int {
	affected, err := wpdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wpdu *WakatimeProjectDurationUpdate) Exec(ctx context.Context) error {
	_, err := wpdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wpdu *WakatimeProjectDurationUpdate) ExecX(ctx context.Context) {
	if err := wpdu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wpdu *WakatimeProjectDurationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(wakatimeprojectduration.Table, wakatimeprojectduration.Columns, sqlgraph.NewFieldSpec(wakatimeprojectduration.FieldID, field.TypeInt))
	if ps := wpdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wpdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wakatimeprojectduration.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wpdu.mutation.done = true
	return n, nil
}

// WakatimeProjectDurationUpdateOne is the builder for updating a single WakatimeProjectDuration entity.
type WakatimeProjectDurationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WakatimeProjectDurationMutation
}

// Mutation returns the WakatimeProjectDurationMutation object of the builder.
func (wpduo *WakatimeProjectDurationUpdateOne) Mutation() *WakatimeProjectDurationMutation {
	return wpduo.mutation
}

// Where appends a list predicates to the WakatimeProjectDurationUpdate builder.
func (wpduo *WakatimeProjectDurationUpdateOne) Where(ps ...predicate.WakatimeProjectDuration) *WakatimeProjectDurationUpdateOne {
	wpduo.mutation.Where(ps...)
	return wpduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wpduo *WakatimeProjectDurationUpdateOne) Select(field string, fields ...string) *WakatimeProjectDurationUpdateOne {
	wpduo.fields = append([]string{field}, fields...)
	return wpduo
}

// Save executes the query and returns the updated WakatimeProjectDuration entity.
func (wpduo *WakatimeProjectDurationUpdateOne) Save(ctx context.Context) (*WakatimeProjectDuration, error) {
	return withHooks(ctx, wpduo.sqlSave, wpduo.mutation, wpduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wpduo *WakatimeProjectDurationUpdateOne) SaveX(ctx context.Context) *WakatimeProjectDuration {
	node, err := wpduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wpduo *WakatimeProjectDurationUpdateOne) Exec(ctx context.Context) error {
	_, err := wpduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wpduo *WakatimeProjectDurationUpdateOne) ExecX(ctx context.Context) {
	if err := wpduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wpduo *WakatimeProjectDurationUpdateOne) sqlSave(ctx context.Context) (_node *WakatimeProjectDuration, err error) {
	_spec := sqlgraph.NewUpdateSpec(wakatimeprojectduration.Table, wakatimeprojectduration.Columns, sqlgraph.NewFieldSpec(wakatimeprojectduration.FieldID, field.TypeInt))
	id, ok := wpduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "WakatimeProjectDuration.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wpduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wakatimeprojectduration.FieldID)
		for _, f := range fields {
			if !wakatimeprojectduration.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != wakatimeprojectduration.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wpduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &WakatimeProjectDuration{config: wpduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wpduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wakatimeprojectduration.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wpduo.mutation.done = true
	return _node, nil
}
