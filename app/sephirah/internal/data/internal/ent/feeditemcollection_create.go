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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feeditem"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feeditemcollection"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifysource"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// FeedItemCollectionCreate is the builder for creating a FeedItemCollection entity.
type FeedItemCollectionCreate struct {
	config
	mutation *FeedItemCollectionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (ficc *FeedItemCollectionCreate) SetName(s string) *FeedItemCollectionCreate {
	ficc.mutation.SetName(s)
	return ficc
}

// SetDescription sets the "description" field.
func (ficc *FeedItemCollectionCreate) SetDescription(s string) *FeedItemCollectionCreate {
	ficc.mutation.SetDescription(s)
	return ficc
}

// SetCategory sets the "category" field.
func (ficc *FeedItemCollectionCreate) SetCategory(s string) *FeedItemCollectionCreate {
	ficc.mutation.SetCategory(s)
	return ficc
}

// SetUpdatedAt sets the "updated_at" field.
func (ficc *FeedItemCollectionCreate) SetUpdatedAt(t time.Time) *FeedItemCollectionCreate {
	ficc.mutation.SetUpdatedAt(t)
	return ficc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ficc *FeedItemCollectionCreate) SetNillableUpdatedAt(t *time.Time) *FeedItemCollectionCreate {
	if t != nil {
		ficc.SetUpdatedAt(*t)
	}
	return ficc
}

// SetCreatedAt sets the "created_at" field.
func (ficc *FeedItemCollectionCreate) SetCreatedAt(t time.Time) *FeedItemCollectionCreate {
	ficc.mutation.SetCreatedAt(t)
	return ficc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ficc *FeedItemCollectionCreate) SetNillableCreatedAt(t *time.Time) *FeedItemCollectionCreate {
	if t != nil {
		ficc.SetCreatedAt(*t)
	}
	return ficc
}

// SetID sets the "id" field.
func (ficc *FeedItemCollectionCreate) SetID(mi model.InternalID) *FeedItemCollectionCreate {
	ficc.mutation.SetID(mi)
	return ficc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (ficc *FeedItemCollectionCreate) SetOwnerID(id model.InternalID) *FeedItemCollectionCreate {
	ficc.mutation.SetOwnerID(id)
	return ficc
}

// SetOwner sets the "owner" edge to the User entity.
func (ficc *FeedItemCollectionCreate) SetOwner(u *User) *FeedItemCollectionCreate {
	return ficc.SetOwnerID(u.ID)
}

// AddFeedItemIDs adds the "feed_item" edge to the FeedItem entity by IDs.
func (ficc *FeedItemCollectionCreate) AddFeedItemIDs(ids ...model.InternalID) *FeedItemCollectionCreate {
	ficc.mutation.AddFeedItemIDs(ids...)
	return ficc
}

// AddFeedItem adds the "feed_item" edges to the FeedItem entity.
func (ficc *FeedItemCollectionCreate) AddFeedItem(f ...*FeedItem) *FeedItemCollectionCreate {
	ids := make([]model.InternalID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ficc.AddFeedItemIDs(ids...)
}

// AddNotifySourceIDs adds the "notify_source" edge to the NotifySource entity by IDs.
func (ficc *FeedItemCollectionCreate) AddNotifySourceIDs(ids ...model.InternalID) *FeedItemCollectionCreate {
	ficc.mutation.AddNotifySourceIDs(ids...)
	return ficc
}

// AddNotifySource adds the "notify_source" edges to the NotifySource entity.
func (ficc *FeedItemCollectionCreate) AddNotifySource(n ...*NotifySource) *FeedItemCollectionCreate {
	ids := make([]model.InternalID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ficc.AddNotifySourceIDs(ids...)
}

// Mutation returns the FeedItemCollectionMutation object of the builder.
func (ficc *FeedItemCollectionCreate) Mutation() *FeedItemCollectionMutation {
	return ficc.mutation
}

// Save creates the FeedItemCollection in the database.
func (ficc *FeedItemCollectionCreate) Save(ctx context.Context) (*FeedItemCollection, error) {
	ficc.defaults()
	return withHooks(ctx, ficc.sqlSave, ficc.mutation, ficc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ficc *FeedItemCollectionCreate) SaveX(ctx context.Context) *FeedItemCollection {
	v, err := ficc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ficc *FeedItemCollectionCreate) Exec(ctx context.Context) error {
	_, err := ficc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ficc *FeedItemCollectionCreate) ExecX(ctx context.Context) {
	if err := ficc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ficc *FeedItemCollectionCreate) defaults() {
	if _, ok := ficc.mutation.UpdatedAt(); !ok {
		v := feeditemcollection.DefaultUpdatedAt()
		ficc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ficc.mutation.CreatedAt(); !ok {
		v := feeditemcollection.DefaultCreatedAt()
		ficc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ficc *FeedItemCollectionCreate) check() error {
	if _, ok := ficc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "FeedItemCollection.name"`)}
	}
	if _, ok := ficc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "FeedItemCollection.description"`)}
	}
	if _, ok := ficc.mutation.Category(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required field "FeedItemCollection.category"`)}
	}
	if _, ok := ficc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "FeedItemCollection.updated_at"`)}
	}
	if _, ok := ficc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FeedItemCollection.created_at"`)}
	}
	if len(ficc.mutation.OwnerIDs()) == 0 {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "FeedItemCollection.owner"`)}
	}
	return nil
}

func (ficc *FeedItemCollectionCreate) sqlSave(ctx context.Context) (*FeedItemCollection, error) {
	if err := ficc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ficc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ficc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = model.InternalID(id)
	}
	ficc.mutation.id = &_node.ID
	ficc.mutation.done = true
	return _node, nil
}

func (ficc *FeedItemCollectionCreate) createSpec() (*FeedItemCollection, *sqlgraph.CreateSpec) {
	var (
		_node = &FeedItemCollection{config: ficc.config}
		_spec = sqlgraph.NewCreateSpec(feeditemcollection.Table, sqlgraph.NewFieldSpec(feeditemcollection.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = ficc.conflict
	if id, ok := ficc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ficc.mutation.Name(); ok {
		_spec.SetField(feeditemcollection.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ficc.mutation.Description(); ok {
		_spec.SetField(feeditemcollection.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ficc.mutation.Category(); ok {
		_spec.SetField(feeditemcollection.FieldCategory, field.TypeString, value)
		_node.Category = value
	}
	if value, ok := ficc.mutation.UpdatedAt(); ok {
		_spec.SetField(feeditemcollection.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ficc.mutation.CreatedAt(); ok {
		_spec.SetField(feeditemcollection.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := ficc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   feeditemcollection.OwnerTable,
			Columns: []string{feeditemcollection.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_feed_item_collection = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ficc.mutation.FeedItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   feeditemcollection.FeedItemTable,
			Columns: feeditemcollection.FeedItemPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feeditem.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ficc.mutation.NotifySourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feeditemcollection.NotifySourceTable,
			Columns: []string{feeditemcollection.NotifySourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifysource.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FeedItemCollection.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FeedItemCollectionUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ficc *FeedItemCollectionCreate) OnConflict(opts ...sql.ConflictOption) *FeedItemCollectionUpsertOne {
	ficc.conflict = opts
	return &FeedItemCollectionUpsertOne{
		create: ficc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FeedItemCollection.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ficc *FeedItemCollectionCreate) OnConflictColumns(columns ...string) *FeedItemCollectionUpsertOne {
	ficc.conflict = append(ficc.conflict, sql.ConflictColumns(columns...))
	return &FeedItemCollectionUpsertOne{
		create: ficc,
	}
}

type (
	// FeedItemCollectionUpsertOne is the builder for "upsert"-ing
	//  one FeedItemCollection node.
	FeedItemCollectionUpsertOne struct {
		create *FeedItemCollectionCreate
	}

	// FeedItemCollectionUpsert is the "OnConflict" setter.
	FeedItemCollectionUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *FeedItemCollectionUpsert) SetName(v string) *FeedItemCollectionUpsert {
	u.Set(feeditemcollection.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FeedItemCollectionUpsert) UpdateName() *FeedItemCollectionUpsert {
	u.SetExcluded(feeditemcollection.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *FeedItemCollectionUpsert) SetDescription(v string) *FeedItemCollectionUpsert {
	u.Set(feeditemcollection.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *FeedItemCollectionUpsert) UpdateDescription() *FeedItemCollectionUpsert {
	u.SetExcluded(feeditemcollection.FieldDescription)
	return u
}

// SetCategory sets the "category" field.
func (u *FeedItemCollectionUpsert) SetCategory(v string) *FeedItemCollectionUpsert {
	u.Set(feeditemcollection.FieldCategory, v)
	return u
}

// UpdateCategory sets the "category" field to the value that was provided on create.
func (u *FeedItemCollectionUpsert) UpdateCategory() *FeedItemCollectionUpsert {
	u.SetExcluded(feeditemcollection.FieldCategory)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FeedItemCollectionUpsert) SetUpdatedAt(v time.Time) *FeedItemCollectionUpsert {
	u.Set(feeditemcollection.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FeedItemCollectionUpsert) UpdateUpdatedAt() *FeedItemCollectionUpsert {
	u.SetExcluded(feeditemcollection.FieldUpdatedAt)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FeedItemCollectionUpsert) SetCreatedAt(v time.Time) *FeedItemCollectionUpsert {
	u.Set(feeditemcollection.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FeedItemCollectionUpsert) UpdateCreatedAt() *FeedItemCollectionUpsert {
	u.SetExcluded(feeditemcollection.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.FeedItemCollection.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(feeditemcollection.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FeedItemCollectionUpsertOne) UpdateNewValues() *FeedItemCollectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(feeditemcollection.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FeedItemCollection.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FeedItemCollectionUpsertOne) Ignore() *FeedItemCollectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FeedItemCollectionUpsertOne) DoNothing() *FeedItemCollectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FeedItemCollectionCreate.OnConflict
// documentation for more info.
func (u *FeedItemCollectionUpsertOne) Update(set func(*FeedItemCollectionUpsert)) *FeedItemCollectionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FeedItemCollectionUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *FeedItemCollectionUpsertOne) SetName(v string) *FeedItemCollectionUpsertOne {
	return u.Update(func(s *FeedItemCollectionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FeedItemCollectionUpsertOne) UpdateName() *FeedItemCollectionUpsertOne {
	return u.Update(func(s *FeedItemCollectionUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *FeedItemCollectionUpsertOne) SetDescription(v string) *FeedItemCollectionUpsertOne {
	return u.Update(func(s *FeedItemCollectionUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *FeedItemCollectionUpsertOne) UpdateDescription() *FeedItemCollectionUpsertOne {
	return u.Update(func(s *FeedItemCollectionUpsert) {
		s.UpdateDescription()
	})
}

// SetCategory sets the "category" field.
func (u *FeedItemCollectionUpsertOne) SetCategory(v string) *FeedItemCollectionUpsertOne {
	return u.Update(func(s *FeedItemCollectionUpsert) {
		s.SetCategory(v)
	})
}

// UpdateCategory sets the "category" field to the value that was provided on create.
func (u *FeedItemCollectionUpsertOne) UpdateCategory() *FeedItemCollectionUpsertOne {
	return u.Update(func(s *FeedItemCollectionUpsert) {
		s.UpdateCategory()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FeedItemCollectionUpsertOne) SetUpdatedAt(v time.Time) *FeedItemCollectionUpsertOne {
	return u.Update(func(s *FeedItemCollectionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FeedItemCollectionUpsertOne) UpdateUpdatedAt() *FeedItemCollectionUpsertOne {
	return u.Update(func(s *FeedItemCollectionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *FeedItemCollectionUpsertOne) SetCreatedAt(v time.Time) *FeedItemCollectionUpsertOne {
	return u.Update(func(s *FeedItemCollectionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FeedItemCollectionUpsertOne) UpdateCreatedAt() *FeedItemCollectionUpsertOne {
	return u.Update(func(s *FeedItemCollectionUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *FeedItemCollectionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FeedItemCollectionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FeedItemCollectionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FeedItemCollectionUpsertOne) ID(ctx context.Context) (id model.InternalID, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FeedItemCollectionUpsertOne) IDX(ctx context.Context) model.InternalID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FeedItemCollectionCreateBulk is the builder for creating many FeedItemCollection entities in bulk.
type FeedItemCollectionCreateBulk struct {
	config
	err      error
	builders []*FeedItemCollectionCreate
	conflict []sql.ConflictOption
}

// Save creates the FeedItemCollection entities in the database.
func (ficcb *FeedItemCollectionCreateBulk) Save(ctx context.Context) ([]*FeedItemCollection, error) {
	if ficcb.err != nil {
		return nil, ficcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ficcb.builders))
	nodes := make([]*FeedItemCollection, len(ficcb.builders))
	mutators := make([]Mutator, len(ficcb.builders))
	for i := range ficcb.builders {
		func(i int, root context.Context) {
			builder := ficcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FeedItemCollectionMutation)
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
					_, err = mutators[i+1].Mutate(root, ficcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ficcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ficcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = model.InternalID(id)
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
		if _, err := mutators[0].Mutate(ctx, ficcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ficcb *FeedItemCollectionCreateBulk) SaveX(ctx context.Context) []*FeedItemCollection {
	v, err := ficcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ficcb *FeedItemCollectionCreateBulk) Exec(ctx context.Context) error {
	_, err := ficcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ficcb *FeedItemCollectionCreateBulk) ExecX(ctx context.Context) {
	if err := ficcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FeedItemCollection.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FeedItemCollectionUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ficcb *FeedItemCollectionCreateBulk) OnConflict(opts ...sql.ConflictOption) *FeedItemCollectionUpsertBulk {
	ficcb.conflict = opts
	return &FeedItemCollectionUpsertBulk{
		create: ficcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FeedItemCollection.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ficcb *FeedItemCollectionCreateBulk) OnConflictColumns(columns ...string) *FeedItemCollectionUpsertBulk {
	ficcb.conflict = append(ficcb.conflict, sql.ConflictColumns(columns...))
	return &FeedItemCollectionUpsertBulk{
		create: ficcb,
	}
}

// FeedItemCollectionUpsertBulk is the builder for "upsert"-ing
// a bulk of FeedItemCollection nodes.
type FeedItemCollectionUpsertBulk struct {
	create *FeedItemCollectionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.FeedItemCollection.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(feeditemcollection.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FeedItemCollectionUpsertBulk) UpdateNewValues() *FeedItemCollectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(feeditemcollection.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FeedItemCollection.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FeedItemCollectionUpsertBulk) Ignore() *FeedItemCollectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FeedItemCollectionUpsertBulk) DoNothing() *FeedItemCollectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FeedItemCollectionCreateBulk.OnConflict
// documentation for more info.
func (u *FeedItemCollectionUpsertBulk) Update(set func(*FeedItemCollectionUpsert)) *FeedItemCollectionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FeedItemCollectionUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *FeedItemCollectionUpsertBulk) SetName(v string) *FeedItemCollectionUpsertBulk {
	return u.Update(func(s *FeedItemCollectionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FeedItemCollectionUpsertBulk) UpdateName() *FeedItemCollectionUpsertBulk {
	return u.Update(func(s *FeedItemCollectionUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *FeedItemCollectionUpsertBulk) SetDescription(v string) *FeedItemCollectionUpsertBulk {
	return u.Update(func(s *FeedItemCollectionUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *FeedItemCollectionUpsertBulk) UpdateDescription() *FeedItemCollectionUpsertBulk {
	return u.Update(func(s *FeedItemCollectionUpsert) {
		s.UpdateDescription()
	})
}

// SetCategory sets the "category" field.
func (u *FeedItemCollectionUpsertBulk) SetCategory(v string) *FeedItemCollectionUpsertBulk {
	return u.Update(func(s *FeedItemCollectionUpsert) {
		s.SetCategory(v)
	})
}

// UpdateCategory sets the "category" field to the value that was provided on create.
func (u *FeedItemCollectionUpsertBulk) UpdateCategory() *FeedItemCollectionUpsertBulk {
	return u.Update(func(s *FeedItemCollectionUpsert) {
		s.UpdateCategory()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FeedItemCollectionUpsertBulk) SetUpdatedAt(v time.Time) *FeedItemCollectionUpsertBulk {
	return u.Update(func(s *FeedItemCollectionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FeedItemCollectionUpsertBulk) UpdateUpdatedAt() *FeedItemCollectionUpsertBulk {
	return u.Update(func(s *FeedItemCollectionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *FeedItemCollectionUpsertBulk) SetCreatedAt(v time.Time) *FeedItemCollectionUpsertBulk {
	return u.Update(func(s *FeedItemCollectionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FeedItemCollectionUpsertBulk) UpdateCreatedAt() *FeedItemCollectionUpsertBulk {
	return u.Update(func(s *FeedItemCollectionUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *FeedItemCollectionUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FeedItemCollectionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FeedItemCollectionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FeedItemCollectionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
