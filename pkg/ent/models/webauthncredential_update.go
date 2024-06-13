// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/webauthncredential"
	"errors"
	"fmt"
	"frame/util/timeutil"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WebAuthnCredentialUpdate is the builder for updating WebAuthnCredential entities.
type WebAuthnCredentialUpdate struct {
	config
	hooks    []Hook
	mutation *WebAuthnCredentialMutation
}

// Where appends a list predicates to the WebAuthnCredentialUpdate builder.
func (wacu *WebAuthnCredentialUpdate) Where(ps ...predicate.WebAuthnCredential) *WebAuthnCredentialUpdate {
	wacu.mutation.Where(ps...)
	return wacu
}

// SetName sets the "name" field.
func (wacu *WebAuthnCredentialUpdate) SetName(s string) *WebAuthnCredentialUpdate {
	wacu.mutation.SetName(s)
	return wacu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (wacu *WebAuthnCredentialUpdate) SetNillableName(s *string) *WebAuthnCredentialUpdate {
	if s != nil {
		wacu.SetName(*s)
	}
	return wacu
}

// SetLowerName sets the "lower_name" field.
func (wacu *WebAuthnCredentialUpdate) SetLowerName(s string) *WebAuthnCredentialUpdate {
	wacu.mutation.SetLowerName(s)
	return wacu
}

// SetNillableLowerName sets the "lower_name" field if the given value is not nil.
func (wacu *WebAuthnCredentialUpdate) SetNillableLowerName(s *string) *WebAuthnCredentialUpdate {
	if s != nil {
		wacu.SetLowerName(*s)
	}
	return wacu
}

// SetUserID sets the "user_id" field.
func (wacu *WebAuthnCredentialUpdate) SetUserID(i int64) *WebAuthnCredentialUpdate {
	wacu.mutation.ResetUserID()
	wacu.mutation.SetUserID(i)
	return wacu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wacu *WebAuthnCredentialUpdate) SetNillableUserID(i *int64) *WebAuthnCredentialUpdate {
	if i != nil {
		wacu.SetUserID(*i)
	}
	return wacu
}

// AddUserID adds i to the "user_id" field.
func (wacu *WebAuthnCredentialUpdate) AddUserID(i int64) *WebAuthnCredentialUpdate {
	wacu.mutation.AddUserID(i)
	return wacu
}

// SetCredentialID sets the "credential_id" field.
func (wacu *WebAuthnCredentialUpdate) SetCredentialID(b []byte) *WebAuthnCredentialUpdate {
	wacu.mutation.SetCredentialID(b)
	return wacu
}

// SetPublicKey sets the "public_key" field.
func (wacu *WebAuthnCredentialUpdate) SetPublicKey(b []byte) *WebAuthnCredentialUpdate {
	wacu.mutation.SetPublicKey(b)
	return wacu
}

// SetAttestationType sets the "attestation_type" field.
func (wacu *WebAuthnCredentialUpdate) SetAttestationType(s string) *WebAuthnCredentialUpdate {
	wacu.mutation.SetAttestationType(s)
	return wacu
}

// SetNillableAttestationType sets the "attestation_type" field if the given value is not nil.
func (wacu *WebAuthnCredentialUpdate) SetNillableAttestationType(s *string) *WebAuthnCredentialUpdate {
	if s != nil {
		wacu.SetAttestationType(*s)
	}
	return wacu
}

// SetAAGUID sets the "AAGUID" field.
func (wacu *WebAuthnCredentialUpdate) SetAAGUID(b []byte) *WebAuthnCredentialUpdate {
	wacu.mutation.SetAAGUID(b)
	return wacu
}

// SetSignCount sets the "sign_count" field.
func (wacu *WebAuthnCredentialUpdate) SetSignCount(u uint32) *WebAuthnCredentialUpdate {
	wacu.mutation.ResetSignCount()
	wacu.mutation.SetSignCount(u)
	return wacu
}

// SetNillableSignCount sets the "sign_count" field if the given value is not nil.
func (wacu *WebAuthnCredentialUpdate) SetNillableSignCount(u *uint32) *WebAuthnCredentialUpdate {
	if u != nil {
		wacu.SetSignCount(*u)
	}
	return wacu
}

// AddSignCount adds u to the "sign_count" field.
func (wacu *WebAuthnCredentialUpdate) AddSignCount(u int32) *WebAuthnCredentialUpdate {
	wacu.mutation.AddSignCount(u)
	return wacu
}

// SetCloneWarning sets the "clone_warning" field.
func (wacu *WebAuthnCredentialUpdate) SetCloneWarning(b bool) *WebAuthnCredentialUpdate {
	wacu.mutation.SetCloneWarning(b)
	return wacu
}

// SetNillableCloneWarning sets the "clone_warning" field if the given value is not nil.
func (wacu *WebAuthnCredentialUpdate) SetNillableCloneWarning(b *bool) *WebAuthnCredentialUpdate {
	if b != nil {
		wacu.SetCloneWarning(*b)
	}
	return wacu
}

// SetUpdateTime sets the "update_time" field.
func (wacu *WebAuthnCredentialUpdate) SetUpdateTime(ts timeutil.TimeStamp) *WebAuthnCredentialUpdate {
	wacu.mutation.ResetUpdateTime()
	wacu.mutation.SetUpdateTime(ts)
	return wacu
}

// AddUpdateTime adds ts to the "update_time" field.
func (wacu *WebAuthnCredentialUpdate) AddUpdateTime(ts timeutil.TimeStamp) *WebAuthnCredentialUpdate {
	wacu.mutation.AddUpdateTime(ts)
	return wacu
}

// Mutation returns the WebAuthnCredentialMutation object of the builder.
func (wacu *WebAuthnCredentialUpdate) Mutation() *WebAuthnCredentialMutation {
	return wacu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wacu *WebAuthnCredentialUpdate) Save(ctx context.Context) (int, error) {
	wacu.defaults()
	return withHooks(ctx, wacu.sqlSave, wacu.mutation, wacu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wacu *WebAuthnCredentialUpdate) SaveX(ctx context.Context) int {
	affected, err := wacu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wacu *WebAuthnCredentialUpdate) Exec(ctx context.Context) error {
	_, err := wacu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wacu *WebAuthnCredentialUpdate) ExecX(ctx context.Context) {
	if err := wacu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wacu *WebAuthnCredentialUpdate) defaults() {
	if _, ok := wacu.mutation.UpdateTime(); !ok {
		v := webauthncredential.UpdateDefaultUpdateTime()
		wacu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wacu *WebAuthnCredentialUpdate) check() error {
	if v, ok := wacu.mutation.CredentialID(); ok {
		if err := webauthncredential.CredentialIDValidator(v); err != nil {
			return &ValidationError{Name: "credential_id", err: fmt.Errorf(`models: validator failed for field "WebAuthnCredential.credential_id": %w`, err)}
		}
	}
	return nil
}

func (wacu *WebAuthnCredentialUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wacu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(webauthncredential.Table, webauthncredential.Columns, sqlgraph.NewFieldSpec(webauthncredential.FieldID, field.TypeInt64))
	if ps := wacu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wacu.mutation.Name(); ok {
		_spec.SetField(webauthncredential.FieldName, field.TypeString, value)
	}
	if value, ok := wacu.mutation.LowerName(); ok {
		_spec.SetField(webauthncredential.FieldLowerName, field.TypeString, value)
	}
	if value, ok := wacu.mutation.UserID(); ok {
		_spec.SetField(webauthncredential.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := wacu.mutation.AddedUserID(); ok {
		_spec.AddField(webauthncredential.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := wacu.mutation.CredentialID(); ok {
		_spec.SetField(webauthncredential.FieldCredentialID, field.TypeBytes, value)
	}
	if value, ok := wacu.mutation.PublicKey(); ok {
		_spec.SetField(webauthncredential.FieldPublicKey, field.TypeBytes, value)
	}
	if value, ok := wacu.mutation.AttestationType(); ok {
		_spec.SetField(webauthncredential.FieldAttestationType, field.TypeString, value)
	}
	if value, ok := wacu.mutation.AAGUID(); ok {
		_spec.SetField(webauthncredential.FieldAAGUID, field.TypeBytes, value)
	}
	if value, ok := wacu.mutation.SignCount(); ok {
		_spec.SetField(webauthncredential.FieldSignCount, field.TypeUint32, value)
	}
	if value, ok := wacu.mutation.AddedSignCount(); ok {
		_spec.AddField(webauthncredential.FieldSignCount, field.TypeUint32, value)
	}
	if value, ok := wacu.mutation.CloneWarning(); ok {
		_spec.SetField(webauthncredential.FieldCloneWarning, field.TypeBool, value)
	}
	if value, ok := wacu.mutation.UpdateTime(); ok {
		_spec.SetField(webauthncredential.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := wacu.mutation.AddedUpdateTime(); ok {
		_spec.AddField(webauthncredential.FieldUpdateTime, field.TypeInt64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wacu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{webauthncredential.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wacu.mutation.done = true
	return n, nil
}

// WebAuthnCredentialUpdateOne is the builder for updating a single WebAuthnCredential entity.
type WebAuthnCredentialUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WebAuthnCredentialMutation
}

// SetName sets the "name" field.
func (wacuo *WebAuthnCredentialUpdateOne) SetName(s string) *WebAuthnCredentialUpdateOne {
	wacuo.mutation.SetName(s)
	return wacuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (wacuo *WebAuthnCredentialUpdateOne) SetNillableName(s *string) *WebAuthnCredentialUpdateOne {
	if s != nil {
		wacuo.SetName(*s)
	}
	return wacuo
}

// SetLowerName sets the "lower_name" field.
func (wacuo *WebAuthnCredentialUpdateOne) SetLowerName(s string) *WebAuthnCredentialUpdateOne {
	wacuo.mutation.SetLowerName(s)
	return wacuo
}

// SetNillableLowerName sets the "lower_name" field if the given value is not nil.
func (wacuo *WebAuthnCredentialUpdateOne) SetNillableLowerName(s *string) *WebAuthnCredentialUpdateOne {
	if s != nil {
		wacuo.SetLowerName(*s)
	}
	return wacuo
}

// SetUserID sets the "user_id" field.
func (wacuo *WebAuthnCredentialUpdateOne) SetUserID(i int64) *WebAuthnCredentialUpdateOne {
	wacuo.mutation.ResetUserID()
	wacuo.mutation.SetUserID(i)
	return wacuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wacuo *WebAuthnCredentialUpdateOne) SetNillableUserID(i *int64) *WebAuthnCredentialUpdateOne {
	if i != nil {
		wacuo.SetUserID(*i)
	}
	return wacuo
}

// AddUserID adds i to the "user_id" field.
func (wacuo *WebAuthnCredentialUpdateOne) AddUserID(i int64) *WebAuthnCredentialUpdateOne {
	wacuo.mutation.AddUserID(i)
	return wacuo
}

// SetCredentialID sets the "credential_id" field.
func (wacuo *WebAuthnCredentialUpdateOne) SetCredentialID(b []byte) *WebAuthnCredentialUpdateOne {
	wacuo.mutation.SetCredentialID(b)
	return wacuo
}

// SetPublicKey sets the "public_key" field.
func (wacuo *WebAuthnCredentialUpdateOne) SetPublicKey(b []byte) *WebAuthnCredentialUpdateOne {
	wacuo.mutation.SetPublicKey(b)
	return wacuo
}

// SetAttestationType sets the "attestation_type" field.
func (wacuo *WebAuthnCredentialUpdateOne) SetAttestationType(s string) *WebAuthnCredentialUpdateOne {
	wacuo.mutation.SetAttestationType(s)
	return wacuo
}

// SetNillableAttestationType sets the "attestation_type" field if the given value is not nil.
func (wacuo *WebAuthnCredentialUpdateOne) SetNillableAttestationType(s *string) *WebAuthnCredentialUpdateOne {
	if s != nil {
		wacuo.SetAttestationType(*s)
	}
	return wacuo
}

// SetAAGUID sets the "AAGUID" field.
func (wacuo *WebAuthnCredentialUpdateOne) SetAAGUID(b []byte) *WebAuthnCredentialUpdateOne {
	wacuo.mutation.SetAAGUID(b)
	return wacuo
}

// SetSignCount sets the "sign_count" field.
func (wacuo *WebAuthnCredentialUpdateOne) SetSignCount(u uint32) *WebAuthnCredentialUpdateOne {
	wacuo.mutation.ResetSignCount()
	wacuo.mutation.SetSignCount(u)
	return wacuo
}

// SetNillableSignCount sets the "sign_count" field if the given value is not nil.
func (wacuo *WebAuthnCredentialUpdateOne) SetNillableSignCount(u *uint32) *WebAuthnCredentialUpdateOne {
	if u != nil {
		wacuo.SetSignCount(*u)
	}
	return wacuo
}

// AddSignCount adds u to the "sign_count" field.
func (wacuo *WebAuthnCredentialUpdateOne) AddSignCount(u int32) *WebAuthnCredentialUpdateOne {
	wacuo.mutation.AddSignCount(u)
	return wacuo
}

// SetCloneWarning sets the "clone_warning" field.
func (wacuo *WebAuthnCredentialUpdateOne) SetCloneWarning(b bool) *WebAuthnCredentialUpdateOne {
	wacuo.mutation.SetCloneWarning(b)
	return wacuo
}

// SetNillableCloneWarning sets the "clone_warning" field if the given value is not nil.
func (wacuo *WebAuthnCredentialUpdateOne) SetNillableCloneWarning(b *bool) *WebAuthnCredentialUpdateOne {
	if b != nil {
		wacuo.SetCloneWarning(*b)
	}
	return wacuo
}

// SetUpdateTime sets the "update_time" field.
func (wacuo *WebAuthnCredentialUpdateOne) SetUpdateTime(ts timeutil.TimeStamp) *WebAuthnCredentialUpdateOne {
	wacuo.mutation.ResetUpdateTime()
	wacuo.mutation.SetUpdateTime(ts)
	return wacuo
}

// AddUpdateTime adds ts to the "update_time" field.
func (wacuo *WebAuthnCredentialUpdateOne) AddUpdateTime(ts timeutil.TimeStamp) *WebAuthnCredentialUpdateOne {
	wacuo.mutation.AddUpdateTime(ts)
	return wacuo
}

// Mutation returns the WebAuthnCredentialMutation object of the builder.
func (wacuo *WebAuthnCredentialUpdateOne) Mutation() *WebAuthnCredentialMutation {
	return wacuo.mutation
}

// Where appends a list predicates to the WebAuthnCredentialUpdate builder.
func (wacuo *WebAuthnCredentialUpdateOne) Where(ps ...predicate.WebAuthnCredential) *WebAuthnCredentialUpdateOne {
	wacuo.mutation.Where(ps...)
	return wacuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wacuo *WebAuthnCredentialUpdateOne) Select(field string, fields ...string) *WebAuthnCredentialUpdateOne {
	wacuo.fields = append([]string{field}, fields...)
	return wacuo
}

// Save executes the query and returns the updated WebAuthnCredential entity.
func (wacuo *WebAuthnCredentialUpdateOne) Save(ctx context.Context) (*WebAuthnCredential, error) {
	wacuo.defaults()
	return withHooks(ctx, wacuo.sqlSave, wacuo.mutation, wacuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wacuo *WebAuthnCredentialUpdateOne) SaveX(ctx context.Context) *WebAuthnCredential {
	node, err := wacuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wacuo *WebAuthnCredentialUpdateOne) Exec(ctx context.Context) error {
	_, err := wacuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wacuo *WebAuthnCredentialUpdateOne) ExecX(ctx context.Context) {
	if err := wacuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wacuo *WebAuthnCredentialUpdateOne) defaults() {
	if _, ok := wacuo.mutation.UpdateTime(); !ok {
		v := webauthncredential.UpdateDefaultUpdateTime()
		wacuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wacuo *WebAuthnCredentialUpdateOne) check() error {
	if v, ok := wacuo.mutation.CredentialID(); ok {
		if err := webauthncredential.CredentialIDValidator(v); err != nil {
			return &ValidationError{Name: "credential_id", err: fmt.Errorf(`models: validator failed for field "WebAuthnCredential.credential_id": %w`, err)}
		}
	}
	return nil
}

func (wacuo *WebAuthnCredentialUpdateOne) sqlSave(ctx context.Context) (_node *WebAuthnCredential, err error) {
	if err := wacuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(webauthncredential.Table, webauthncredential.Columns, sqlgraph.NewFieldSpec(webauthncredential.FieldID, field.TypeInt64))
	id, ok := wacuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "WebAuthnCredential.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wacuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, webauthncredential.FieldID)
		for _, f := range fields {
			if !webauthncredential.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != webauthncredential.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wacuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wacuo.mutation.Name(); ok {
		_spec.SetField(webauthncredential.FieldName, field.TypeString, value)
	}
	if value, ok := wacuo.mutation.LowerName(); ok {
		_spec.SetField(webauthncredential.FieldLowerName, field.TypeString, value)
	}
	if value, ok := wacuo.mutation.UserID(); ok {
		_spec.SetField(webauthncredential.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := wacuo.mutation.AddedUserID(); ok {
		_spec.AddField(webauthncredential.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := wacuo.mutation.CredentialID(); ok {
		_spec.SetField(webauthncredential.FieldCredentialID, field.TypeBytes, value)
	}
	if value, ok := wacuo.mutation.PublicKey(); ok {
		_spec.SetField(webauthncredential.FieldPublicKey, field.TypeBytes, value)
	}
	if value, ok := wacuo.mutation.AttestationType(); ok {
		_spec.SetField(webauthncredential.FieldAttestationType, field.TypeString, value)
	}
	if value, ok := wacuo.mutation.AAGUID(); ok {
		_spec.SetField(webauthncredential.FieldAAGUID, field.TypeBytes, value)
	}
	if value, ok := wacuo.mutation.SignCount(); ok {
		_spec.SetField(webauthncredential.FieldSignCount, field.TypeUint32, value)
	}
	if value, ok := wacuo.mutation.AddedSignCount(); ok {
		_spec.AddField(webauthncredential.FieldSignCount, field.TypeUint32, value)
	}
	if value, ok := wacuo.mutation.CloneWarning(); ok {
		_spec.SetField(webauthncredential.FieldCloneWarning, field.TypeBool, value)
	}
	if value, ok := wacuo.mutation.UpdateTime(); ok {
		_spec.SetField(webauthncredential.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := wacuo.mutation.AddedUpdateTime(); ok {
		_spec.AddField(webauthncredential.FieldUpdateTime, field.TypeInt64, value)
	}
	_node = &WebAuthnCredential{config: wacuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wacuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{webauthncredential.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wacuo.mutation.done = true
	return _node, nil
}
