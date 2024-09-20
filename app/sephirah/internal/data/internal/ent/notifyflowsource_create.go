// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflow"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflowsource"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifysource"
	"github.com/tuihub/librarian/internal/model"
)

// NotifyFlowSourceCreate is the builder for creating a NotifyFlowSource entity.
type NotifyFlowSourceCreate struct {
	config
	mutation *NotifyFlowSourceMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetNotifyFlowID sets the "notify_flow_id" field.
func (nfsc *NotifyFlowSourceCreate) SetNotifyFlowID(mi model.InternalID) *NotifyFlowSourceCreate {
	nfsc.mutation.SetNotifyFlowID(mi)
	return nfsc
}

// SetNotifySourceID sets the "notify_source_id" field.
func (nfsc *NotifyFlowSourceCreate) SetNotifySourceID(mi model.InternalID) *NotifyFlowSourceCreate {
	nfsc.mutation.SetNotifySourceID(mi)
	return nfsc
}

// SetFilterIncludeKeywords sets the "filter_include_keywords" field.
func (nfsc *NotifyFlowSourceCreate) SetFilterIncludeKeywords(s []string) *NotifyFlowSourceCreate {
	nfsc.mutation.SetFilterIncludeKeywords(s)
	return nfsc
}

// SetFilterExcludeKeywords sets the "filter_exclude_keywords" field.
func (nfsc *NotifyFlowSourceCreate) SetFilterExcludeKeywords(s []string) *NotifyFlowSourceCreate {
	nfsc.mutation.SetFilterExcludeKeywords(s)
	return nfsc
}

// SetUpdatedAt sets the "updated_at" field.
func (nfsc *NotifyFlowSourceCreate) SetUpdatedAt(t time.Time) *NotifyFlowSourceCreate {
	nfsc.mutation.SetUpdatedAt(t)
	return nfsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nfsc *NotifyFlowSourceCreate) SetNillableUpdatedAt(t *time.Time) *NotifyFlowSourceCreate {
	if t != nil {
		nfsc.SetUpdatedAt(*t)
	}
	return nfsc
}

// SetCreatedAt sets the "created_at" field.
func (nfsc *NotifyFlowSourceCreate) SetCreatedAt(t time.Time) *NotifyFlowSourceCreate {
	nfsc.mutation.SetCreatedAt(t)
	return nfsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nfsc *NotifyFlowSourceCreate) SetNillableCreatedAt(t *time.Time) *NotifyFlowSourceCreate {
	if t != nil {
		nfsc.SetCreatedAt(*t)
	}
	return nfsc
}

// SetNotifyFlow sets the "notify_flow" edge to the NotifyFlow entity.
func (nfsc *NotifyFlowSourceCreate) SetNotifyFlow(n *NotifyFlow) *NotifyFlowSourceCreate {
	return nfsc.SetNotifyFlowID(n.ID)
}

// SetNotifySource sets the "notify_source" edge to the NotifySource entity.
func (nfsc *NotifyFlowSourceCreate) SetNotifySource(n *NotifySource) *NotifyFlowSourceCreate {
	return nfsc.SetNotifySourceID(n.ID)
}

// Mutation returns the NotifyFlowSourceMutation object of the builder.
func (nfsc *NotifyFlowSourceCreate) Mutation() *NotifyFlowSourceMutation {
	return nfsc.mutation
}

// Save creates the NotifyFlowSource in the database.
func (nfsc *NotifyFlowSourceCreate) Save(ctx context.Context) (*NotifyFlowSource, error) {
	nfsc.defaults()
	return withHooks(ctx, nfsc.sqlSave, nfsc.mutation, nfsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nfsc *NotifyFlowSourceCreate) SaveX(ctx context.Context) *NotifyFlowSource {
	v, err := nfsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nfsc *NotifyFlowSourceCreate) Exec(ctx context.Context) error {
	_, err := nfsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nfsc *NotifyFlowSourceCreate) ExecX(ctx context.Context) {
	if err := nfsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nfsc *NotifyFlowSourceCreate) defaults() {
	if _, ok := nfsc.mutation.UpdatedAt(); !ok {
		v := notifyflowsource.DefaultUpdatedAt()
		nfsc.mutation.SetUpdatedAt(v)
	}
	if _, ok := nfsc.mutation.CreatedAt(); !ok {
		v := notifyflowsource.DefaultCreatedAt()
		nfsc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nfsc *NotifyFlowSourceCreate) check() error {
	if _, ok := nfsc.mutation.NotifyFlowID(); !ok {
		return &ValidationError{Name: "notify_flow_id", err: errors.New(`ent: missing required field "NotifyFlowSource.notify_flow_id"`)}
	}
	if _, ok := nfsc.mutation.NotifySourceID(); !ok {
		return &ValidationError{Name: "notify_source_id", err: errors.New(`ent: missing required field "NotifyFlowSource.notify_source_id"`)}
	}
	if _, ok := nfsc.mutation.FilterIncludeKeywords(); !ok {
		return &ValidationError{Name: "filter_include_keywords", err: errors.New(`ent: missing required field "NotifyFlowSource.filter_include_keywords"`)}
	}
	if _, ok := nfsc.mutation.FilterExcludeKeywords(); !ok {
		return &ValidationError{Name: "filter_exclude_keywords", err: errors.New(`ent: missing required field "NotifyFlowSource.filter_exclude_keywords"`)}
	}
	if _, ok := nfsc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "NotifyFlowSource.updated_at"`)}
	}
	if _, ok := nfsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "NotifyFlowSource.created_at"`)}
	}
	if len(nfsc.mutation.NotifyFlowIDs()) == 0 {
		return &ValidationError{Name: "notify_flow", err: errors.New(`ent: missing required edge "NotifyFlowSource.notify_flow"`)}
	}
	if len(nfsc.mutation.NotifySourceIDs()) == 0 {
		return &ValidationError{Name: "notify_source", err: errors.New(`ent: missing required edge "NotifyFlowSource.notify_source"`)}
	}
	return nil
}

