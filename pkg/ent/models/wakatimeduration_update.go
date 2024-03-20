// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/wakatimeduration"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WakatimeDurationUpdate is the builder for updating WakatimeDuration entities.
type WakatimeDurationUpdate struct {
	config
	hooks    []Hook
	mutation *WakatimeDurationMutation
}

// Where appends a list predicates to the WakatimeDurationUpdate builder.
func (wdu *WakatimeDurationUpdate) Where(ps ...predicate.WakatimeDuration) *WakatimeDurationUpdate {
	wdu.mutation.Where(ps...)
	return wdu
}

// Mutation returns the WakatimeDurationMutation object of the builder.
func (wdu *WakatimeDurationUpdate) Mutation() *WakatimeDurationMutation {
	return wdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wdu *WakatimeDurationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, wdu.sqlSave, wdu.mutation, wdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wdu *WakatimeDurationUpdate) SaveX(ctx context.Context) int {
	affected, err := wdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wdu *WakatimeDurationUpdate) Exec(ctx context.Context) error {
	_, err := wdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wdu *WakatimeDurationUpdate) ExecX(ctx context.Context) {
	if err := wdu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wdu *WakatimeDurationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(wakatimeduration.Table, wakatimeduration.Columns, sqlgraph.NewFieldSpec(wakatimeduration.FieldID, field.TypeInt))
	if ps := wdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wakatimeduration.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wdu.mutation.done = true
	return n, nil
}

// WakatimeDurationUpdateOne is the builder for updating a single WakatimeDuration entity.
type WakatimeDurationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WakatimeDurationMutation
}

// Mutation returns the WakatimeDurationMutation object of the builder.
func (wduo *WakatimeDurationUpdateOne) Mutation() *WakatimeDurationMutation {
	return wduo.mutation
}

// Where appends a list predicates to the WakatimeDurationUpdate builder.
func (wduo *WakatimeDurationUpdateOne) Where(ps ...predicate.WakatimeDuration) *WakatimeDurationUpdateOne {
	wduo.mutation.Where(ps...)
	return wduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wduo *WakatimeDurationUpdateOne) Select(field string, fields ...string) *WakatimeDurationUpdateOne {
	wduo.fields = append([]string{field}, fields...)
	return wduo
}

// Save executes the query and returns the updated WakatimeDuration entity.
func (wduo *WakatimeDurationUpdateOne) Save(ctx context.Context) (*WakatimeDuration, error) {
	return withHooks(ctx, wduo.sqlSave, wduo.mutation, wduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wduo *WakatimeDurationUpdateOne) SaveX(ctx context.Context) *WakatimeDuration {
	node, err := wduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wduo *WakatimeDurationUpdateOne) Exec(ctx context.Context) error {
	_, err := wduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wduo *WakatimeDurationUpdateOne) ExecX(ctx context.Context) {
	if err := wduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wduo *WakatimeDurationUpdateOne) sqlSave(ctx context.Context) (_node *WakatimeDuration, err error) {
	_spec := sqlgraph.NewUpdateSpec(wakatimeduration.Table, wakatimeduration.Columns, sqlgraph.NewFieldSpec(wakatimeduration.FieldID, field.TypeInt))
	id, ok := wduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "WakatimeDuration.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wakatimeduration.FieldID)
		for _, f := range fields {
			if !wakatimeduration.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != wakatimeduration.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &WakatimeDuration{config: wduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wakatimeduration.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wduo.mutation.done = true
	return _node, nil
}
