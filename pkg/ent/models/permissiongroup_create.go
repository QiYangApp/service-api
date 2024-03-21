// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/enums/state"
	"ent/models/permissiongroup"
	"ent/utils/timeutil"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PermissionGroupCreate is the builder for creating a PermissionGroup entity.
type PermissionGroupCreate struct {
	config
	mutation *PermissionGroupMutation
	hooks    []Hook
}

// SetPermissionName sets the "permission_name" field.
func (pgc *PermissionGroupCreate) SetPermissionName(s string) *PermissionGroupCreate {
	pgc.mutation.SetPermissionName(s)
	return pgc
}

// SetNillablePermissionName sets the "permission_name" field if the given value is not nil.
func (pgc *PermissionGroupCreate) SetNillablePermissionName(s *string) *PermissionGroupCreate {
	if s != nil {
		pgc.SetPermissionName(*s)
	}
	return pgc
}

// SetIoc sets the "ioc" field.
func (pgc *PermissionGroupCreate) SetIoc(s string) *PermissionGroupCreate {
	pgc.mutation.SetIoc(s)
	return pgc
}

// SetNillableIoc sets the "ioc" field if the given value is not nil.
func (pgc *PermissionGroupCreate) SetNillableIoc(s *string) *PermissionGroupCreate {
	if s != nil {
		pgc.SetIoc(*s)
	}
	return pgc
}

// SetSort sets the "sort" field.
func (pgc *PermissionGroupCreate) SetSort(i int32) *PermissionGroupCreate {
	pgc.mutation.SetSort(i)
	return pgc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (pgc *PermissionGroupCreate) SetNillableSort(i *int32) *PermissionGroupCreate {
	if i != nil {
		pgc.SetSort(*i)
	}
	return pgc
}

// SetLeft sets the "left" field.
func (pgc *PermissionGroupCreate) SetLeft(i int32) *PermissionGroupCreate {
	pgc.mutation.SetLeft(i)
	return pgc
}

// SetNillableLeft sets the "left" field if the given value is not nil.
func (pgc *PermissionGroupCreate) SetNillableLeft(i *int32) *PermissionGroupCreate {
	if i != nil {
		pgc.SetLeft(*i)
	}
	return pgc
}

// SetRight sets the "right" field.
func (pgc *PermissionGroupCreate) SetRight(i int32) *PermissionGroupCreate {
	pgc.mutation.SetRight(i)
	return pgc
}

// SetNillableRight sets the "right" field if the given value is not nil.
func (pgc *PermissionGroupCreate) SetNillableRight(i *int32) *PermissionGroupCreate {
	if i != nil {
		pgc.SetRight(*i)
	}
	return pgc
}

// SetState sets the "state" field.
func (pgc *PermissionGroupCreate) SetState(ss state.SwitchState) *PermissionGroupCreate {
	pgc.mutation.SetState(ss)
	return pgc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (pgc *PermissionGroupCreate) SetNillableState(ss *state.SwitchState) *PermissionGroupCreate {
	if ss != nil {
		pgc.SetState(*ss)
	}
	return pgc
}

// SetCreateTime sets the "create_time" field.
func (pgc *PermissionGroupCreate) SetCreateTime(ts timeutil.TimeStamp) *PermissionGroupCreate {
	pgc.mutation.SetCreateTime(ts)
	return pgc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pgc *PermissionGroupCreate) SetNillableCreateTime(ts *timeutil.TimeStamp) *PermissionGroupCreate {
	if ts != nil {
		pgc.SetCreateTime(*ts)
	}
	return pgc
}

// SetUpdateTime sets the "update_time" field.
func (pgc *PermissionGroupCreate) SetUpdateTime(ts timeutil.TimeStamp) *PermissionGroupCreate {
	pgc.mutation.SetUpdateTime(ts)
	return pgc
}

// SetID sets the "id" field.
func (pgc *PermissionGroupCreate) SetID(i int64) *PermissionGroupCreate {
	pgc.mutation.SetID(i)
	return pgc
}

// Mutation returns the PermissionGroupMutation object of the builder.
func (pgc *PermissionGroupCreate) Mutation() *PermissionGroupMutation {
	return pgc.mutation
}

// Save creates the PermissionGroup in the database.
func (pgc *PermissionGroupCreate) Save(ctx context.Context) (*PermissionGroup, error) {
	pgc.defaults()
	return withHooks(ctx, pgc.sqlSave, pgc.mutation, pgc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pgc *PermissionGroupCreate) SaveX(ctx context.Context) *PermissionGroup {
	v, err := pgc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pgc *PermissionGroupCreate) Exec(ctx context.Context) error {
	_, err := pgc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pgc *PermissionGroupCreate) ExecX(ctx context.Context) {
	if err := pgc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pgc *PermissionGroupCreate) defaults() {
	if _, ok := pgc.mutation.PermissionName(); !ok {
		v := permissiongroup.DefaultPermissionName
		pgc.mutation.SetPermissionName(v)
	}
	if _, ok := pgc.mutation.Ioc(); !ok {
		v := permissiongroup.DefaultIoc
		pgc.mutation.SetIoc(v)
	}
	if _, ok := pgc.mutation.Sort(); !ok {
		v := permissiongroup.DefaultSort
		pgc.mutation.SetSort(v)
	}
	if _, ok := pgc.mutation.Left(); !ok {
		v := permissiongroup.DefaultLeft
		pgc.mutation.SetLeft(v)
	}
	if _, ok := pgc.mutation.Right(); !ok {
		v := permissiongroup.DefaultRight
		pgc.mutation.SetRight(v)
	}
	if _, ok := pgc.mutation.State(); !ok {
		v := permissiongroup.DefaultState
		pgc.mutation.SetState(v)
	}
	if _, ok := pgc.mutation.CreateTime(); !ok {
		v := permissiongroup.DefaultCreateTime
		pgc.mutation.SetCreateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pgc *PermissionGroupCreate) check() error {
	if _, ok := pgc.mutation.PermissionName(); !ok {
		return &ValidationError{Name: "permission_name", err: errors.New(`models: missing required field "PermissionGroup.permission_name"`)}
	}
	if v, ok := pgc.mutation.PermissionName(); ok {
		if err := permissiongroup.PermissionNameValidator(v); err != nil {
			return &ValidationError{Name: "permission_name", err: fmt.Errorf(`models: validator failed for field "PermissionGroup.permission_name": %w`, err)}
		}
	}
	if _, ok := pgc.mutation.Ioc(); !ok {
		return &ValidationError{Name: "ioc", err: errors.New(`models: missing required field "PermissionGroup.ioc"`)}
	}
	if v, ok := pgc.mutation.Ioc(); ok {
		if err := permissiongroup.IocValidator(v); err != nil {
			return &ValidationError{Name: "ioc", err: fmt.Errorf(`models: validator failed for field "PermissionGroup.ioc": %w`, err)}
		}
	}
	if _, ok := pgc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`models: missing required field "PermissionGroup.sort"`)}
	}
	if _, ok := pgc.mutation.Left(); !ok {
		return &ValidationError{Name: "left", err: errors.New(`models: missing required field "PermissionGroup.left"`)}
	}
	if _, ok := pgc.mutation.Right(); !ok {
		return &ValidationError{Name: "right", err: errors.New(`models: missing required field "PermissionGroup.right"`)}
	}
	if _, ok := pgc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`models: missing required field "PermissionGroup.state"`)}
	}
	if _, ok := pgc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`models: missing required field "PermissionGroup.create_time"`)}
	}
	if _, ok := pgc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`models: missing required field "PermissionGroup.update_time"`)}
	}
	return nil
}