func (nfsc *NotifyFlowSourceCreate) sqlSave(ctx context.Context) (*NotifyFlowSource, error) {
	if err := nfsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nfsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nfsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	nfsc.mutation.id = &_node.ID
	nfsc.mutation.done = true
	return _node, nil
}

func (nfsc *NotifyFlowSourceCreate) createSpec() (*NotifyFlowSource, *sqlgraph.CreateSpec) {
	var (
		_node = &NotifyFlowSource{config: nfsc.config}
		_spec = sqlgraph.NewCreateSpec(notifyflowsource.Table, sqlgraph.NewFieldSpec(notifyflowsource.FieldID, field.TypeInt))
	)
	_spec.OnConflict = nfsc.conflict
	if value, ok := nfsc.mutation.FilterIncludeKeywords(); ok {
		_spec.SetField(notifyflowsource.FieldFilterIncludeKeywords, field.TypeJSON, value)
		_node.FilterIncludeKeywords = value
	}
	if value, ok := nfsc.mutation.FilterExcludeKeywords(); ok {
		_spec.SetField(notifyflowsource.FieldFilterExcludeKeywords, field.TypeJSON, value)
		_node.FilterExcludeKeywords = value
	}
	if value, ok := nfsc.mutation.UpdatedAt(); ok {
		_spec.SetField(notifyflowsource.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := nfsc.mutation.CreatedAt(); ok {
		_spec.SetField(notifyflowsource.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := nfsc.mutation.NotifyFlowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   notifyflowsource.NotifyFlowTable,
			Columns: []string{notifyflowsource.NotifyFlowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.NotifyFlowID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nfsc.mutation.NotifySourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   notifyflowsource.NotifySourceTable,
			Columns: []string{notifyflowsource.NotifySourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifysource.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.NotifySourceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.NotifyFlowSource.Create().
//		SetNotifyFlowID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NotifyFlowSourceUpsert) {
//			SetNotifyFlowID(v+v).
//		}).
//		Exec(ctx)
func (nfsc *NotifyFlowSourceCreate) OnConflict(opts ...sql.ConflictOption) *NotifyFlowSourceUpsertOne {
	nfsc.conflict = opts
	return &NotifyFlowSourceUpsertOne{
		create: nfsc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.NotifyFlowSource.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (nfsc *NotifyFlowSourceCreate) OnConflictColumns(columns ...string) *NotifyFlowSourceUpsertOne {
	nfsc.conflict = append(nfsc.conflict, sql.ConflictColumns(columns...))
	return &NotifyFlowSourceUpsertOne{
		create: nfsc,
	}
}

type (
	// NotifyFlowSourceUpsertOne is the builder for "upsert"-ing
	//  one NotifyFlowSource node.
	NotifyFlowSourceUpsertOne struct {
		create *NotifyFlowSourceCreate
	}

	// NotifyFlowSourceUpsert is the "OnConflict" setter.
	NotifyFlowSourceUpsert struct {
		*sql.UpdateSet
	}
)

// SetNotifyFlowID sets the "notify_flow_id" field.
func (u *NotifyFlowSourceUpsert) SetNotifyFlowID(v model.InternalID) *NotifyFlowSourceUpsert {
	u.Set(notifyflowsource.FieldNotifyFlowID, v)
	return u
}

// UpdateNotifyFlowID sets the "notify_flow_id" field to the value that was provided on create.
func (u *NotifyFlowSourceUpsert) UpdateNotifyFlowID() *NotifyFlowSourceUpsert {
	u.SetExcluded(notifyflowsource.FieldNotifyFlowID)
	return u
}

// SetNotifySourceID sets the "notify_source_id" field.
func (u *NotifyFlowSourceUpsert) SetNotifySourceID(v model.InternalID) *NotifyFlowSourceUpsert {
	u.Set(notifyflowsource.FieldNotifySourceID, v)
	return u
}

// UpdateNotifySourceID sets the "notify_source_id" field to the value that was provided on create.
func (u *NotifyFlowSourceUpsert) UpdateNotifySourceID() *NotifyFlowSourceUpsert {
	u.SetExcluded(notifyflowsource.FieldNotifySourceID)
	return u
}

// SetFilterIncludeKeywords sets the "filter_include_keywords" field.
func (u *NotifyFlowSourceUpsert) SetFilterIncludeKeywords(v []string) *NotifyFlowSourceUpsert {
	u.Set(notifyflowsource.FieldFilterIncludeKeywords, v)
	return u
}

// UpdateFilterIncludeKeywords sets the "filter_include_keywords" field to the value that was provided on create.
func (u *NotifyFlowSourceUpsert) UpdateFilterIncludeKeywords() *NotifyFlowSourceUpsert {
	u.SetExcluded(notifyflowsource.FieldFilterIncludeKeywords)
	return u
}

// SetFilterExcludeKeywords sets the "filter_exclude_keywords" field.
func (u *NotifyFlowSourceUpsert) SetFilterExcludeKeywords(v []string) *NotifyFlowSourceUpsert {
	u.Set(notifyflowsource.FieldFilterExcludeKeywords, v)
	return u
}

// UpdateFilterExcludeKeywords sets the "filter_exclude_keywords" field to the value that was provided on create.
func (u *NotifyFlowSourceUpsert) UpdateFilterExcludeKeywords() *NotifyFlowSourceUpsert {
	u.SetExcluded(notifyflowsource.FieldFilterExcludeKeywords)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NotifyFlowSourceUpsert) SetUpdatedAt(v time.Time) *NotifyFlowSourceUpsert {
	u.Set(notifyflowsource.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NotifyFlowSourceUpsert) UpdateUpdatedAt() *NotifyFlowSourceUpsert {
	u.SetExcluded(notifyflowsource.FieldUpdatedAt)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *NotifyFlowSourceUpsert) SetCreatedAt(v time.Time) *NotifyFlowSourceUpsert {
	u.Set(notifyflowsource.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *NotifyFlowSourceUpsert) UpdateCreatedAt() *NotifyFlowSourceUpsert {
	u.SetExcluded(notifyflowsource.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.NotifyFlowSource.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *NotifyFlowSourceUpsertOne) UpdateNewValues() *NotifyFlowSourceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.NotifyFlowSource.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *NotifyFlowSourceUpsertOne) Ignore() *NotifyFlowSourceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NotifyFlowSourceUpsertOne) DoNothing() *NotifyFlowSourceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NotifyFlowSourceCreate.OnConflict
// documentation for more info.
func (u *NotifyFlowSourceUpsertOne) Update(set func(*NotifyFlowSourceUpsert)) *NotifyFlowSourceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NotifyFlowSourceUpsert{UpdateSet: update})
	}))
	return u
}

// SetNotifyFlowID sets the "notify_flow_id" field.
func (u *NotifyFlowSourceUpsertOne) SetNotifyFlowID(v model.InternalID) *NotifyFlowSourceUpsertOne {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.SetNotifyFlowID(v)
	})
}

// UpdateNotifyFlowID sets the "notify_flow_id" field to the value that was provided on create.
func (u *NotifyFlowSourceUpsertOne) UpdateNotifyFlowID() *NotifyFlowSourceUpsertOne {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.UpdateNotifyFlowID()
	})
}

// SetNotifySourceID sets the "notify_source_id" field.
func (u *NotifyFlowSourceUpsertOne) SetNotifySourceID(v model.InternalID) *NotifyFlowSourceUpsertOne {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.SetNotifySourceID(v)
	})
}

