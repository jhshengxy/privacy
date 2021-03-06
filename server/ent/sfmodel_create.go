// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/kallydev/privacy/ent/sfmodel"
)

// SFModelCreate is the builder for creating a SFModel entity.
type SFModelCreate struct {
	config
	mutation *SFModelMutation
	hooks    []Hook
}

// SetName sets the name field.
func (smc *SFModelCreate) SetName(s string) *SFModelCreate {
	smc.mutation.SetName(s)
	return smc
}

// SetPhoneNumber sets the phone_number field.
func (smc *SFModelCreate) SetPhoneNumber(i int64) *SFModelCreate {
	smc.mutation.SetPhoneNumber(i)
	return smc
}

// SetAddress sets the address field.
func (smc *SFModelCreate) SetAddress(s string) *SFModelCreate {
	smc.mutation.SetAddress(s)
	return smc
}

// Mutation returns the SFModelMutation object of the builder.
func (smc *SFModelCreate) Mutation() *SFModelMutation {
	return smc.mutation
}

// Save creates the SFModel in the database.
func (smc *SFModelCreate) Save(ctx context.Context) (*SFModel, error) {
	var (
		err  error
		node *SFModel
	)
	if len(smc.hooks) == 0 {
		if err = smc.check(); err != nil {
			return nil, err
		}
		node, err = smc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SFModelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = smc.check(); err != nil {
				return nil, err
			}
			smc.mutation = mutation
			node, err = smc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(smc.hooks) - 1; i >= 0; i-- {
			mut = smc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (smc *SFModelCreate) SaveX(ctx context.Context) *SFModel {
	v, err := smc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (smc *SFModelCreate) check() error {
	if _, ok := smc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := smc.mutation.PhoneNumber(); !ok {
		return &ValidationError{Name: "phone_number", err: errors.New("ent: missing required field \"phone_number\"")}
	}
	if _, ok := smc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New("ent: missing required field \"address\"")}
	}
	return nil
}

func (smc *SFModelCreate) sqlSave(ctx context.Context) (*SFModel, error) {
	_node, _spec := smc.createSpec()
	if err := sqlgraph.CreateNode(ctx, smc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (smc *SFModelCreate) createSpec() (*SFModel, *sqlgraph.CreateSpec) {
	var (
		_node = &SFModel{config: smc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sfmodel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sfmodel.FieldID,
			},
		}
	)
	if value, ok := smc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sfmodel.FieldName,
		})
		_node.Name = value
	}
	if value, ok := smc.mutation.PhoneNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sfmodel.FieldPhoneNumber,
		})
		_node.PhoneNumber = value
	}
	if value, ok := smc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sfmodel.FieldAddress,
		})
		_node.Address = value
	}
	return _node, _spec
}

// SFModelCreateBulk is the builder for creating a bulk of SFModel entities.
type SFModelCreateBulk struct {
	config
	builders []*SFModelCreate
}

// Save creates the SFModel entities in the database.
func (smcb *SFModelCreateBulk) Save(ctx context.Context) ([]*SFModel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(smcb.builders))
	nodes := make([]*SFModel, len(smcb.builders))
	mutators := make([]Mutator, len(smcb.builders))
	for i := range smcb.builders {
		func(i int, root context.Context) {
			builder := smcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SFModelMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, smcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, smcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, smcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (smcb *SFModelCreateBulk) SaveX(ctx context.Context) []*SFModel {
	v, err := smcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