func (pgc *PermissionGroupCreate) sqlSave(ctx context.Context) (*PermissionGroup, error) {
	if err := pgc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pgc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pgc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	pgc.mutation.id = &_node.ID
	pgc.mutation.done = true
	return _node, nil
}

func (pgc *PermissionGroupCreate) createSpec() (*PermissionGroup, *sqlgraph.CreateSpec) {
	var (
		_node = &PermissionGroup{config: pgc.config}
		_spec = sqlgraph.NewCreateSpec(permissiongroup.Table, sqlgraph.NewFieldSpec(permissiongroup.FieldID, field.TypeInt64))
	)
	if id, ok := pgc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pgc.mutation.PermissionName(); ok {
		_spec.SetField(permissiongroup.FieldPermissionName, field.TypeString, value)
		_node.PermissionName = value
	}
	if value, ok := pgc.mutation.Ioc(); ok {
		_spec.SetField(permissiongroup.FieldIoc, field.TypeString, value)
		_node.Ioc = value
	}
	if value, ok := pgc.mutation.Sort(); ok {
		_spec.SetField(permissiongroup.FieldSort, field.TypeInt32, value)
		_node.Sort = value
	}
	if value, ok := pgc.mutation.Left(); ok {
		_spec.SetField(permissiongroup.FieldLeft, field.TypeInt32, value)
		_node.Left = value
	}
	if value, ok := pgc.mutation.Right(); ok {
		_spec.SetField(permissiongroup.FieldRight, field.TypeInt32, value)
		_node.Right = value
	}
	if value, ok := pgc.mutation.State(); ok {
		_spec.SetField(permissiongroup.FieldState, field.TypeInt, value)
		_node.State = value
	}
	if value, ok := pgc.mutation.CreateTime(); ok {
		_spec.SetField(permissiongroup.FieldCreateTime, field.TypeInt64, value)
		_node.CreateTime = value
	}
	if value, ok := pgc.mutation.UpdateTime(); ok {
		_spec.SetField(permissiongroup.FieldUpdateTime, field.TypeInt64, value)
		_node.UpdateTime = value
	}
	return _node, _spec
}

// PermissionGroupCreateBulk is the builder for creating many PermissionGroup entities in bulk.
type PermissionGroupCreateBulk struct {
	config
	err      error
	builders []*PermissionGroupCreate
}

// Save creates the PermissionGroup entities in the database.
func (pgcb *PermissionGroupCreateBulk) Save(ctx context.Context) ([]*PermissionGroup, error) {
	if pgcb.err != nil {
		return nil, pgcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pgcb.builders))
	nodes := make([]*PermissionGroup, len(pgcb.builders))
	mutators := make([]Mutator, len(pgcb.builders))
	for i := range pgcb.builders {
		func(i int, root context.Context) {
			builder := pgcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PermissionGroupMutation)
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
					_, err = mutators[i+1].Mutate(root, pgcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pgcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, pgcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pgcb *PermissionGroupCreateBulk) SaveX(ctx context.Context) []*PermissionGroup {
	v, err := pgcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pgcb *PermissionGroupCreateBulk) Exec(ctx context.Context) error {
	_, err := pgcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pgcb *PermissionGroupCreateBulk) ExecX(ctx context.Context) {
	if err := pgcb.Exec(ctx); err != nil {
		panic(err)
	}
}
