// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"service-api/src/models/ent/router"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RouterCreate is the builder for creating a Router entity.
type RouterCreate struct {
	config
	mutation *RouterMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (rc *RouterCreate) SetCreateTime(t time.Time) *RouterCreate {
	rc.mutation.SetCreateTime(t)
	return rc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (rc *RouterCreate) SetNillableCreateTime(t *time.Time) *RouterCreate {
	if t != nil {
		rc.SetCreateTime(*t)
	}
	return rc
}

// SetUpdateTime sets the "update_time" field.
func (rc *RouterCreate) SetUpdateTime(t time.Time) *RouterCreate {
	rc.mutation.SetUpdateTime(t)
	return rc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (rc *RouterCreate) SetNillableUpdateTime(t *time.Time) *RouterCreate {
	if t != nil {
		rc.SetUpdateTime(*t)
	}
	return rc
}

// SetRouteName sets the "route_name" field.
func (rc *RouterCreate) SetRouteName(s string) *RouterCreate {
	rc.mutation.SetRouteName(s)
	return rc
}

// SetNillableRouteName sets the "route_name" field if the given value is not nil.
func (rc *RouterCreate) SetNillableRouteName(s *string) *RouterCreate {
	if s != nil {
		rc.SetRouteName(*s)
	}
	return rc
}

// SetRoute sets the "route" field.
func (rc *RouterCreate) SetRoute(s string) *RouterCreate {
	rc.mutation.SetRoute(s)
	return rc
}

// SetNillableRoute sets the "route" field if the given value is not nil.
func (rc *RouterCreate) SetNillableRoute(s *string) *RouterCreate {
	if s != nil {
		rc.SetRoute(*s)
	}
	return rc
}

// SetDescription sets the "description" field.
func (rc *RouterCreate) SetDescription(s string) *RouterCreate {
	rc.mutation.SetDescription(s)
	return rc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rc *RouterCreate) SetNillableDescription(s *string) *RouterCreate {
	if s != nil {
		rc.SetDescription(*s)
	}
	return rc
}

// SetState sets the "state" field.
func (rc *RouterCreate) SetState(s string) *RouterCreate {
	rc.mutation.SetState(s)
	return rc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (rc *RouterCreate) SetNillableState(s *string) *RouterCreate {
	if s != nil {
		rc.SetState(*s)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *RouterCreate) SetID(u uuid.UUID) *RouterCreate {
	rc.mutation.SetID(u)
	return rc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rc *RouterCreate) SetNillableID(u *uuid.UUID) *RouterCreate {
	if u != nil {
		rc.SetID(*u)
	}
	return rc
}

// Mutation returns the RouterMutation object of the builder.
func (rc *RouterCreate) Mutation() *RouterMutation {
	return rc.mutation
}

// Save creates the Router in the database.
func (rc *RouterCreate) Save(ctx context.Context) (*Router, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RouterCreate) SaveX(ctx context.Context) *Router {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RouterCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RouterCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RouterCreate) defaults() {
	if _, ok := rc.mutation.CreateTime(); !ok {
		v := router.DefaultCreateTime()
		rc.mutation.SetCreateTime(v)
	}
	if _, ok := rc.mutation.UpdateTime(); !ok {
		v := router.DefaultUpdateTime()
		rc.mutation.SetUpdateTime(v)
	}
	if _, ok := rc.mutation.RouteName(); !ok {
		v := router.DefaultRouteName
		rc.mutation.SetRouteName(v)
	}
	if _, ok := rc.mutation.Route(); !ok {
		v := router.DefaultRoute
		rc.mutation.SetRoute(v)
	}
	if _, ok := rc.mutation.Description(); !ok {
		v := router.DefaultDescription
		rc.mutation.SetDescription(v)
	}
	if _, ok := rc.mutation.State(); !ok {
		v := router.DefaultState
		rc.mutation.SetState(v)
	}
	if _, ok := rc.mutation.ID(); !ok {
		v := router.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RouterCreate) check() error {
	if _, ok := rc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Router.create_time"`)}
	}
	if _, ok := rc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Router.update_time"`)}
	}
	if _, ok := rc.mutation.RouteName(); !ok {
		return &ValidationError{Name: "route_name", err: errors.New(`ent: missing required field "Router.route_name"`)}
	}
	if v, ok := rc.mutation.RouteName(); ok {
		if err := router.RouteNameValidator(v); err != nil {
			return &ValidationError{Name: "route_name", err: fmt.Errorf(`ent: validator failed for field "Router.route_name": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Route(); !ok {
		return &ValidationError{Name: "route", err: errors.New(`ent: missing required field "Router.route"`)}
	}
	if v, ok := rc.mutation.Route(); ok {
		if err := router.RouteValidator(v); err != nil {
			return &ValidationError{Name: "route", err: fmt.Errorf(`ent: validator failed for field "Router.route": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Router.description"`)}
	}
	if v, ok := rc.mutation.Description(); ok {
		if err := router.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Router.description": %w`, err)}
		}
	}
	if _, ok := rc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "Router.state"`)}
	}
	if v, ok := rc.mutation.State(); ok {
		if err := router.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Router.state": %w`, err)}
		}
	}
	return nil
}

func (rc *RouterCreate) sqlSave(ctx context.Context) (*Router, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RouterCreate) createSpec() (*Router, *sqlgraph.CreateSpec) {
	var (
		_node = &Router{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(router.Table, sqlgraph.NewFieldSpec(router.FieldID, field.TypeUUID))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rc.mutation.CreateTime(); ok {
		_spec.SetField(router.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := rc.mutation.UpdateTime(); ok {
		_spec.SetField(router.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := rc.mutation.RouteName(); ok {
		_spec.SetField(router.FieldRouteName, field.TypeString, value)
		_node.RouteName = value
	}
	if value, ok := rc.mutation.Route(); ok {
		_spec.SetField(router.FieldRoute, field.TypeString, value)
		_node.Route = value
	}
	if value, ok := rc.mutation.Description(); ok {
		_spec.SetField(router.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := rc.mutation.State(); ok {
		_spec.SetField(router.FieldState, field.TypeString, value)
		_node.State = value
	}
	return _node, _spec
}

// RouterCreateBulk is the builder for creating many Router entities in bulk.
type RouterCreateBulk struct {
	config
	err      error
	builders []*RouterCreate
}

// Save creates the Router entities in the database.
func (rcb *RouterCreateBulk) Save(ctx context.Context) ([]*Router, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Router, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RouterMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RouterCreateBulk) SaveX(ctx context.Context) []*Router {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RouterCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RouterCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
