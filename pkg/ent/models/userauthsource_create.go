// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/userauthsource"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserAuthSourceCreate is the builder for creating a UserAuthSource entity.
type UserAuthSourceCreate struct {
	config
	mutation *UserAuthSourceMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (uasc *UserAuthSourceCreate) SetUserID(u uuid.UUID) *UserAuthSourceCreate {
	uasc.mutation.SetUserID(u)
	return uasc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uasc *UserAuthSourceCreate) SetNillableUserID(u *uuid.UUID) *UserAuthSourceCreate {
	if u != nil {
		uasc.SetUserID(*u)
	}
	return uasc
}

// SetToken sets the "token" field.
func (uasc *UserAuthSourceCreate) SetToken(s string) *UserAuthSourceCreate {
	uasc.mutation.SetToken(s)
	return uasc
}

// SetChannel sets the "channel" field.
func (uasc *UserAuthSourceCreate) SetChannel(s string) *UserAuthSourceCreate {
	uasc.mutation.SetChannel(s)
	return uasc
}

// SetDevice sets the "device" field.
func (uasc *UserAuthSourceCreate) SetDevice(s string) *UserAuthSourceCreate {
	uasc.mutation.SetDevice(s)
	return uasc
}

// SetDeviceDetail sets the "device_detail" field.
func (uasc *UserAuthSourceCreate) SetDeviceDetail(s string) *UserAuthSourceCreate {
	uasc.mutation.SetDeviceDetail(s)
	return uasc
}

// SetClientIP sets the "client_ip" field.
func (uasc *UserAuthSourceCreate) SetClientIP(s string) *UserAuthSourceCreate {
	uasc.mutation.SetClientIP(s)
	return uasc
}

// SetRemoteIP sets the "remote_ip" field.
func (uasc *UserAuthSourceCreate) SetRemoteIP(s string) *UserAuthSourceCreate {
	uasc.mutation.SetRemoteIP(s)
	return uasc
}

// SetSnapshot sets the "snapshot" field.
func (uasc *UserAuthSourceCreate) SetSnapshot(s string) *UserAuthSourceCreate {
	uasc.mutation.SetSnapshot(s)
	return uasc
}

// SetLoginName sets the "login_name" field.
func (uasc *UserAuthSourceCreate) SetLoginName(s string) *UserAuthSourceCreate {
	uasc.mutation.SetLoginName(s)
	return uasc
}

// SetLoginSource sets the "login_source" field.
func (uasc *UserAuthSourceCreate) SetLoginSource(i int) *UserAuthSourceCreate {
	uasc.mutation.SetLoginSource(i)
	return uasc
}

// SetNillableLoginSource sets the "login_source" field if the given value is not nil.
func (uasc *UserAuthSourceCreate) SetNillableLoginSource(i *int) *UserAuthSourceCreate {
	if i != nil {
		uasc.SetLoginSource(*i)
	}
	return uasc
}

// SetLoginType sets the "login_type" field.
func (uasc *UserAuthSourceCreate) SetLoginType(i int) *UserAuthSourceCreate {
	uasc.mutation.SetLoginType(i)
	return uasc
}

// SetNillableLoginType sets the "login_type" field if the given value is not nil.
func (uasc *UserAuthSourceCreate) SetNillableLoginType(i *int) *UserAuthSourceCreate {
	if i != nil {
		uasc.SetLoginType(*i)
	}
	return uasc
}

// SetCreateTime sets the "create_time" field.
func (uasc *UserAuthSourceCreate) SetCreateTime(t time.Time) *UserAuthSourceCreate {
	uasc.mutation.SetCreateTime(t)
	return uasc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (uasc *UserAuthSourceCreate) SetNillableCreateTime(t *time.Time) *UserAuthSourceCreate {
	if t != nil {
		uasc.SetCreateTime(*t)
	}
	return uasc
}

// SetUpdateTime sets the "update_time" field.
func (uasc *UserAuthSourceCreate) SetUpdateTime(t time.Time) *UserAuthSourceCreate {
	uasc.mutation.SetUpdateTime(t)
	return uasc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (uasc *UserAuthSourceCreate) SetNillableUpdateTime(t *time.Time) *UserAuthSourceCreate {
	if t != nil {
		uasc.SetUpdateTime(*t)
	}
	return uasc
}

// SetID sets the "id" field.
func (uasc *UserAuthSourceCreate) SetID(i int64) *UserAuthSourceCreate {
	uasc.mutation.SetID(i)
	return uasc
}

// Mutation returns the UserAuthSourceMutation object of the builder.
func (uasc *UserAuthSourceCreate) Mutation() *UserAuthSourceMutation {
	return uasc.mutation
}

// Save creates the UserAuthSource in the database.
func (uasc *UserAuthSourceCreate) Save(ctx context.Context) (*UserAuthSource, error) {
	uasc.defaults()
	return withHooks(ctx, uasc.sqlSave, uasc.mutation, uasc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uasc *UserAuthSourceCreate) SaveX(ctx context.Context) *UserAuthSource {
	v, err := uasc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uasc *UserAuthSourceCreate) Exec(ctx context.Context) error {
	_, err := uasc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uasc *UserAuthSourceCreate) ExecX(ctx context.Context) {
	if err := uasc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uasc *UserAuthSourceCreate) defaults() {
	if _, ok := uasc.mutation.UserID(); !ok {
		v := userauthsource.DefaultUserID()
		uasc.mutation.SetUserID(v)
	}
	if _, ok := uasc.mutation.LoginSource(); !ok {
		v := userauthsource.DefaultLoginSource
		uasc.mutation.SetLoginSource(v)
	}
	if _, ok := uasc.mutation.LoginType(); !ok {
		v := userauthsource.DefaultLoginType
		uasc.mutation.SetLoginType(v)
	}
	if _, ok := uasc.mutation.CreateTime(); !ok {
		v := userauthsource.DefaultCreateTime()
		uasc.mutation.SetCreateTime(v)
	}
	if _, ok := uasc.mutation.UpdateTime(); !ok {
		v := userauthsource.DefaultUpdateTime()
		uasc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uasc *UserAuthSourceCreate) check() error {
	if _, ok := uasc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`models: missing required field "UserAuthSource.user_id"`)}
	}
	if _, ok := uasc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`models: missing required field "UserAuthSource.token"`)}
	}
	if v, ok := uasc.mutation.Token(); ok {
		if err := userauthsource.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.token": %w`, err)}
		}
	}
	if _, ok := uasc.mutation.Channel(); !ok {
		return &ValidationError{Name: "channel", err: errors.New(`models: missing required field "UserAuthSource.channel"`)}
	}
	if v, ok := uasc.mutation.Channel(); ok {
		if err := userauthsource.ChannelValidator(v); err != nil {
			return &ValidationError{Name: "channel", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.channel": %w`, err)}
		}
	}
	if _, ok := uasc.mutation.Device(); !ok {
		return &ValidationError{Name: "device", err: errors.New(`models: missing required field "UserAuthSource.device"`)}
	}
	if v, ok := uasc.mutation.Device(); ok {
		if err := userauthsource.DeviceValidator(v); err != nil {
			return &ValidationError{Name: "device", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.device": %w`, err)}
		}
	}
	if _, ok := uasc.mutation.DeviceDetail(); !ok {
		return &ValidationError{Name: "device_detail", err: errors.New(`models: missing required field "UserAuthSource.device_detail"`)}
	}
	if v, ok := uasc.mutation.DeviceDetail(); ok {
		if err := userauthsource.DeviceDetailValidator(v); err != nil {
			return &ValidationError{Name: "device_detail", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.device_detail": %w`, err)}
		}
	}
	if _, ok := uasc.mutation.ClientIP(); !ok {
		return &ValidationError{Name: "client_ip", err: errors.New(`models: missing required field "UserAuthSource.client_ip"`)}
	}
	if v, ok := uasc.mutation.ClientIP(); ok {
		if err := userauthsource.ClientIPValidator(v); err != nil {
			return &ValidationError{Name: "client_ip", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.client_ip": %w`, err)}
		}
	}
	if _, ok := uasc.mutation.RemoteIP(); !ok {
		return &ValidationError{Name: "remote_ip", err: errors.New(`models: missing required field "UserAuthSource.remote_ip"`)}
	}
	if v, ok := uasc.mutation.RemoteIP(); ok {
		if err := userauthsource.RemoteIPValidator(v); err != nil {
			return &ValidationError{Name: "remote_ip", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.remote_ip": %w`, err)}
		}
	}
	if _, ok := uasc.mutation.Snapshot(); !ok {
		return &ValidationError{Name: "snapshot", err: errors.New(`models: missing required field "UserAuthSource.snapshot"`)}
	}
	if v, ok := uasc.mutation.Snapshot(); ok {
		if err := userauthsource.SnapshotValidator(v); err != nil {
			return &ValidationError{Name: "snapshot", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.snapshot": %w`, err)}
		}
	}
	if _, ok := uasc.mutation.LoginName(); !ok {
		return &ValidationError{Name: "login_name", err: errors.New(`models: missing required field "UserAuthSource.login_name"`)}
	}
	if v, ok := uasc.mutation.LoginName(); ok {
		if err := userauthsource.LoginNameValidator(v); err != nil {
			return &ValidationError{Name: "login_name", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.login_name": %w`, err)}
		}
	}
	if _, ok := uasc.mutation.LoginSource(); !ok {
		return &ValidationError{Name: "login_source", err: errors.New(`models: missing required field "UserAuthSource.login_source"`)}
	}
	if _, ok := uasc.mutation.LoginType(); !ok {
		return &ValidationError{Name: "login_type", err: errors.New(`models: missing required field "UserAuthSource.login_type"`)}
	}
	if _, ok := uasc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`models: missing required field "UserAuthSource.create_time"`)}
	}
	if _, ok := uasc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`models: missing required field "UserAuthSource.update_time"`)}
	}
	return nil
}