// UpdateNotifySourceID sets the "notify_source_id" field to the value that was provided on create.
func (u *NotifyFlowSourceUpsertOne) UpdateNotifySourceID() *NotifyFlowSourceUpsertOne {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.UpdateNotifySourceID()
	})
}

// SetFilterIncludeKeywords sets the "filter_include_keywords" field.
func (u *NotifyFlowSourceUpsertOne) SetFilterIncludeKeywords(v []string) *NotifyFlowSourceUpsertOne {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.SetFilterIncludeKeywords(v)
	})
}

// UpdateFilterIncludeKeywords sets the "filter_include_keywords" field to the value that was provided on create.
func (u *NotifyFlowSourceUpsertOne) UpdateFilterIncludeKeywords() *NotifyFlowSourceUpsertOne {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.UpdateFilterIncludeKeywords()
	})
}

// SetFilterExcludeKeywords sets the "filter_exclude_keywords" field.
func (u *NotifyFlowSourceUpsertOne) SetFilterExcludeKeywords(v []string) *NotifyFlowSourceUpsertOne {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.SetFilterExcludeKeywords(v)
	})
}

// UpdateFilterExcludeKeywords sets the "filter_exclude_keywords" field to the value that was provided on create.
func (u *NotifyFlowSourceUpsertOne) UpdateFilterExcludeKeywords() *NotifyFlowSourceUpsertOne {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.UpdateFilterExcludeKeywords()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NotifyFlowSourceUpsertOne) SetUpdatedAt(v time.Time) *NotifyFlowSourceUpsertOne {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NotifyFlowSourceUpsertOne) UpdateUpdatedAt() *NotifyFlowSourceUpsertOne {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *NotifyFlowSourceUpsertOne) SetCreatedAt(v time.Time) *NotifyFlowSourceUpsertOne {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *NotifyFlowSourceUpsertOne) UpdateCreatedAt() *NotifyFlowSourceUpsertOne {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *NotifyFlowSourceUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NotifyFlowSourceCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NotifyFlowSourceUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *NotifyFlowSourceUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *NotifyFlowSourceUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// NotifyFlowSourceCreateBulk is the builder for creating many NotifyFlowSource entities in bulk.
type NotifyFlowSourceCreateBulk struct {
	config
	err      error
	builders []*NotifyFlowSourceCreate
	conflict []sql.ConflictOption
}

// Save creates the NotifyFlowSource entities in the database.
func (nfscb *NotifyFlowSourceCreateBulk) Save(ctx context.Context) ([]*NotifyFlowSource, error) {
	if nfscb.err != nil {
		return nil, nfscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(nfscb.builders))
	nodes := make([]*NotifyFlowSource, len(nfscb.builders))
	mutators := make([]Mutator, len(nfscb.builders))
	for i := range nfscb.builders {
		func(i int, root context.Context) {
			builder := nfscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NotifyFlowSourceMutation)
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
					_, err = mutators[i+1].Mutate(root, nfscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = nfscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, nfscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, nfscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (nfscb *NotifyFlowSourceCreateBulk) SaveX(ctx context.Context) []*NotifyFlowSource {
	v, err := nfscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nfscb *NotifyFlowSourceCreateBulk) Exec(ctx context.Context) error {
	_, err := nfscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nfscb *NotifyFlowSourceCreateBulk) ExecX(ctx context.Context) {
	if err := nfscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.NotifyFlowSource.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NotifyFlowSourceUpsert) {
//			SetNotifyFlowID(v+v).
//		}).
//		Exec(ctx)
func (nfscb *NotifyFlowSourceCreateBulk) OnConflict(opts ...sql.ConflictOption) *NotifyFlowSourceUpsertBulk {
	nfscb.conflict = opts
	return &NotifyFlowSourceUpsertBulk{
		create: nfscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.NotifyFlowSource.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (nfscb *NotifyFlowSourceCreateBulk) OnConflictColumns(columns ...string) *NotifyFlowSourceUpsertBulk {
	nfscb.conflict = append(nfscb.conflict, sql.ConflictColumns(columns...))
	return &NotifyFlowSourceUpsertBulk{
		create: nfscb,
	}
}

// NotifyFlowSourceUpsertBulk is the builder for "upsert"-ing
// a bulk of NotifyFlowSource nodes.
type NotifyFlowSourceUpsertBulk struct {
	create *NotifyFlowSourceCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.NotifyFlowSource.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *NotifyFlowSourceUpsertBulk) UpdateNewValues() *NotifyFlowSourceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.NotifyFlowSource.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *NotifyFlowSourceUpsertBulk) Ignore() *NotifyFlowSourceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NotifyFlowSourceUpsertBulk) DoNothing() *NotifyFlowSourceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NotifyFlowSourceCreateBulk.OnConflict
// documentation for more info.
func (u *NotifyFlowSourceUpsertBulk) Update(set func(*NotifyFlowSourceUpsert)) *NotifyFlowSourceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NotifyFlowSourceUpsert{UpdateSet: update})
	}))
	return u
}

