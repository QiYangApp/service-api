// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/user"
	"ent/types/auth"
	"errors"
	"fmt"
	"frame/util/timeutil"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetAvatar sets the "avatar" field.
func (uc *UserCreate) SetAvatar(s string) *UserCreate {
	uc.mutation.SetAvatar(s)
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetLowerName sets the "lower_name" field.
func (uc *UserCreate) SetLowerName(s string) *UserCreate {
	uc.mutation.SetLowerName(s)
	return uc
}

// SetFullName sets the "full_name" field.
func (uc *UserCreate) SetFullName(s string) *UserCreate {
	uc.mutation.SetFullName(s)
	return uc
}

// SetPasswdSalt sets the "passwd_salt" field.
func (uc *UserCreate) SetPasswdSalt(s string) *UserCreate {
	uc.mutation.SetPasswdSalt(s)
	return uc
}

// SetNillablePasswdSalt sets the "passwd_salt" field if the given value is not nil.
func (uc *UserCreate) SetNillablePasswdSalt(s *string) *UserCreate {
	if s != nil {
		uc.SetPasswdSalt(*s)
	}
	return uc
}

// SetPasswdHashAlgo sets the "passwd_hash_algo" field.
func (uc *UserCreate) SetPasswdHashAlgo(s string) *UserCreate {
	uc.mutation.SetPasswdHashAlgo(s)
	return uc
}

// SetPasswd sets the "passwd" field.
func (uc *UserCreate) SetPasswd(s string) *UserCreate {
	uc.mutation.SetPasswd(s)
	return uc
}

// SetLanguage sets the "language" field.
func (uc *UserCreate) SetLanguage(s string) *UserCreate {
	uc.mutation.SetLanguage(s)
	return uc
}

// SetTheme sets the "theme" field.
func (uc *UserCreate) SetTheme(s string) *UserCreate {
	uc.mutation.SetTheme(s)
	return uc
}

// SetLoginName sets the "login_name" field.
func (uc *UserCreate) SetLoginName(s string) *UserCreate {
	uc.mutation.SetLoginName(s)
	return uc
}

// SetLoginSource sets the "login_source" field.
func (uc *UserCreate) SetLoginSource(i int64) *UserCreate {
	uc.mutation.SetLoginSource(i)
	return uc
}

// SetNillableLoginSource sets the "login_source" field if the given value is not nil.
func (uc *UserCreate) SetNillableLoginSource(i *int64) *UserCreate {
	if i != nil {
		uc.SetLoginSource(*i)
	}
	return uc
}

// SetLoginType sets the "login_type" field.
func (uc *UserCreate) SetLoginType(a auth.Type) *UserCreate {
	uc.mutation.SetLoginType(a)
	return uc
}

// SetIsRestricted sets the "is_restricted" field.
func (uc *UserCreate) SetIsRestricted(b bool) *UserCreate {
	uc.mutation.SetIsRestricted(b)
	return uc
}

// SetNillableIsRestricted sets the "is_restricted" field if the given value is not nil.
func (uc *UserCreate) SetNillableIsRestricted(b *bool) *UserCreate {
	if b != nil {
		uc.SetIsRestricted(*b)
	}
	return uc
}

// SetIsActive sets the "is_active" field.
func (uc *UserCreate) SetIsActive(b bool) *UserCreate {
	uc.mutation.SetIsActive(b)
	return uc
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (uc *UserCreate) SetNillableIsActive(b *bool) *UserCreate {
	if b != nil {
		uc.SetIsActive(*b)
	}
	return uc
}

// SetProhibitLogin sets the "prohibit_login" field.
func (uc *UserCreate) SetProhibitLogin(b bool) *UserCreate {
	uc.mutation.SetProhibitLogin(b)
	return uc
}

// SetNillableProhibitLogin sets the "prohibit_login" field if the given value is not nil.
func (uc *UserCreate) SetNillableProhibitLogin(b *bool) *UserCreate {
	if b != nil {
		uc.SetProhibitLogin(*b)
	}
	return uc
}

// SetCreateTime sets the "create_time" field.
func (uc *UserCreate) SetCreateTime(ts timeutil.TimeStamp) *UserCreate {
	uc.mutation.SetCreateTime(ts)
	return uc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreateTime(ts *timeutil.TimeStamp) *UserCreate {
	if ts != nil {
		uc.SetCreateTime(*ts)
	}
	return uc
}

// SetUpdateTime sets the "update_time" field.
func (uc *UserCreate) SetUpdateTime(ts timeutil.TimeStamp) *UserCreate {
	uc.mutation.SetUpdateTime(ts)
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(i int64) *UserCreate {
	uc.mutation.SetID(i)
	return uc
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.PasswdSalt(); !ok {
		v := user.DefaultPasswdSalt
		uc.mutation.SetPasswdSalt(v)
	}
	if _, ok := uc.mutation.LoginSource(); !ok {
		v := user.DefaultLoginSource
		uc.mutation.SetLoginSource(v)
	}
	if _, ok := uc.mutation.IsRestricted(); !ok {
		v := user.DefaultIsRestricted
		uc.mutation.SetIsRestricted(v)
	}
	if _, ok := uc.mutation.IsActive(); !ok {
		v := user.DefaultIsActive
		uc.mutation.SetIsActive(v)
	}
	if _, ok := uc.mutation.ProhibitLogin(); !ok {
		v := user.DefaultProhibitLogin
		uc.mutation.SetProhibitLogin(v)
	}
	if _, ok := uc.mutation.CreateTime(); !ok {
		v := user.DefaultCreateTime
		uc.mutation.SetCreateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Avatar(); !ok {
		return &ValidationError{Name: "avatar", err: errors.New(`models: missing required field "User.avatar"`)}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`models: missing required field "User.email"`)}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`models: validator failed for field "User.email": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`models: missing required field "User.name"`)}
	}
	if v, ok := uc.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`models: validator failed for field "User.name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.LowerName(); !ok {
		return &ValidationError{Name: "lower_name", err: errors.New(`models: missing required field "User.lower_name"`)}
	}
	if v, ok := uc.mutation.LowerName(); ok {
		if err := user.LowerNameValidator(v); err != nil {
			return &ValidationError{Name: "lower_name", err: fmt.Errorf(`models: validator failed for field "User.lower_name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.FullName(); !ok {
		return &ValidationError{Name: "full_name", err: errors.New(`models: missing required field "User.full_name"`)}
	}
	if v, ok := uc.mutation.FullName(); ok {
		if err := user.FullNameValidator(v); err != nil {
			return &ValidationError{Name: "full_name", err: fmt.Errorf(`models: validator failed for field "User.full_name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.PasswdSalt(); !ok {
		return &ValidationError{Name: "passwd_salt", err: errors.New(`models: missing required field "User.passwd_salt"`)}
	}
	if v, ok := uc.mutation.PasswdSalt(); ok {
		if err := user.PasswdSaltValidator(v); err != nil {
			return &ValidationError{Name: "passwd_salt", err: fmt.Errorf(`models: validator failed for field "User.passwd_salt": %w`, err)}
		}
	}
	if _, ok := uc.mutation.PasswdHashAlgo(); !ok {
		return &ValidationError{Name: "passwd_hash_algo", err: errors.New(`models: missing required field "User.passwd_hash_algo"`)}
	}
	if v, ok := uc.mutation.PasswdHashAlgo(); ok {
		if err := user.PasswdHashAlgoValidator(v); err != nil {
			return &ValidationError{Name: "passwd_hash_algo", err: fmt.Errorf(`models: validator failed for field "User.passwd_hash_algo": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Passwd(); !ok {
		return &ValidationError{Name: "passwd", err: errors.New(`models: missing required field "User.passwd"`)}
	}
	if v, ok := uc.mutation.Passwd(); ok {
		if err := user.PasswdValidator(v); err != nil {
			return &ValidationError{Name: "passwd", err: fmt.Errorf(`models: validator failed for field "User.passwd": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Language(); !ok {
		return &ValidationError{Name: "language", err: errors.New(`models: missing required field "User.language"`)}
	}
	if v, ok := uc.mutation.Language(); ok {
		if err := user.LanguageValidator(v); err != nil {
			return &ValidationError{Name: "language", err: fmt.Errorf(`models: validator failed for field "User.language": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Theme(); !ok {
		return &ValidationError{Name: "theme", err: errors.New(`models: missing required field "User.theme"`)}
	}
	if v, ok := uc.mutation.Theme(); ok {
		if err := user.ThemeValidator(v); err != nil {
			return &ValidationError{Name: "theme", err: fmt.Errorf(`models: validator failed for field "User.theme": %w`, err)}
		}
	}
	if _, ok := uc.mutation.LoginName(); !ok {
		return &ValidationError{Name: "login_name", err: errors.New(`models: missing required field "User.login_name"`)}
	}
	if v, ok := uc.mutation.LoginName(); ok {
		if err := user.LoginNameValidator(v); err != nil {
			return &ValidationError{Name: "login_name", err: fmt.Errorf(`models: validator failed for field "User.login_name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.LoginSource(); !ok {
		return &ValidationError{Name: "login_source", err: errors.New(`models: missing required field "User.login_source"`)}
	}
	if _, ok := uc.mutation.LoginType(); !ok {
		return &ValidationError{Name: "login_type", err: errors.New(`models: missing required field "User.login_type"`)}
	}
	if _, ok := uc.mutation.IsRestricted(); !ok {
		return &ValidationError{Name: "is_restricted", err: errors.New(`models: missing required field "User.is_restricted"`)}
	}
	if _, ok := uc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`models: missing required field "User.is_active"`)}
	}
	if _, ok := uc.mutation.ProhibitLogin(); !ok {
		return &ValidationError{Name: "prohibit_login", err: errors.New(`models: missing required field "User.prohibit_login"`)}
	}
	if _, ok := uc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`models: missing required field "User.create_time"`)}
	}
	if _, ok := uc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`models: missing required field "User.update_time"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
		_node.Avatar = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.LowerName(); ok {
		_spec.SetField(user.FieldLowerName, field.TypeString, value)
		_node.LowerName = value
	}
	if value, ok := uc.mutation.FullName(); ok {
		_spec.SetField(user.FieldFullName, field.TypeString, value)
		_node.FullName = value
	}
	if value, ok := uc.mutation.PasswdSalt(); ok {
		_spec.SetField(user.FieldPasswdSalt, field.TypeString, value)
		_node.PasswdSalt = value
	}
	if value, ok := uc.mutation.PasswdHashAlgo(); ok {
		_spec.SetField(user.FieldPasswdHashAlgo, field.TypeString, value)
		_node.PasswdHashAlgo = value
	}
	if value, ok := uc.mutation.Passwd(); ok {
		_spec.SetField(user.FieldPasswd, field.TypeString, value)
		_node.Passwd = value
	}
	if value, ok := uc.mutation.Language(); ok {
		_spec.SetField(user.FieldLanguage, field.TypeString, value)
		_node.Language = value
	}
	if value, ok := uc.mutation.Theme(); ok {
		_spec.SetField(user.FieldTheme, field.TypeString, value)
		_node.Theme = value
	}
	if value, ok := uc.mutation.LoginName(); ok {
		_spec.SetField(user.FieldLoginName, field.TypeString, value)
		_node.LoginName = value
	}
	if value, ok := uc.mutation.LoginSource(); ok {
		_spec.SetField(user.FieldLoginSource, field.TypeInt64, value)
		_node.LoginSource = value
	}
	if value, ok := uc.mutation.LoginType(); ok {
		_spec.SetField(user.FieldLoginType, field.TypeInt, value)
		_node.LoginType = value
	}
	if value, ok := uc.mutation.IsRestricted(); ok {
		_spec.SetField(user.FieldIsRestricted, field.TypeBool, value)
		_node.IsRestricted = value
	}
	if value, ok := uc.mutation.IsActive(); ok {
		_spec.SetField(user.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := uc.mutation.ProhibitLogin(); ok {
		_spec.SetField(user.FieldProhibitLogin, field.TypeBool, value)
		_node.ProhibitLogin = value
	}
	if value, ok := uc.mutation.CreateTime(); ok {
		_spec.SetField(user.FieldCreateTime, field.TypeInt64, value)
		_node.CreateTime = value
	}
	if value, ok := uc.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeInt64, value)
		_node.UpdateTime = value
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
