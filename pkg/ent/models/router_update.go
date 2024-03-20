// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/enums/state"
	"ent/models/predicate"
	"ent/models/router"
	"ent/utils/timeutil"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RouterUpdate is the builder for updating Router entities.
type RouterUpdate struct {
	config
	hooks    []Hook
	mutation *RouterMutation
}

// Where appends a list predicates to the RouterUpdate builder.
func (ru *RouterUpdate) Where(ps ...predicate.Router) *RouterUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetRouteName sets the "route_name" field.
func (ru *RouterUpdate) SetRouteName(s string) *RouterUpdate {
	ru.mutation.SetRouteName(s)
	return ru
}

// SetNillableRouteName sets the "route_name" field if the given value is not nil.
func (ru *RouterUpdate) SetNillableRouteName(s *string) *RouterUpdate {
	if s != nil {
		ru.SetRouteName(*s)
	}
	return ru
}

// SetRoute sets the "route" field.
func (ru *RouterUpdate) SetRoute(s string) *RouterUpdate {
	ru.mutation.SetRoute(s)
	return ru
}

// SetNillableRoute sets the "route" field if the given value is not nil.
func (ru *RouterUpdate) SetNillableRoute(s *string) *RouterUpdate {
	if s != nil {
		ru.SetRoute(*s)
	}
	return ru
}

// SetDescription sets the "description" field.
func (ru *RouterUpdate) SetDescription(s string) *RouterUpdate {
	ru.mutation.SetDescription(s)
	return ru
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ru *RouterUpdate) SetNillableDescription(s *string) *RouterUpdate {
	if s != nil {
		ru.SetDescription(*s)
	}
	return ru
}

// SetState sets the "state" field.
func (ru *RouterUpdate) SetState(ss state.SwitchState) *RouterUpdate {
	ru.mutation.ResetState()
	ru.mutation.SetState(ss)
	return ru
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ru *RouterUpdate) SetNillableState(ss *state.SwitchState) *RouterUpdate {
	if ss != nil {
		ru.SetState(*ss)
	}
	return ru
}

// AddState adds ss to the "state" field.
func (ru *RouterUpdate) AddState(ss state.SwitchState) *RouterUpdate {
	ru.mutation.AddState(ss)
	return ru
}

// SetUpdateTime sets the "update_time" field.
func (ru *RouterUpdate) SetUpdateTime(ts timeutil.TimeStamp) *RouterUpdate {
	ru.mutation.ResetUpdateTime()
	ru.mutation.SetUpdateTime(ts)
	return ru
}

// AddUpdateTime adds ts to the "update_time" field.
func (ru *RouterUpdate) AddUpdateTime(ts timeutil.TimeStamp) *RouterUpdate {
	ru.mutation.AddUpdateTime(ts)
	return ru
}

