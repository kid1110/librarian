// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedactionset"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// FeedActionSetQuery is the builder for querying FeedActionSet entities.
type FeedActionSetQuery struct {
	config
	ctx            *QueryContext
	order          []feedactionset.OrderOption
	inters         []Interceptor
	predicates     []predicate.FeedActionSet
	withOwner      *UserQuery
	withFeedConfig *FeedConfigQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FeedActionSetQuery builder.
func (fasq *FeedActionSetQuery) Where(ps ...predicate.FeedActionSet) *FeedActionSetQuery {
	fasq.predicates = append(fasq.predicates, ps...)
	return fasq
}

// Limit the number of records to be returned by this query.
func (fasq *FeedActionSetQuery) Limit(limit int) *FeedActionSetQuery {
	fasq.ctx.Limit = &limit
	return fasq
}

// Offset to start from.
func (fasq *FeedActionSetQuery) Offset(offset int) *FeedActionSetQuery {
	fasq.ctx.Offset = &offset
	return fasq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fasq *FeedActionSetQuery) Unique(unique bool) *FeedActionSetQuery {
	fasq.ctx.Unique = &unique
	return fasq
}

// Order specifies how the records should be ordered.
func (fasq *FeedActionSetQuery) Order(o ...feedactionset.OrderOption) *FeedActionSetQuery {
	fasq.order = append(fasq.order, o...)
	return fasq
}

// QueryOwner chains the current query on the "owner" edge.
func (fasq *FeedActionSetQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: fasq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fasq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fasq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(feedactionset.Table, feedactionset.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, feedactionset.OwnerTable, feedactionset.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(fasq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFeedConfig chains the current query on the "feed_config" edge.
func (fasq *FeedActionSetQuery) QueryFeedConfig() *FeedConfigQuery {
	query := (&FeedConfigClient{config: fasq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fasq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fasq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(feedactionset.Table, feedactionset.FieldID, selector),
			sqlgraph.To(feedconfig.Table, feedconfig.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, feedactionset.FeedConfigTable, feedactionset.FeedConfigPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(fasq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FeedActionSet entity from the query.
// Returns a *NotFoundError when no FeedActionSet was found.
func (fasq *FeedActionSetQuery) First(ctx context.Context) (*FeedActionSet, error) {
	nodes, err := fasq.Limit(1).All(setContextOp(ctx, fasq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{feedactionset.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fasq *FeedActionSetQuery) FirstX(ctx context.Context) *FeedActionSet {
	node, err := fasq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FeedActionSet ID from the query.
// Returns a *NotFoundError when no FeedActionSet ID was found.
func (fasq *FeedActionSetQuery) FirstID(ctx context.Context) (id model.InternalID, err error) {
	var ids []model.InternalID
	if ids, err = fasq.Limit(1).IDs(setContextOp(ctx, fasq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{feedactionset.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fasq *FeedActionSetQuery) FirstIDX(ctx context.Context) model.InternalID {
	id, err := fasq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FeedActionSet entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FeedActionSet entity is found.
// Returns a *NotFoundError when no FeedActionSet entities are found.
func (fasq *FeedActionSetQuery) Only(ctx context.Context) (*FeedActionSet, error) {
	nodes, err := fasq.Limit(2).All(setContextOp(ctx, fasq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{feedactionset.Label}
	default:
		return nil, &NotSingularError{feedactionset.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fasq *FeedActionSetQuery) OnlyX(ctx context.Context) *FeedActionSet {
	node, err := fasq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FeedActionSet ID in the query.
// Returns a *NotSingularError when more than one FeedActionSet ID is found.
// Returns a *NotFoundError when no entities are found.
func (fasq *FeedActionSetQuery) OnlyID(ctx context.Context) (id model.InternalID, err error) {
	var ids []model.InternalID
	if ids, err = fasq.Limit(2).IDs(setContextOp(ctx, fasq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{feedactionset.Label}
	default:
		err = &NotSingularError{feedactionset.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fasq *FeedActionSetQuery) OnlyIDX(ctx context.Context) model.InternalID {
	id, err := fasq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FeedActionSets.
func (fasq *FeedActionSetQuery) All(ctx context.Context) ([]*FeedActionSet, error) {
	ctx = setContextOp(ctx, fasq.ctx, ent.OpQueryAll)
	if err := fasq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FeedActionSet, *FeedActionSetQuery]()
	return withInterceptors[[]*FeedActionSet](ctx, fasq, qr, fasq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fasq *FeedActionSetQuery) AllX(ctx context.Context) []*FeedActionSet {
	nodes, err := fasq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FeedActionSet IDs.
func (fasq *FeedActionSetQuery) IDs(ctx context.Context) (ids []model.InternalID, err error) {
	if fasq.ctx.Unique == nil && fasq.path != nil {
		fasq.Unique(true)
	}
	ctx = setContextOp(ctx, fasq.ctx, ent.OpQueryIDs)
	if err = fasq.Select(feedactionset.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fasq *FeedActionSetQuery) IDsX(ctx context.Context) []model.InternalID {
	ids, err := fasq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fasq *FeedActionSetQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fasq.ctx, ent.OpQueryCount)
	if err := fasq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fasq, querierCount[*FeedActionSetQuery](), fasq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fasq *FeedActionSetQuery) CountX(ctx context.Context) int {
	count, err := fasq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fasq *FeedActionSetQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fasq.ctx, ent.OpQueryExist)
	switch _, err := fasq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fasq *FeedActionSetQuery) ExistX(ctx context.Context) bool {
	exist, err := fasq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FeedActionSetQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fasq *FeedActionSetQuery) Clone() *FeedActionSetQuery {
	if fasq == nil {
		return nil
	}
	return &FeedActionSetQuery{
		config:         fasq.config,
		ctx:            fasq.ctx.Clone(),
		order:          append([]feedactionset.OrderOption{}, fasq.order...),
		inters:         append([]Interceptor{}, fasq.inters...),
		predicates:     append([]predicate.FeedActionSet{}, fasq.predicates...),
		withOwner:      fasq.withOwner.Clone(),
		withFeedConfig: fasq.withFeedConfig.Clone(),
		// clone intermediate query.
		sql:  fasq.sql.Clone(),
		path: fasq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (fasq *FeedActionSetQuery) WithOwner(opts ...func(*UserQuery)) *FeedActionSetQuery {
	query := (&UserClient{config: fasq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fasq.withOwner = query
	return fasq
}

// WithFeedConfig tells the query-builder to eager-load the nodes that are connected to
// the "feed_config" edge. The optional arguments are used to configure the query builder of the edge.
func (fasq *FeedActionSetQuery) WithFeedConfig(opts ...func(*FeedConfigQuery)) *FeedActionSetQuery {
	query := (&FeedConfigClient{config: fasq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fasq.withFeedConfig = query
	return fasq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FeedActionSet.Query().
//		GroupBy(feedactionset.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fasq *FeedActionSetQuery) GroupBy(field string, fields ...string) *FeedActionSetGroupBy {
	fasq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FeedActionSetGroupBy{build: fasq}
	grbuild.flds = &fasq.ctx.Fields
	grbuild.label = feedactionset.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.FeedActionSet.Query().
//		Select(feedactionset.FieldName).
//		Scan(ctx, &v)
func (fasq *FeedActionSetQuery) Select(fields ...string) *FeedActionSetSelect {
	fasq.ctx.Fields = append(fasq.ctx.Fields, fields...)
	sbuild := &FeedActionSetSelect{FeedActionSetQuery: fasq}
	sbuild.label = feedactionset.Label
	sbuild.flds, sbuild.scan = &fasq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FeedActionSetSelect configured with the given aggregations.
func (fasq *FeedActionSetQuery) Aggregate(fns ...AggregateFunc) *FeedActionSetSelect {
	return fasq.Select().Aggregate(fns...)
}

func (fasq *FeedActionSetQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fasq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fasq); err != nil {
				return err
			}
		}
	}
	for _, f := range fasq.ctx.Fields {
		if !feedactionset.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fasq.path != nil {
		prev, err := fasq.path(ctx)
		if err != nil {
			return err
		}
		fasq.sql = prev
	}
	return nil
}

func (fasq *FeedActionSetQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FeedActionSet, error) {
	var (
		nodes       = []*FeedActionSet{}
		withFKs     = fasq.withFKs
		_spec       = fasq.querySpec()
		loadedTypes = [2]bool{
			fasq.withOwner != nil,
			fasq.withFeedConfig != nil,
		}
	)
	if fasq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, feedactionset.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FeedActionSet).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FeedActionSet{config: fasq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fasq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fasq.withOwner; query != nil {
		if err := fasq.loadOwner(ctx, query, nodes, nil,
			func(n *FeedActionSet, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	if query := fasq.withFeedConfig; query != nil {
		if err := fasq.loadFeedConfig(ctx, query, nodes,
			func(n *FeedActionSet) { n.Edges.FeedConfig = []*FeedConfig{} },
			func(n *FeedActionSet, e *FeedConfig) { n.Edges.FeedConfig = append(n.Edges.FeedConfig, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fasq *FeedActionSetQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*FeedActionSet, init func(*FeedActionSet), assign func(*FeedActionSet, *User)) error {
	ids := make([]model.InternalID, 0, len(nodes))
	nodeids := make(map[model.InternalID][]*FeedActionSet)
	for i := range nodes {
		if nodes[i].user_feed_action_set == nil {
			continue
		}
		fk := *nodes[i].user_feed_action_set
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_feed_action_set" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (fasq *FeedActionSetQuery) loadFeedConfig(ctx context.Context, query *FeedConfigQuery, nodes []*FeedActionSet, init func(*FeedActionSet), assign func(*FeedActionSet, *FeedConfig)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[model.InternalID]*FeedActionSet)
	nids := make(map[model.InternalID]map[*FeedActionSet]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(feedactionset.FeedConfigTable)
		s.Join(joinT).On(s.C(feedconfig.FieldID), joinT.C(feedactionset.FeedConfigPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(feedactionset.FeedConfigPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(feedactionset.FeedConfigPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := model.InternalID(values[0].(*sql.NullInt64).Int64)
				inValue := model.InternalID(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*FeedActionSet]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*FeedConfig](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "feed_config" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (fasq *FeedActionSetQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fasq.querySpec()
	_spec.Node.Columns = fasq.ctx.Fields
	if len(fasq.ctx.Fields) > 0 {
		_spec.Unique = fasq.ctx.Unique != nil && *fasq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fasq.driver, _spec)
}

func (fasq *FeedActionSetQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(feedactionset.Table, feedactionset.Columns, sqlgraph.NewFieldSpec(feedactionset.FieldID, field.TypeInt64))
	_spec.From = fasq.sql
	if unique := fasq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fasq.path != nil {
		_spec.Unique = true
	}
	if fields := fasq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, feedactionset.FieldID)
		for i := range fields {
			if fields[i] != feedactionset.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fasq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fasq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fasq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fasq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fasq *FeedActionSetQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fasq.driver.Dialect())
	t1 := builder.Table(feedactionset.Table)
	columns := fasq.ctx.Fields
	if len(columns) == 0 {
		columns = feedactionset.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fasq.sql != nil {
		selector = fasq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fasq.ctx.Unique != nil && *fasq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fasq.predicates {
		p(selector)
	}
	for _, p := range fasq.order {
		p(selector)
	}
	if offset := fasq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fasq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FeedActionSetGroupBy is the group-by builder for FeedActionSet entities.
type FeedActionSetGroupBy struct {
	selector
	build *FeedActionSetQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fasgb *FeedActionSetGroupBy) Aggregate(fns ...AggregateFunc) *FeedActionSetGroupBy {
	fasgb.fns = append(fasgb.fns, fns...)
	return fasgb
}

// Scan applies the selector query and scans the result into the given value.
func (fasgb *FeedActionSetGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fasgb.build.ctx, ent.OpQueryGroupBy)
	if err := fasgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FeedActionSetQuery, *FeedActionSetGroupBy](ctx, fasgb.build, fasgb, fasgb.build.inters, v)
}

func (fasgb *FeedActionSetGroupBy) sqlScan(ctx context.Context, root *FeedActionSetQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fasgb.fns))
	for _, fn := range fasgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fasgb.flds)+len(fasgb.fns))
		for _, f := range *fasgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fasgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fasgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FeedActionSetSelect is the builder for selecting fields of FeedActionSet entities.
type FeedActionSetSelect struct {
	*FeedActionSetQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fass *FeedActionSetSelect) Aggregate(fns ...AggregateFunc) *FeedActionSetSelect {
	fass.fns = append(fass.fns, fns...)
	return fass
}

// Scan applies the selector query and scans the result into the given value.
func (fass *FeedActionSetSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fass.ctx, ent.OpQuerySelect)
	if err := fass.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FeedActionSetQuery, *FeedActionSetSelect](ctx, fass.FeedActionSetQuery, fass, fass.inters, v)
}

func (fass *FeedActionSetSelect) sqlScan(ctx context.Context, root *FeedActionSetQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fass.fns))
	for _, fn := range fass.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fass.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fass.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
