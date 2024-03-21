// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/accounts"
	"ent/utils/timeutil"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountsCreate is the builder for creating a Accounts entity.
type AccountsCreate struct {
	config
	mutation *AccountsMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (ac *AccountsCreate) SetUserID(i int64) *AccountsCreate {
	ac.mutation.SetUserID(i)
	return ac
}

// SetAccount sets the "account" field.
func (ac *AccountsCreate) SetAccount(s string) *AccountsCreate {
	ac.mutation.SetAccount(s)
	return ac
}

// SetType sets the "type" field.
func (ac *AccountsCreate) SetType(u uint8) *AccountsCreate {
	ac.mutation.SetType(u)
	return ac
}

// SetDesc sets the "desc" field.
func (ac *AccountsCreate) SetDesc(s string) *AccountsCreate {
	ac.mutation.SetDesc(s)
	return ac
}

// SetIsPrivate sets the "is_private" field.
func (ac *AccountsCreate) SetIsPrivate(b bool) *AccountsCreate {
	ac.mutation.SetIsPrivate(b)
	return ac
}

// SetNillableIsPrivate sets the "is_private" field if the given value is not nil.
func (ac *AccountsCreate) SetNillableIsPrivate(b *bool) *AccountsCreate {
	if b != nil {
		ac.SetIsPrivate(*b)
	}
	return ac
}

// SetIsActivated sets the "is_activated" field.
func (ac *AccountsCreate) SetIsActivated(b bool) *AccountsCreate {
	ac.mutation.SetIsActivated(b)
	return ac
}

// SetIsPrimary sets the "is_primary" field.
func (ac *AccountsCreate) SetIsPrimary(b bool) *AccountsCreate {
	ac.mutation.SetIsPrimary(b)
	return ac
}

// SetNillableIsPrimary sets the "is_primary" field if the given value is not nil.
func (ac *AccountsCreate) SetNillableIsPrimary(b *bool) *AccountsCreate {
	if b != nil {
		ac.SetIsPrimary(*b)
	}
	return ac
}

// SetCreateTime sets the "create_time" field.
func (ac *AccountsCreate) SetCreateTime(ts timeutil.TimeStamp) *AccountsCreate {
	ac.mutation.SetCreateTime(ts)
	return ac
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ac *AccountsCreate) SetNillableCreateTime(ts *timeutil.TimeStamp) *AccountsCreate {
	if ts != nil {
		ac.SetCreateTime(*ts)
	}
	return ac
}

// SetUpdateTime sets the "update_time" field.
func (ac *AccountsCreate) SetUpdateTime(ts timeutil.TimeStamp) *AccountsCreate {
	ac.mutation.SetUpdateTime(ts)
	return ac
}

// SetID sets the "id" field.
func (ac *AccountsCreate) SetID(i int64) *AccountsCreate {
	ac.mutation.SetID(i)
	return ac
}

// Mutation returns the AccountsMutation object of the builder.
func (ac *AccountsCreate) Mutation() *AccountsMutation {
	return ac.mutation
}

// Save creates the Accounts in the database.
func (ac *AccountsCreate) Save(ctx context.Context) (*Accounts, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccountsCreate) SaveX(ctx context.Context) *Accounts {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AccountsCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AccountsCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AccountsCreate) defaults() {
	if _, ok := ac.mutation.IsPrivate(); !ok {
		v := accounts.DefaultIsPrivate
		ac.mutation.SetIsPrivate(v)
	}
	if _, ok := ac.mutation.IsPrimary(); !ok {
		v := accounts.DefaultIsPrimary
		ac.mutation.SetIsPrimary(v)
	}
	if _, ok := ac.mutation.CreateTime(); !ok {
		v := accounts.DefaultCreateTime
		ac.mutation.SetCreateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AccountsCreate) check() error {
	if _, ok := ac.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`models: missing required field "Accounts.user_id"`)}
	}
	if _, ok := ac.mutation.Account(); !ok {
		return &ValidationError{Name: "account", err: errors.New(`models: missing required field "Accounts.account"`)}
	}
	if _, ok := ac.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`models: missing required field "Accounts.type"`)}
	}
	if _, ok := ac.mutation.Desc(); !ok {
		return &ValidationError{Name: "desc", err: errors.New(`models: missing required field "Accounts.desc"`)}
	}
	if _, ok := ac.mutation.IsPrivate(); !ok {
		return &ValidationError{Name: "is_private", err: errors.New(`models: missing required field "Accounts.is_private"`)}
	}
	if _, ok := ac.mutation.IsActivated(); !ok {
		return &ValidationError{Name: "is_activated", err: errors.New(`models: missing required field "Accounts.is_activated"`)}
	}
	if _, ok := ac.mutation.IsPrimary(); !ok {
		return &ValidationError{Name: "is_primary", err: errors.New(`models: missing required field "Accounts.is_primary"`)}
	}
	if _, ok := ac.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`models: missing required field "Accounts.create_time"`)}
	}
	if _, ok := ac.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`models: missing required field "Accounts.update_time"`)}
	}
	return nil
}

func (ac *AccountsCreate) sqlSave(ctx context.Context) (*Accounts, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AccountsCreate) createSpec() (*Accounts, *sqlgraph.CreateSpec) {
	var (
		_node = &Accounts{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(accounts.Table, sqlgraph.NewFieldSpec(accounts.FieldID, field.TypeInt64))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.UserID(); ok {
		_spec.SetField(accounts.FieldUserID, field.TypeInt64, value)
		_node.UserID = value
	}
	if value, ok := ac.mutation.Account(); ok {
		_spec.SetField(accounts.FieldAccount, field.TypeString, value)
		_node.Account = value
	}
	if value, ok := ac.mutation.GetType(); ok {
		_spec.SetField(accounts.FieldType, field.TypeUint8, value)
		_node.Type = value
	}
	if value, ok := ac.mutation.Desc(); ok {
		_spec.SetField(accounts.FieldDesc, field.TypeString, value)
		_node.Desc = value
	}
	if value, ok := ac.mutation.IsPrivate(); ok {
		_spec.SetField(accounts.FieldIsPrivate, field.TypeBool, value)
		_node.IsPrivate = value
	}
	if value, ok := ac.mutation.IsActivated(); ok {
		_spec.SetField(accounts.FieldIsActivated, field.TypeBool, value)
		_node.IsActivated = value
	}
	if value, ok := ac.mutation.IsPrimary(); ok {
		_spec.SetField(accounts.FieldIsPrimary, field.TypeBool, value)
		_node.IsPrimary = value
	}
	if value, ok := ac.mutation.CreateTime(); ok {
		_spec.SetField(accounts.FieldCreateTime, field.TypeInt64, value)
		_node.CreateTime = value
	}
	if value, ok := ac.mutation.UpdateTime(); ok {
		_spec.SetField(accounts.FieldUpdateTime, field.TypeInt64, value)
		_node.UpdateTime = value
	}
	return _node, _spec
}

// AccountsCreateBulk is the builder for creating many Accounts entities in bulk.
type AccountsCreateBulk struct {
	config
	err      error
	builders []*AccountsCreate
}

// Save creates the Accounts entities in the database.
func (acb *AccountsCreateBulk) Save(ctx context.Context) ([]*Accounts, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Accounts, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountsMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AccountsCreateBulk) SaveX(ctx context.Context) []*Accounts {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AccountsCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AccountsCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}