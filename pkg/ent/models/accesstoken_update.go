// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/accesstoken"
	"ent/models/predicate"
	"ent/utils/timeutil"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccessTokenUpdate is the builder for updating AccessToken entities.
type AccessTokenUpdate struct {
	config
	hooks    []Hook
	mutation *AccessTokenMutation
}

// Where appends a list predicates to the AccessTokenUpdate builder.
func (atu *AccessTokenUpdate) Where(ps ...predicate.AccessToken) *AccessTokenUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetUserID sets the "user_id" field.
func (atu *AccessTokenUpdate) SetUserID(i int64) *AccessTokenUpdate {
	atu.mutation.ResetUserID()
	atu.mutation.SetUserID(i)
	return atu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (atu *AccessTokenUpdate) SetNillableUserID(i *int64) *AccessTokenUpdate {
	if i != nil {
		atu.SetUserID(*i)
	}
	return atu
}

// AddUserID adds i to the "user_id" field.
func (atu *AccessTokenUpdate) AddUserID(i int64) *AccessTokenUpdate {
	atu.mutation.AddUserID(i)
	return atu
}

// SetName sets the "name" field.
func (atu *AccessTokenUpdate) SetName(s string) *AccessTokenUpdate {
	atu.mutation.SetName(s)
	return atu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (atu *AccessTokenUpdate) SetNillableName(s *string) *AccessTokenUpdate {
	if s != nil {
		atu.SetName(*s)
	}
	return atu
}

// SetToken sets the "token" field.
func (atu *AccessTokenUpdate) SetToken(s string) *AccessTokenUpdate {
	atu.mutation.SetToken(s)
	return atu
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (atu *AccessTokenUpdate) SetNillableToken(s *string) *AccessTokenUpdate {
	if s != nil {
		atu.SetToken(*s)
	}
	return atu
}

// SetTokenHash sets the "token_hash" field.
func (atu *AccessTokenUpdate) SetTokenHash(s string) *AccessTokenUpdate {
	atu.mutation.SetTokenHash(s)
	return atu
}

// SetNillableTokenHash sets the "token_hash" field if the given value is not nil.
func (atu *AccessTokenUpdate) SetNillableTokenHash(s *string) *AccessTokenUpdate {
	if s != nil {
		atu.SetTokenHash(*s)
	}
	return atu
}

// SetTokenSalt sets the "token_salt" field.
func (atu *AccessTokenUpdate) SetTokenSalt(s string) *AccessTokenUpdate {
	atu.mutation.SetTokenSalt(s)
	return atu
}

// SetNillableTokenSalt sets the "token_salt" field if the given value is not nil.
func (atu *AccessTokenUpdate) SetNillableTokenSalt(s *string) *AccessTokenUpdate {
	if s != nil {
		atu.SetTokenSalt(*s)
	}
	return atu
}

// SetTokenLastEight sets the "token_last_eight" field.
func (atu *AccessTokenUpdate) SetTokenLastEight(s string) *AccessTokenUpdate {
	atu.mutation.SetTokenLastEight(s)
	return atu
}

// SetNillableTokenLastEight sets the "token_last_eight" field if the given value is not nil.
func (atu *AccessTokenUpdate) SetNillableTokenLastEight(s *string) *AccessTokenUpdate {
	if s != nil {
		atu.SetTokenLastEight(*s)
	}
	return atu
}

// SetScope sets the "scope" field.
func (atu *AccessTokenUpdate) SetScope(s string) *AccessTokenUpdate {
	atu.mutation.SetScope(s)
	return atu
}

// SetNillableScope sets the "scope" field if the given value is not nil.
func (atu *AccessTokenUpdate) SetNillableScope(s *string) *AccessTokenUpdate {
	if s != nil {
		atu.SetScope(*s)
	}
	return atu
}

// SetHasRecentActivity sets the "has_recent_activity" field.
func (atu *AccessTokenUpdate) SetHasRecentActivity(s string) *AccessTokenUpdate {
	atu.mutation.SetHasRecentActivity(s)
	return atu
}

// SetNillableHasRecentActivity sets the "has_recent_activity" field if the given value is not nil.
func (atu *AccessTokenUpdate) SetNillableHasRecentActivity(s *string) *AccessTokenUpdate {
	if s != nil {
		atu.SetHasRecentActivity(*s)
	}
	return atu
}

// SetHasUsed sets the "has_used" field.
func (atu *AccessTokenUpdate) SetHasUsed(s string) *AccessTokenUpdate {
	atu.mutation.SetHasUsed(s)
	return atu
}

// SetNillableHasUsed sets the "has_used" field if the given value is not nil.
func (atu *AccessTokenUpdate) SetNillableHasUsed(s *string) *AccessTokenUpdate {
	if s != nil {
		atu.SetHasUsed(*s)
	}
	return atu
}

// SetUpdateTime sets the "update_time" field.
func (atu *AccessTokenUpdate) SetUpdateTime(ts timeutil.TimeStamp) *AccessTokenUpdate {
	atu.mutation.ResetUpdateTime()
	atu.mutation.SetUpdateTime(ts)
	return atu
}

// AddUpdateTime adds ts to the "update_time" field.
func (atu *AccessTokenUpdate) AddUpdateTime(ts timeutil.TimeStamp) *AccessTokenUpdate {
	atu.mutation.AddUpdateTime(ts)
	return atu
}

// Mutation returns the AccessTokenMutation object of the builder.
func (atu *AccessTokenUpdate) Mutation() *AccessTokenMutation {
	return atu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *AccessTokenUpdate) Save(ctx context.Context) (int, error) {
	atu.defaults()
	return withHooks(ctx, atu.sqlSave, atu.mutation, atu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atu *AccessTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *AccessTokenUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *AccessTokenUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atu *AccessTokenUpdate) defaults() {
	if _, ok := atu.mutation.UpdateTime(); !ok {
		v := accesstoken.UpdateDefaultUpdateTime()
		atu.mutation.SetUpdateTime(v)
	}
}

func (atu *AccessTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(accesstoken.Table, accesstoken.Columns, sqlgraph.NewFieldSpec(accesstoken.FieldID, field.TypeInt64))
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atu.mutation.UserID(); ok {
		_spec.SetField(accesstoken.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := atu.mutation.AddedUserID(); ok {
		_spec.AddField(accesstoken.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := atu.mutation.Name(); ok {
		_spec.SetField(accesstoken.FieldName, field.TypeString, value)
	}
	if value, ok := atu.mutation.Token(); ok {
		_spec.SetField(accesstoken.FieldToken, field.TypeString, value)
	}
	if value, ok := atu.mutation.TokenHash(); ok {
		_spec.SetField(accesstoken.FieldTokenHash, field.TypeString, value)
	}
	if value, ok := atu.mutation.TokenSalt(); ok {
		_spec.SetField(accesstoken.FieldTokenSalt, field.TypeString, value)
	}
	if value, ok := atu.mutation.TokenLastEight(); ok {
		_spec.SetField(accesstoken.FieldTokenLastEight, field.TypeString, value)
	}
	if value, ok := atu.mutation.Scope(); ok {
		_spec.SetField(accesstoken.FieldScope, field.TypeString, value)
	}
	if value, ok := atu.mutation.HasRecentActivity(); ok {
		_spec.SetField(accesstoken.FieldHasRecentActivity, field.TypeString, value)
	}
	if value, ok := atu.mutation.HasUsed(); ok {
		_spec.SetField(accesstoken.FieldHasUsed, field.TypeString, value)
	}
	if value, ok := atu.mutation.UpdateTime(); ok {
		_spec.SetField(accesstoken.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := atu.mutation.AddedUpdateTime(); ok {
		_spec.AddField(accesstoken.FieldUpdateTime, field.TypeInt64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accesstoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	atu.mutation.done = true
	return n, nil
}

// AccessTokenUpdateOne is the builder for updating a single AccessToken entity.
type AccessTokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccessTokenMutation
}

// SetUserID sets the "user_id" field.
func (atuo *AccessTokenUpdateOne) SetUserID(i int64) *AccessTokenUpdateOne {
	atuo.mutation.ResetUserID()
	atuo.mutation.SetUserID(i)
	return atuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (atuo *AccessTokenUpdateOne) SetNillableUserID(i *int64) *AccessTokenUpdateOne {
	if i != nil {
		atuo.SetUserID(*i)
	}
	return atuo
}

// AddUserID adds i to the "user_id" field.
func (atuo *AccessTokenUpdateOne) AddUserID(i int64) *AccessTokenUpdateOne {
	atuo.mutation.AddUserID(i)
	return atuo
}

// SetName sets the "name" field.
func (atuo *AccessTokenUpdateOne) SetName(s string) *AccessTokenUpdateOne {
	atuo.mutation.SetName(s)
	return atuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (atuo *AccessTokenUpdateOne) SetNillableName(s *string) *AccessTokenUpdateOne {
	if s != nil {
		atuo.SetName(*s)
	}
	return atuo
}

// SetToken sets the "token" field.
func (atuo *AccessTokenUpdateOne) SetToken(s string) *AccessTokenUpdateOne {
	atuo.mutation.SetToken(s)
	return atuo
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (atuo *AccessTokenUpdateOne) SetNillableToken(s *string) *AccessTokenUpdateOne {
	if s != nil {
		atuo.SetToken(*s)
	}
	return atuo
}

// SetTokenHash sets the "token_hash" field.
func (atuo *AccessTokenUpdateOne) SetTokenHash(s string) *AccessTokenUpdateOne {
	atuo.mutation.SetTokenHash(s)
	return atuo
}

// SetNillableTokenHash sets the "token_hash" field if the given value is not nil.
func (atuo *AccessTokenUpdateOne) SetNillableTokenHash(s *string) *AccessTokenUpdateOne {
	if s != nil {
		atuo.SetTokenHash(*s)
	}
	return atuo
}

// SetTokenSalt sets the "token_salt" field.
func (atuo *AccessTokenUpdateOne) SetTokenSalt(s string) *AccessTokenUpdateOne {
	atuo.mutation.SetTokenSalt(s)
	return atuo
}

// SetNillableTokenSalt sets the "token_salt" field if the given value is not nil.
func (atuo *AccessTokenUpdateOne) SetNillableTokenSalt(s *string) *AccessTokenUpdateOne {
	if s != nil {
		atuo.SetTokenSalt(*s)
	}
	return atuo
}

// SetTokenLastEight sets the "token_last_eight" field.
func (atuo *AccessTokenUpdateOne) SetTokenLastEight(s string) *AccessTokenUpdateOne {
	atuo.mutation.SetTokenLastEight(s)
	return atuo
}

// SetNillableTokenLastEight sets the "token_last_eight" field if the given value is not nil.
func (atuo *AccessTokenUpdateOne) SetNillableTokenLastEight(s *string) *AccessTokenUpdateOne {
	if s != nil {
		atuo.SetTokenLastEight(*s)
	}
	return atuo
}

// SetScope sets the "scope" field.
func (atuo *AccessTokenUpdateOne) SetScope(s string) *AccessTokenUpdateOne {
	atuo.mutation.SetScope(s)
	return atuo
}

// SetNillableScope sets the "scope" field if the given value is not nil.
func (atuo *AccessTokenUpdateOne) SetNillableScope(s *string) *AccessTokenUpdateOne {
	if s != nil {
		atuo.SetScope(*s)
	}
	return atuo
}

// SetHasRecentActivity sets the "has_recent_activity" field.
func (atuo *AccessTokenUpdateOne) SetHasRecentActivity(s string) *AccessTokenUpdateOne {
	atuo.mutation.SetHasRecentActivity(s)
	return atuo
}

// SetNillableHasRecentActivity sets the "has_recent_activity" field if the given value is not nil.
func (atuo *AccessTokenUpdateOne) SetNillableHasRecentActivity(s *string) *AccessTokenUpdateOne {
	if s != nil {
		atuo.SetHasRecentActivity(*s)
	}
	return atuo
}

// SetHasUsed sets the "has_used" field.
func (atuo *AccessTokenUpdateOne) SetHasUsed(s string) *AccessTokenUpdateOne {
	atuo.mutation.SetHasUsed(s)
	return atuo
}

// SetNillableHasUsed sets the "has_used" field if the given value is not nil.
func (atuo *AccessTokenUpdateOne) SetNillableHasUsed(s *string) *AccessTokenUpdateOne {
	if s != nil {
		atuo.SetHasUsed(*s)
	}
	return atuo
}

// SetUpdateTime sets the "update_time" field.
func (atuo *AccessTokenUpdateOne) SetUpdateTime(ts timeutil.TimeStamp) *AccessTokenUpdateOne {
	atuo.mutation.ResetUpdateTime()
	atuo.mutation.SetUpdateTime(ts)
	return atuo
}

// AddUpdateTime adds ts to the "update_time" field.
func (atuo *AccessTokenUpdateOne) AddUpdateTime(ts timeutil.TimeStamp) *AccessTokenUpdateOne {
	atuo.mutation.AddUpdateTime(ts)
	return atuo
}

// Mutation returns the AccessTokenMutation object of the builder.
func (atuo *AccessTokenUpdateOne) Mutation() *AccessTokenMutation {
	return atuo.mutation
}

// Where appends a list predicates to the AccessTokenUpdate builder.
func (atuo *AccessTokenUpdateOne) Where(ps ...predicate.AccessToken) *AccessTokenUpdateOne {
	atuo.mutation.Where(ps...)
	return atuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *AccessTokenUpdateOne) Select(field string, fields ...string) *AccessTokenUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated AccessToken entity.
func (atuo *AccessTokenUpdateOne) Save(ctx context.Context) (*AccessToken, error) {
	atuo.defaults()
	return withHooks(ctx, atuo.sqlSave, atuo.mutation, atuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *AccessTokenUpdateOne) SaveX(ctx context.Context) *AccessToken {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *AccessTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *AccessTokenUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atuo *AccessTokenUpdateOne) defaults() {
	if _, ok := atuo.mutation.UpdateTime(); !ok {
		v := accesstoken.UpdateDefaultUpdateTime()
		atuo.mutation.SetUpdateTime(v)
	}
}

func (atuo *AccessTokenUpdateOne) sqlSave(ctx context.Context) (_node *AccessToken, err error) {
	_spec := sqlgraph.NewUpdateSpec(accesstoken.Table, accesstoken.Columns, sqlgraph.NewFieldSpec(accesstoken.FieldID, field.TypeInt64))
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "AccessToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accesstoken.FieldID)
		for _, f := range fields {
			if !accesstoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != accesstoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atuo.mutation.UserID(); ok {
		_spec.SetField(accesstoken.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := atuo.mutation.AddedUserID(); ok {
		_spec.AddField(accesstoken.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := atuo.mutation.Name(); ok {
		_spec.SetField(accesstoken.FieldName, field.TypeString, value)
	}
	if value, ok := atuo.mutation.Token(); ok {
		_spec.SetField(accesstoken.FieldToken, field.TypeString, value)
	}
	if value, ok := atuo.mutation.TokenHash(); ok {
		_spec.SetField(accesstoken.FieldTokenHash, field.TypeString, value)
	}
	if value, ok := atuo.mutation.TokenSalt(); ok {
		_spec.SetField(accesstoken.FieldTokenSalt, field.TypeString, value)
	}
	if value, ok := atuo.mutation.TokenLastEight(); ok {
		_spec.SetField(accesstoken.FieldTokenLastEight, field.TypeString, value)
	}
	if value, ok := atuo.mutation.Scope(); ok {
		_spec.SetField(accesstoken.FieldScope, field.TypeString, value)
	}
	if value, ok := atuo.mutation.HasRecentActivity(); ok {
		_spec.SetField(accesstoken.FieldHasRecentActivity, field.TypeString, value)
	}
	if value, ok := atuo.mutation.HasUsed(); ok {
		_spec.SetField(accesstoken.FieldHasUsed, field.TypeString, value)
	}
	if value, ok := atuo.mutation.UpdateTime(); ok {
		_spec.SetField(accesstoken.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := atuo.mutation.AddedUpdateTime(); ok {
		_spec.AddField(accesstoken.FieldUpdateTime, field.TypeInt64, value)
	}
	_node = &AccessToken{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accesstoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	atuo.mutation.done = true
	return _node, nil
}