// Mutation returns the RouterMutation object of the builder.
func (ru *RouterUpdate) Mutation() *RouterMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RouterUpdate) Save(ctx context.Context) (int, error) {
	ru.defaults()
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RouterUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RouterUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RouterUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RouterUpdate) defaults() {
	if _, ok := ru.mutation.UpdateTime(); !ok {
		v := router.UpdateDefaultUpdateTime()
		ru.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RouterUpdate) check() error {
	if v, ok := ru.mutation.RouteName(); ok {
		if err := router.RouteNameValidator(v); err != nil {
			return &ValidationError{Name: "route_name", err: fmt.Errorf(`models: validator failed for field "Router.route_name": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Route(); ok {
		if err := router.RouteValidator(v); err != nil {
			return &ValidationError{Name: "route", err: fmt.Errorf(`models: validator failed for field "Router.route": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Description(); ok {
		if err := router.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`models: validator failed for field "Router.description": %w`, err)}
		}
	}
	return nil
}

func (ru *RouterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(router.Table, router.Columns, sqlgraph.NewFieldSpec(router.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.RouteName(); ok {
		_spec.SetField(router.FieldRouteName, field.TypeString, value)
	}
	if value, ok := ru.mutation.Route(); ok {
		_spec.SetField(router.FieldRoute, field.TypeString, value)
	}
	if value, ok := ru.mutation.Description(); ok {
		_spec.SetField(router.FieldDescription, field.TypeString, value)
	}
	if value, ok := ru.mutation.State(); ok {
		_spec.SetField(router.FieldState, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedState(); ok {
		_spec.AddField(router.FieldState, field.TypeInt, value)
	}
	if value, ok := ru.mutation.UpdateTime(); ok {
		_spec.SetField(router.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.AddedUpdateTime(); ok {
		_spec.AddField(router.FieldUpdateTime, field.TypeInt64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{router.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RouterUpdateOne is the builder for updating a single Router entity.
type RouterUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RouterMutation
}

// SetRouteName sets the "route_name" field.
func (ruo *RouterUpdateOne) SetRouteName(s string) *RouterUpdateOne {
	ruo.mutation.SetRouteName(s)
	return ruo
}

// SetNillableRouteName sets the "route_name" field if the given value is not nil.
func (ruo *RouterUpdateOne) SetNillableRouteName(s *string) *RouterUpdateOne {
	if s != nil {
		ruo.SetRouteName(*s)
	}
	return ruo
}

// SetRoute sets the "route" field.
func (ruo *RouterUpdateOne) SetRoute(s string) *RouterUpdateOne {
	ruo.mutation.SetRoute(s)
	return ruo
}

// SetNillableRoute sets the "route" field if the given value is not nil.
func (ruo *RouterUpdateOne) SetNillableRoute(s *string) *RouterUpdateOne {
	if s != nil {
		ruo.SetRoute(*s)
	}
	return ruo
}

// SetDescription sets the "description" field.
func (ruo *RouterUpdateOne) SetDescription(s string) *RouterUpdateOne {
	ruo.mutation.SetDescription(s)
	return ruo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ruo *RouterUpdateOne) SetNillableDescription(s *string) *RouterUpdateOne {
	if s != nil {
		ruo.SetDescription(*s)
	}
	return ruo
}

// SetState sets the "state" field.
func (ruo *RouterUpdateOne) SetState(ss state.SwitchState) *RouterUpdateOne {
	ruo.mutation.ResetState()
	ruo.mutation.SetState(ss)
	return ruo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ruo *RouterUpdateOne) SetNillableState(ss *state.SwitchState) *RouterUpdateOne {
	if ss != nil {
		ruo.SetState(*ss)
	}
	return ruo
}

// AddState adds ss to the "state" field.
func (ruo *RouterUpdateOne) AddState(ss state.SwitchState) *RouterUpdateOne {
	ruo.mutation.AddState(ss)
	return ruo
}

// SetUpdateTime sets the "update_time" field.
func (ruo *RouterUpdateOne) SetUpdateTime(ts timeutil.TimeStamp) *RouterUpdateOne {
	ruo.mutation.ResetUpdateTime()
	ruo.mutation.SetUpdateTime(ts)
	return ruo
}

// AddUpdateTime adds ts to the "update_time" field.
func (ruo *RouterUpdateOne) AddUpdateTime(ts timeutil.TimeStamp) *RouterUpdateOne {
	ruo.mutation.AddUpdateTime(ts)
	return ruo
}

// Mutation returns the RouterMutation object of the builder.
func (ruo *RouterUpdateOne) Mutation() *RouterMutation {
	return ruo.mutation
}

// Where appends a list predicates to the RouterUpdate builder.
func (ruo *RouterUpdateOne) Where(ps ...predicate.Router) *RouterUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RouterUpdateOne) Select(field string, fields ...string) *RouterUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Router entity.
func (ruo *RouterUpdateOne) Save(ctx context.Context) (*Router, error) {
	ruo.defaults()
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RouterUpdateOne) SaveX(ctx context.Context) *Router {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RouterUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RouterUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RouterUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdateTime(); !ok {
		v := router.UpdateDefaultUpdateTime()
		ruo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RouterUpdateOne) check() error {
	if v, ok := ruo.mutation.RouteName(); ok {
		if err := router.RouteNameValidator(v); err != nil {
			return &ValidationError{Name: "route_name", err: fmt.Errorf(`models: validator failed for field "Router.route_name": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Route(); ok {
		if err := router.RouteValidator(v); err != nil {
			return &ValidationError{Name: "route", err: fmt.Errorf(`models: validator failed for field "Router.route": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Description(); ok {
		if err := router.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`models: validator failed for field "Router.description": %w`, err)}
		}
	}
	return nil
}

func (ruo *RouterUpdateOne) sqlSave(ctx context.Context) (_node *Router, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(router.Table, router.Columns, sqlgraph.NewFieldSpec(router.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "Router.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, router.FieldID)
		for _, f := range fields {
			if !router.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != router.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.RouteName(); ok {
		_spec.SetField(router.FieldRouteName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Route(); ok {
		_spec.SetField(router.FieldRoute, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Description(); ok {
		_spec.SetField(router.FieldDescription, field.TypeString, value)
	}
	if value, ok := ruo.mutation.State(); ok {
		_spec.SetField(router.FieldState, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedState(); ok {
		_spec.AddField(router.FieldState, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.UpdateTime(); ok {
		_spec.SetField(router.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.AddedUpdateTime(); ok {
		_spec.AddField(router.FieldUpdateTime, field.TypeInt64, value)
	}
	_node = &Router{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{router.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
