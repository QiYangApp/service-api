// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/userrelatedrole"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserRelatedRoleUpdate is the builder for updating UserRelatedRole entities.
type UserRelatedRoleUpdate struct {
	config
	hooks    []Hook
	mutation *UserRelatedRoleMutation
}

// Where appends a list predicates to the UserRelatedRoleUpdate builder.
func (urru *UserRelatedRoleUpdate) Where(ps ...predicate.UserRelatedRole) *UserRelatedRoleUpdate {
	urru.mutation.Where(ps...)
	return urru
}

// SetUserID sets the "user_id" field.
func (urru *UserRelatedRoleUpdate) SetUserID(u uuid.UUID) *UserRelatedRoleUpdate {
	urru.mutation.SetUserID(u)
	return urru
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (urru *UserRelatedRoleUpdate) SetNillableUserID(u *uuid.UUID) *UserRelatedRoleUpdate {
	if u != nil {
		urru.SetUserID(*u)
	}
	return urru
}

// SetRoleID sets the "role_id" field.
func (urru *UserRelatedRoleUpdate) SetRoleID(u uuid.UUID) *UserRelatedRoleUpdate {
	urru.mutation.SetRoleID(u)
	return urru
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (urru *UserRelatedRoleUpdate) SetNillableRoleID(u *uuid.UUID) *UserRelatedRoleUpdate {
	if u != nil {
		urru.SetRoleID(*u)
	}
	return urru
}

// SetUpdateTime sets the "update_time" field.
func (urru *UserRelatedRoleUpdate) SetUpdateTime(t time.Time) *UserRelatedRoleUpdate {
	urru.mutation.SetUpdateTime(t)
	return urru
}

// Mutation returns the UserRelatedRoleMutation object of the builder.
func (urru *UserRelatedRoleUpdate) Mutation() *UserRelatedRoleMutation {
	return urru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (urru *UserRelatedRoleUpdate) Save(ctx context.Context) (int, error) {
	urru.defaults()
	return withHooks(ctx, urru.sqlSave, urru.mutation, urru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (urru *UserRelatedRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := urru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (urru *UserRelatedRoleUpdate) Exec(ctx context.Context) error {
	_, err := urru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urru *UserRelatedRoleUpdate) ExecX(ctx context.Context) {
	if err := urru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (urru *UserRelatedRoleUpdate) defaults() {
	if _, ok := urru.mutation.UpdateTime(); !ok {
		v := userrelatedrole.UpdateDefaultUpdateTime()
		urru.mutation.SetUpdateTime(v)
	}
}

func (urru *UserRelatedRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(userrelatedrole.Table, userrelatedrole.Columns, sqlgraph.NewFieldSpec(userrelatedrole.FieldID, field.TypeInt))
	if ps := urru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := urru.mutation.UserID(); ok {
		_spec.SetField(userrelatedrole.FieldUserID, field.TypeUUID, value)
	}
	if value, ok := urru.mutation.RoleID(); ok {
		_spec.SetField(userrelatedrole.FieldRoleID, field.TypeUUID, value)
	}
	if value, ok := urru.mutation.UpdateTime(); ok {
		_spec.SetField(userrelatedrole.FieldUpdateTime, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, urru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userrelatedrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	urru.mutation.done = true
	return n, nil
}

// UserRelatedRoleUpdateOne is the builder for updating a single UserRelatedRole entity.
type UserRelatedRoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserRelatedRoleMutation
}

// SetUserID sets the "user_id" field.
func (urruo *UserRelatedRoleUpdateOne) SetUserID(u uuid.UUID) *UserRelatedRoleUpdateOne {
	urruo.mutation.SetUserID(u)
	return urruo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (urruo *UserRelatedRoleUpdateOne) SetNillableUserID(u *uuid.UUID) *UserRelatedRoleUpdateOne {
	if u != nil {
		urruo.SetUserID(*u)
	}
	return urruo
}

// SetRoleID sets the "role_id" field.
func (urruo *UserRelatedRoleUpdateOne) SetRoleID(u uuid.UUID) *UserRelatedRoleUpdateOne {
	urruo.mutation.SetRoleID(u)
	return urruo
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (urruo *UserRelatedRoleUpdateOne) SetNillableRoleID(u *uuid.UUID) *UserRelatedRoleUpdateOne {
	if u != nil {
		urruo.SetRoleID(*u)
	}
	return urruo
}

// SetUpdateTime sets the "update_time" field.
func (urruo *UserRelatedRoleUpdateOne) SetUpdateTime(t time.Time) *UserRelatedRoleUpdateOne {
	urruo.mutation.SetUpdateTime(t)
	return urruo
}

// Mutation returns the UserRelatedRoleMutation object of the builder.
func (urruo *UserRelatedRoleUpdateOne) Mutation() *UserRelatedRoleMutation {
	return urruo.mutation
}

// Where appends a list predicates to the UserRelatedRoleUpdate builder.
func (urruo *UserRelatedRoleUpdateOne) Where(ps ...predicate.UserRelatedRole) *UserRelatedRoleUpdateOne {
	urruo.mutation.Where(ps...)
	return urruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (urruo *UserRelatedRoleUpdateOne) Select(field string, fields ...string) *UserRelatedRoleUpdateOne {
	urruo.fields = append([]string{field}, fields...)
	return urruo
}

// Save executes the query and returns the updated UserRelatedRole entity.
func (urruo *UserRelatedRoleUpdateOne) Save(ctx context.Context) (*UserRelatedRole, error) {
	urruo.defaults()
	return withHooks(ctx, urruo.sqlSave, urruo.mutation, urruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (urruo *UserRelatedRoleUpdateOne) SaveX(ctx context.Context) *UserRelatedRole {
	node, err := urruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (urruo *UserRelatedRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := urruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urruo *UserRelatedRoleUpdateOne) ExecX(ctx context.Context) {
	if err := urruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (urruo *UserRelatedRoleUpdateOne) defaults() {
	if _, ok := urruo.mutation.UpdateTime(); !ok {
		v := userrelatedrole.UpdateDefaultUpdateTime()
		urruo.mutation.SetUpdateTime(v)
	}
}

func (urruo *UserRelatedRoleUpdateOne) sqlSave(ctx context.Context) (_node *UserRelatedRole, err error) {
	_spec := sqlgraph.NewUpdateSpec(userrelatedrole.Table, userrelatedrole.Columns, sqlgraph.NewFieldSpec(userrelatedrole.FieldID, field.TypeInt))
	id, ok := urruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "UserRelatedRole.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := urruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userrelatedrole.FieldID)
		for _, f := range fields {
			if !userrelatedrole.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != userrelatedrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := urruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := urruo.mutation.UserID(); ok {
		_spec.SetField(userrelatedrole.FieldUserID, field.TypeUUID, value)
	}
	if value, ok := urruo.mutation.RoleID(); ok {
		_spec.SetField(userrelatedrole.FieldRoleID, field.TypeUUID, value)
	}
	if value, ok := urruo.mutation.UpdateTime(); ok {
		_spec.SetField(userrelatedrole.FieldUpdateTime, field.TypeTime, value)
	}
	_node = &UserRelatedRole{config: urruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, urruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userrelatedrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	urruo.mutation.done = true
	return _node, nil
}