// SetNotifyFlowID sets the "notify_flow_id" field.
func (u *NotifyFlowSourceUpsertBulk) SetNotifyFlowID(v model.InternalID) *NotifyFlowSourceUpsertBulk {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.SetNotifyFlowID(v)
	})
}

// UpdateNotifyFlowID sets the "notify_flow_id" field to the value that was provided on create.
func (u *NotifyFlowSourceUpsertBulk) UpdateNotifyFlowID() *NotifyFlowSourceUpsertBulk {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.UpdateNotifyFlowID()
	})
}

// SetNotifySourceID sets the "notify_source_id" field.
func (u *NotifyFlowSourceUpsertBulk) SetNotifySourceID(v model.InternalID) *NotifyFlowSourceUpsertBulk {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.SetNotifySourceID(v)
	})
}

// UpdateNotifySourceID sets the "notify_source_id" field to the value that was provided on create.
func (u *NotifyFlowSourceUpsertBulk) UpdateNotifySourceID() *NotifyFlowSourceUpsertBulk {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.UpdateNotifySourceID()
	})
}

// SetFilterIncludeKeywords sets the "filter_include_keywords" field.
func (u *NotifyFlowSourceUpsertBulk) SetFilterIncludeKeywords(v []string) *NotifyFlowSourceUpsertBulk {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.SetFilterIncludeKeywords(v)
	})
}

