// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"service-api/src/models/ent/memberrole"
	"service-api/src/models/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberRoleUpdate is the builder for updating MemberRole entities.
type MemberRoleUpdate struct {
	config
	hooks    []Hook
	mutation *MemberRoleMutation
}

// Where appends a list predicates to the MemberRoleUpdate builder.
func (mru *MemberRoleUpdate) Where(ps ...predicate.MemberRole) *MemberRoleUpdate {
	mru.mutation.Where(ps...)
	return mru
}

// SetUpdateTime sets the "update_time" field.
func (mru *MemberRoleUpdate) SetUpdateTime(t time.Time) *MemberRoleUpdate {
	mru.mutation.SetUpdateTime(t)
	return mru
}

// SetRoleName sets the "role_name" field.
func (mru *MemberRoleUpdate) SetRoleName(s string) *MemberRoleUpdate {
	mru.mutation.SetRoleName(s)
	return mru
}

// SetNillableRoleName sets the "role_name" field if the given value is not nil.
func (mru *MemberRoleUpdate) SetNillableRoleName(s *string) *MemberRoleUpdate {
	if s != nil {
		mru.SetRoleName(*s)
	}
	return mru
}

// SetState sets the "state" field.
func (mru *MemberRoleUpdate) SetState(s string) *MemberRoleUpdate {
	mru.mutation.SetState(s)
	return mru
}

// SetNillableState sets the "state" field if the given value is not nil.
func (mru *MemberRoleUpdate) SetNillableState(s *string) *MemberRoleUpdate {
	if s != nil {
		mru.SetState(*s)
	}
	return mru
}

// Mutation returns the MemberRoleMutation object of the builder.
func (mru *MemberRoleUpdate) Mutation() *MemberRoleMutation {
	return mru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mru *MemberRoleUpdate) Save(ctx context.Context) (int, error) {
	mru.defaults()
	return withHooks(ctx, mru.sqlSave, mru.mutation, mru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mru *MemberRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := mru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mru *MemberRoleUpdate) Exec(ctx context.Context) error {
	_, err := mru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mru *MemberRoleUpdate) ExecX(ctx context.Context) {
	if err := mru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mru *MemberRoleUpdate) defaults() {
	if _, ok := mru.mutation.UpdateTime(); !ok {
		v := memberrole.UpdateDefaultUpdateTime()
		mru.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mru *MemberRoleUpdate) check() error {
	if v, ok := mru.mutation.RoleName(); ok {
		if err := memberrole.RoleNameValidator(v); err != nil {
			return &ValidationError{Name: "role_name", err: fmt.Errorf(`ent: validator failed for field "MemberRole.role_name": %w`, err)}
		}
	}
	if v, ok := mru.mutation.State(); ok {
		if err := memberrole.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "MemberRole.state": %w`, err)}
		}
	}
	return nil
}

func (mru *MemberRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(memberrole.Table, memberrole.Columns, sqlgraph.NewFieldSpec(memberrole.FieldID, field.TypeUUID))
	if ps := mru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mru.mutation.UpdateTime(); ok {
		_spec.SetField(memberrole.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := mru.mutation.RoleName(); ok {
		_spec.SetField(memberrole.FieldRoleName, field.TypeString, value)
	}
	if value, ok := mru.mutation.State(); ok {
		_spec.SetField(memberrole.FieldState, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{memberrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mru.mutation.done = true
	return n, nil
}

// MemberRoleUpdateOne is the builder for updating a single MemberRole entity.
type MemberRoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MemberRoleMutation
}

// SetUpdateTime sets the "update_time" field.
func (mruo *MemberRoleUpdateOne) SetUpdateTime(t time.Time) *MemberRoleUpdateOne {
	mruo.mutation.SetUpdateTime(t)
	return mruo
}

// SetRoleName sets the "role_name" field.
func (mruo *MemberRoleUpdateOne) SetRoleName(s string) *MemberRoleUpdateOne {
	mruo.mutation.SetRoleName(s)
	return mruo
}

// SetNillableRoleName sets the "role_name" field if the given value is not nil.
func (mruo *MemberRoleUpdateOne) SetNillableRoleName(s *string) *MemberRoleUpdateOne {
	if s != nil {
		mruo.SetRoleName(*s)
	}
	return mruo
}

// SetState sets the "state" field.
func (mruo *MemberRoleUpdateOne) SetState(s string) *MemberRoleUpdateOne {
	mruo.mutation.SetState(s)
	return mruo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (mruo *MemberRoleUpdateOne) SetNillableState(s *string) *MemberRoleUpdateOne {
	if s != nil {
		mruo.SetState(*s)
	}
	return mruo
}

// Mutation returns the MemberRoleMutation object of the builder.
func (mruo *MemberRoleUpdateOne) Mutation() *MemberRoleMutation {
	return mruo.mutation
}

// Where appends a list predicates to the MemberRoleUpdate builder.
func (mruo *MemberRoleUpdateOne) Where(ps ...predicate.MemberRole) *MemberRoleUpdateOne {
	mruo.mutation.Where(ps...)
	return mruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mruo *MemberRoleUpdateOne) Select(field string, fields ...string) *MemberRoleUpdateOne {
	mruo.fields = append([]string{field}, fields...)
	return mruo
}

// Save executes the query and returns the updated MemberRole entity.
func (mruo *MemberRoleUpdateOne) Save(ctx context.Context) (*MemberRole, error) {
	mruo.defaults()
	return withHooks(ctx, mruo.sqlSave, mruo.mutation, mruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mruo *MemberRoleUpdateOne) SaveX(ctx context.Context) *MemberRole {
	node, err := mruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mruo *MemberRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := mruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mruo *MemberRoleUpdateOne) ExecX(ctx context.Context) {
	if err := mruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mruo *MemberRoleUpdateOne) defaults() {
	if _, ok := mruo.mutation.UpdateTime(); !ok {
		v := memberrole.UpdateDefaultUpdateTime()
		mruo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mruo *MemberRoleUpdateOne) check() error {
	if v, ok := mruo.mutation.RoleName(); ok {
		if err := memberrole.RoleNameValidator(v); err != nil {
			return &ValidationError{Name: "role_name", err: fmt.Errorf(`ent: validator failed for field "MemberRole.role_name": %w`, err)}
		}
	}
	if v, ok := mruo.mutation.State(); ok {
		if err := memberrole.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "MemberRole.state": %w`, err)}
		}
	}
	return nil
}

func (mruo *MemberRoleUpdateOne) sqlSave(ctx context.Context) (_node *MemberRole, err error) {
	if err := mruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(memberrole.Table, memberrole.Columns, sqlgraph.NewFieldSpec(memberrole.FieldID, field.TypeUUID))
	id, ok := mruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MemberRole.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, memberrole.FieldID)
		for _, f := range fields {
			if !memberrole.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != memberrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mruo.mutation.UpdateTime(); ok {
		_spec.SetField(memberrole.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := mruo.mutation.RoleName(); ok {
		_spec.SetField(memberrole.FieldRoleName, field.TypeString, value)
	}
	if value, ok := mruo.mutation.State(); ok {
		_spec.SetField(memberrole.FieldState, field.TypeString, value)
	}
	_node = &MemberRole{config: mruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{memberrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mruo.mutation.done = true
	return _node, nil
}
