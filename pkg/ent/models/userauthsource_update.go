// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/userauthsource"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserAuthSourceUpdate is the builder for updating UserAuthSource entities.
type UserAuthSourceUpdate struct {
	config
	hooks    []Hook
	mutation *UserAuthSourceMutation
}

// Where appends a list predicates to the UserAuthSourceUpdate builder.
func (uasu *UserAuthSourceUpdate) Where(ps ...predicate.UserAuthSource) *UserAuthSourceUpdate {
	uasu.mutation.Where(ps...)
	return uasu
}

// SetUserID sets the "user_id" field.
func (uasu *UserAuthSourceUpdate) SetUserID(u uuid.UUID) *UserAuthSourceUpdate {
	uasu.mutation.SetUserID(u)
	return uasu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uasu *UserAuthSourceUpdate) SetNillableUserID(u *uuid.UUID) *UserAuthSourceUpdate {
	if u != nil {
		uasu.SetUserID(*u)
	}
	return uasu
}

// SetToken sets the "token" field.
func (uasu *UserAuthSourceUpdate) SetToken(s string) *UserAuthSourceUpdate {
	uasu.mutation.SetToken(s)
	return uasu
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (uasu *UserAuthSourceUpdate) SetNillableToken(s *string) *UserAuthSourceUpdate {
	if s != nil {
		uasu.SetToken(*s)
	}
	return uasu
}

// SetChannel sets the "channel" field.
func (uasu *UserAuthSourceUpdate) SetChannel(s string) *UserAuthSourceUpdate {
	uasu.mutation.SetChannel(s)
	return uasu
}

// SetNillableChannel sets the "channel" field if the given value is not nil.
func (uasu *UserAuthSourceUpdate) SetNillableChannel(s *string) *UserAuthSourceUpdate {
	if s != nil {
		uasu.SetChannel(*s)
	}
	return uasu
}

// SetDevice sets the "device" field.
func (uasu *UserAuthSourceUpdate) SetDevice(s string) *UserAuthSourceUpdate {
	uasu.mutation.SetDevice(s)
	return uasu
}

// SetNillableDevice sets the "device" field if the given value is not nil.
func (uasu *UserAuthSourceUpdate) SetNillableDevice(s *string) *UserAuthSourceUpdate {
	if s != nil {
		uasu.SetDevice(*s)
	}
	return uasu
}

// SetDeviceDetail sets the "device_detail" field.
func (uasu *UserAuthSourceUpdate) SetDeviceDetail(s string) *UserAuthSourceUpdate {
	uasu.mutation.SetDeviceDetail(s)
	return uasu
}

// SetNillableDeviceDetail sets the "device_detail" field if the given value is not nil.
func (uasu *UserAuthSourceUpdate) SetNillableDeviceDetail(s *string) *UserAuthSourceUpdate {
	if s != nil {
		uasu.SetDeviceDetail(*s)
	}
	return uasu
}

// SetClientIP sets the "client_ip" field.
func (uasu *UserAuthSourceUpdate) SetClientIP(s string) *UserAuthSourceUpdate {
	uasu.mutation.SetClientIP(s)
	return uasu
}

// SetNillableClientIP sets the "client_ip" field if the given value is not nil.
func (uasu *UserAuthSourceUpdate) SetNillableClientIP(s *string) *UserAuthSourceUpdate {
	if s != nil {
		uasu.SetClientIP(*s)
	}
	return uasu
}

// SetRemoteIP sets the "remote_ip" field.
func (uasu *UserAuthSourceUpdate) SetRemoteIP(s string) *UserAuthSourceUpdate {
	uasu.mutation.SetRemoteIP(s)
	return uasu
}

// SetNillableRemoteIP sets the "remote_ip" field if the given value is not nil.
func (uasu *UserAuthSourceUpdate) SetNillableRemoteIP(s *string) *UserAuthSourceUpdate {
	if s != nil {
		uasu.SetRemoteIP(*s)
	}
	return uasu
}

// SetSnapshot sets the "snapshot" field.
func (uasu *UserAuthSourceUpdate) SetSnapshot(s string) *UserAuthSourceUpdate {
	uasu.mutation.SetSnapshot(s)
	return uasu
}

// SetNillableSnapshot sets the "snapshot" field if the given value is not nil.
func (uasu *UserAuthSourceUpdate) SetNillableSnapshot(s *string) *UserAuthSourceUpdate {
	if s != nil {
		uasu.SetSnapshot(*s)
	}
	return uasu
}

// SetLoginName sets the "login_name" field.
func (uasu *UserAuthSourceUpdate) SetLoginName(s string) *UserAuthSourceUpdate {
	uasu.mutation.SetLoginName(s)
	return uasu
}

// SetNillableLoginName sets the "login_name" field if the given value is not nil.
func (uasu *UserAuthSourceUpdate) SetNillableLoginName(s *string) *UserAuthSourceUpdate {
	if s != nil {
		uasu.SetLoginName(*s)
	}
	return uasu
}

// SetLoginSource sets the "login_source" field.
func (uasu *UserAuthSourceUpdate) SetLoginSource(i int) *UserAuthSourceUpdate {
	uasu.mutation.ResetLoginSource()
	uasu.mutation.SetLoginSource(i)
	return uasu
}

// SetNillableLoginSource sets the "login_source" field if the given value is not nil.
func (uasu *UserAuthSourceUpdate) SetNillableLoginSource(i *int) *UserAuthSourceUpdate {
	if i != nil {
		uasu.SetLoginSource(*i)
	}
	return uasu
}

// AddLoginSource adds i to the "login_source" field.
func (uasu *UserAuthSourceUpdate) AddLoginSource(i int) *UserAuthSourceUpdate {
	uasu.mutation.AddLoginSource(i)
	return uasu
}

// SetLoginType sets the "login_type" field.
func (uasu *UserAuthSourceUpdate) SetLoginType(i int) *UserAuthSourceUpdate {
	uasu.mutation.ResetLoginType()
	uasu.mutation.SetLoginType(i)
	return uasu
}

// SetNillableLoginType sets the "login_type" field if the given value is not nil.
func (uasu *UserAuthSourceUpdate) SetNillableLoginType(i *int) *UserAuthSourceUpdate {
	if i != nil {
		uasu.SetLoginType(*i)
	}
	return uasu
}

// AddLoginType adds i to the "login_type" field.
func (uasu *UserAuthSourceUpdate) AddLoginType(i int) *UserAuthSourceUpdate {
	uasu.mutation.AddLoginType(i)
	return uasu
}

// SetUpdateTime sets the "update_time" field.
func (uasu *UserAuthSourceUpdate) SetUpdateTime(t time.Time) *UserAuthSourceUpdate {
	uasu.mutation.SetUpdateTime(t)
	return uasu
}

// Mutation returns the UserAuthSourceMutation object of the builder.
func (uasu *UserAuthSourceUpdate) Mutation() *UserAuthSourceMutation {
	return uasu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uasu *UserAuthSourceUpdate) Save(ctx context.Context) (int, error) {
	uasu.defaults()
	return withHooks(ctx, uasu.sqlSave, uasu.mutation, uasu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uasu *UserAuthSourceUpdate) SaveX(ctx context.Context) int {
	affected, err := uasu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uasu *UserAuthSourceUpdate) Exec(ctx context.Context) error {
	_, err := uasu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uasu *UserAuthSourceUpdate) ExecX(ctx context.Context) {
	if err := uasu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uasu *UserAuthSourceUpdate) defaults() {
	if _, ok := uasu.mutation.UpdateTime(); !ok {
		v := userauthsource.UpdateDefaultUpdateTime()
		uasu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uasu *UserAuthSourceUpdate) check() error {
	if v, ok := uasu.mutation.Token(); ok {
		if err := userauthsource.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.token": %w`, err)}
		}
	}
	if v, ok := uasu.mutation.Channel(); ok {
		if err := userauthsource.ChannelValidator(v); err != nil {
			return &ValidationError{Name: "channel", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.channel": %w`, err)}
		}
	}
	if v, ok := uasu.mutation.Device(); ok {
		if err := userauthsource.DeviceValidator(v); err != nil {
			return &ValidationError{Name: "device", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.device": %w`, err)}
		}
	}
	if v, ok := uasu.mutation.DeviceDetail(); ok {
		if err := userauthsource.DeviceDetailValidator(v); err != nil {
			return &ValidationError{Name: "device_detail", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.device_detail": %w`, err)}
		}
	}
	if v, ok := uasu.mutation.ClientIP(); ok {
		if err := userauthsource.ClientIPValidator(v); err != nil {
			return &ValidationError{Name: "client_ip", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.client_ip": %w`, err)}
		}
	}
	if v, ok := uasu.mutation.RemoteIP(); ok {
		if err := userauthsource.RemoteIPValidator(v); err != nil {
			return &ValidationError{Name: "remote_ip", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.remote_ip": %w`, err)}
		}
	}
	if v, ok := uasu.mutation.Snapshot(); ok {
		if err := userauthsource.SnapshotValidator(v); err != nil {
			return &ValidationError{Name: "snapshot", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.snapshot": %w`, err)}
		}
	}
	if v, ok := uasu.mutation.LoginName(); ok {
		if err := userauthsource.LoginNameValidator(v); err != nil {
			return &ValidationError{Name: "login_name", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.login_name": %w`, err)}
		}
	}
	return nil
}

func (uasu *UserAuthSourceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uasu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userauthsource.Table, userauthsource.Columns, sqlgraph.NewFieldSpec(userauthsource.FieldID, field.TypeInt))
	if ps := uasu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uasu.mutation.UserID(); ok {
		_spec.SetField(userauthsource.FieldUserID, field.TypeUUID, value)
	}
	if value, ok := uasu.mutation.Token(); ok {
		_spec.SetField(userauthsource.FieldToken, field.TypeString, value)
	}
	if value, ok := uasu.mutation.Channel(); ok {
		_spec.SetField(userauthsource.FieldChannel, field.TypeString, value)
	}
	if value, ok := uasu.mutation.Device(); ok {
		_spec.SetField(userauthsource.FieldDevice, field.TypeString, value)
	}
	if value, ok := uasu.mutation.DeviceDetail(); ok {
		_spec.SetField(userauthsource.FieldDeviceDetail, field.TypeString, value)
	}
	if value, ok := uasu.mutation.ClientIP(); ok {
		_spec.SetField(userauthsource.FieldClientIP, field.TypeString, value)
	}
	if value, ok := uasu.mutation.RemoteIP(); ok {
		_spec.SetField(userauthsource.FieldRemoteIP, field.TypeString, value)
	}
	if value, ok := uasu.mutation.Snapshot(); ok {
		_spec.SetField(userauthsource.FieldSnapshot, field.TypeString, value)
	}
	if value, ok := uasu.mutation.LoginName(); ok {
		_spec.SetField(userauthsource.FieldLoginName, field.TypeString, value)
	}
	if value, ok := uasu.mutation.LoginSource(); ok {
		_spec.SetField(userauthsource.FieldLoginSource, field.TypeInt, value)
	}
	if value, ok := uasu.mutation.AddedLoginSource(); ok {
		_spec.AddField(userauthsource.FieldLoginSource, field.TypeInt, value)
	}
	if value, ok := uasu.mutation.LoginType(); ok {
		_spec.SetField(userauthsource.FieldLoginType, field.TypeInt, value)
	}
	if value, ok := uasu.mutation.AddedLoginType(); ok {
		_spec.AddField(userauthsource.FieldLoginType, field.TypeInt, value)
	}
	if value, ok := uasu.mutation.UpdateTime(); ok {
		_spec.SetField(userauthsource.FieldUpdateTime, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uasu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userauthsource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uasu.mutation.done = true
	return n, nil
}

// UserAuthSourceUpdateOne is the builder for updating a single UserAuthSource entity.
type UserAuthSourceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserAuthSourceMutation
}

// SetUserID sets the "user_id" field.
func (uasuo *UserAuthSourceUpdateOne) SetUserID(u uuid.UUID) *UserAuthSourceUpdateOne {
	uasuo.mutation.SetUserID(u)
	return uasuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uasuo *UserAuthSourceUpdateOne) SetNillableUserID(u *uuid.UUID) *UserAuthSourceUpdateOne {
	if u != nil {
		uasuo.SetUserID(*u)
	}
	return uasuo
}

// SetToken sets the "token" field.
func (uasuo *UserAuthSourceUpdateOne) SetToken(s string) *UserAuthSourceUpdateOne {
	uasuo.mutation.SetToken(s)
	return uasuo
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (uasuo *UserAuthSourceUpdateOne) SetNillableToken(s *string) *UserAuthSourceUpdateOne {
	if s != nil {
		uasuo.SetToken(*s)
	}
	return uasuo
}

// SetChannel sets the "channel" field.
func (uasuo *UserAuthSourceUpdateOne) SetChannel(s string) *UserAuthSourceUpdateOne {
	uasuo.mutation.SetChannel(s)
	return uasuo
}

// SetNillableChannel sets the "channel" field if the given value is not nil.
func (uasuo *UserAuthSourceUpdateOne) SetNillableChannel(s *string) *UserAuthSourceUpdateOne {
	if s != nil {
		uasuo.SetChannel(*s)
	}
	return uasuo
}

// SetDevice sets the "device" field.
func (uasuo *UserAuthSourceUpdateOne) SetDevice(s string) *UserAuthSourceUpdateOne {
	uasuo.mutation.SetDevice(s)
	return uasuo
}

// SetNillableDevice sets the "device" field if the given value is not nil.
func (uasuo *UserAuthSourceUpdateOne) SetNillableDevice(s *string) *UserAuthSourceUpdateOne {
	if s != nil {
		uasuo.SetDevice(*s)
	}
	return uasuo
}

// SetDeviceDetail sets the "device_detail" field.
func (uasuo *UserAuthSourceUpdateOne) SetDeviceDetail(s string) *UserAuthSourceUpdateOne {
	uasuo.mutation.SetDeviceDetail(s)
	return uasuo
}

// SetNillableDeviceDetail sets the "device_detail" field if the given value is not nil.
func (uasuo *UserAuthSourceUpdateOne) SetNillableDeviceDetail(s *string) *UserAuthSourceUpdateOne {
	if s != nil {
		uasuo.SetDeviceDetail(*s)
	}
	return uasuo
}

// SetClientIP sets the "client_ip" field.
func (uasuo *UserAuthSourceUpdateOne) SetClientIP(s string) *UserAuthSourceUpdateOne {
	uasuo.mutation.SetClientIP(s)
	return uasuo
}

// SetNillableClientIP sets the "client_ip" field if the given value is not nil.
func (uasuo *UserAuthSourceUpdateOne) SetNillableClientIP(s *string) *UserAuthSourceUpdateOne {
	if s != nil {
		uasuo.SetClientIP(*s)
	}
	return uasuo
}

// SetRemoteIP sets the "remote_ip" field.
func (uasuo *UserAuthSourceUpdateOne) SetRemoteIP(s string) *UserAuthSourceUpdateOne {
	uasuo.mutation.SetRemoteIP(s)
	return uasuo
}

// SetNillableRemoteIP sets the "remote_ip" field if the given value is not nil.
func (uasuo *UserAuthSourceUpdateOne) SetNillableRemoteIP(s *string) *UserAuthSourceUpdateOne {
	if s != nil {
		uasuo.SetRemoteIP(*s)
	}
	return uasuo
}

// SetSnapshot sets the "snapshot" field.
func (uasuo *UserAuthSourceUpdateOne) SetSnapshot(s string) *UserAuthSourceUpdateOne {
	uasuo.mutation.SetSnapshot(s)
	return uasuo
}

// SetNillableSnapshot sets the "snapshot" field if the given value is not nil.
func (uasuo *UserAuthSourceUpdateOne) SetNillableSnapshot(s *string) *UserAuthSourceUpdateOne {
	if s != nil {
		uasuo.SetSnapshot(*s)
	}
	return uasuo
}

// SetLoginName sets the "login_name" field.
func (uasuo *UserAuthSourceUpdateOne) SetLoginName(s string) *UserAuthSourceUpdateOne {
	uasuo.mutation.SetLoginName(s)
	return uasuo
}

// SetNillableLoginName sets the "login_name" field if the given value is not nil.
func (uasuo *UserAuthSourceUpdateOne) SetNillableLoginName(s *string) *UserAuthSourceUpdateOne {
	if s != nil {
		uasuo.SetLoginName(*s)
	}
	return uasuo
}

// SetLoginSource sets the "login_source" field.
func (uasuo *UserAuthSourceUpdateOne) SetLoginSource(i int) *UserAuthSourceUpdateOne {
	uasuo.mutation.ResetLoginSource()
	uasuo.mutation.SetLoginSource(i)
	return uasuo
}

// SetNillableLoginSource sets the "login_source" field if the given value is not nil.
func (uasuo *UserAuthSourceUpdateOne) SetNillableLoginSource(i *int) *UserAuthSourceUpdateOne {
	if i != nil {
		uasuo.SetLoginSource(*i)
	}
	return uasuo
}

// AddLoginSource adds i to the "login_source" field.
func (uasuo *UserAuthSourceUpdateOne) AddLoginSource(i int) *UserAuthSourceUpdateOne {
	uasuo.mutation.AddLoginSource(i)
	return uasuo
}

// SetLoginType sets the "login_type" field.
func (uasuo *UserAuthSourceUpdateOne) SetLoginType(i int) *UserAuthSourceUpdateOne {
	uasuo.mutation.ResetLoginType()
	uasuo.mutation.SetLoginType(i)
	return uasuo
}

// SetNillableLoginType sets the "login_type" field if the given value is not nil.
func (uasuo *UserAuthSourceUpdateOne) SetNillableLoginType(i *int) *UserAuthSourceUpdateOne {
	if i != nil {
		uasuo.SetLoginType(*i)
	}
	return uasuo
}

// AddLoginType adds i to the "login_type" field.
func (uasuo *UserAuthSourceUpdateOne) AddLoginType(i int) *UserAuthSourceUpdateOne {
	uasuo.mutation.AddLoginType(i)
	return uasuo
}

// SetUpdateTime sets the "update_time" field.
func (uasuo *UserAuthSourceUpdateOne) SetUpdateTime(t time.Time) *UserAuthSourceUpdateOne {
	uasuo.mutation.SetUpdateTime(t)
	return uasuo
}

// Mutation returns the UserAuthSourceMutation object of the builder.
func (uasuo *UserAuthSourceUpdateOne) Mutation() *UserAuthSourceMutation {
	return uasuo.mutation
}

// Where appends a list predicates to the UserAuthSourceUpdate builder.
func (uasuo *UserAuthSourceUpdateOne) Where(ps ...predicate.UserAuthSource) *UserAuthSourceUpdateOne {
	uasuo.mutation.Where(ps...)
	return uasuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uasuo *UserAuthSourceUpdateOne) Select(field string, fields ...string) *UserAuthSourceUpdateOne {
	uasuo.fields = append([]string{field}, fields...)
	return uasuo
}

// Save executes the query and returns the updated UserAuthSource entity.
func (uasuo *UserAuthSourceUpdateOne) Save(ctx context.Context) (*UserAuthSource, error) {
	uasuo.defaults()
	return withHooks(ctx, uasuo.sqlSave, uasuo.mutation, uasuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uasuo *UserAuthSourceUpdateOne) SaveX(ctx context.Context) *UserAuthSource {
	node, err := uasuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uasuo *UserAuthSourceUpdateOne) Exec(ctx context.Context) error {
	_, err := uasuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uasuo *UserAuthSourceUpdateOne) ExecX(ctx context.Context) {
	if err := uasuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uasuo *UserAuthSourceUpdateOne) defaults() {
	if _, ok := uasuo.mutation.UpdateTime(); !ok {
		v := userauthsource.UpdateDefaultUpdateTime()
		uasuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uasuo *UserAuthSourceUpdateOne) check() error {
	if v, ok := uasuo.mutation.Token(); ok {
		if err := userauthsource.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.token": %w`, err)}
		}
	}
	if v, ok := uasuo.mutation.Channel(); ok {
		if err := userauthsource.ChannelValidator(v); err != nil {
			return &ValidationError{Name: "channel", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.channel": %w`, err)}
		}
	}
	if v, ok := uasuo.mutation.Device(); ok {
		if err := userauthsource.DeviceValidator(v); err != nil {
			return &ValidationError{Name: "device", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.device": %w`, err)}
		}
	}
	if v, ok := uasuo.mutation.DeviceDetail(); ok {
		if err := userauthsource.DeviceDetailValidator(v); err != nil {
			return &ValidationError{Name: "device_detail", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.device_detail": %w`, err)}
		}
	}
	if v, ok := uasuo.mutation.ClientIP(); ok {
		if err := userauthsource.ClientIPValidator(v); err != nil {
			return &ValidationError{Name: "client_ip", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.client_ip": %w`, err)}
		}
	}
	if v, ok := uasuo.mutation.RemoteIP(); ok {
		if err := userauthsource.RemoteIPValidator(v); err != nil {
			return &ValidationError{Name: "remote_ip", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.remote_ip": %w`, err)}
		}
	}
	if v, ok := uasuo.mutation.Snapshot(); ok {
		if err := userauthsource.SnapshotValidator(v); err != nil {
			return &ValidationError{Name: "snapshot", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.snapshot": %w`, err)}
		}
	}
	if v, ok := uasuo.mutation.LoginName(); ok {
		if err := userauthsource.LoginNameValidator(v); err != nil {
			return &ValidationError{Name: "login_name", err: fmt.Errorf(`models: validator failed for field "UserAuthSource.login_name": %w`, err)}
		}
	}
	return nil
}

func (uasuo *UserAuthSourceUpdateOne) sqlSave(ctx context.Context) (_node *UserAuthSource, err error) {
	if err := uasuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userauthsource.Table, userauthsource.Columns, sqlgraph.NewFieldSpec(userauthsource.FieldID, field.TypeInt))
	id, ok := uasuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "UserAuthSource.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uasuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userauthsource.FieldID)
		for _, f := range fields {
			if !userauthsource.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != userauthsource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uasuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uasuo.mutation.UserID(); ok {
		_spec.SetField(userauthsource.FieldUserID, field.TypeUUID, value)
	}
	if value, ok := uasuo.mutation.Token(); ok {
		_spec.SetField(userauthsource.FieldToken, field.TypeString, value)
	}
	if value, ok := uasuo.mutation.Channel(); ok {
		_spec.SetField(userauthsource.FieldChannel, field.TypeString, value)
	}
	if value, ok := uasuo.mutation.Device(); ok {
		_spec.SetField(userauthsource.FieldDevice, field.TypeString, value)
	}
	if value, ok := uasuo.mutation.DeviceDetail(); ok {
		_spec.SetField(userauthsource.FieldDeviceDetail, field.TypeString, value)
	}
	if value, ok := uasuo.mutation.ClientIP(); ok {
		_spec.SetField(userauthsource.FieldClientIP, field.TypeString, value)
	}
	if value, ok := uasuo.mutation.RemoteIP(); ok {
		_spec.SetField(userauthsource.FieldRemoteIP, field.TypeString, value)
	}
	if value, ok := uasuo.mutation.Snapshot(); ok {
		_spec.SetField(userauthsource.FieldSnapshot, field.TypeString, value)
	}
	if value, ok := uasuo.mutation.LoginName(); ok {
		_spec.SetField(userauthsource.FieldLoginName, field.TypeString, value)
	}
	if value, ok := uasuo.mutation.LoginSource(); ok {
		_spec.SetField(userauthsource.FieldLoginSource, field.TypeInt, value)
	}
	if value, ok := uasuo.mutation.AddedLoginSource(); ok {
		_spec.AddField(userauthsource.FieldLoginSource, field.TypeInt, value)
	}
	if value, ok := uasuo.mutation.LoginType(); ok {
		_spec.SetField(userauthsource.FieldLoginType, field.TypeInt, value)
	}
	if value, ok := uasuo.mutation.AddedLoginType(); ok {
		_spec.AddField(userauthsource.FieldLoginType, field.TypeInt, value)
	}
	if value, ok := uasuo.mutation.UpdateTime(); ok {
		_spec.SetField(userauthsource.FieldUpdateTime, field.TypeTime, value)
	}
	_node = &UserAuthSource{config: uasuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uasuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userauthsource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uasuo.mutation.done = true
	return _node, nil
}