// UpdateFilterIncludeKeywords sets the "filter_include_keywords" field to the value that was provided on create.
func (u *NotifyFlowSourceUpsertBulk) UpdateFilterIncludeKeywords() *NotifyFlowSourceUpsertBulk {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.UpdateFilterIncludeKeywords()
	})
}

// SetFilterExcludeKeywords sets the "filter_exclude_keywords" field.
func (u *NotifyFlowSourceUpsertBulk) SetFilterExcludeKeywords(v []string) *NotifyFlowSourceUpsertBulk {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.SetFilterExcludeKeywords(v)
	})
}

// UpdateFilterExcludeKeywords sets the "filter_exclude_keywords" field to the value that was provided on create.
func (u *NotifyFlowSourceUpsertBulk) UpdateFilterExcludeKeywords() *NotifyFlowSourceUpsertBulk {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.UpdateFilterExcludeKeywords()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NotifyFlowSourceUpsertBulk) SetUpdatedAt(v time.Time) *NotifyFlowSourceUpsertBulk {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NotifyFlowSourceUpsertBulk) UpdateUpdatedAt() *NotifyFlowSourceUpsertBulk {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *NotifyFlowSourceUpsertBulk) SetCreatedAt(v time.Time) *NotifyFlowSourceUpsertBulk {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *NotifyFlowSourceUpsertBulk) UpdateCreatedAt() *NotifyFlowSourceUpsertBulk {
	return u.Update(func(s *NotifyFlowSourceUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *NotifyFlowSourceUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the NotifyFlowSourceCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NotifyFlowSourceCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NotifyFlowSourceUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
