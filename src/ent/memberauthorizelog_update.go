// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"service-api/src/ent/memberauthorizelog"
	"service-api/src/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MemberAuthorizeLogUpdate is the builder for updating MemberAuthorizeLog entities.
type MemberAuthorizeLogUpdate struct {
	config
	hooks    []Hook
	mutation *MemberAuthorizeLogMutation
}

// Where appends a list predicates to the MemberAuthorizeLogUpdate builder.
func (malu *MemberAuthorizeLogUpdate) Where(ps ...predicate.MemberAuthorizeLog) *MemberAuthorizeLogUpdate {
	malu.mutation.Where(ps...)
	return malu
}

// SetUpdateTime sets the "update_time" field.
func (malu *MemberAuthorizeLogUpdate) SetUpdateTime(t time.Time) *MemberAuthorizeLogUpdate {
	malu.mutation.SetUpdateTime(t)
	return malu
}

// SetMemberID sets the "member_id" field.
func (malu *MemberAuthorizeLogUpdate) SetMemberID(u uuid.UUID) *MemberAuthorizeLogUpdate {
	malu.mutation.SetMemberID(u)
	return malu
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (malu *MemberAuthorizeLogUpdate) SetNillableMemberID(u *uuid.UUID) *MemberAuthorizeLogUpdate {
	if u != nil {
		malu.SetMemberID(*u)
	}
	return malu
}

// SetChannel sets the "channel" field.
func (malu *MemberAuthorizeLogUpdate) SetChannel(s string) *MemberAuthorizeLogUpdate {
	malu.mutation.SetChannel(s)
	return malu
}

// SetDevice sets the "device" field.
func (malu *MemberAuthorizeLogUpdate) SetDevice(s string) *MemberAuthorizeLogUpdate {
	malu.mutation.SetDevice(s)
	return malu
}

// SetDeviceDetail sets the "device_detail" field.
func (malu *MemberAuthorizeLogUpdate) SetDeviceDetail(s string) *MemberAuthorizeLogUpdate {
	malu.mutation.SetDeviceDetail(s)
	return malu
}

// SetIpv4 sets the "ipv4" field.
func (malu *MemberAuthorizeLogUpdate) SetIpv4(s string) *MemberAuthorizeLogUpdate {
	malu.mutation.SetIpv4(s)
	return malu
}

// SetIpv6 sets the "ipv6" field.
func (malu *MemberAuthorizeLogUpdate) SetIpv6(s string) *MemberAuthorizeLogUpdate {
	malu.mutation.SetIpv6(s)
	return malu
}

// SetSnapshot sets the "snapshot" field.
func (malu *MemberAuthorizeLogUpdate) SetSnapshot(s string) *MemberAuthorizeLogUpdate {
	malu.mutation.SetSnapshot(s)
	return malu
}

// Mutation returns the MemberAuthorizeLogMutation object of the builder.
func (malu *MemberAuthorizeLogUpdate) Mutation() *MemberAuthorizeLogMutation {
	return malu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (malu *MemberAuthorizeLogUpdate) Save(ctx context.Context) (int, error) {
	malu.defaults()
	return withHooks[int, MemberAuthorizeLogMutation](ctx, malu.sqlSave, malu.mutation, malu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (malu *MemberAuthorizeLogUpdate) SaveX(ctx context.Context) int {
	affected, err := malu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (malu *MemberAuthorizeLogUpdate) Exec(ctx context.Context) error {
	_, err := malu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (malu *MemberAuthorizeLogUpdate) ExecX(ctx context.Context) {
	if err := malu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (malu *MemberAuthorizeLogUpdate) defaults() {
	if _, ok := malu.mutation.UpdateTime(); !ok {
		v := memberauthorizelog.UpdateDefaultUpdateTime()
		malu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (malu *MemberAuthorizeLogUpdate) check() error {
	if v, ok := malu.mutation.Device(); ok {
		if err := memberauthorizelog.DeviceValidator(v); err != nil {
			return &ValidationError{Name: "device", err: fmt.Errorf(`ent: validator failed for field "MemberAuthorizeLog.device": %w`, err)}
		}
	}
	if v, ok := malu.mutation.DeviceDetail(); ok {
		if err := memberauthorizelog.DeviceDetailValidator(v); err != nil {
			return &ValidationError{Name: "device_detail", err: fmt.Errorf(`ent: validator failed for field "MemberAuthorizeLog.device_detail": %w`, err)}
		}
	}
	if v, ok := malu.mutation.Ipv4(); ok {
		if err := memberauthorizelog.Ipv4Validator(v); err != nil {
			return &ValidationError{Name: "ipv4", err: fmt.Errorf(`ent: validator failed for field "MemberAuthorizeLog.ipv4": %w`, err)}
		}
	}
	if v, ok := malu.mutation.Ipv6(); ok {
		if err := memberauthorizelog.Ipv6Validator(v); err != nil {
			return &ValidationError{Name: "ipv6", err: fmt.Errorf(`ent: validator failed for field "MemberAuthorizeLog.ipv6": %w`, err)}
		}
	}
	if v, ok := malu.mutation.Snapshot(); ok {
		if err := memberauthorizelog.SnapshotValidator(v); err != nil {
			return &ValidationError{Name: "snapshot", err: fmt.Errorf(`ent: validator failed for field "MemberAuthorizeLog.snapshot": %w`, err)}
		}
	}
	return nil
}

func (malu *MemberAuthorizeLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := malu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(memberauthorizelog.Table, memberauthorizelog.Columns, sqlgraph.NewFieldSpec(memberauthorizelog.FieldID, field.TypeUUID))
	if ps := malu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := malu.mutation.UpdateTime(); ok {
		_spec.SetField(memberauthorizelog.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := malu.mutation.MemberID(); ok {
		_spec.SetField(memberauthorizelog.FieldMemberID, field.TypeUUID, value)
	}
	if value, ok := malu.mutation.Channel(); ok {
		_spec.SetField(memberauthorizelog.FieldChannel, field.TypeString, value)
	}
	if value, ok := malu.mutation.Device(); ok {
		_spec.SetField(memberauthorizelog.FieldDevice, field.TypeString, value)
	}
	if value, ok := malu.mutation.DeviceDetail(); ok {
		_spec.SetField(memberauthorizelog.FieldDeviceDetail, field.TypeString, value)
	}
	if value, ok := malu.mutation.Ipv4(); ok {
		_spec.SetField(memberauthorizelog.FieldIpv4, field.TypeString, value)
	}
	if value, ok := malu.mutation.Ipv6(); ok {
		_spec.SetField(memberauthorizelog.FieldIpv6, field.TypeString, value)
	}
	if value, ok := malu.mutation.Snapshot(); ok {
		_spec.SetField(memberauthorizelog.FieldSnapshot, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, malu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{memberauthorizelog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	malu.mutation.done = true
	return n, nil
}

// MemberAuthorizeLogUpdateOne is the builder for updating a single MemberAuthorizeLog entity.
type MemberAuthorizeLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MemberAuthorizeLogMutation
}

// SetUpdateTime sets the "update_time" field.
func (maluo *MemberAuthorizeLogUpdateOne) SetUpdateTime(t time.Time) *MemberAuthorizeLogUpdateOne {
	maluo.mutation.SetUpdateTime(t)
	return maluo
}

// SetMemberID sets the "member_id" field.
func (maluo *MemberAuthorizeLogUpdateOne) SetMemberID(u uuid.UUID) *MemberAuthorizeLogUpdateOne {
	maluo.mutation.SetMemberID(u)
	return maluo
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (maluo *MemberAuthorizeLogUpdateOne) SetNillableMemberID(u *uuid.UUID) *MemberAuthorizeLogUpdateOne {
	if u != nil {
		maluo.SetMemberID(*u)
	}
	return maluo
}

// SetChannel sets the "channel" field.
func (maluo *MemberAuthorizeLogUpdateOne) SetChannel(s string) *MemberAuthorizeLogUpdateOne {
	maluo.mutation.SetChannel(s)
	return maluo
}

// SetDevice sets the "device" field.
func (maluo *MemberAuthorizeLogUpdateOne) SetDevice(s string) *MemberAuthorizeLogUpdateOne {
	maluo.mutation.SetDevice(s)
	return maluo
}

// SetDeviceDetail sets the "device_detail" field.
func (maluo *MemberAuthorizeLogUpdateOne) SetDeviceDetail(s string) *MemberAuthorizeLogUpdateOne {
	maluo.mutation.SetDeviceDetail(s)
	return maluo
}

// SetIpv4 sets the "ipv4" field.
func (maluo *MemberAuthorizeLogUpdateOne) SetIpv4(s string) *MemberAuthorizeLogUpdateOne {
	maluo.mutation.SetIpv4(s)
	return maluo
}

// SetIpv6 sets the "ipv6" field.
func (maluo *MemberAuthorizeLogUpdateOne) SetIpv6(s string) *MemberAuthorizeLogUpdateOne {
	maluo.mutation.SetIpv6(s)
	return maluo
}

// SetSnapshot sets the "snapshot" field.
func (maluo *MemberAuthorizeLogUpdateOne) SetSnapshot(s string) *MemberAuthorizeLogUpdateOne {
	maluo.mutation.SetSnapshot(s)
	return maluo
}

// Mutation returns the MemberAuthorizeLogMutation object of the builder.
func (maluo *MemberAuthorizeLogUpdateOne) Mutation() *MemberAuthorizeLogMutation {
	return maluo.mutation
}

// Where appends a list predicates to the MemberAuthorizeLogUpdate builder.
func (maluo *MemberAuthorizeLogUpdateOne) Where(ps ...predicate.MemberAuthorizeLog) *MemberAuthorizeLogUpdateOne {
	maluo.mutation.Where(ps...)
	return maluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (maluo *MemberAuthorizeLogUpdateOne) Select(field string, fields ...string) *MemberAuthorizeLogUpdateOne {
	maluo.fields = append([]string{field}, fields...)
	return maluo
}

// Save executes the query and returns the updated MemberAuthorizeLog entity.
func (maluo *MemberAuthorizeLogUpdateOne) Save(ctx context.Context) (*MemberAuthorizeLog, error) {
	maluo.defaults()
	return withHooks[*MemberAuthorizeLog, MemberAuthorizeLogMutation](ctx, maluo.sqlSave, maluo.mutation, maluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (maluo *MemberAuthorizeLogUpdateOne) SaveX(ctx context.Context) *MemberAuthorizeLog {
	node, err := maluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (maluo *MemberAuthorizeLogUpdateOne) Exec(ctx context.Context) error {
	_, err := maluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (maluo *MemberAuthorizeLogUpdateOne) ExecX(ctx context.Context) {
	if err := maluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (maluo *MemberAuthorizeLogUpdateOne) defaults() {
	if _, ok := maluo.mutation.UpdateTime(); !ok {
		v := memberauthorizelog.UpdateDefaultUpdateTime()
		maluo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (maluo *MemberAuthorizeLogUpdateOne) check() error {
	if v, ok := maluo.mutation.Device(); ok {
		if err := memberauthorizelog.DeviceValidator(v); err != nil {
			return &ValidationError{Name: "device", err: fmt.Errorf(`ent: validator failed for field "MemberAuthorizeLog.device": %w`, err)}
		}
	}
	if v, ok := maluo.mutation.DeviceDetail(); ok {
		if err := memberauthorizelog.DeviceDetailValidator(v); err != nil {
			return &ValidationError{Name: "device_detail", err: fmt.Errorf(`ent: validator failed for field "MemberAuthorizeLog.device_detail": %w`, err)}
		}
	}
	if v, ok := maluo.mutation.Ipv4(); ok {
		if err := memberauthorizelog.Ipv4Validator(v); err != nil {
			return &ValidationError{Name: "ipv4", err: fmt.Errorf(`ent: validator failed for field "MemberAuthorizeLog.ipv4": %w`, err)}
		}
	}
	if v, ok := maluo.mutation.Ipv6(); ok {
		if err := memberauthorizelog.Ipv6Validator(v); err != nil {
			return &ValidationError{Name: "ipv6", err: fmt.Errorf(`ent: validator failed for field "MemberAuthorizeLog.ipv6": %w`, err)}
		}
	}
	if v, ok := maluo.mutation.Snapshot(); ok {
		if err := memberauthorizelog.SnapshotValidator(v); err != nil {
			return &ValidationError{Name: "snapshot", err: fmt.Errorf(`ent: validator failed for field "MemberAuthorizeLog.snapshot": %w`, err)}
		}
	}
	return nil
}

func (maluo *MemberAuthorizeLogUpdateOne) sqlSave(ctx context.Context) (_node *MemberAuthorizeLog, err error) {
	if err := maluo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(memberauthorizelog.Table, memberauthorizelog.Columns, sqlgraph.NewFieldSpec(memberauthorizelog.FieldID, field.TypeUUID))
	id, ok := maluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MemberAuthorizeLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := maluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, memberauthorizelog.FieldID)
		for _, f := range fields {
			if !memberauthorizelog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != memberauthorizelog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := maluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := maluo.mutation.UpdateTime(); ok {
		_spec.SetField(memberauthorizelog.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := maluo.mutation.MemberID(); ok {
		_spec.SetField(memberauthorizelog.FieldMemberID, field.TypeUUID, value)
	}
	if value, ok := maluo.mutation.Channel(); ok {
		_spec.SetField(memberauthorizelog.FieldChannel, field.TypeString, value)
	}
	if value, ok := maluo.mutation.Device(); ok {
		_spec.SetField(memberauthorizelog.FieldDevice, field.TypeString, value)
	}
	if value, ok := maluo.mutation.DeviceDetail(); ok {
		_spec.SetField(memberauthorizelog.FieldDeviceDetail, field.TypeString, value)
	}
	if value, ok := maluo.mutation.Ipv4(); ok {
		_spec.SetField(memberauthorizelog.FieldIpv4, field.TypeString, value)
	}
	if value, ok := maluo.mutation.Ipv6(); ok {
		_spec.SetField(memberauthorizelog.FieldIpv6, field.TypeString, value)
	}
	if value, ok := maluo.mutation.Snapshot(); ok {
		_spec.SetField(memberauthorizelog.FieldSnapshot, field.TypeString, value)
	}
	_node = &MemberAuthorizeLog{config: maluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, maluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{memberauthorizelog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	maluo.mutation.done = true
	return _node, nil
}
