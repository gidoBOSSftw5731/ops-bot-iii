// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.ritsec.cloud/1nv8rZim/ops-bot-iii/ent/predicate"
	"gitlab.ritsec.cloud/1nv8rZim/ops-bot-iii/ent/voteresult"
)

// VoteResultQuery is the builder for querying VoteResult entities.
type VoteResultQuery struct {
	config
	ctx        *QueryContext
	order      []voteresult.OrderOption
	inters     []Interceptor
	predicates []predicate.VoteResult
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VoteResultQuery builder.
func (vrq *VoteResultQuery) Where(ps ...predicate.VoteResult) *VoteResultQuery {
	vrq.predicates = append(vrq.predicates, ps...)
	return vrq
}

// Limit the number of records to be returned by this query.
func (vrq *VoteResultQuery) Limit(limit int) *VoteResultQuery {
	vrq.ctx.Limit = &limit
	return vrq
}

// Offset to start from.
func (vrq *VoteResultQuery) Offset(offset int) *VoteResultQuery {
	vrq.ctx.Offset = &offset
	return vrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vrq *VoteResultQuery) Unique(unique bool) *VoteResultQuery {
	vrq.ctx.Unique = &unique
	return vrq
}

// Order specifies how the records should be ordered.
func (vrq *VoteResultQuery) Order(o ...voteresult.OrderOption) *VoteResultQuery {
	vrq.order = append(vrq.order, o...)
	return vrq
}

// First returns the first VoteResult entity from the query.
// Returns a *NotFoundError when no VoteResult was found.
func (vrq *VoteResultQuery) First(ctx context.Context) (*VoteResult, error) {
	nodes, err := vrq.Limit(1).All(setContextOp(ctx, vrq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{voteresult.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vrq *VoteResultQuery) FirstX(ctx context.Context) *VoteResult {
	node, err := vrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VoteResult ID from the query.
// Returns a *NotFoundError when no VoteResult ID was found.
func (vrq *VoteResultQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vrq.Limit(1).IDs(setContextOp(ctx, vrq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{voteresult.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vrq *VoteResultQuery) FirstIDX(ctx context.Context) int {
	id, err := vrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VoteResult entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VoteResult entity is found.
// Returns a *NotFoundError when no VoteResult entities are found.
func (vrq *VoteResultQuery) Only(ctx context.Context) (*VoteResult, error) {
	nodes, err := vrq.Limit(2).All(setContextOp(ctx, vrq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{voteresult.Label}
	default:
		return nil, &NotSingularError{voteresult.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vrq *VoteResultQuery) OnlyX(ctx context.Context) *VoteResult {
	node, err := vrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VoteResult ID in the query.
// Returns a *NotSingularError when more than one VoteResult ID is found.
// Returns a *NotFoundError when no entities are found.
func (vrq *VoteResultQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vrq.Limit(2).IDs(setContextOp(ctx, vrq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{voteresult.Label}
	default:
		err = &NotSingularError{voteresult.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vrq *VoteResultQuery) OnlyIDX(ctx context.Context) int {
	id, err := vrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VoteResults.
func (vrq *VoteResultQuery) All(ctx context.Context) ([]*VoteResult, error) {
	ctx = setContextOp(ctx, vrq.ctx, "All")
	if err := vrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*VoteResult, *VoteResultQuery]()
	return withInterceptors[[]*VoteResult](ctx, vrq, qr, vrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vrq *VoteResultQuery) AllX(ctx context.Context) []*VoteResult {
	nodes, err := vrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VoteResult IDs.
func (vrq *VoteResultQuery) IDs(ctx context.Context) (ids []int, err error) {
	if vrq.ctx.Unique == nil && vrq.path != nil {
		vrq.Unique(true)
	}
	ctx = setContextOp(ctx, vrq.ctx, "IDs")
	if err = vrq.Select(voteresult.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vrq *VoteResultQuery) IDsX(ctx context.Context) []int {
	ids, err := vrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vrq *VoteResultQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vrq.ctx, "Count")
	if err := vrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vrq, querierCount[*VoteResultQuery](), vrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vrq *VoteResultQuery) CountX(ctx context.Context) int {
	count, err := vrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vrq *VoteResultQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vrq.ctx, "Exist")
	switch _, err := vrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vrq *VoteResultQuery) ExistX(ctx context.Context) bool {
	exist, err := vrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VoteResultQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vrq *VoteResultQuery) Clone() *VoteResultQuery {
	if vrq == nil {
		return nil
	}
	return &VoteResultQuery{
		config:     vrq.config,
		ctx:        vrq.ctx.Clone(),
		order:      append([]voteresult.OrderOption{}, vrq.order...),
		inters:     append([]Interceptor{}, vrq.inters...),
		predicates: append([]predicate.VoteResult{}, vrq.predicates...),
		// clone intermediate query.
		sql:  vrq.sql.Clone(),
		path: vrq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		HTML string `json:"html,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.VoteResult.Query().
//		GroupBy(voteresult.FieldHTML).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vrq *VoteResultQuery) GroupBy(field string, fields ...string) *VoteResultGroupBy {
	vrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VoteResultGroupBy{build: vrq}
	grbuild.flds = &vrq.ctx.Fields
	grbuild.label = voteresult.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		HTML string `json:"html,omitempty"`
//	}
//
//	client.VoteResult.Query().
//		Select(voteresult.FieldHTML).
//		Scan(ctx, &v)
func (vrq *VoteResultQuery) Select(fields ...string) *VoteResultSelect {
	vrq.ctx.Fields = append(vrq.ctx.Fields, fields...)
	sbuild := &VoteResultSelect{VoteResultQuery: vrq}
	sbuild.label = voteresult.Label
	sbuild.flds, sbuild.scan = &vrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VoteResultSelect configured with the given aggregations.
func (vrq *VoteResultQuery) Aggregate(fns ...AggregateFunc) *VoteResultSelect {
	return vrq.Select().Aggregate(fns...)
}

func (vrq *VoteResultQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vrq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vrq); err != nil {
				return err
			}
		}
	}
	for _, f := range vrq.ctx.Fields {
		if !voteresult.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vrq.path != nil {
		prev, err := vrq.path(ctx)
		if err != nil {
			return err
		}
		vrq.sql = prev
	}
	return nil
}

func (vrq *VoteResultQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VoteResult, error) {
	var (
		nodes = []*VoteResult{}
		_spec = vrq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*VoteResult).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &VoteResult{config: vrq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (vrq *VoteResultQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vrq.querySpec()
	_spec.Node.Columns = vrq.ctx.Fields
	if len(vrq.ctx.Fields) > 0 {
		_spec.Unique = vrq.ctx.Unique != nil && *vrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vrq.driver, _spec)
}

func (vrq *VoteResultQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(voteresult.Table, voteresult.Columns, sqlgraph.NewFieldSpec(voteresult.FieldID, field.TypeInt))
	_spec.From = vrq.sql
	if unique := vrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vrq.path != nil {
		_spec.Unique = true
	}
	if fields := vrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, voteresult.FieldID)
		for i := range fields {
			if fields[i] != voteresult.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vrq *VoteResultQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vrq.driver.Dialect())
	t1 := builder.Table(voteresult.Table)
	columns := vrq.ctx.Fields
	if len(columns) == 0 {
		columns = voteresult.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vrq.sql != nil {
		selector = vrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vrq.ctx.Unique != nil && *vrq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vrq.predicates {
		p(selector)
	}
	for _, p := range vrq.order {
		p(selector)
	}
	if offset := vrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VoteResultGroupBy is the group-by builder for VoteResult entities.
type VoteResultGroupBy struct {
	selector
	build *VoteResultQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vrgb *VoteResultGroupBy) Aggregate(fns ...AggregateFunc) *VoteResultGroupBy {
	vrgb.fns = append(vrgb.fns, fns...)
	return vrgb
}

// Scan applies the selector query and scans the result into the given value.
func (vrgb *VoteResultGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vrgb.build.ctx, "GroupBy")
	if err := vrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VoteResultQuery, *VoteResultGroupBy](ctx, vrgb.build, vrgb, vrgb.build.inters, v)
}

func (vrgb *VoteResultGroupBy) sqlScan(ctx context.Context, root *VoteResultQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vrgb.fns))
	for _, fn := range vrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vrgb.flds)+len(vrgb.fns))
		for _, f := range *vrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VoteResultSelect is the builder for selecting fields of VoteResult entities.
type VoteResultSelect struct {
	*VoteResultQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vrs *VoteResultSelect) Aggregate(fns ...AggregateFunc) *VoteResultSelect {
	vrs.fns = append(vrs.fns, fns...)
	return vrs
}

// Scan applies the selector query and scans the result into the given value.
func (vrs *VoteResultSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vrs.ctx, "Select")
	if err := vrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VoteResultQuery, *VoteResultSelect](ctx, vrs.VoteResultQuery, vrs, vrs.inters, v)
}

func (vrs *VoteResultSelect) sqlScan(ctx context.Context, root *VoteResultQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vrs.fns))
	for _, fn := range vrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
