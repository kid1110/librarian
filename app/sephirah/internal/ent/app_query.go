// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/apppackage"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/predicate"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/user"
)

// AppQuery is the builder for querying App entities.
type AppQuery struct {
	config
	ctx            *QueryContext
	order          []OrderFunc
	inters         []Interceptor
	predicates     []predicate.App
	withUser       *UserQuery
	withAppPackage *AppPackageQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppQuery builder.
func (aq *AppQuery) Where(ps ...predicate.App) *AppQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AppQuery) Limit(limit int) *AppQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AppQuery) Offset(offset int) *AppQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AppQuery) Unique(unique bool) *AppQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AppQuery) Order(o ...OrderFunc) *AppQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryUser chains the current query on the "user" edge.
func (aq *AppQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(app.Table, app.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, app.UserTable, app.UserPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAppPackage chains the current query on the "app_package" edge.
func (aq *AppQuery) QueryAppPackage() *AppPackageQuery {
	query := (&AppPackageClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(app.Table, app.FieldID, selector),
			sqlgraph.To(apppackage.Table, apppackage.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, app.AppPackageTable, app.AppPackageColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first App entity from the query.
// Returns a *NotFoundError when no App was found.
func (aq *AppQuery) First(ctx context.Context) (*App, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{app.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AppQuery) FirstX(ctx context.Context) *App {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first App ID from the query.
// Returns a *NotFoundError when no App ID was found.
func (aq *AppQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{app.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AppQuery) FirstIDX(ctx context.Context) int64 {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single App entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one App entity is found.
// Returns a *NotFoundError when no App entities are found.
func (aq *AppQuery) Only(ctx context.Context) (*App, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{app.Label}
	default:
		return nil, &NotSingularError{app.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AppQuery) OnlyX(ctx context.Context) *App {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only App ID in the query.
// Returns a *NotSingularError when more than one App ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AppQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{app.Label}
	default:
		err = &NotSingularError{app.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AppQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Apps.
func (aq *AppQuery) All(ctx context.Context) ([]*App, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*App, *AppQuery]()
	return withInterceptors[[]*App](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AppQuery) AllX(ctx context.Context) []*App {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of App IDs.
func (aq *AppQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err = aq.Select(app.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AppQuery) IDsX(ctx context.Context) []int64 {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AppQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AppQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AppQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AppQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AppQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AppQuery) Clone() *AppQuery {
	if aq == nil {
		return nil
	}
	return &AppQuery{
		config:         aq.config,
		ctx:            aq.ctx.Clone(),
		order:          append([]OrderFunc{}, aq.order...),
		inters:         append([]Interceptor{}, aq.inters...),
		predicates:     append([]predicate.App{}, aq.predicates...),
		withUser:       aq.withUser.Clone(),
		withAppPackage: aq.withAppPackage.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AppQuery) WithUser(opts ...func(*UserQuery)) *AppQuery {
	query := (&UserClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withUser = query
	return aq
}

// WithAppPackage tells the query-builder to eager-load the nodes that are connected to
// the "app_package" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AppQuery) WithAppPackage(opts ...func(*AppPackageQuery)) *AppQuery {
	query := (&AppPackageClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withAppPackage = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Source app.Source `json:"source,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.App.Query().
//		GroupBy(app.FieldSource).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AppQuery) GroupBy(field string, fields ...string) *AppGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AppGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = app.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Source app.Source `json:"source,omitempty"`
//	}
//
//	client.App.Query().
//		Select(app.FieldSource).
//		Scan(ctx, &v)
func (aq *AppQuery) Select(fields ...string) *AppSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AppSelect{AppQuery: aq}
	sbuild.label = app.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AppSelect configured with the given aggregations.
func (aq *AppQuery) Aggregate(fns ...AggregateFunc) *AppSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AppQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !app.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AppQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*App, error) {
	var (
		nodes       = []*App{}
		_spec       = aq.querySpec()
		loadedTypes = [2]bool{
			aq.withUser != nil,
			aq.withAppPackage != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*App).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &App{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withUser; query != nil {
		if err := aq.loadUser(ctx, query, nodes,
			func(n *App) { n.Edges.User = []*User{} },
			func(n *App, e *User) { n.Edges.User = append(n.Edges.User, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withAppPackage; query != nil {
		if err := aq.loadAppPackage(ctx, query, nodes,
			func(n *App) { n.Edges.AppPackage = []*AppPackage{} },
			func(n *App, e *AppPackage) { n.Edges.AppPackage = append(n.Edges.AppPackage, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AppQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*App, init func(*App), assign func(*App, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int64]*App)
	nids := make(map[int64]map[*App]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(app.UserTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(app.UserPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(app.UserPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(app.UserPrimaryKey[1]))
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
				outValue := values[0].(*sql.NullInt64).Int64
				inValue := values[1].(*sql.NullInt64).Int64
				if nids[inValue] == nil {
					nids[inValue] = map[*App]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "user" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (aq *AppQuery) loadAppPackage(ctx context.Context, query *AppPackageQuery, nodes []*App, init func(*App), assign func(*App, *AppPackage)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*App)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.InValues(app.AppPackageColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.app_app_package
		if fk == nil {
			return fmt.Errorf(`foreign-key "app_app_package" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "app_app_package" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (aq *AppQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AppQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(app.Table, app.Columns, sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt64))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, app.FieldID)
		for i := range fields {
			if fields[i] != app.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AppQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(app.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = app.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppGroupBy is the group-by builder for App entities.
type AppGroupBy struct {
	selector
	build *AppQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AppGroupBy) Aggregate(fns ...AggregateFunc) *AppGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AppGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppQuery, *AppGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AppGroupBy) sqlScan(ctx context.Context, root *AppQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AppSelect is the builder for selecting fields of App entities.
type AppSelect struct {
	*AppQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AppSelect) Aggregate(fns ...AggregateFunc) *AppSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AppSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppQuery, *AppSelect](ctx, as.AppQuery, as, as.inters, v)
}

func (as *AppSelect) sqlScan(ctx context.Context, root *AppQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