func (uasc *UserAuthSourceCreate) sqlSave(ctx context.Context) (*UserAuthSource, error) {
	if err := uasc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uasc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uasc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	uasc.mutation.id = &_node.ID
	uasc.mutation.done = true
	return _node, nil
}

func (uasc *UserAuthSourceCreate) createSpec() (*UserAuthSource, *sqlgraph.CreateSpec) {
	var (
		_node = &UserAuthSource{config: uasc.config}
		_spec = sqlgraph.NewCreateSpec(userauthsource.Table, sqlgraph.NewFieldSpec(userauthsource.FieldID, field.TypeInt64))
	)
	if id, ok := uasc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uasc.mutation.UserID(); ok {
		_spec.SetField(userauthsource.FieldUserID, field.TypeUUID, value)
		_node.UserID = value
	}
	if value, ok := uasc.mutation.Token(); ok {
		_spec.SetField(userauthsource.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := uasc.mutation.Channel(); ok {
		_spec.SetField(userauthsource.FieldChannel, field.TypeString, value)
		_node.Channel = value
	}
	if value, ok := uasc.mutation.Device(); ok {
		_spec.SetField(userauthsource.FieldDevice, field.TypeString, value)
		_node.Device = value
	}
	if value, ok := uasc.mutation.DeviceDetail(); ok {
		_spec.SetField(userauthsource.FieldDeviceDetail, field.TypeString, value)
		_node.DeviceDetail = value
	}
	if value, ok := uasc.mutation.ClientIP(); ok {
		_spec.SetField(userauthsource.FieldClientIP, field.TypeString, value)
		_node.ClientIP = value
	}
	if value, ok := uasc.mutation.RemoteIP(); ok {
		_spec.SetField(userauthsource.FieldRemoteIP, field.TypeString, value)
		_node.RemoteIP = value
	}
	if value, ok := uasc.mutation.Snapshot(); ok {
		_spec.SetField(userauthsource.FieldSnapshot, field.TypeString, value)
		_node.Snapshot = value
	}
	if value, ok := uasc.mutation.LoginName(); ok {
		_spec.SetField(userauthsource.FieldLoginName, field.TypeString, value)
		_node.LoginName = value
	}
	if value, ok := uasc.mutation.LoginSource(); ok {
		_spec.SetField(userauthsource.FieldLoginSource, field.TypeInt, value)
		_node.LoginSource = value
	}
	if value, ok := uasc.mutation.LoginType(); ok {
		_spec.SetField(userauthsource.FieldLoginType, field.TypeInt, value)
		_node.LoginType = value
	}
	if value, ok := uasc.mutation.CreateTime(); ok {
		_spec.SetField(userauthsource.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := uasc.mutation.UpdateTime(); ok {
		_spec.SetField(userauthsource.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	return _node, _spec
}

// UserAuthSourceCreateBulk is the builder for creating many UserAuthSource entities in bulk.
type UserAuthSourceCreateBulk struct {
	config
	err      error
	builders []*UserAuthSourceCreate
}

// Save creates the UserAuthSource entities in the database.
func (uascb *UserAuthSourceCreateBulk) Save(ctx context.Context) ([]*UserAuthSource, error) {
	if uascb.err != nil {
		return nil, uascb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uascb.builders))
	nodes := make([]*UserAuthSource, len(uascb.builders))
	mutators := make([]Mutator, len(uascb.builders))
	for i := range uascb.builders {
		func(i int, root context.Context) {
			builder := uascb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserAuthSourceMutation)
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
					_, err = mutators[i+1].Mutate(root, uascb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uascb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uascb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uascb *UserAuthSourceCreateBulk) SaveX(ctx context.Context) []*UserAuthSource {
	v, err := uascb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uascb *UserAuthSourceCreateBulk) Exec(ctx context.Context) error {
	_, err := uascb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uascb *UserAuthSourceCreateBulk) ExecX(ctx context.Context) {
	if err := uascb.Exec(ctx); err != nil {
		panic(err)
	}
}
