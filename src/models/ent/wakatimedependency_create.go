// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"service-api/src/models/ent/wakatimedependency"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WakatimeDependencyCreate is the builder for creating a WakatimeDependency entity.
type WakatimeDependencyCreate struct {
	config
	mutation *WakatimeDependencyMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (wdc *WakatimeDependencyCreate) SetCreateTime(t time.Time) *WakatimeDependencyCreate {
	wdc.mutation.SetCreateTime(t)
	return wdc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (wdc *WakatimeDependencyCreate) SetNillableCreateTime(t *time.Time) *WakatimeDependencyCreate {
	if t != nil {
		wdc.SetCreateTime(*t)
	}
	return wdc
}

// SetUpdateTime sets the "update_time" field.
func (wdc *WakatimeDependencyCreate) SetUpdateTime(t time.Time) *WakatimeDependencyCreate {
	wdc.mutation.SetUpdateTime(t)
	return wdc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (wdc *WakatimeDependencyCreate) SetNillableUpdateTime(t *time.Time) *WakatimeDependencyCreate {
	if t != nil {
		wdc.SetUpdateTime(*t)
	}
	return wdc
}

// SetWakatimeID sets the "wakatime_id" field.
func (wdc *WakatimeDependencyCreate) SetWakatimeID(u uuid.UUID) *WakatimeDependencyCreate {
	wdc.mutation.SetWakatimeID(u)
	return wdc
}

// SetMemberID sets the "member_id" field.
func (wdc *WakatimeDependencyCreate) SetMemberID(u uuid.UUID) *WakatimeDependencyCreate {
	wdc.mutation.SetMemberID(u)
	return wdc
}

// SetName sets the "name" field.
func (wdc *WakatimeDependencyCreate) SetName(s string) *WakatimeDependencyCreate {
	wdc.mutation.SetName(s)
	return wdc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (wdc *WakatimeDependencyCreate) SetNillableName(s *string) *WakatimeDependencyCreate {
	if s != nil {
		wdc.SetName(*s)
	}
	return wdc
}

// SetTotalSeconds sets the "total_seconds" field.
func (wdc *WakatimeDependencyCreate) SetTotalSeconds(i int64) *WakatimeDependencyCreate {
	wdc.mutation.SetTotalSeconds(i)
	return wdc
}

// SetNillableTotalSeconds sets the "total_seconds" field if the given value is not nil.
func (wdc *WakatimeDependencyCreate) SetNillableTotalSeconds(i *int64) *WakatimeDependencyCreate {
	if i != nil {
		wdc.SetTotalSeconds(*i)
	}
	return wdc
}

// SetID sets the "id" field.
func (wdc *WakatimeDependencyCreate) SetID(u uuid.UUID) *WakatimeDependencyCreate {
	wdc.mutation.SetID(u)
	return wdc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wdc *WakatimeDependencyCreate) SetNillableID(u *uuid.UUID) *WakatimeDependencyCreate {
	if u != nil {
		wdc.SetID(*u)
	}
	return wdc
}

// Mutation returns the WakatimeDependencyMutation object of the builder.
func (wdc *WakatimeDependencyCreate) Mutation() *WakatimeDependencyMutation {
	return wdc.mutation
}

// Save creates the WakatimeDependency in the database.
func (wdc *WakatimeDependencyCreate) Save(ctx context.Context) (*WakatimeDependency, error) {
	wdc.defaults()
	return withHooks(ctx, wdc.sqlSave, wdc.mutation, wdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wdc *WakatimeDependencyCreate) SaveX(ctx context.Context) *WakatimeDependency {
	v, err := wdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wdc *WakatimeDependencyCreate) Exec(ctx context.Context) error {
	_, err := wdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wdc *WakatimeDependencyCreate) ExecX(ctx context.Context) {
	if err := wdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wdc *WakatimeDependencyCreate) defaults() {
	if _, ok := wdc.mutation.CreateTime(); !ok {
		v := wakatimedependency.DefaultCreateTime()
		wdc.mutation.SetCreateTime(v)
	}
	if _, ok := wdc.mutation.UpdateTime(); !ok {
		v := wakatimedependency.DefaultUpdateTime()
		wdc.mutation.SetUpdateTime(v)
	}
	if _, ok := wdc.mutation.Name(); !ok {
		v := wakatimedependency.DefaultName
		wdc.mutation.SetName(v)
	}
	if _, ok := wdc.mutation.TotalSeconds(); !ok {
		v := wakatimedependency.DefaultTotalSeconds
		wdc.mutation.SetTotalSeconds(v)
	}
	if _, ok := wdc.mutation.ID(); !ok {
		v := wakatimedependency.DefaultID()
		wdc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wdc *WakatimeDependencyCreate) check() error {
	if _, ok := wdc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "WakatimeDependency.create_time"`)}
	}
	if _, ok := wdc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "WakatimeDependency.update_time"`)}
	}
	if _, ok := wdc.mutation.WakatimeID(); !ok {
		return &ValidationError{Name: "wakatime_id", err: errors.New(`ent: missing required field "WakatimeDependency.wakatime_id"`)}
	}
	if _, ok := wdc.mutation.MemberID(); !ok {
		return &ValidationError{Name: "member_id", err: errors.New(`ent: missing required field "WakatimeDependency.member_id"`)}
	}
	if _, ok := wdc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "WakatimeDependency.name"`)}
	}
	if v, ok := wdc.mutation.Name(); ok {
		if err := wakatimedependency.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "WakatimeDependency.name": %w`, err)}
		}
	}
	if _, ok := wdc.mutation.TotalSeconds(); !ok {
		return &ValidationError{Name: "total_seconds", err: errors.New(`ent: missing required field "WakatimeDependency.total_seconds"`)}
	}
	return nil
}

