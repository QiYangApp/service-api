// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"service-api/src/models/ent/member"
	"service-api/src/models/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberUpdate is the builder for updating Member entities.
type MemberUpdate struct {
	config
	hooks    []Hook
	mutation *MemberMutation
}

// Where appends a list predicates to the MemberUpdate builder.
func (mu *MemberUpdate) Where(ps ...predicate.Member) *MemberUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUpdateTime sets the "update_time" field.
func (mu *MemberUpdate) SetUpdateTime(t time.Time) *MemberUpdate {
	mu.mutation.SetUpdateTime(t)
	return mu
}

// SetEmail sets the "email" field.
func (mu *MemberUpdate) SetEmail(s string) *MemberUpdate {
	mu.mutation.SetEmail(s)
	return mu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableEmail(s *string) *MemberUpdate {
	if s != nil {
		mu.SetEmail(*s)
	}
	return mu
}

// SetAvatar sets the "avatar" field.
func (mu *MemberUpdate) SetAvatar(s string) *MemberUpdate {
	mu.mutation.SetAvatar(s)
	return mu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableAvatar(s *string) *MemberUpdate {
	if s != nil {
		mu.SetAvatar(*s)
	}
	return mu
}

// SetPasswordSing sets the "password_sing" field.
func (mu *MemberUpdate) SetPasswordSing(s string) *MemberUpdate {
	mu.mutation.SetPasswordSing(s)
	return mu
}

// SetNillablePasswordSing sets the "password_sing" field if the given value is not nil.
func (mu *MemberUpdate) SetNillablePasswordSing(s *string) *MemberUpdate {
	if s != nil {
		mu.SetPasswordSing(*s)
	}
	return mu
}

// SetPassword sets the "password" field.
func (mu *MemberUpdate) SetPassword(s string) *MemberUpdate {
	mu.mutation.SetPassword(s)
	return mu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (mu *MemberUpdate) SetNillablePassword(s *string) *MemberUpdate {
	if s != nil {
		mu.SetPassword(*s)
	}
	return mu
}

// SetMobile sets the "mobile" field.
func (mu *MemberUpdate) SetMobile(s string) *MemberUpdate {
	mu.mutation.SetMobile(s)
	return mu
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableMobile(s *string) *MemberUpdate {
	if s != nil {
		mu.SetMobile(*s)
	}
	return mu
}

// SetNickname sets the "nickname" field.
func (mu *MemberUpdate) SetNickname(s string) *MemberUpdate {
	mu.mutation.SetNickname(s)
	return mu
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableNickname(s *string) *MemberUpdate {
	if s != nil {
		mu.SetNickname(*s)
	}
	return mu
}

// SetState sets the "state" field.
func (mu *MemberUpdate) SetState(s string) *MemberUpdate {
	mu.mutation.SetState(s)
	return mu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableState(s *string) *MemberUpdate {
	if s != nil {
		mu.SetState(*s)
	}
	return mu
}

// Mutation returns the MemberMutation object of the builder.
func (mu *MemberUpdate) Mutation() *MemberMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MemberUpdate) Save(ctx context.Context) (int, error) {
	mu.defaults()
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MemberUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MemberUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MemberUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MemberUpdate) defaults() {
	if _, ok := mu.mutation.UpdateTime(); !ok {
		v := member.UpdateDefaultUpdateTime()
		mu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MemberUpdate) check() error {
	if v, ok := mu.mutation.Email(); ok {
		if err := member.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Member.email": %w`, err)}
		}
	}
	if v, ok := mu.mutation.PasswordSing(); ok {
		if err := member.PasswordSingValidator(v); err != nil {
			return &ValidationError{Name: "password_sing", err: fmt.Errorf(`ent: validator failed for field "Member.password_sing": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Password(); ok {
		if err := member.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Member.password": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Mobile(); ok {
		if err := member.MobileValidator(v); err != nil {
			return &ValidationError{Name: "mobile", err: fmt.Errorf(`ent: validator failed for field "Member.mobile": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Nickname(); ok {
		if err := member.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "Member.nickname": %w`, err)}
		}
	}
	if v, ok := mu.mutation.State(); ok {
		if err := member.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Member.state": %w`, err)}
		}
	}
	return nil
}

func (mu *MemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(member.Table, member.Columns, sqlgraph.NewFieldSpec(member.FieldID, field.TypeUUID))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.UpdateTime(); ok {
		_spec.SetField(member.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := mu.mutation.Email(); ok {
		_spec.SetField(member.FieldEmail, field.TypeString, value)
	}
	if value, ok := mu.mutation.Avatar(); ok {
		_spec.SetField(member.FieldAvatar, field.TypeString, value)
	}
	if value, ok := mu.mutation.PasswordSing(); ok {
		_spec.SetField(member.FieldPasswordSing, field.TypeString, value)
	}
	if value, ok := mu.mutation.Password(); ok {
		_spec.SetField(member.FieldPassword, field.TypeString, value)
	}
	if value, ok := mu.mutation.Mobile(); ok {
		_spec.SetField(member.FieldMobile, field.TypeString, value)
	}
	if value, ok := mu.mutation.Nickname(); ok {
		_spec.SetField(member.FieldNickname, field.TypeString, value)
	}
	if value, ok := mu.mutation.State(); ok {
		_spec.SetField(member.FieldState, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MemberUpdateOne is the builder for updating a single Member entity.
type MemberUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MemberMutation
}

// SetUpdateTime sets the "update_time" field.
func (muo *MemberUpdateOne) SetUpdateTime(t time.Time) *MemberUpdateOne {
	muo.mutation.SetUpdateTime(t)
	return muo
}

// SetEmail sets the "email" field.
func (muo *MemberUpdateOne) SetEmail(s string) *MemberUpdateOne {
	muo.mutation.SetEmail(s)
	return muo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableEmail(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetEmail(*s)
	}
	return muo
}

// SetAvatar sets the "avatar" field.
func (muo *MemberUpdateOne) SetAvatar(s string) *MemberUpdateOne {
	muo.mutation.SetAvatar(s)
	return muo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableAvatar(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetAvatar(*s)
	}
	return muo
}

// SetPasswordSing sets the "password_sing" field.
func (muo *MemberUpdateOne) SetPasswordSing(s string) *MemberUpdateOne {
	muo.mutation.SetPasswordSing(s)
	return muo
}

// SetNillablePasswordSing sets the "password_sing" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillablePasswordSing(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetPasswordSing(*s)
	}
	return muo
}

// SetPassword sets the "password" field.
func (muo *MemberUpdateOne) SetPassword(s string) *MemberUpdateOne {
	muo.mutation.SetPassword(s)
	return muo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillablePassword(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetPassword(*s)
	}
	return muo
}

// SetMobile sets the "mobile" field.
func (muo *MemberUpdateOne) SetMobile(s string) *MemberUpdateOne {
	muo.mutation.SetMobile(s)
	return muo
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableMobile(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetMobile(*s)
	}
	return muo
}

// SetNickname sets the "nickname" field.
func (muo *MemberUpdateOne) SetNickname(s string) *MemberUpdateOne {
	muo.mutation.SetNickname(s)
	return muo
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableNickname(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetNickname(*s)
	}
	return muo
}

// SetState sets the "state" field.
func (muo *MemberUpdateOne) SetState(s string) *MemberUpdateOne {
	muo.mutation.SetState(s)
	return muo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableState(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetState(*s)
	}
	return muo
}

// Mutation returns the MemberMutation object of the builder.
func (muo *MemberUpdateOne) Mutation() *MemberMutation {
	return muo.mutation
}

// Where appends a list predicates to the MemberUpdate builder.
func (muo *MemberUpdateOne) Where(ps ...predicate.Member) *MemberUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MemberUpdateOne) Select(field string, fields ...string) *MemberUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Member entity.
func (muo *MemberUpdateOne) Save(ctx context.Context) (*Member, error) {
	muo.defaults()
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MemberUpdateOne) SaveX(ctx context.Context) *Member {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MemberUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MemberUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MemberUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdateTime(); !ok {
		v := member.UpdateDefaultUpdateTime()
		muo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MemberUpdateOne) check() error {
	if v, ok := muo.mutation.Email(); ok {
		if err := member.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Member.email": %w`, err)}
		}
	}
	if v, ok := muo.mutation.PasswordSing(); ok {
		if err := member.PasswordSingValidator(v); err != nil {
			return &ValidationError{Name: "password_sing", err: fmt.Errorf(`ent: validator failed for field "Member.password_sing": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Password(); ok {
		if err := member.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Member.password": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Mobile(); ok {
		if err := member.MobileValidator(v); err != nil {
			return &ValidationError{Name: "mobile", err: fmt.Errorf(`ent: validator failed for field "Member.mobile": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Nickname(); ok {
		if err := member.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "Member.nickname": %w`, err)}
		}
	}
	if v, ok := muo.mutation.State(); ok {
		if err := member.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Member.state": %w`, err)}
		}
	}
	return nil
}

func (muo *MemberUpdateOne) sqlSave(ctx context.Context) (_node *Member, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(member.Table, member.Columns, sqlgraph.NewFieldSpec(member.FieldID, field.TypeUUID))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Member.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, member.FieldID)
		for _, f := range fields {
			if !member.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != member.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.UpdateTime(); ok {
		_spec.SetField(member.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := muo.mutation.Email(); ok {
		_spec.SetField(member.FieldEmail, field.TypeString, value)
	}
	if value, ok := muo.mutation.Avatar(); ok {
		_spec.SetField(member.FieldAvatar, field.TypeString, value)
	}
	if value, ok := muo.mutation.PasswordSing(); ok {
		_spec.SetField(member.FieldPasswordSing, field.TypeString, value)
	}
	if value, ok := muo.mutation.Password(); ok {
		_spec.SetField(member.FieldPassword, field.TypeString, value)
	}
	if value, ok := muo.mutation.Mobile(); ok {
		_spec.SetField(member.FieldMobile, field.TypeString, value)
	}
	if value, ok := muo.mutation.Nickname(); ok {
		_spec.SetField(member.FieldNickname, field.TypeString, value)
	}
	if value, ok := muo.mutation.State(); ok {
		_spec.SetField(member.FieldState, field.TypeString, value)
	}
	_node = &Member{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
