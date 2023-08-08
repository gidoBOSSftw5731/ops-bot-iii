// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.ritsec.cloud/1nv8rZim/ops-bot-iii/ent/predicate"
	"gitlab.ritsec.cloud/1nv8rZim/ops-bot-iii/ent/voteresult"
)

// VoteResultUpdate is the builder for updating VoteResult entities.
type VoteResultUpdate struct {
	config
	hooks    []Hook
	mutation *VoteResultMutation
}

// Where appends a list predicates to the VoteResultUpdate builder.
func (vru *VoteResultUpdate) Where(ps ...predicate.VoteResult) *VoteResultUpdate {
	vru.mutation.Where(ps...)
	return vru
}

// SetHTML sets the "html" field.
func (vru *VoteResultUpdate) SetHTML(s string) *VoteResultUpdate {
	vru.mutation.SetHTML(s)
	return vru
}

// SetPlain sets the "plain" field.
func (vru *VoteResultUpdate) SetPlain(s string) *VoteResultUpdate {
	vru.mutation.SetPlain(s)
	return vru
}

// SetVoteID sets the "vote_id" field.
func (vru *VoteResultUpdate) SetVoteID(s string) *VoteResultUpdate {
	vru.mutation.SetVoteID(s)
	return vru
}

// Mutation returns the VoteResultMutation object of the builder.
func (vru *VoteResultUpdate) Mutation() *VoteResultMutation {
	return vru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vru *VoteResultUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, VoteResultMutation](ctx, vru.sqlSave, vru.mutation, vru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vru *VoteResultUpdate) SaveX(ctx context.Context) int {
	affected, err := vru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vru *VoteResultUpdate) Exec(ctx context.Context) error {
	_, err := vru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vru *VoteResultUpdate) ExecX(ctx context.Context) {
	if err := vru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vru *VoteResultUpdate) check() error {
	if v, ok := vru.mutation.VoteID(); ok {
		if err := voteresult.VoteIDValidator(v); err != nil {
			return &ValidationError{Name: "vote_id", err: fmt.Errorf(`ent: validator failed for field "VoteResult.vote_id": %w`, err)}
		}
	}
	return nil
}

func (vru *VoteResultUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := vru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(voteresult.Table, voteresult.Columns, sqlgraph.NewFieldSpec(voteresult.FieldID, field.TypeInt))
	if ps := vru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vru.mutation.HTML(); ok {
		_spec.SetField(voteresult.FieldHTML, field.TypeString, value)
	}
	if value, ok := vru.mutation.Plain(); ok {
		_spec.SetField(voteresult.FieldPlain, field.TypeString, value)
	}
	if value, ok := vru.mutation.VoteID(); ok {
		_spec.SetField(voteresult.FieldVoteID, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{voteresult.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vru.mutation.done = true
	return n, nil
}

// VoteResultUpdateOne is the builder for updating a single VoteResult entity.
type VoteResultUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VoteResultMutation
}

// SetHTML sets the "html" field.
func (vruo *VoteResultUpdateOne) SetHTML(s string) *VoteResultUpdateOne {
	vruo.mutation.SetHTML(s)
	return vruo
}

// SetPlain sets the "plain" field.
func (vruo *VoteResultUpdateOne) SetPlain(s string) *VoteResultUpdateOne {
	vruo.mutation.SetPlain(s)
	return vruo
}

// SetVoteID sets the "vote_id" field.
func (vruo *VoteResultUpdateOne) SetVoteID(s string) *VoteResultUpdateOne {
	vruo.mutation.SetVoteID(s)
	return vruo
}

// Mutation returns the VoteResultMutation object of the builder.
func (vruo *VoteResultUpdateOne) Mutation() *VoteResultMutation {
	return vruo.mutation
}

// Where appends a list predicates to the VoteResultUpdate builder.
func (vruo *VoteResultUpdateOne) Where(ps ...predicate.VoteResult) *VoteResultUpdateOne {
	vruo.mutation.Where(ps...)
	return vruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vruo *VoteResultUpdateOne) Select(field string, fields ...string) *VoteResultUpdateOne {
	vruo.fields = append([]string{field}, fields...)
	return vruo
}

// Save executes the query and returns the updated VoteResult entity.
func (vruo *VoteResultUpdateOne) Save(ctx context.Context) (*VoteResult, error) {
	return withHooks[*VoteResult, VoteResultMutation](ctx, vruo.sqlSave, vruo.mutation, vruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vruo *VoteResultUpdateOne) SaveX(ctx context.Context) *VoteResult {
	node, err := vruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vruo *VoteResultUpdateOne) Exec(ctx context.Context) error {
	_, err := vruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vruo *VoteResultUpdateOne) ExecX(ctx context.Context) {
	if err := vruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vruo *VoteResultUpdateOne) check() error {
	if v, ok := vruo.mutation.VoteID(); ok {
		if err := voteresult.VoteIDValidator(v); err != nil {
			return &ValidationError{Name: "vote_id", err: fmt.Errorf(`ent: validator failed for field "VoteResult.vote_id": %w`, err)}
		}
	}
	return nil
}

func (vruo *VoteResultUpdateOne) sqlSave(ctx context.Context) (_node *VoteResult, err error) {
	if err := vruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(voteresult.Table, voteresult.Columns, sqlgraph.NewFieldSpec(voteresult.FieldID, field.TypeInt))
	id, ok := vruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "VoteResult.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, voteresult.FieldID)
		for _, f := range fields {
			if !voteresult.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != voteresult.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vruo.mutation.HTML(); ok {
		_spec.SetField(voteresult.FieldHTML, field.TypeString, value)
	}
	if value, ok := vruo.mutation.Plain(); ok {
		_spec.SetField(voteresult.FieldPlain, field.TypeString, value)
	}
	if value, ok := vruo.mutation.VoteID(); ok {
		_spec.SetField(voteresult.FieldVoteID, field.TypeString, value)
	}
	_node = &VoteResult{config: vruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{voteresult.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vruo.mutation.done = true
	return _node, nil
}