func (wdc *WakatimeDependencyCreate) sqlSave(ctx context.Context) (*WakatimeDependency, error) {
	if err := wdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	wdc.mutation.id = &_node.ID
	wdc.mutation.done = true
	return _node, nil
}

func (wdc *WakatimeDependencyCreate) createSpec() (*WakatimeDependency, *sqlgraph.CreateSpec) {
	var (
		_node = &WakatimeDependency{config: wdc.config}
		_spec = sqlgraph.NewCreateSpec(wakatimedependency.Table, sqlgraph.NewFieldSpec(wakatimedependency.FieldID, field.TypeUUID))
	)
	if id, ok := wdc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wdc.mutation.CreateTime(); ok {
		_spec.SetField(wakatimedependency.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := wdc.mutation.UpdateTime(); ok {
		_spec.SetField(wakatimedependency.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := wdc.mutation.WakatimeID(); ok {
		_spec.SetField(wakatimedependency.FieldWakatimeID, field.TypeUUID, value)
		_node.WakatimeID = value
	}
	if value, ok := wdc.mutation.MemberID(); ok {
		_spec.SetField(wakatimedependency.FieldMemberID, field.TypeUUID, value)
		_node.MemberID = value
	}
	if value, ok := wdc.mutation.Name(); ok {
		_spec.SetField(wakatimedependency.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := wdc.mutation.TotalSeconds(); ok {
		_spec.SetField(wakatimedependency.FieldTotalSeconds, field.TypeInt64, value)
		_node.TotalSeconds = value
	}
	return _node, _spec
}

// WakatimeDependencyCreateBulk is the builder for creating many WakatimeDependency entities in bulk.
type WakatimeDependencyCreateBulk struct {
	config
	err      error
	builders []*WakatimeDependencyCreate
}

// Save creates the WakatimeDependency entities in the database.
func (wdcb *WakatimeDependencyCreateBulk) Save(ctx context.Context) ([]*WakatimeDependency, error) {
	if wdcb.err != nil {
		return nil, wdcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wdcb.builders))
	nodes := make([]*WakatimeDependency, len(wdcb.builders))
	mutators := make([]Mutator, len(wdcb.builders))
	for i := range wdcb.builders {
		func(i int, root context.Context) {
			builder := wdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WakatimeDependencyMutation)
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
					_, err = mutators[i+1].Mutate(root, wdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wdcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, wdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wdcb *WakatimeDependencyCreateBulk) SaveX(ctx context.Context) []*WakatimeDependency {
	v, err := wdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wdcb *WakatimeDependencyCreateBulk) Exec(ctx context.Context) error {
	_, err := wdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wdcb *WakatimeDependencyCreateBulk) ExecX(ctx context.Context) {
	if err := wdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
