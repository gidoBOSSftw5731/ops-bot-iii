// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.ritsec.cloud/1nv8rZim/ops-bot-iii/ent/signin"
	"gitlab.ritsec.cloud/1nv8rZim/ops-bot-iii/ent/user"
)

// SigninCreate is the builder for creating a Signin entity.
type SigninCreate struct {
	config
	mutation *SigninMutation
	hooks    []Hook
}

// SetTimestamp sets the "timestamp" field.
func (sc *SigninCreate) SetTimestamp(t time.Time) *SigninCreate {
	sc.mutation.SetTimestamp(t)
	return sc
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (sc *SigninCreate) SetNillableTimestamp(t *time.Time) *SigninCreate {
	if t != nil {
		sc.SetTimestamp(*t)
	}
	return sc
}

// SetType sets the "type" field.
func (sc *SigninCreate) SetType(s signin.Type) *SigninCreate {
	sc.mutation.SetType(s)
	return sc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (sc *SigninCreate) SetUserID(id string) *SigninCreate {
	sc.mutation.SetUserID(id)
	return sc
}

// SetUser sets the "user" edge to the User entity.
func (sc *SigninCreate) SetUser(u *User) *SigninCreate {
	return sc.SetUserID(u.ID)
}

// Mutation returns the SigninMutation object of the builder.
func (sc *SigninCreate) Mutation() *SigninMutation {
	return sc.mutation
}

// Save creates the Signin in the database.
func (sc *SigninCreate) Save(ctx context.Context) (*Signin, error) {
	sc.defaults()
	return withHooks[*Signin, SigninMutation](ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SigninCreate) SaveX(ctx context.Context) *Signin {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SigninCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SigninCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SigninCreate) defaults() {
	if _, ok := sc.mutation.Timestamp(); !ok {
		v := signin.DefaultTimestamp
		sc.mutation.SetTimestamp(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SigninCreate) check() error {
	if _, ok := sc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "timestamp", err: errors.New(`ent: missing required field "Signin.timestamp"`)}
	}
	if _, ok := sc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Signin.type"`)}
	}
	if v, ok := sc.mutation.GetType(); ok {
		if err := signin.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Signin.type": %w`, err)}
		}
	}
	if _, ok := sc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Signin.user"`)}
	}
	return nil
}

func (sc *SigninCreate) sqlSave(ctx context.Context) (*Signin, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SigninCreate) createSpec() (*Signin, *sqlgraph.CreateSpec) {
	var (
		_node = &Signin{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(signin.Table, sqlgraph.NewFieldSpec(signin.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.Timestamp(); ok {
		_spec.SetField(signin.FieldTimestamp, field.TypeTime, value)
		_node.Timestamp = value
	}
	if value, ok := sc.mutation.GetType(); ok {
		_spec.SetField(signin.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if nodes := sc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   signin.UserTable,
			Columns: []string{signin.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_signins = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SigninCreateBulk is the builder for creating many Signin entities in bulk.
type SigninCreateBulk struct {
	config
	builders []*SigninCreate
}

// Save creates the Signin entities in the database.
func (scb *SigninCreateBulk) Save(ctx context.Context) ([]*Signin, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Signin, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SigninMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SigninCreateBulk) SaveX(ctx context.Context) []*Signin {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SigninCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SigninCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
