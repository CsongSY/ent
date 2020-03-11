// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv1

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv1/car"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv1/user"
	"github.com/facebookincubator/ent/schema/field"
)

// CarCreate is the builder for creating a Car entity.
type CarCreate struct {
	config
	mutation *CarMutation
	hooks    []Hook
}

// SetOwnerID sets the owner edge to User by id.
func (cc *CarCreate) SetOwnerID(id int) *CarCreate {
	cc.mutation.SetOwnerID(id)
	return cc
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (cc *CarCreate) SetNillableOwnerID(id *int) *CarCreate {
	if id != nil {
		cc = cc.SetOwnerID(*id)
	}
	return cc
}

// SetOwner sets the owner edge to User.
func (cc *CarCreate) SetOwner(u *User) *CarCreate {
	return cc.SetOwnerID(u.ID)
}

// Save creates the Car in the database.
func (cc *CarCreate) Save(ctx context.Context) (*Car, error) {
	var (
		err  error
		node *Car
	)
	if len(cc.hooks) == 0 {
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CarMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			return node, err
		})
		for _, hook := range cc.hooks {
			mut = hook(mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CarCreate) SaveX(ctx context.Context) *Car {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CarCreate) sqlSave(ctx context.Context) (*Car, error) {
	var (
		c     = &Car{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: car.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: car.FieldID,
			},
		}
	)
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	c.ID = int(id)
	return c, nil
}
